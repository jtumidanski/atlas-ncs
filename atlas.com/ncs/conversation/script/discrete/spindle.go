package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Spindle is located in Omega Sector - Omega Sector (221000000)
type Spindle struct {
}

func (r Spindle) NPCId() uint32 {
	return npc.Spindle
}

func (r Spindle) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Have you ever come to know about the card game based on MapleStory, the MapleStory iTCG?")
	return script.SendOk(l, span, c, m.String())
}
