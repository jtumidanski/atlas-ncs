package discrete

import (
	"atlas-ncs/buff"
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
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

func (r PalaceOasis) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3900) && quest.ProgressInt(l)(c.CharacterId, 3900, 0) != 5 {
		r.Refreshed(l, span, c)
	}
	if quest.IsCompleted(l)(c.CharacterId, 3938) {
		return r.LockOfHair(l, span, c)
	}
	if quest.IsStarted(l)(c.CharacterId, 3934) || (quest.IsCompleted(l)(c.CharacterId, 3934) && !quest.IsCompleted(l)(c.CharacterId, 3935)) {
		return r.FloatingInRiver(l, span, c)
	}
	return script.Exit()(l, span, c)
}

func (r PalaceOasis) FloatingInRiver(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.TigunTransformationBundle) {
		m := message.NewBuilder().AddText("You found a strange flask floating on the river. But you decided to ignore it since you don't have a USE slot available.")
		script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
	}
	character.GainItem(l, span)(c.CharacterId, item.TigunTransformationBundle, 1)
	m := message.NewBuilder().
		AddText("You managed to find a strange flask floating on the river. It seems like a transformation bottle mimicking one of the guards of the castle, maybe with it you will be able to roam inside freely.")
	return script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}

func (r PalaceOasis) LockOfHair(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.TigunTransformationBundle) {
		m := message.NewBuilder().AddText("You don't have a USE slot available.")
		script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
	}
	if !character.HasItem(l, span)(c.CharacterId, item.TigunTransformationBundle) && !IsTigunMorphed(l)(c.CharacterId) {
		character.GainItem(l, span)(c.CharacterId, item.TigunTransformationBundle, 1)
		m := message.NewBuilder().
			AddText("You found a lock of hair (probably Tigun's) floating by the water and caught it. Remembering how ").
			BlueText().AddText("Jano").
			BlackText().AddText(" made it last time, you crafted a new ").ShowItemName1(item.TigunTransformationBundle)
		return script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
	}
	return script.Exit()(l, span, c)
}

func (r PalaceOasis) Refreshed(l logrus.FieldLogger, span opentracing.Span, c script.Context) {
	quest.SetProgress(l)(c.CharacterId, 3900, 0, 5)
	m := message.NewBuilder().BlueText().AddText("(You drink the water from the oasis and feel refreshed.)")
	script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}
