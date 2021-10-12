package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Papulatus is located in Hidden Street - Origin of the Clock Tower (922020300)
type Papulatus struct {
}

func (r Papulatus) NPCId() uint32 {
	return npc.Papulatus
}

func (r Papulatus) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't belong to this world... Return now.")
	return script.SendNext(l, span, c, m.String(), r.Warp)
}

func (r Papulatus) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.DeepInsideTheClocktower, 0)(l, span, c)
}
