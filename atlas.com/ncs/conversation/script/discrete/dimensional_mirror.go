package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DimensionalMirror is located in all towns
type DimensionalMirror struct {
}

func (r DimensionalMirror) NPCId() uint32 {
	return npc.DimensionalMirror
}

func (r DimensionalMirror) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.IsLevel(l)(c.CharacterId, 20) {
		m := message.NewBuilder().AddText("There is no place for you to transport to from here.")
		return script.SendDimensionalMirror(l, c, m.String(), r.Selection)
	}

	m := message.NewBuilder()
	if character.MeetsCriteria(l)(c.CharacterId, character.IsLevelBetweenCriteria(20, 30)) {
		m = m.DimensionalMirrorOption(0, "Ariant Coliseum")
	}
	if character.IsLevel(l)(c.CharacterId, 25) {
		m = m.DimensionalMirrorOption(1, "Mu Lung Dojo")
	}
	if character.MeetsCriteria(l)(c.CharacterId, character.IsLevelBetweenCriteria(30, 50)) {
		m = m.DimensionalMirrorOption(2, "Monster Carnival 1")
	}
	if character.MeetsCriteria(l)(c.CharacterId, character.IsLevelBetweenCriteria(51, 70)) {
		m = m.DimensionalMirrorOption(3, "Monster Carnival 2")
	}
	//TODO what is 4?
	if character.IsLevel(l)(c.CharacterId, 40) {
		m = m.DimensionalMirrorOption(5, "Nett's Pyramid")
	}
	if character.MeetsCriteria(l)(c.CharacterId, character.IsLevelBetweenCriteria(25, 30)) {
		m = m.DimensionalMirrorOption(6, "Construction Site")
	}
	return script.SendDimensionalMirror(l, c, m.String(), r.Selection)
}

func (r DimensionalMirror) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return script.WarpById(_map.BattleArenaLobby, 3)
	case 1:
		return script.WarpById(_map.MuLungDojoEntrance, 0)
	case 2:
		return r.SaveAndWarpById("MONSTER_CARNIVAL", _map.SpiegelmannsOffice, 3)
	case 3:
		return r.SaveAndWarpById("MONSTER_CARNIVAL", _map.SpiegelmannsOffice2, 3)
	case 5:
		return script.WarpById(_map.PyramidDunes, 4)
	case 6:
		return script.WarpById(_map.AbandonedSubwayStation, 2)

	}
	return nil
}

func (r DimensionalMirror) SaveAndWarpById(location string, mapId uint32, portalId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		character.SaveLocation(l)(c.CharacterId, location)
		return script.WarpById(mapId, portalId)(l, c)
	}
}
