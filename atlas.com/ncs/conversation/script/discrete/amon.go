package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/event"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Amon is located in 
type Amon struct {
}

func (r Amon) NPCId() uint32 {
	return npc.Amon
}

func (r Amon) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId != _map.ZakumsAltar {
		return r.LeaveNow(l, span, c)
	}

	if !event.Cleared(l)(c.CharacterId) {
		return r.LeaveNow(l, span, c)
	}

	return r.Congratulations(l, span, c)
}

func (r Amon) LeaveNow(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you leave now, you'll have to start over. Are you sure you want to leave?")
	return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.TheDoorToZakum, 0), script.Exit())
}

func (r Amon) Congratulations(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You guys finally overthrew Zakum, what a superb feat! Congratulations! Are you sure you want to leave now?")
	return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.TheDoorToZakum, 0), script.Exit())
}
