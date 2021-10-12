package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Kiridu is located in Empress' Road - Kiridu's Hatchery (130010220)
type Kiridu struct {
}

func (r Kiridu) NPCId() uint32 {
	return npc.Kiridu
}

func (r Kiridu) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Yo. I am ").
		ShowNPC(npc.Kiridu).
		AddText(", in charge of mount raising and training for the Cygnus Knights' of Ereve!")
	return script.SendOk(l, span, c, m.String())
}
