package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Slyn is located in Ariant - Before Takeoff <To Orbis> (260000110)
type Slyn struct {
}

func (r Slyn) NPCId() uint32 {
	return npc.Slyn
}

func (r Slyn) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to leave the genie?")
	return SendYesNo(l, c, m.String(), r.NextTime, Exit())
}

func (r Slyn) NextTime(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r Slyn) Warp(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.AriantStationPlatform, 0)(l, c)
}
