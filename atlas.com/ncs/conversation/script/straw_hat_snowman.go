package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// StrawHatSnowman is located in Hidden Street - Happyville (209000000)
type StrawHatSnowman struct {
}

func (r StrawHatSnowman) NPCId() uint32 {
	return npc.StrawHatSnowman
}

func (r StrawHatSnowman) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("We have a beautiful christmas tree.").NewLine().
		AddText("Do you want to see/decorate it?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r StrawHatSnowman) Warp(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.TheHillOfChristmas3, 0)(l, c)
}
