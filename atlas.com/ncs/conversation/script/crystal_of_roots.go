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
		return r.WarpToInBetween(l, c)
	} else {
		return r.WarpToEntrance(l, c)
	}
}

func (r CrystalOfRoots) WarpToInBetween(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.CaveTheRoadInBetween, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.CaveTheRoadInBetween, c.NPCId)
	}
	return Exit()(l, c)
}

func (r CrystalOfRoots) WarpToEntrance(l logrus.FieldLogger, c Context) State {
	err := npc.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.CaveOfLifeEntrance, "out00")
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.CaveOfLifeEntrance, c.NPCId)
	}
	return Exit()(l, c)
}
