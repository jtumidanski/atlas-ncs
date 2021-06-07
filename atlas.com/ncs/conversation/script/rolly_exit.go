package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// RollyExit is located in Ludibrium - Exit of the Maze (809050017)
type RollyExit struct {
}

func (r RollyExit) NPCId() uint32 {
	return npc.RollyExit
}

func (r RollyExit) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Come this way to return to Ludibrium.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r RollyExit) Warp(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.Ludibrium, 0)(l, c)
}
