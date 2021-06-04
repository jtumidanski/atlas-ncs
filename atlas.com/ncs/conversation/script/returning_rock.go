package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type ReturningRock struct {
}

func (r ReturningRock) NPCId() uint32 {
	return npc.ReturningRock
}

func (r ReturningRock) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to exit the Guild Quest?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r ReturningRock) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.Warp(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ExcavationSiteCamp)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ExcavationSiteCamp, c.NPCId)
	}
	return Exit()(l, c)
}
