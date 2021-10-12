package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// DangerZoneTaxi is located in El Nath - El Nath (211000000), Ludibrium - Ludibrium (220000000), Omega Sector - Omega Sector (221000000), Leafre - Leafre (240000000)
type DangerZoneTaxi struct {
}

func (r DangerZoneTaxi) NPCId() uint32 {
	return npc.DangerZoneTaxi
}

func (r DangerZoneTaxi) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	destinationId := r.DestinationMap(c.MapId)
	destinationCost := r.DestinationCost(c.MapId)

	m := message.NewBuilder().
		AddText("Hello there! This taxi will take you to dangerous places in Ossyria faster than an arrow! We go from ").
		ShowMap(c.MapId).AddText(" to ").
		BlueText().ShowMap(destinationId).
		BlackText().AddText(" on this Ossyria Continent! It'll cost you ").
		BlueText().AddText(fmt.Sprintf("%d mesos", destinationCost)).
		BlackText().AddText(". I know it's a bit expensive, but it's well worth passing all the dangerous areas!")
	return script.SendNext(l, span, c, m.String(), r.Confirm(destinationId, destinationCost))
}

func (r DangerZoneTaxi) DestinationMap(mapId uint32) uint32 {
	switch mapId {
	case _map.ElNath:
		return _map.IceValleyII
	case _map.Ludibrium:
		return _map.PathOfTime
	case _map.OmegaSector:
		//TODO this doesn't seem right
		return _map.Ludibrium
	case _map.Leafre:
		return _map.EntranceToDragonForest
	}
	return 0
}

func (r DangerZoneTaxi) DestinationCost(mapId uint32) uint32 {
	switch mapId {
	case _map.ElNath:
		return 10000
	case _map.Ludibrium:
		return 25000
	case _map.OmegaSector:
		return 25000
	case _map.Leafre:
		return 65000
	}
	return 0
}

func (r DangerZoneTaxi) Confirm(mapId uint32, cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Would you like to pay ").
			BlueText().AddText(fmt.Sprintf("%d mesos", cost)).
			BlackText().AddText(" to travel to the ").
			BlueText().ShowMap(mapId).
			BlackText().AddText("?")
		return script.SendYesNo(l, span, c, m.String(), r.Validate(mapId, cost), r.NotCheap)
	}
}

func (r DangerZoneTaxi) Validate(mapId uint32, cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, cost) {
			return r.NotEnoughMesos(l, span, c)
		}
		return r.Process(mapId, cost)(l, span, c)
	}
}

func (r DangerZoneTaxi) Process(mapId uint32, cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainMeso(l, span)(c.CharacterId, -int32(cost))
		return script.WarpById(mapId, 0)(l, span, c)
	}
}

func (r DangerZoneTaxi) NotEnoughMesos(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't seem to have enough mesos. I am terribly sorry, but I cannot help you unless you pay up. Bring in the mesos by hunting more and come back when you have enough.")
	return script.SendOk(l, span, c, m.String())
}

func (r DangerZoneTaxi) NotCheap(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hmm, please think this over. It's not cheap, but you will NOT be disappointed with our premier service!")
	return script.SendOk(l, span, c, m.String())
}
