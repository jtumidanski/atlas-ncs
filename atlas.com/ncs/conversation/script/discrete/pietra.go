package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Pietra is located in Hidden Street - Leaving the Event (109050001)
type Pietra struct {
}

func (r Pietra) NPCId() uint32 {
	return npc.Pietra
}

func (r Pietra) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I'm sorry but I'm afraid you didn't win the event. Try it again some other time. You can return to where you were through me.")
	return script.SendOk(l, span, c, m.String())
}
