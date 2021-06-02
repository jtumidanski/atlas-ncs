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
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r MachineApparatus) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.DeepInsideTheClocktower, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.DeepInsideTheClocktower, c.NPCId)
	}
	return Exit()(l, c)
}
