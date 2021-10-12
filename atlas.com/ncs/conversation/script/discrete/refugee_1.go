package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Refugee1 is located in Black Road - Ready to Leave (914000100)
type Refugee1 struct {
}

func (r Refugee1) NPCId() uint32 {
	return npc.Refugee1
}

func (r Refugee1) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The ").
		RedText().AddText("Black Magician").
		BlackText().AddText("'s forces approaches here in an unstoppable pace... We have no other way than to flee this area now, leaving our home behind. Oh, the tragedy!")
	return script.SendOk(l, span, c, m.String())
}
