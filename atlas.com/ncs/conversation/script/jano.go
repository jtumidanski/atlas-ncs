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
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.Stage1MagikMirror, 3)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.Stage1MagikMirror, c.NPCId)
	}
	return Exit()(l, c)
}
