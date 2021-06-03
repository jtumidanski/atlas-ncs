package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ControlDevice is located in Hidden Street - Authorized Personnel Only (261020401)
type ControlDevice struct {
}

func (r ControlDevice) NPCId() uint32 {
	return npc.ControlDevice
}

func (r ControlDevice) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("This control device seems to be monitoring something...")
	return SendOk(l, c, m.String())
}
