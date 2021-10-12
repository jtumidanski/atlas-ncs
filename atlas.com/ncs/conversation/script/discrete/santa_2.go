package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Santa2 is located in Hidden Street - Happyville (209000000)
type Santa2 struct {
}

func (r Santa2) NPCId() uint32 {
	return npc.Santa2
}

func (r Santa2) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Welcome to ").
		BlueText().AddText("Happyville").
		BlackText().AddText(", young traveler. Do you have any wishes?")
	return script.SendOk(l, span, c, m.String())
}
