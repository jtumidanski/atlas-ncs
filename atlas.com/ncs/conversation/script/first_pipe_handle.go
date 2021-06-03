package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// FirstPipeHandle is located in Magatia - Home of the Missing Alchemist (261000001)
type FirstPipeHandle struct {
}

func (r FirstPipeHandle) NPCId() uint32 {
	return npc.FirstPipeHandle
}

func (r FirstPipeHandle) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3339) {
		return r.Progress(l, c)
	}
	if character.QuestCompleted(l)(c.CharacterId, 3339) {
		return r.WarpInMap(l, c)
	}
	return Exit()(l, c)
}

func (r FirstPipeHandle) GetPassword(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The pipe reacts as the water starts flowing. A secret compartment with a keypad shows up. ").
		BlueText().AddText("Password").
		BlackText().AddText("!")
	return SendGetText(l, c, m.String(), r.Validate)
}

func (r FirstPipeHandle) WarpInMap(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.HomeOfTheMissingAlchemist, 1)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.HomeOfTheMissingAlchemist, c.NPCId)
	}
	return Exit()(l, c)
}

func (r FirstPipeHandle) Progress(l logrus.FieldLogger, c Context) State {
	progress := character.QuestProgressInt(l)(c.CharacterId, 23339, 1)
	if progress == 3 {
		return r.GetPassword(l, c)
	} else if progress == 0 {
		character.SetQuestProgress(l)(c.CharacterId, 23339, 1, 1)
		return Exit()(l, c)
	} else if progress < 3 {
		character.SetQuestProgress(l)(c.CharacterId, 23339, 1, 0)
		return Exit()(l, c)
	} else {
		return r.WarpInMap(l, c)
	}
}

func (r FirstPipeHandle) Validate(text string) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if text != "my love Phyllia" {
			m := message.NewBuilder().RedText().AddText("Wrong!")
			return SendOk(l, c, m.String())
		}
		character.SetQuestProgress(l)(c.CharacterId, 23339, 1, 4)
		return r.WarpInMap(l, c)
	}
}
