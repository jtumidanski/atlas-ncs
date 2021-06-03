package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// QueensCabinet is located in Ariant Castle - King's Room (260000303)
type QueensCabinet struct {
}

func (r QueensCabinet) NPCId() uint32 {
	return npc.QueensCabinet
}

func (r QueensCabinet) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3923) && !character.HasItem(l)(c.CharacterId, item.QueensRing) {
		if !character.CanHold(l)(c.CharacterId, item.QueensRing) {
			return r.MakeRoom(l, c)
		}
		return r.ClearOut(l, c)
	}
	return Exit()(l, c)
}

func (r QueensCabinet) ClearOut(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.QueensRing, 1)
	m := message.NewBuilder().AddText("You have just swiped the ring. Clear the area asap!")
	return SendOk(l, c, m.String(), AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}

func (r QueensCabinet) MakeRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You don't have a ETC slot available.")
	return SendOk(l, c, m.String(), AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}
