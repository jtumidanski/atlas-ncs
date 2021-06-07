package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// DangerZoneTaxi is located in El Nath - El Nath (211000000), Ludibrium - Ludibrium (220000000), Omega Sector - Omega Sector (221000000), Leafre - Leafre (240000000)
type DangerZoneTaxi struct {
}

func (r DangerZoneTaxi) NPCId() uint32 {
	return npc.DangerZoneTaxi
}

func (r DangerZoneTaxi) Initial(l logrus.FieldLogger, c Context) State {
	destinationId := r.DestinationMap(c.MapId)
	destinationCost := r.DestinationCost(c.MapId)

	m := message.NewBuilder().
		AddText("Hello there! This taxi will take you to dangerous places in Ossyria faster than an arrow! We go from ").
		ShowMap(c.MapId).AddText(" to ").
		BlueText().ShowMap(destinationId).
		BlackText().AddText(" on this Ossyria Continent! It'll cost you ").
		BlueText().AddText(fmt.Sprintf("%d mesos", destinationCost)).
		BlackText().AddText(". I know it's a bit expensive, but it's well worth passing all the dangerous areas!")
	return SendNext(l, c, m.String(), r.Confirm(destinationId, destinationCost))
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

func (r DangerZoneTaxi) Confirm(mapId uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("Would you like to pay ").
			BlueText().AddText(fmt.Sprintf("%d mesos", cost)).
			BlackText().AddText(" to travel to the ").
			BlueText().ShowMap(mapId).
			BlackText().AddText("?")
		return SendYesNo(l, c, m.String(), r.Validate(mapId, cost), r.NotCheap)
	}
}

func (r DangerZoneTaxi) Validate(mapId uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, cost) {
			return r.NotEnoughMesos(l, c)
		}
		return r.Process(mapId, cost)(l, c)
	}
}

func (r DangerZoneTaxi) Process(mapId uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -int32(cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
		}
		return WarpById(mapId, 0)(l, c)
	}
}

func (r DangerZoneTaxi) NotEnoughMesos(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You don't seem to have enough mesos. I am terribly sorry, but I cannot help you unless you pay up. Bring in the mesos by hunting more and come back when you have enough.")
	return SendOk(l, c, m.String())
}

func (r DangerZoneTaxi) NotCheap(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hmm, please think this over. It's not cheap, but you will NOT be disappointed with our premier service!")
	return SendOk(l, c, m.String())
}
