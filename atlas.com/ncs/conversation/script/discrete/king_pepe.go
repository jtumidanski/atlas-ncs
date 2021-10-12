package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// KingPepe is located in Mushroom Castle - Wedding Hall (106021600)
type KingPepe struct {
}

func (r KingPepe) NPCId() uint32 {
	return npc.KingPepe
}

func (r KingPepe) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Let the ceremony begins, we cannot let the masses waiting! Hem~hem~heeh~~")
	return script.SendOk(l, span, c, m.String())
}
