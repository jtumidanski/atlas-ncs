package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// QueensCabinet is located in Ariant Castle - King's Room (260000303)
type QueensCabinet struct {
}

func (r QueensCabinet) NPCId() uint32 {
	return npc.QueensCabinet
}

func (r QueensCabinet) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3923) && !character.HasItem(l, span)(c.CharacterId, item.QueensRing) {
		if !character.CanHold(l)(c.CharacterId, item.QueensRing) {
			return r.MakeRoom(l, span, c)
		}
		return r.ClearOut(l, span, c)
	}
	return script.Exit()(l, span, c)
}

func (r QueensCabinet) ClearOut(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.QueensRing, 1)
	m := message.NewBuilder().AddText("You have just swiped the ring. Clear the area asap!")
	return script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}

func (r QueensCabinet) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't have a ETC slot available.")
	return script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
}
