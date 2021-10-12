package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Maed is located in Magatia - Alcadno Society (261000020)
type Maed struct {
}

func (r Maed) NPCId() uint32 {
	return npc.Maed
}

func (r Maed) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Zenumist......I know what they say. They don't like combination the of life with machine. But it is about being fearful of machine only. Seeking Pure Alchemy won't achieve anything.")
	return script.SendOk(l, span, c, m.String())
}
