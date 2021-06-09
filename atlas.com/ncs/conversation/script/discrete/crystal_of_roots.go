package discrete

import (
	"atlas-ncs/conversation/script"
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

func (r CrystalOfRoots) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to leave?")
	return script.SendYesNo(l, c, m.String(), r.Warp, script.Exit())
}

func (r CrystalOfRoots) Warp(l logrus.FieldLogger, c script.Context) script.State {
	if c.MapId > _map.EntranceToHorntailsCave {
		return script.WarpById(_map.CaveTheRoadInBetween, 0)(l, c)
	} else {
		return script.WarpByName(_map.CaveOfLifeEntrance, "out00")(l, c)
	}
}
