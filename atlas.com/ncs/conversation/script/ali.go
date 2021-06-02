package script

import (
	"atlas-ncs/character"
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

func (r Ali) Initial(l logrus.FieldLogger, c Context) State {
	character.RemoveAll(l)(c.CharacterId, item.PaperDocument)
	character.RemoveAll(l)(c.CharacterId, item.TheKey)
	character.RemoveAll(l)(c.CharacterId, item.FireOre)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.TheDoorToZakum, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.TheDoorToZakum, c.NPCId)
	}
	return Exit()(l, c)
}
