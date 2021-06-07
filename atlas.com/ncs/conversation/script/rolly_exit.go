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
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.Ludibrium, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.Ludibrium, c.NPCId)
	}
	return Exit()(l, c)
}
