package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Schegerazade is located in Ariant Castle - King's Room (260000303)
type Schegerazade struct {
}

func (r Schegerazade) NPCId() uint32 {
	return npc.Schegerazade
}

func (r Schegerazade) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The King and Queen are so bossy and demanding lately. I only get to see my family every Sunday or whenever they come visit. But like me, they're poor and are in need of mesos... for some reason Tigun doesn't allow the poor to enter.")
	return script.SendOk(l, span, c, m.String())
}
