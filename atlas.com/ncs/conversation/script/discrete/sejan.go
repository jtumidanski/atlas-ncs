package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Sejan is located in Ariant - The Town of Ariant (260000200)
type Sejan struct {
}

func (r Sejan) NPCId() uint32 {
	return npc.Sejan
}

func (r Sejan) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The light and dark always coexist...")
	return script.SendOk(l, span, c, m.String())
}
