package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ScarfSnowman is located in Hidden Street - The Hill of Christmas (209000001-209000015)
type ScarfSnowman struct {
}

func (r ScarfSnowman) NPCId() uint32 {
	return npc.ScarfSnowman
}

func (r ScarfSnowman) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, are you ready to head out of here?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r ScarfSnowman) Warp(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.Happyville, 5)(l, c)
}
