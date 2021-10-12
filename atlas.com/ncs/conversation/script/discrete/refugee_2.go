package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Refugee2 is located in Black Road - Ready to Leave (914000100)
type Refugee2 struct {
}

func (r Refugee2) NPCId() uint32 {
	return npc.Refugee2
}

func (r Refugee2) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please calm down, uncle. We are embarking to ").
		BlueText().AddText("Victoria Island").
		BlackText().AddText(", we will be safe once we reach there. So, come on!")
	return script.SendOk(l, span, c, m.String())
}
