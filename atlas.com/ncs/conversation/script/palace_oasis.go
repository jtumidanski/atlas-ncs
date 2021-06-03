package script

import (
	"atlas-ncs/buff"
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// PalaceOasis is located in Ariant - Ariant Castle (260000300)
type PalaceOasis struct {
}

func (r PalaceOasis) NPCId() uint32 {
	return npc.PalaceOasis
}

func IsTigunMorphed(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return character.BuffSource(l)(characterId, buff.BuffMorph) == item.TigunTransformationBundle
	}
}

func (r PalaceOasis) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3900) && character.QuestProgressInt(l)(c.CharacterId, 3900, 0) != 5 {
		r.Refreshed(l, c)
	}
	if character.QuestCompleted(l)(c.CharacterId, 3938) {
		return r.LockOfHair(l, c)
	}
	if character.QuestStarted(l)(c.CharacterId, 3934) || (character.QuestCompleted(l)(c.CharacterId, 3934) && !character.QuestCompleted(l)(c.CharacterId, 3935)) {
		return r.FloatingInRiver(l, c)
	}
	return Exit()(l, c)
}

func (r PalaceOasis) FloatingInRiver(l logrus.FieldLogger, c Context) State {
	if !character.CanHold(l)(c.CharacterId, item.TigunTransformationBundle) {
		m := message.NewBuilder().AddText("You found a strange flask floating on the river. But you decided to ignore it since you don't have a USE slot available.")
		SendOk(l, c, m.String(), AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
	}
	character.GainItem(l)(c.CharacterId, item.TigunTransformationBundle, 1)
	m := message.NewBuilder().
		AddText("You managed to find a strange flask floating on the river. It seems like a transformation bottle mimicking one of the guards of the castle, maybe with it you will be able to roam inside freely.")
	return SendOk(l, c, m.String(), AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}

func (r PalaceOasis) LockOfHair(l logrus.FieldLogger, c Context) State {
	if !character.CanHold(l)(c.CharacterId, item.TigunTransformationBundle) {
		m := message.NewBuilder().AddText("You don't have a USE slot available.")
		SendOk(l, c, m.String(), AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
	}
	if !character.HasItem(l)(c.CharacterId, item.TigunTransformationBundle) && !IsTigunMorphed(l)(c.CharacterId) {
		character.GainItem(l)(c.CharacterId, item.TigunTransformationBundle, 1)
		m := message.NewBuilder().
			AddText("You found a lock of hair (probably Tigun's) floating by the water and caught it. Remembering how ").
			BlueText().AddText("Jano").
			BlackText().AddText(" made it last time, you crafted a new ").ShowItemName1(item.TigunTransformationBundle)
		return SendOk(l, c, m.String(), AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
	}
	return Exit()(l, c)
}

func (r PalaceOasis) Refreshed(l logrus.FieldLogger, c Context) {
	character.SetQuestProgress(l)(c.CharacterId, 3900, 0, 5)
	m := message.NewBuilder().BlueText().AddText("(You drink the water from the oasis and feel refreshed.)")
	SendOk(l, c, m.String(), AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}
