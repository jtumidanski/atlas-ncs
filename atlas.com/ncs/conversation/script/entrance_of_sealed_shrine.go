package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// EntranceOfSealedShrine is located in Mu Lung - Practice Field : Advanced Level (250020300)
type EntranceOfSealedShrine struct {
}

func (r EntranceOfSealedShrine) NPCId() uint32 {
	return npc.EntranceOfSealedShrine
}

func (r EntranceOfSealedShrine) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The entrance of the Sealed Shrine... ").
		BlueText().AddText("Password").
		BlackText().AddText("!")
	return SendGetText(l, c, m.String(), r.ProcessPassword)
}

func (r EntranceOfSealedShrine) ProcessPassword(text string) StateProducer {
	if text != "Actions speak louder than words" {
		return r.Wrong
	}
	return func(l logrus.FieldLogger, c Context) State {
		if _map.CharacterCount(l)(c.WorldId, c.ChannelId, _map.SealedTemple) > 0 {
			return r.AlreadyAttending(l, c)
		}
		if character.QuestStarted(l)(c.CharacterId, 21747) && character.QuestProgressInt(l)(c.CharacterId, 21747, 9300351) == 0 {
			return r.WarpToTemple(l, c)
		} else {
			return r.CorrectButMissingPrerequisites(l, c)
		}
	}
}

func (r EntranceOfSealedShrine) CorrectButMissingPrerequisites(l logrus.FieldLogger, c Context) State {
	character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "Although you said the right answer, some mysterious forces are blocking the way in.")
	return Exit()(l, c)
}

func (r EntranceOfSealedShrine) WarpToTemple(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.SealedTemple, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.SealedTemple, c.NPCId)
	}
	return Exit()(l, c)
}

func (r EntranceOfSealedShrine) AlreadyAttending(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Someone is already attending the Sealed Shrine.")
	return SendOk(l, c, m.String())
}

func (r EntranceOfSealedShrine) Wrong(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().RedText().AddText("Wrong!")
	return SendOk(l, c, m.String())
}
