package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// Gate is located in Crimsonwood  Keep - Hall of Mastery (610030010)
type Gate struct {
}

func (r Gate) NPCId() uint32 {
	return npc.Gate
}

func (r Gate) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.CrimsonwoodKeystone) {
		character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "The giant gate of iron will not budge no matter what, however there is a visible key-shaped socket.")
		return Exit()(l, c)
	}
	return WarpByName(_map.HallToInnerSanctum, "out00")(l, c)
}
