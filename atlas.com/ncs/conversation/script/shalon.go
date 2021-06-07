package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Shalon is located in Singapore - Changi Airport (540010000)
type Shalon struct {
}

func (r Shalon) NPCId() uint32 {
	return npc.Shalon
}

func (r Shalon) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hello, I am Shalon from Singapore Airport. I can assist you in getting you to Kerning City in no time. Do you want to go to Kerning City?").NewLine().
		OpenItem(0).BlueText().AddText("I would like to buy a plane ticket to Kerning City").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Let me go in to the departure point.").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r Shalon) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Buy
	case 1:
		return r.Warp
	}
	return nil
}

func (r Shalon) Buy(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The ticket will cost you 5,000 mesos. Will you purchase the ticket?")
	return SendYesNo(l, c, m.String(), r.Validate, Exit())
}

func (r Shalon) Warp(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to go in now? You will lose your ticket once you go in! Thank you for choosing Wizet Airline.")
	return SendYesNo(l, c, m.String(), r.ValidateWarp, Exit())
}

func (r Shalon) Validate(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 5000) || character.HasItem(l)(c.CharacterId, item.TicketToKerningCityFromCBD) {
		m := message.NewBuilder().AddText("You do not have enough mesos or you've already purchased a ticket.")
		return SendOk(l, c, m.String())
	}

	if !character.CanHold(l)(c.CharacterId, item.TicketToKerningCityFromCBD) {
		m := message.NewBuilder().AddText("You don't have a free slot on your ETC inventory for the ticket, please make a room beforehand.")
		return SendOk(l, c, m.String())
	}

	return r.ProcessPurchase(l, c)
}

func (r Shalon) ProcessPurchase(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -5000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process purchase for character %d.", c.CharacterId)
	}
	character.GainItem(l)(c.CharacterId, item.TicketToKerningCityFromCBD, 1)
	return r.ThankYou(l, c)
}

func (r Shalon) ThankYou(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Thank you for choosing Wizet Airline! Enjoy your flight!")
	return SendOk(l, c, m.String())
}

func (r Shalon) ValidateWarp(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToKerningCityFromCBD) {
		return r.MissingTicket(l, c)
	}
	if !character.TransportBoarding(l)(c.CharacterId, _map.BeforeDepartureToKerningCity, _map.KerningCity) {
		m := message.NewBuilder().AddText("Sorry the plane has taken off, please wait a few minutes.")
		return SendOk(l, c, m.String())
	}
	character.GainItem(l)(c.CharacterId, item.TicketToKerningCityFromCBD, -1)
	err := npc.Warp(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.BeforeDepartureToKerningCity)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.BeforeDepartureToKerningCity, c.NPCId)
	}
	return Exit()(l, c)
}

func (r Shalon) MissingTicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You need a ").
		BlueText().ShowItemName1(item.TicketToKerningCityFromCBD).
		BlackText().AddText(" to get on the plane!")
	return SendOk(l, c, m.String())
}
