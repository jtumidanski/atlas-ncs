package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Sunny is located in Orbis - Station <Ludibrium> (200000121)
type Sunny struct {
}

func (r Sunny) NPCId() uint32 {
	return npc.Sunny
}

func (r Sunny) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToLudibriumRegular) {
		return r.MissingTicket(l, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.OrbisStationLudibrium, _map.LudibriumStationOrbis) {
		return r.AlreadyTraveling(l, c)
	}

	return r.Confirm(l, c)
}

func (r Sunny) MissingTicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Make sure you got a Ludibrium ticket to travel in this train. Check your inventory.")
	return SendOk(l, c, m.String())
}

func (r Sunny) AlreadyTraveling(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The train to Ludibrium is already travelling, please be patient for the next one.")
	return SendOk(l, c, m.String())
}

func (r Sunny) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you want to go to Ludibrium?")
	return SendYesNo(l, c, m.String(), r.Validate, r.ChangeYourMind)
}

func (r Sunny) ChangeYourMind(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return SendOk(l, c, m.String())
}

func (r Sunny) Validate(l logrus.FieldLogger, c Context) State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.OrbisStationLudibrium, _map.LudibriumStationOrbis) {
		return r.BePatient(l, c)
	}
	return r.Process(l, c)
}

func (r Sunny) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.TicketToLudibriumRegular, -1)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.BeforeTheDepartureLudibrium, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.BeforeTheDepartureLudibrium, c.NPCId)
	}
	return Exit()(l, c)
}

func (r Sunny) BePatient(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The train to Ludibrium is ready to take off, please be patient for the next one.")
	return SendOk(l, c, m.String())
}
