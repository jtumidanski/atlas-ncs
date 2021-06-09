package discrete

import (
	"atlas-ncs/conversation/script"
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

func (r OldManTom) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return script.WarpById(_map.HauntedHouseFoyer, 0)(l, c)
}
