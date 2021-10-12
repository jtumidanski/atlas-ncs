package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Jano is located in Town of Ariant - An Old, Empty House (260000201)
type Jano struct {
}

func (r Jano) NPCId() uint32 {
	return npc.Jano
}

func (r Jano) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.ChangeMusic(l)(c.CharacterId, "Bgm14/Ariant")
	return script.WarpById(_map.Stage1MagikMirror, 3)(l, span, c)
}
