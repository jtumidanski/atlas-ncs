package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Harry is located in Leafre - Before Takeoff <To Orbis> (240000111)
type Harry struct {
}

func (r Harry) NPCId() uint32 {
	return npc.Harry
}

func (r Harry) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to leave the flight?")
	return SendYesNo(l, c, m.String(), r.SeeYouNextTime, Exit())
}

func (r Harry) SeeYouNextTime(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return SendNext(l, c, m.String(), WarpById(_map.Station, 0))
}
