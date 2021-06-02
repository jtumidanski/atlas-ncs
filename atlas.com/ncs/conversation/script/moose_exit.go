package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MooseExit is located in Hidden Street - On the Way to the Practice Field (924000000), Hidden Street - Moose's Practice Field (924000001), and Hidden Street - Exiting the Practice Field (924000002)
type MooseExit struct {
}

func (r MooseExit) NPCId() uint32 {
	return npc.MooseExit
}

func (r MooseExit) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you want to exit the area? If you quit, you will need to start this task from the scratch.")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r MooseExit) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ForestCrossroad, "st00")
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ForestCrossroad, c.NPCId)
	}
	return Exit()(l, c)
}
