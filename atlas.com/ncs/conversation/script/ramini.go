package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Ramini is located in  Orbis - Cabin <To Leafre> (200000131)
type Ramini struct {
}

func (r Ramini) NPCId() uint32 {
	return npc.Ramini
}

func (r Ramini) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToLeafreRegular) {
		return r.NeedTicket(l, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.CabinToLeafre, _map.LeafreStationEntrance) {
		return r.NotArrived(l, c)
	}

	return r.Confirm(l, c)
}

func (r Ramini) NeedTicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Make sure you got a Leafre ticket to travel in this flight. Check your inventory.")
	return SendOk(l, c, m.String())
}

func (r Ramini) NotArrived(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The flight has not arrived yet. Come back soon.")
	return SendOk(l, c, m.String())
}

func (r Ramini) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to board the flight?")
	return SendYesNo(l, c, m.String(), r.Validate, r.ChangeYourMind)
}

func (r Ramini) Validate(l logrus.FieldLogger, c Context) State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.CabinToLeafre, _map.LeafreStationEntrance) {
		return r.NotArrived(l, c)
	}
	return r.Process(l, c)
}

func (r Ramini) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.TicketToLeafreRegular, -1)
	return WarpById(_map.CabinToLeafre2, 0)(l, c)
}

func (r Ramini) ChangeYourMind(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return SendOk(l, c, m.String())
}
