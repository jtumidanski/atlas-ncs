package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// StrangeLookingStatue is located in Hidden Street - Puppeteer's Secret Passage (910510100)
type StrangeLookingStatue struct {
}

func (r StrangeLookingStatue) NPCId() uint32 {
	return npc.StrangeLookingStatue
}

func (r StrangeLookingStatue) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r StrangeLookingStatue) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Will you exit this trial?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r StrangeLookingStatue) Warp(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.PuppeteersHidingPlace, 2)(l, c)
}
