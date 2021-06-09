package discrete

import (
	"atlas-ncs/conversation/script"
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

func (r Harry) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to leave the flight?")
	return script.SendYesNo(l, c, m.String(), r.SeeYouNextTime, script.Exit())
}

func (r Harry) SeeYouNextTime(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return script.SendNext(l, c, m.String(), script.WarpById(_map.Station, 0))
}
