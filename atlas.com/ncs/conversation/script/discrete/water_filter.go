package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// WaterFilter is located in The Nautilus - Bedroom (120000202)
type WaterFilter struct {
}

func (r WaterFilter) NPCId() uint32 {
	return npc.WaterFilter
}

func (r WaterFilter) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm bored! Someone come play with me!")
	return script.SendOk(l, span, c, m.String())
}
