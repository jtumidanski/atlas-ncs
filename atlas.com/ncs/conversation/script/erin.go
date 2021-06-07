package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Erin is located in Orbis - Before Takeoff <To Ellinia> (200000112)
type Erin struct {
}

func (r Erin) NPCId() uint32 {
	return npc.Erin
}

func (r Erin) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to leave the boat?")
	return SendYesNo(l, c, m.String(), r.Alright, r.GoodChoice)
}

func (r Erin) GoodChoice(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Good choice.")
	return SendOk(l, c, m.String())
}

func (r Erin) Alright(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return SendNext(l, c, m.String(), WarpById(_map.StationToEllinia, 0))
}
