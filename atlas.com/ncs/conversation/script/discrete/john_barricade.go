package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// JohnBarricade is located in Bigger Ben - Lobby (600020000)
type JohnBarricade struct {
}

func (r JohnBarricade) NPCId() uint32 {
	return npc.JohnBarricade
}

func (r JohnBarricade) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The patrol in New Leaf City is always ready. No creatures are able to break through to the city.")
	return script.SendOk(l, span, c, m.String())
}
