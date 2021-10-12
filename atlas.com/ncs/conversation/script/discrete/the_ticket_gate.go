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

// TheTicketGate is located in Victoria Road - Subway Ticketing Booth (103000100)
type TheTicketGate struct {
}

func (r TheTicketGate) NPCId() uint32 {
	return npc.TheTicketGate
}

func (r TheTicketGate) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r TheTicketGate) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Pick your destination.").NewLine().
		OpenItem(0).BlueText().AddText("Kerning Square Shopping Center").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Enter Construction Site").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("New Leaf City").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.PickDestination)
}

func (r TheTicketGate) PickDestination(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.KerningSquare
	case 1:
		return r.ConstructionSite
	case 2:
		return r.NewLeafCity
	}
	return nil
}

func (r TheTicketGate) KerningSquare(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.SubwayTicketingBooth, _map.KerningSquareLobby) {
		return r.AlreadyFull(l, span, c)
	}
	//TODO next steps
	return script.Exit()(l, span, c)
}

func (r TheTicketGate) ConstructionSite(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasAnyItem(l, span)(c.CharacterId, item.TicketToConstructionSiteB1, item.TicketToConstructionSiteB2, item.TicketToConstructionSiteB3) {
		return r.NeedATicket(l, span, c)
	}

	m := message.NewBuilder().
		AddText("Here's the ticket reader. You will be brought in immediately. Which ticket you would like to use?").NewLine()
	for i, itemId := range []uint32{item.TicketToConstructionSiteB1, item.TicketToConstructionSiteB2, item.TicketToConstructionSiteB3} {
		if character.HasItem(l, span)(c.CharacterId, itemId) {
			m = m.OpenItem(i).BlueText().ShowItemName1(itemId).CloseItem().NewLine()
		}
	}
	return script.SendListSelection(l, span, c, m.String(), r.PickConstructionTicket)
}

func (r TheTicketGate) NewLeafCity(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.SubwayTicketingBooth && !character.HasItem(l, span)(c.CharacterId, item.SubwayTicketToNLCRegular) {
		return r.BuyATicketFromBell(l, span, c)
	}
	if !character.TransportBoarding(l)(c.CharacterId, _map.SubwayTicketingBooth, _map.NLCSubwayStation) {
		return r.BePatient(l, span, c)
	}
	return r.DoYouWantToGetOn(l, span, c)
}

func (r TheTicketGate) AlreadyFull(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The passenger wagon is already full. Try again a bit later.")
	return script.SendOk(l, span, c, m.String())
}

func (r TheTicketGate) NeedATicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It seems as though you don't have a ticket!")
	return script.SendOk(l, span, c, m.String())
}

func (r TheTicketGate) PickConstructionTicket(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.UseConstructionTicket(item.TicketToConstructionSiteB1, _map.Line3ConstructionSiteB1Area1)
	case 1:
		return r.UseConstructionTicket(item.TicketToConstructionSiteB2, _map.Line3ConstructionSiteB2Area1)
	case 2:
		return r.UseConstructionTicket(item.TicketToConstructionSiteB3, _map.Line3ConstructionSiteB3Area1)
	}
	return nil
}

func (r TheTicketGate) UseConstructionTicket(itemId uint32, mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainItem(l, span)(c.CharacterId, itemId, -1)
		return script.WarpById(mapId, 0)(l, span, c)
	}
}

func (r TheTicketGate) BuyATicketFromBell(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It seems you don't have a ticket! You can buy one from Bell.")
	return script.SendOk(l, span, c, m.String())
}

func (r TheTicketGate) BePatient(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("We will begin boarding 1 minute before the takeoff. Please be patient and wait for a few minutes. Be aware that the subway will take off right on time, and we stop receiving tickets 1 minute before that, so please make sure to be here on time.")
	return script.SendOk(l, span, c, m.String())
}

func (r TheTicketGate) DoYouWantToGetOn(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It looks like there's plenty of room for this ride. Please have your ticket ready so I can let you in. The ride will be long, but you'll get to your destination just fine. What do you think? Do you want to get on this ride?")
	return script.SendYesNo(l, span, c, m.String(), r.ProcessBoarding, script.Exit())
}

func (r TheTicketGate) ProcessBoarding(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.SubwayTicketingBooth, _map.NLCSubwayStation) {
		return r.BePatient(l, span, c)
	}
	character.GainItem(l, span)(c.CharacterId, item.SubwayTicketToNLCRegular, -1)
	return script.WarpById(_map.WaitingRoomFromKerningCityToNewLeafCity, 0)(l, span, c)
}
