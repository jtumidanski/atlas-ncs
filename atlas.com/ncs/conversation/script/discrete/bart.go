package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Bart is located in Victoria Road - Nautilus Harbor (120000000)
type Bart struct {
}

func (r Bart) NPCId() uint32 {
	return npc.Bart
}

func (r Bart) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I need to keep my eyes wide open to look for the enemy although my sea gull friends help me out so it's not all that bad.")
	return script.SendOk(l, span, c, m.String())
}
