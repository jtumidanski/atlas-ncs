package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CrystalOfRoots struct {
}

func (r CrystalOfRoots) NPCId() uint32 {
	return npc.CrystalOfRoots
}

func (r CrystalOfRoots) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to leave?")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, script.Exit())
}

func (r CrystalOfRoots) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId > _map.EntranceToHorntailsCave {
		return script.WarpById(_map.CaveTheRoadInBetween, 0)(l, span, c)
	} else {
		return script.WarpByName(_map.CaveOfLifeEntrance, "out00")(l, span, c)
	}
}
