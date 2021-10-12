package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Ardin is located in Ariant - The Town of Ariant (260000200)
type Ardin struct {
}

func (r Ardin) NPCId() uint32 {
	return npc.Ardin
}

func (r Ardin) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hey hey, don't try to start trouble with anyone. I want nothing to do with you.")
	return script.SendOk(l, span, c, m.String())
}
