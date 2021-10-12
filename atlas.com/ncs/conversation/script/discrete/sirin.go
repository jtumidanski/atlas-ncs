package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Sirin is located in Ariant - The Town of Ariant (260000200)
type Sirin struct {
}

func (r Sirin) NPCId() uint32 {
	return npc.Sirin
}

func (r Sirin) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Just dancing well is not enough for me. I want to do a marvelous brilliant dance!")
	return script.SendOk(l, span, c, m.String())
}
