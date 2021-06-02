package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Rini is located in Orbis - Station<To Ellinia> (200000111)
type Rini struct {
}

func (r Rini) NPCId() uint32 {
	return npc.Rini
}

func (r Rini) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToElliniaRegular) {
		return r.NeedATicket(l, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.StationToEllinia, _map.ElliniaStation) {
		return r.AlreadyTraveling(l, c)
	}

	return r.Confirm(l, c)
}

func (r Rini) NeedATicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Make sure you got a Ellinia ticket to travel in this boat. Check your inventory.")
	return SendOk(l, c, m.String())
}

func (r Rini) AlreadyTraveling(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The boat to Ellinia is already travelling, please be patient for the next one.")
	return SendOk(l, c, m.String())
}

func (r Rini) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you want to go to Ellinia?")
	return SendYesNo(l, c, m.String(), r.Validate, r.ChangeYourMind)
}

func (r Rini) Validate(l logrus.FieldLogger, c Context) State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.StationToEllinia, _map.ElliniaStation) {
		return r.NextOne(l, c)
	}
	return r.Perform(l, c)
}

func (r Rini) Perform(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.TicketToElliniaRegular, -1)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.BeforeTakeoffToEllinia, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.BeforeTakeoffToEllinia, c.NPCId)
	}
	return nil
}

func (r Rini) ChangeYourMind(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return SendOk(l, c, m.String())
}

func (r Rini) NextOne(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The boat to Ellinia is ready to take off, please be patient for the next one.")
	return SendOk(l, c, m.String())
}
