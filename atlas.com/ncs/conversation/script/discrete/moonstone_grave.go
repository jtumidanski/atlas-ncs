package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// MoonstoneGrave is located in MesoGears - Ice Chamber (600020500)
type MoonstoneGrave struct {
}

func (r MoonstoneGrave) NPCId() uint32 {
	return npc.MoonstoneGrave
}

func (r MoonstoneGrave) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("(This tombstone keeps emitting ever odder light waves the more I stare to it...)")
	return script.SendOk(l, span, c, m.String())
}
