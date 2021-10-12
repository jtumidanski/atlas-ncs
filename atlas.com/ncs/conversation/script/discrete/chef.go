package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Chef is located in Victoria Road - Lith Harbor (104000000)
type Chef struct {
}

func (r Chef) NPCId() uint32 {
	return npc.Chef
}

func (r Chef) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r Chef) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hi, I'm ").
		BlueText().ShowNPC(r.NPCId()).
		BlackText().AddText(". Nice to meet you.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}
