package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Ria is located in Victoria Road - Ellinia (101000000)
type Ria struct {
}

func (r Ria) NPCId() uint32 {
	return npc.Ria
}

func (r Ria) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hi, I'm ").ShowNPC(npc.Ria).AddText(".")
	return script.SendOk(l, span, c, m.String())
}
