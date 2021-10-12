package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Slyn is located in Ariant - Before Takeoff <To Orbis> (260000110)
type Slyn struct {
}

func (r Slyn) NPCId() uint32 {
	return npc.Slyn
}

func (r Slyn) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to leave the genie?")
	return script.SendYesNo(l, span, c, m.String(), r.NextTime, script.Exit())
}

func (r Slyn) NextTime(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return script.SendNext(l, span, c, m.String(), r.Warp)
}

func (r Slyn) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.AriantStationPlatform, 0)(l, span, c)
}
