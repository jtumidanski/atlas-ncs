package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Shulynch2 is located in Hidden Street - Looking for Delli 1 (925010000)
type Shulynch2 struct {
}

func (r Shulynch2) NPCId() uint32 {
	return npc.Shulynch2
}

func (r Shulynch2) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		BlueText().ShowNPC(npc.Delli).
		BlackText().AddText(" must be some way up this cliff, according to our latest reports... Or are you saying you want to ").
		RedText().AddText("leave here").
		BlackText().AddText("?")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r Shulynch2) Warp(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.TrainingRoom, 0)(l, c)
}
