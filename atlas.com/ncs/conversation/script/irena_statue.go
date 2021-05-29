package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// IrenaStatue is located in Hidden Street - Quiet Ereve (913030000)
type IrenaStatue struct {
}

func (r IrenaStatue) NPCId() uint32 {
	return npc.IrenaStatue
}

func (r IrenaStatue) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("...")
	return SendOk(l, c, m.String())
}
