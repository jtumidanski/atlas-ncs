package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type CrystalOfRoots struct {
}

func (r CrystalOfRoots) NPCId() uint32 {
	return npc.CrystalOfRoots
}

func (r CrystalOfRoots) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to leave?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r CrystalOfRoots) Warp(l logrus.FieldLogger, c Context) State {
	if c.MapId > _map.EntranceToHorntailsCave {
		return WarpById(_map.CaveTheRoadInBetween, 0)(l, c)
	} else {
		return WarpByName(_map.CaveOfLifeEntrance, "out00")(l, c)
	}
}
