package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Kiriko is located in Hidden Street - The 2nd Drill Hall (108000600)
type Kiriko struct {
}

func (r Kiriko) NPCId() uint32 {
	return npc.Kiriko
}

func (r Kiriko) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Would you like to exit the drill hall?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r Kiriko) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.EntranceToTheDrillHall, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.EntranceToTheDrillHall, c.NPCId)
	}
	return nil
}
