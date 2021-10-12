package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// RollyExit is located in Ludibrium - Exit of the Maze (809050017)
type RollyExit struct {
}

func (r RollyExit) NPCId() uint32 {
	return npc.RollyExit
}

func (r RollyExit) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Come this way to return to Ludibrium.")
	return script.SendNext(l, span, c, m.String(), r.Warp)
}

func (r RollyExit) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.Ludibrium, 0)(l, span, c)
}
