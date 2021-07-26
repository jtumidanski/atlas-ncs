package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// QueensCabinet is located in Ariant Castle - King's Room (260000303)
type QueensCabinet struct {
}

func (r QueensCabinet) NPCId() uint32 {
	return npc.QueensCabinet
}

func (r QueensCabinet) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3923) && !character.HasItem(l)(c.CharacterId, item.QueensRing) {
		if !character.CanHold(l)(c.CharacterId, item.QueensRing) {
			return r.MakeRoom(l, c)
		}
		return r.ClearOut(l, c)
	}
	return script.Exit()(l, c)
}

func (r QueensCabinet) ClearOut(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.QueensRing, 1)
	m := message.NewBuilder().AddText("You have just swiped the ring. Clear the area asap!")
	return script.SendOk(l, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}

func (r QueensCabinet) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't have a ETC slot available.")
	return script.SendOk(l, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}
