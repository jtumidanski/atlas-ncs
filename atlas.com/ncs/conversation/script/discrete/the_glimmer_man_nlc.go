package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// TheGlimmerManNLC is located in New Leaf City Town Street - New Leaf City - Town Center (600000000)
type TheGlimmerManNLC struct {
}

func (r TheGlimmerManNLC) NPCId() uint32 {
	return npc.TheGlimmerManNLC
}

func (r TheGlimmerManNLC) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The patrol in New Leaf City is always ready. No creatures are able to break through to the city.")
	return script.SendOk(l, span, c, m.String())
}
