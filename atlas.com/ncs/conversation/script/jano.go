package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// Jano is located in Town of Ariant - An Old, Empty House (260000201)
type Jano struct {
}

func (r Jano) NPCId() uint32 {
	return npc.Jano
}

func (r Jano) Initial(l logrus.FieldLogger, c Context) State {
	character.ChangeMusic(l)(c.CharacterId, "Bgm14/Ariant")
	return WarpById(_map.Stage1MagikMirror, 3)(l, c)
}
