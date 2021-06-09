package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// Ali is located in Adobis's Mission I - The Room of Tragedy (280090000)
type Ali struct {
}

func (r Ali) NPCId() uint32 {
	return npc.Ali
}

func (r Ali) Initial(l logrus.FieldLogger, c script.Context) script.State {
	character.RemoveAll(l)(c.CharacterId, item.PaperDocument)
	character.RemoveAll(l)(c.CharacterId, item.TheKey)
	character.RemoveAll(l)(c.CharacterId, item.FireOre)
	return script.WarpById(_map.TheDoorToZakum, 0)(l, c)
}
