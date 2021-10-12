package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Rosey is located in Orbis - Before the Departure <Ludibrium> (200000122)
type Rosey struct {
}

func (r Rosey) NPCId() uint32 {
	return npc.Rosey
}

func (r Rosey) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to leave the train?")
	return script.SendYesNo(l, span, c, m.String(), r.SeeYouNextTime, script.Exit())
}

func (r Rosey) SeeYouNextTime(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return script.SendNext(l, span, c, m.String(), r.Warp)
}

func (r Rosey) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	var mapId uint32
	if c.MapId == _map.BeforeTheDepartureLudibrium {
		mapId = _map.StationLudibrium
	} else {
		mapId = _map.StationOrbis
	}
	return script.WarpById(mapId, 0)(l, span, c)
}
