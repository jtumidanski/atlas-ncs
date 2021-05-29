package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// EckhartStatue is located in Hidden Street - Quiet Ereve (913030000)
type EckhartStatue struct {
}

func (r EckhartStatue) NPCId() uint32 {
	return npc.EckhartStatue
}

func (r EckhartStatue) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("...")
	return SendOk(l, c, m.String())
}
