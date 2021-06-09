package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Dolphin is located in Aquarium - Aquarium (230000000)
type Dolphin struct {
}

func (r Dolphin) NPCId() uint32 {
	return npc.Dolphin
}

func (r Dolphin) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.HasItem(l)(c.CharacterId, item.DolphinTaxiCoupon) {
		if c.MapId == _map.PierOnTheBeach {
			return r.WithTicketInHerbTown(l, c)
		} else {
			return r.WithTicketInAquarium(l, c)
		}
	} else {
		if c.MapId == _map.PierOnTheBeach {
			return r.WithoutTicketInHerbTown(l, c)
		} else {
			return r.WithoutTicketInAquarium(l, c)
		}
	}
}

func (r Dolphin) WithTicketInHerbTown(l logrus.FieldLogger, c script.Context) script.State {
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
	return script.SendListSelection(l, c, m.String(), r.Selection(true, _map.Aquarium))
}

func (r Dolphin) WithTicketInAquarium(l logrus.FieldLogger, c script.Context) script.State {
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
	return script.SendListSelection(l, c, m.String(), r.Selection(true, _map.HerbTown))
}

func (r Dolphin) WithoutTicketInHerbTown(l logrus.FieldLogger, c script.Context) script.State {
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
	return script.SendListSelection(l, c, m.String(), r.Selection(false, _map.Aquarium))
}

func (r Dolphin) WithoutTicketInAquarium(l logrus.FieldLogger, c script.Context) script.State {
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
	return script.SendListSelection(l, c, m.String(), r.Selection(false, _map.HerbTown))
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
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if !ticket && !character.HasMeso(l)(c.CharacterId, 1000) {
			return r.NotEnoughMeso(l, c)
		}

		return r.ProcessSharpUnknown(ticket)(l, c)
	}
}

func (r Dolphin) ProcessSharpUnknown(ticket bool) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if ticket {
			character.GainItem(l)(c.CharacterId, item.DolphinTaxiCoupon, -1)
		} else {
			err := character.GainMeso(l)(c.CharacterId, -1000)
			if err != nil {
				l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
			}
		}
		return script.WarpById(_map.TheSharpUnknown, 2)(l, c)
	}
}

func (r Dolphin) NotEnoughMeso(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I don't think you have enough money...")
	return script.SendOk(l, c, m.String())
}

func (r Dolphin) ValidateTown(destination uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if !character.HasMeso(l)(c.CharacterId, 10000) {
			return r.NotEnoughMeso(l, c)
		}
		return r.ProcessTown(destination)(l, c)
	}
}

func (r Dolphin) ProcessTown(destination uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		err := character.GainMeso(l)(c.CharacterId, -10000)
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
		}
		return script.WarpById(destination, 0)(l, c)
	}
}
