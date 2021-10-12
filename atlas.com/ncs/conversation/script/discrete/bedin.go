package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Bedin is located in Zenumist Research Institute - Lab - 1st Floor Hallway (261010000)
type Bedin struct {
}

func (r Bedin) NPCId() uint32 {
	return npc.Bedin
}

func (r Bedin) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Make it clear of your position! Are you Zenumist or Alcadno?")
	return script.SendOk(l, span, c, m.String())
}
