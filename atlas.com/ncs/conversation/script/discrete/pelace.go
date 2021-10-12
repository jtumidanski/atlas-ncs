package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Pelace is located in  Orbis - Cabin <To Leafre> (200000132)
type Pelace struct {
}

func (r Pelace) NPCId() uint32 {
	return npc.Pelace
}

func (r Pelace) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to leave the flight?")
	return script.SendYesNo(l, span, c, m.String(), r.Alright, script.Exit())
}

func (r Pelace) Alright(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return script.SendNext(l, span, c, m.String(), r.Warp)
}

func (r Pelace) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.CabinToLeafre, 0)(l, span, c)
}
