package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// OldManTom is located in 
type OldManTom struct {
}

func (r OldManTom) NPCId() uint32 {
	return npc.OldManTom
}

func (r OldManTom) Initial(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.HauntedHouseFoyer, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.HauntedHouseFoyer, c.NPCId)
	}
	return Exit()(l, c)
}
