package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Irene is located in 
type Irene struct {
}

func (r Irene) NPCId() uint32 {
	return npc.Irene
}

func (r Irene) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello, I am Irene from Singapore Airport. I can assist you in getting you to Singapore in no time. Do you want to go to Singapore?").NewLine().
		OpenItem(0).BlueText().AddText("I would like to buy a plane ticket to Singapore").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Let me go in to the departure point.").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r Irene) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Buy
	case 1:
		return r.Warp
	}
	return nil
}

func (r Irene) Buy(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The ticket will cost you 5,000 mesos. Will you purchase the ticket?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, script.Exit())
}

func (r Irene) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Would you like to go in now? You will lose your ticket once you go in! Thank you for choosing Wizet Airline.")
	return script.SendYesNo(l, span, c, m.String(), r.ValidateWarp, script.Exit())
}

func (r Irene) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 5000) || character.HasItem(l, span)(c.CharacterId, item.TicketToSingaporeFromKerningCity) {
		m := message.NewBuilder().AddText("You do not have enough mesos or you've already purchased a ticket.")
		return script.SendOk(l, span, c, m.String())
	}

	if !character.CanHold(l)(c.CharacterId, item.TicketToSingaporeFromKerningCity) {
		m := message.NewBuilder().AddText("You don't have a free slot on your ETC inventory for the ticket, please make a room beforehand.")
		return script.SendOk(l, span, c, m.String())
	}

	return r.ProcessPurchase(l, span, c)
}

func (r Irene) ProcessPurchase(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -5000)
	character.GainItem(l, span)(c.CharacterId, item.TicketToSingaporeFromKerningCity, 1)
	return r.ThankYou(l, span, c)
}

func (r Irene) ThankYou(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Thank you for choosing Wizet Airline! Enjoy your flight!")
	return script.SendOk(l, span, c, m.String())
}

func (r Irene) ValidateWarp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.TicketToSingaporeFromKerningCity) {
		return r.MissingTicket(l, span, c)
	}
	if !character.TransportBoarding(l)(c.CharacterId, _map.KerningAirport, _map.ChangiAirport) {
		m := message.NewBuilder().AddText("Sorry the plane has taken off, please wait a few minutes.")
		return script.SendOk(l, span, c, m.String())
	}
	character.GainItem(l, span)(c.CharacterId, item.TicketToSingaporeFromKerningCity, -1)
	return script.Warp(_map.KerningAirport)(l, span, c)
}

func (r Irene) MissingTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You need a ").
		BlueText().ShowItemName1(item.TicketToSingaporeFromKerningCity).
		BlackText().AddText(" to get on the plane!")
	return script.SendOk(l, span, c, m.String())
}