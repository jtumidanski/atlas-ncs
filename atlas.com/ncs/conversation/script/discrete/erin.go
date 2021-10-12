package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Erin is located in Orbis - Before Takeoff <To Ellinia> (200000112)
type Erin struct {
}

func (r Erin) NPCId() uint32 {
	return npc.Erin
}

func (r Erin) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to leave the boat?")
	return script.SendYesNo(l, span, c, m.String(), r.Alright, r.GoodChoice)
}

func (r Erin) GoodChoice(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Good choice.")
	return script.SendOk(l, span, c, m.String())
}

func (r Erin) Alright(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return script.SendNext(l, span, c, m.String(), script.WarpById(_map.StationToEllinia, 0))
}
