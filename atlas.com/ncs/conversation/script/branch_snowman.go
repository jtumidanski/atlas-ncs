package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// BranchSnowman is located in Hidden Street - Happyville (209000000)
type BranchSnowman struct {
}

func (r BranchSnowman) NPCId() uint32 {
	return npc.BranchSnowman
}

func (r BranchSnowman) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("We have a beautiful christmas tree.").NewLine().
		AddText("Do you want to see/decorate it?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r BranchSnowman) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.TheHillOfChristmas1, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.TheHillOfChristmas1, c.NPCId)
	}
	return Exit()(l, c)
}
