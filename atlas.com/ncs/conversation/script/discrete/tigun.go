package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Tigun is located in Ariant - Ariant Castle (260000300)
type Tigun struct {
}

func (r Tigun) NPCId() uint32 {
	return npc.Tigun
}

func (r Tigun) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This isn't much of a ").
		BlueText().AddText("palace").
		BlackText().AddText(" but it'll do until we can get an even better palace built! Anyone's allowed to speak to the king and queen, but don't expect niceness from either of them. Well... maybe King Abdullah VIII, if you catch him in his non-lazy and not-paying-attention moods.")
	return script.SendOk(l, span, c, m.String())
}
