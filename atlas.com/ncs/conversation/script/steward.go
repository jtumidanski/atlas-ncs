package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Steward struct {
}

func (r Steward) NPCId() uint32 {
	return npc.Steward
}

func (r Steward) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("At your service, my friend.")
	return SendOk(l, c, m.String())
}
