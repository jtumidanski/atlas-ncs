package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Purin is located in Victoria Road - Before Takeoff <To Orbis> (101000301)
type Purin struct {
}

func (r Purin) NPCId() uint32 {
	return npc.Purin
}

func (r Purin) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.DoYouWish(l, c)
}

func (r Purin) DoYouWish(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Do you wish to leave the boat?")
	return script.SendYesNo(l, c, m.String(), r.ProcessLeave, script.Exit())
}

func (r Purin) ProcessLeave(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright, see you next time. Take care.")
	return script.SendNext(l, c, m.String(), r.Warp)
}

func (r Purin) Warp(l logrus.FieldLogger, c script.Context) script.State {
	return script.WarpById(_map.ElliniaStation, 0)(l, c)
}
