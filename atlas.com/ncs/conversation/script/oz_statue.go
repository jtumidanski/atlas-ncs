package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// OzStatue is located in Hidden Street - Quiet Ereve (913030000)
type OzStatue struct {
}

func (r OzStatue) NPCId() uint32 {
	return npc.OzStatue
}

func (r OzStatue) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("...")
	return SendOk(l, c, m.String())
}
