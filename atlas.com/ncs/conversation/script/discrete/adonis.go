package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Adonis is located in El Nath - El Nath (211000000)
type Adonis struct {
}

func (r Adonis) NPCId() uint32 {
	return npc.Adonis
}

func (r Adonis) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I came from far-away places looking for people powerful enough to join my expedition against the evil that lays waste on this land. Are you, by any chance, one of those people?")
	return script.SendOk(l, span, c, m.String())
}
