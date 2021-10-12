package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// PrinceGiuseppe is located in Mushroom Castle - Wedding Hall (106021600)
type PrinceGiuseppe struct {
}

func (r PrinceGiuseppe) NPCId() uint32 {
	return npc.PrinceGiuseppe
}

func (r PrinceGiuseppe) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey guys, what's going on? I'm already getting married at this age? But I'm only a child!!! How comes?")
	return script.SendOk(l, span, c, m.String())
}
