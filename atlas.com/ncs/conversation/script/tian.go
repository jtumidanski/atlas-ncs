package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Tian is located in Ludibrium - Station<Orbis> (220000110)
type Tian struct {
}

func (r Tian) NPCId() uint32 {
	return npc.Tian
}

func (r Tian) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.NeedTicket(l, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.StationOrbis, _map.OrbisStationEntrance) {
		return r.AlreadyTraveling(l, c)
	}

	return r.Confirm(l, c)
}

func (r Tian) NeedTicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Make sure you got a Orbis ticket to travel in this train. Check your inventory.")
	return SendOk(l, c, m.String())
}

func (r Tian) AlreadyTraveling(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The train to Orbis is already travelling, please be patient for the next one.")
	return SendOk(l, c, m.String())
}

func (r Tian) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you want to go to Orbis?")
	return SendYesNo(l, c, m.String(), r.Validate, r.TalkToMe)
}

func (r Tian) TalkToMe(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return SendOk(l, c, m.String())
}

func (r Tian) Validate(l logrus.FieldLogger, c Context) State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.StationOrbis, _map.OrbisStationEntrance) {
		return r.BePatient(l, c)
	}
	return r.Warp(l, c)
}

func (r Tian) BePatient(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The train to Orbis is ready to take off, please be patient for the next one.")
	return SendOk(l, c, m.String())
}

func (r Tian) Warp(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.TicketToOrbisRegular, -1)
	return WarpById(_map.BeforeTheDepartureOrbis, 0)(l, c)
}
