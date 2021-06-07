package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Crane is located in Orbis - Cabin <To Mu Lung> (200000141), Herb Town - Herb Town (251000000), and Mu Lung - Mu Lung Temple (250000100)
type Crane struct {
}

func (r Crane) NPCId() uint32 {
	return npc.Crane
}

func (r Crane) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.HerbTown {
		m := message.NewBuilder().
			AddText("Hello there. How's the traveling so far? I've been transporting other travelers like you to ").
			BlueText().AddText("Mu Lung").
			BlackText().AddText(" in no time, and... are you interested? It's not as stable as the ship, so you'll have to hold on tight, but i can get there much faster than the ship. I'll take you there as long as you pay ").
			BlueText().AddText(fmt.Sprintf("%d mesos", 1500)).
			BlackText().AddText(".")
		return SendYesNo(l, c, m.String(), r.Validate(_map.MuLungTemple, 1500), r.LetMeKnow)
	} else if c.MapId == _map.MuLungTemple {
		m := message.NewBuilder().
			AddText("Hello there. How's the traveling so far? I understand that walking on two legs is much harder to cover ground compared to someone like me that can navigate the skies. I've been transporting other travelers like you to other regions in no time, and... are you interested? If so, then select the town you'd like yo head to.").NewLine().
			OpenItem(0).BlueText().AddText(fmt.Sprintf("Orbis (%d mesos)", 1500)).CloseItem().NewLine().
			OpenItem(1).BlueText().AddText(fmt.Sprintf("Herb Town (%d mesos)", 500)).CloseItem()
		return SendListSelection(l, c, m.String(), r.MuLungOriginSelection)
	} else {
		m := message.NewBuilder().
			AddText("Hello there. How's the traveling so far? I've been transporting other travelers like you to other regions in no time, and... are you interested? If so, then select the town you'd like to head to.").NewLine().
			OpenItem(0).BlueText().AddText(fmt.Sprintf("Mu Lung (%d mesos)", 1500)).CloseItem()
		return SendListSelection(l, c, m.String(), r.OrbisOriginSelection)
	}
}

func (r Crane) LetMeKnow(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("OK. If you ever change your mind, please let me know.")
	return SendOk(l, c, m.String())
}

func (r Crane) Validate(mapId uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, cost) {
			return r.NotEnoughMesos(l, c)
		}
		return r.Process(mapId, cost)(l, c)
	}
}

func (r Crane) NotEnoughMesos(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Are you sure you have enough mesos?")
	return SendOk(l, c, m.String())
}

func (r Crane) Process(mapId uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -int32(cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
		}
		return WarpById(mapId, 0)(l, c)
	}
}

func (r Crane) MuLungOriginSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.ConfirmTransport(_map.MuLungTemple, _map.OrbisStationEntrance, _map.DuringTheRideToOrbis, 1500)
	case 1:
		return r.Confirm(_map.HerbTown, 500)
	}
	return nil
}

func (r Crane) OrbisOriginSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.ConfirmTransport(_map.CabinToMuLung, _map.MuLungTemple, _map.DuringTheRideToMuLung, 1500)
	}
	return nil
}

func (r Crane) ValidateTransport(from uint32, to uint32, warp uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.TransportBoarding(l)(c.CharacterId, from, to) {
			return r.TryAgainInABit(l, c)
		}
		if !character.HasMeso(l)(c.CharacterId, cost) {
			return r.NotEnoughMesos(l, c)
		}
		return r.Process(warp, cost)(l, c)
	}
}

func (r Crane) Confirm(mapId uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("Will you move to ").
			BlueText().ShowMap(mapId).
			BlackText().AddText(" now? If you have ").
			BlueText().AddText(fmt.Sprintf("%d mesos", cost)).
			BlackText().AddText(", I'll take you there right now.")
		return SendYesNo(l, c, m.String(), r.Validate(mapId, cost), r.LetMeKnow)
	}
}

func (r Crane) ConfirmTransport(from uint32, to uint32, warp uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.TransportBoarding(l)(c.CharacterId, from, to) {
			return r.TryAgainInABit(l, c)
		}
		m := message.NewBuilder().
			AddText("Will you move to ").
			BlueText().ShowMap(to).
			BlackText().AddText(" now? If you have ").
			BlueText().AddText(fmt.Sprintf("%d mesos", cost)).
			BlackText().AddText(", I'll take you there right now.")
		return SendYesNo(l, c, m.String(), r.ValidateTransport(from, to, warp, cost), r.LetMeKnow)
	}
}

func (r Crane) TryAgainInABit(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Uh... We are currently taking requests from too many maplers right now... Please try again in a bit.")
	return SendOk(l, c, m.String())
}
