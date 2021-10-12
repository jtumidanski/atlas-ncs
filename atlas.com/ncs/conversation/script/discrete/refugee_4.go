package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Refugee4 is located in Black Road - Ready to Leave (914000100)
type Refugee4 struct {
}

func (r Refugee4) NPCId() uint32 {
	return npc.Refugee4
}

func (r Refugee4) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright, embarking in...")
	return script.SendOk(l, span, c, m.String())
}
