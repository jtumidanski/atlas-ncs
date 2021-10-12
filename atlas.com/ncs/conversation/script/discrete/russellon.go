package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Russellon is located in Alcadno Research Institute - Lab - Center Gate (261020000)
type Russellon struct {
}

func (r Russellon) NPCId() uint32 {
	return npc.Russellon
}

func (r Russellon) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Definitely Alcadno is excluded.....Huhuhuhu.....Stupid people....cannot see the real important things..")
	return script.SendOk(l, span, c, m.String())
}
