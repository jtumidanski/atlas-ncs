package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ThirdPipeHandle is located in Magatia - Home of the Missing Alchemist (261000001)
type ThirdPipeHandle struct {
}

func (r ThirdPipeHandle) NPCId() uint32 {
	return npc.ThirdPipeHandle
}

func (r ThirdPipeHandle) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3339) {
		return r.Progress(l, c)
	}
	if character.QuestCompleted(l)(c.CharacterId, 3339) {
		return r.WarpInMap(l, c)
	}
	return Exit()(l, c)
}

func (r ThirdPipeHandle) GetPassword(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The pipe reacts as the water starts flowing. A secret compartment with a keypad shows up. ").
		BlueText().AddText("Password").
		BlackText().AddText("!")
	return SendGetText(l, c, m.String(), r.Validate)
}

func (r ThirdPipeHandle) WarpInMap(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.HomeOfTheMissingAlchemist, 1)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.HomeOfTheMissingAlchemist, c.NPCId)
	}
	return Exit()(l, c)
}

func (r ThirdPipeHandle) Progress(l logrus.FieldLogger, c Context) State {
	progress := character.QuestProgressInt(l)(c.CharacterId, 23339, 1)
	if progress == 3 {
		return r.GetPassword(l, c)
	} else if progress == 1 {
		character.SetQuestProgress(l)(c.CharacterId, 23339, 1, 2)
		return Exit()(l, c)
	} else if progress < 3 {
		character.SetQuestProgress(l)(c.CharacterId, 23339, 1, 0)
		return Exit()(l, c)
	} else {
		return r.WarpInMap(l, c)
	}
}

func (r ThirdPipeHandle) Validate(text string) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if text != "my love Phyllia" {
			m := message.NewBuilder().RedText().AddText("Wrong!")
			return SendOk(l, c, m.String())
		}
		character.SetQuestProgress(l)(c.CharacterId, 23339, 1, 4)
		return r.WarpInMap(l, c)
	}
}
