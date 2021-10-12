package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Vikin is located in Victoria Road - Lith Harbor (104000000)
type Vikin struct {
}

func (r Vikin) NPCId() uint32 {
	return npc.Vikin
}

func (r Vikin) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey hey!!! Find the Treasure Scroll! I lost the map").NewLine().
		AddText("somewhere and I can't leave without it.")
	return script.SendOk(l, span, c, m.String())
}
