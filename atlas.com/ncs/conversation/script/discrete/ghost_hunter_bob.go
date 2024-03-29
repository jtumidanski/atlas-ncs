package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// GhostHunterBob is located in Ludibrium - Forgotten Path of Time<1> (220070000)
type GhostHunterBob struct {
}

func (r GhostHunterBob) NPCId() uint32 {
	return npc.GhostHunterBob
}

func (r GhostHunterBob) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.TimersEgg) {
		return r.Hello(l, span, c)
	}
	return r.TakeIt(l, span, c)
}

func (r GhostHunterBob) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hello there! I'm ").
		BlueText().ShowNPC(npc.GhostHunterBob).
		BlackText().AddText(", in charge of watching and reporting any paranormal activities in this area.")
	return script.SendOk(l, span, c, m.String())
}

func (r GhostHunterBob) TakeIt(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.TimersEgg, -1)
	m := message.NewBuilder().
		AddText("You want to hand the ").
		RedText().ShowItemName1(item.TimersEgg).
		BlackText().AddText(" to me, right? Alright, I'll take it for you.")
	return script.SendOk(l, span, c, m.String())
}
