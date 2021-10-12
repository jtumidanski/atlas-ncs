package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// AbdullahVIII is located in Ariant Castle - King's Room (260000303)
type AbdullahVIII struct {
}

func (r AbdullahVIII) NPCId() uint32 {
	return npc.AbdullahVIII
}

func (r AbdullahVIII) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Yawnnnn~!")
	return script.SendOk(l, span, c, m.String())
}
