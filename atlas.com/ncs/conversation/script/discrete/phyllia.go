package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Phyllia is located in Sunset Road - Magatia (261000000)
type Phyllia struct {
}

func (r Phyllia) NPCId() uint32 {
	return npc.Phyllia
}

func (r Phyllia) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Even though Alchemists and Fairies have antipathy for each other, I can still be in this town because the Zenumist President has been protecting us.")
	return script.SendOk(l, span, c, m.String())
}
