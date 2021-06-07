package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Papulatus is located in Hidden Street - Origin of the Clock Tower (922020300)
type Papulatus struct {
}

func (r Papulatus) NPCId() uint32 {
	return npc.Papulatus
}

func (r Papulatus) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You don't belong to this world... Return now.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r Papulatus) Warp(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.DeepInsideTheClocktower, 0)(l, c)
}
