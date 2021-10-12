package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Neinheart is located in Hidden Street - Quiet Ereve (913030000)
type Neinheart struct {
}

func (r Neinheart) NPCId() uint32 {
	return npc.NeinheartStatue
}

func (r Neinheart) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("...")
	return script.SendOk(l, span, c, m.String())
}
