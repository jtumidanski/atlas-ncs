package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// PracticeChart is located in Mu Lung - Practice Field : Beginner (250020000)
type PracticeChart struct {
}

func (r PracticeChart) NPCId() uint32 {
	return npc.PracticeChart
}

func (r PracticeChart) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Amateurs train on this map. Adepts train on the next. Professionals train on the last, where the boss will be awaiting.")
	return script.SendOk(l, span, c, m.String())
}
