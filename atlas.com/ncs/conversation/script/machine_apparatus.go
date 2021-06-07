package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MachineApparatus is located in Ludibrium - Origin of Clocktower (220080001)
type MachineApparatus struct {
}

func (r MachineApparatus) NPCId() uint32 {
	return npc.MachineApparatus
}

func (r MachineApparatus) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Beep... beep... you can make your escape to a safer place through me. Beep... beep... would you like to leave this place?")
	return SendYesNo(l, c, m.String(), WarpById(_map.DeepInsideTheClocktower, 0), Exit())
}
