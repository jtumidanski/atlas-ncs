package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Gate is located in Crimsonwood  Keep - Hall of Mastery (610030010)
type Gate struct {
}

func (r Gate) NPCId() uint32 {
	return npc.Gate
}

func (r Gate) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.CrimsonwoodKeystone) {
		character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "The giant gate of iron will not budge no matter what, however there is a visible key-shaped socket.")
		return script.Exit()(l, span, c)
	}
	return script.WarpByName(_map.HallToInnerSanctum, "out00")(l, span, c)
}
