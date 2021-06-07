package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Bell is located in 
type Bell struct {
}

func (r Bell) NPCId() uint32 {
	return npc.Bell
}

func (r Bell) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.SubwayTicketingBooth {
		m := message.NewBuilder().
			AddText("The ride to New Leaf City of Masteria takes off every minute, beginning on the hour, and it'll cost you ").
			BlueText().AddText(fmt.Sprintf("%d mesos", 5000)).
			BlackText().AddText(". Are you sure you want to purchase ").
			BlueText().ShowItemName1(item.SubwayTicketToNLCRegular).
			BlackText().AddText("?")
		return SendYesNo(l, c, m.String(), r.ValidatePurchase(item.SubwayTicketToNLCRegular), Exit())
	} else if c.MapId == _map.NLCSubwayStation {
		m := message.NewBuilder().
			AddText("The ride to Kerning City of Victoria Island takes off every minute, beginning on the hour, and it'll cost you ").
			BlueText().AddText(fmt.Sprintf("%d mesos", 5000)).
			BlackText().AddText(". Are you sure you want to purchase ").
			BlueText().ShowItemName1(item.SubwayTicketToKerningCityRegular).
			BlackText().AddText("?")
		return SendYesNo(l, c, m.String(), r.ValidatePurchase(item.SubwayTicketToKerningCityRegular), Exit())
	} else if c.MapId == _map.WaitingRoomFromNewLeafCityToKerningCity {
		m := message.NewBuilder().
			AddText("Do you want to leave before the train start? There will be no refund.")
		return SendYesNo(l, c, m.String(), Warp(_map.NLCSubwayStation), Exit())
	} else if c.MapId == _map.WaitingRoomFromKerningCityToNewLeafCity {
		m := message.NewBuilder().
			AddText("Do you want to leave before the train start? There will be no refund.")
		return SendYesNo(l, c, m.String(), Warp(_map.SubwayTicketingBooth), Exit())
	}
	return Exit()(l, c)
}

func (r Bell) ValidatePurchase(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, 5000) {
			return r.NotEnoughMeso(l, c)
		}
		if !character.CanHold(l)(c.CharacterId, itemId) {
			return r.NeedInventoryRoom(l, c)
		}
		return r.ProcessPurchase(itemId)(l, c)
	}
}

func (r Bell) ProcessPurchase(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -5000)
		if err != nil {
			l.WithError(err).Errorf("Unable to perform purchase for character %d.", c.CharacterId)
		}
		character.GainItem(l)(c.CharacterId, itemId, 1)
		return r.ThereYouGo(l, c)
	}
}

func (r Bell) ThereYouGo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("There you go.")
	return SendOk(l, c, m.String())
}

func (r Bell) NeedInventoryRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You don't have a etc. slot available.")
	return SendOk(l, c, m.String())
}

func (r Bell) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You don't have enough mesos.")
	return SendOk(l, c, m.String())
}
