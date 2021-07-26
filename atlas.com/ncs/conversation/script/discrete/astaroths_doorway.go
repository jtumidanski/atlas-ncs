package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// AstarothsDoorway is located in 
type AstarothsDoorway struct {
}

func (r AstarothsDoorway) NPCId() uint32 {
	return npc.AstarothsDoorway
}

func (r AstarothsDoorway) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if c.MapId == _map.FogForest {
		return r.Exit(l, c)
	}
	if c.MapId == _map.AstarothStrollingPlace {
		return r.MoveTo(_map.AstarothHidingPlace)(l, c)
	}
	if c.MapId == _map.AstarothHidingPlace {
		return r.Exit(l, c)
	}
	if !quest.IsStarted(l)(c.CharacterId, 28283) {
		return r.EntranceBlocked(l, c)
	}
	if !character.HasEquipped(l)(c.CharacterId, item.WildEyesGasMask) {
		return r.EquipMask(l, c)
	}
	return r.MoveTo(_map.FogForest)(l, c)
}

func (r AstarothsDoorway) EquipMask(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The path ahead has a weird stench... Equip the ").
		RedText().AddText("gas mask").
		BlackText().AddText(" before entering.")
	return script.SendOk(l, c, m.String())
}

func (r AstarothsDoorway) EntranceBlocked(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
	return script.SendOk(l, c, m.String())
}

func (r AstarothsDoorway) MoveTo(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Would you like to move to ").
			BlueText().ShowMap(mapId).
			BlackText().AddText("?")
		return script.SendYesNo(l, c, m.String(), script.Warp(mapId), script.Exit())
	}
}

func (r AstarothsDoorway) Exit(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Would you like to ").
		BlueText().AddText("exit this place").
		BlackText().AddText("?")
	return script.SendYesNo(l, c, m.String(), script.Warp(_map.DarkCave), script.Exit())
}
