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

// Dolphin is located in Aquarium - Aquarium (230000000)
type Dolphin struct {
}

func (r Dolphin) NPCId() uint32 {
	return npc.Dolphin
}

func (r Dolphin) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.DolphinTaxiCoupon) {
		if c.MapId == _map.PierOnTheBeach {
			return r.WithTicketInHerbTown(l, span, c)
		} else {
			return r.WithTicketInAquarium(l, span, c)
		}
	} else {
		if c.MapId == _map.PierOnTheBeach {
			return r.WithoutTicketInHerbTown(l, span, c)
		} else {
			return r.WithoutTicketInAquarium(l, span, c)
		}
	}
}

func (r Dolphin) WithTicketInHerbTown(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ocean are all connected to each other. Place you can't reach by foot can easily reached oversea. How about taking ").
		BlueText().AddText("Dolphin Taxi").
		BlackText().AddText(" with us today?").NewLine().
		OpenItem(0).
		BlueText().AddText("I will use ").
		ShowItemName1(item.DolphinTaxiCoupon).
		BlackText().AddText(" to move to ").
		BlueText().ShowMap(_map.TheSharpUnknown).
		BlackText().AddText(".").CloseItem().NewLine().
		OpenItem(1).
		BlueText().AddText("Go to ").
		BlueText().ShowMap(_map.Aquarium).
		BlackText().AddText(" after paying ").
		BlueText().AddText("10000 mesos").
		BlackText().AddText(".").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection(true, _map.Aquarium))
}

func (r Dolphin) WithTicketInAquarium(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ocean are all connected to each other. Place you can't reach by foot can easily reached oversea. How about taking ").
		BlueText().AddText("Dolphin Taxi").
		BlackText().AddText(" with us today?").NewLine().
		OpenItem(0).
		BlueText().AddText("I will use ").
		ShowItemName1(item.DolphinTaxiCoupon).
		BlackText().AddText(" to move to ").
		BlueText().ShowMap(_map.TheSharpUnknown).
		BlackText().AddText(".").CloseItem().NewLine().
		OpenItem(1).
		BlueText().AddText("Go to ").
		BlueText().ShowMap(_map.HerbTown).
		BlackText().AddText(" after paying ").
		BlueText().AddText("10000 mesos").
		BlackText().AddText(".").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection(true, _map.HerbTown))
}

func (r Dolphin) WithoutTicketInHerbTown(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ocean are all connected to each other. Place you can't reach by foot can easily reached oversea. How about taking ").
		BlueText().AddText("Dolphin Taxi").
		BlackText().AddText(" with us today?").NewLine().
		OpenItem(0).
		BlueText().AddText("Go to ").
		BlueText().ShowMap(_map.TheSharpUnknown).
		BlackText().AddText(" after paying ").
		BlueText().AddText("10000 mesos").
		BlackText().AddText(".").CloseItem().NewLine().
		OpenItem(1).
		BlueText().AddText("Go to ").
		BlueText().ShowMap(_map.Aquarium).
		BlackText().AddText(" after paying ").
		BlueText().AddText("10000 mesos").
		BlackText().AddText(".").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection(false, _map.Aquarium))
}

func (r Dolphin) WithoutTicketInAquarium(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ocean are all connected to each other. Place you can't reach by foot can easily reached oversea. How about taking ").
		BlueText().AddText("Dolphin Taxi").
		BlackText().AddText(" with us today?").NewLine().
		OpenItem(0).
		BlueText().AddText("Go to ").
		BlueText().ShowMap(_map.TheSharpUnknown).
		BlackText().AddText(" after paying ").
		BlueText().AddText("10000 mesos").
		BlackText().AddText(".").CloseItem().NewLine().
		OpenItem(1).
		BlueText().AddText("Go to ").
		BlueText().ShowMap(_map.HerbTown).
		BlackText().AddText(" after paying ").
		BlueText().AddText("10000 mesos").
		BlackText().AddText(".").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection(false, _map.HerbTown))
}

func (r Dolphin) Selection(hasTicket bool, secondDestination uint32) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		switch selection {
		case 0:
			return r.ValidateSharpUnknown(hasTicket)
		case 1:
			return r.ValidateTown(secondDestination)
		}
		return nil
	}
}

func (r Dolphin) ValidateSharpUnknown(ticket bool) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !ticket && !character.HasMeso(l, span)(c.CharacterId, 1000) {
			return r.NotEnoughMeso(l, span, c)
		}

		return r.ProcessSharpUnknown(ticket)(l, span, c)
	}
}

func (r Dolphin) ProcessSharpUnknown(ticket bool) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if ticket {
			character.GainItem(l, span)(c.CharacterId, item.DolphinTaxiCoupon, -1)
		} else {
			character.GainMeso(l, span)(c.CharacterId, -1000)
		}
		return script.WarpById(_map.TheSharpUnknown, 2)(l, span, c)
	}
}

func (r Dolphin) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I don't think you have enough money...")
	return script.SendOk(l, span, c, m.String())
}

func (r Dolphin) ValidateTown(destination uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, 10000) {
			return r.NotEnoughMeso(l, span, c)
		}
		return r.ProcessTown(destination)(l, span, c)
	}
}

func (r Dolphin) ProcessTown(destination uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainMeso(l, span)(c.CharacterId, -10000)
		return script.WarpById(destination, 0)(l, span, c)
	}
}
