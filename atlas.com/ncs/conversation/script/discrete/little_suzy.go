package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// LittleSuzy is located in New Leaf City Town Street - New Leaf City - Town Center (600000000)
type LittleSuzy struct {
}

func (r LittleSuzy) NPCId() uint32 {
	return npc.LittleSuzy
}

func (r LittleSuzy) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Have you heard the fantastic Jack Masque appeared around the city these days? That is sooooo nice!")
	return script.SendOk(l, span, c, m.String())
}
