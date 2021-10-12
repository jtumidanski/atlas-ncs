package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Mia is located in Victoria Road - Henesys (100000000)
type Mia struct {
}

func (r Mia) NPCId() uint32 {
	return npc.Mia
}

func (r Mia) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hi, I'm ").ShowNPC(npc.Mia).AddText(".")
	return script.SendOk(l, span, c, m.String())
}
