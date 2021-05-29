package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Neinheart is located in Hidden Street - Quiet Ereve (913030000)
type Neinheart struct {
}

func (r Neinheart) NPCId() uint32 {
	return npc.NeinheartStatue
}

func (r Neinheart) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("...")
	return SendOk(l, c, m.String())
}
