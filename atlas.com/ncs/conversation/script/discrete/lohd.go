package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Lohd is located in Ellin Forest - Mushroom Hill Entrance (300020200)
type Lohd struct {
}

func (r Lohd) NPCId() uint32 {
	return npc.Lohd
}

func (r Lohd) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Call me Dark Lord. I will give thieves a place in society... watch in a few years!")
	return script.SendOk(l, span, c, m.String())
}
