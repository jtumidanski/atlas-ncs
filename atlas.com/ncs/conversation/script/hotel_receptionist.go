package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// HotelReceptionist is located in Dungeon - Sleepywood Hotel (105040400)
type HotelReceptionist struct {
}

func (r HotelReceptionist) NPCId() uint32 {
	return npc.HotelReceptionist
}

func (r HotelReceptionist) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r HotelReceptionist) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Welcome. We're the Sleepywood Hotel. Our hotel works hard to serve you the best at all times. If you are tired and worn out from hunting, how about a relaxing stay at our hotel?")
	return SendNext(l, c, m.String(), r.Choose)
}

func (r HotelReceptionist) Choose(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("We offer two kinds of rooms for our service. Please choose the one of your liking.").NewLine().
		OpenItem(0).BlueText().AddText("Regular sauna (499)").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("VIP sauna (999 mesos per use)").CloseItem()
	return SendListSelection(l, c, m.String(), r.SelectService)
}

func (r HotelReceptionist) SelectService(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.RegularConfirmation
	case 1:
		return r.VIPConfirmation
	}
	return nil
}

func (r HotelReceptionist) RegularConfirmation(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You have chosen the regular sauna. Your HP and MP will recover fast and you can even purchase some items there. Are you sure you want to go in?")
	return SendYesNo(l, c, m.String(), r.Validate(499, _map.RegularSauna), r.ThinkCarefully)
}

func (r HotelReceptionist) ThinkCarefully(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("We offer other kinds of services, too, so please think carefully and then make your decision.")
	return SendNext(l, c, m.String(), Exit())
}

func (r HotelReceptionist) VIPConfirmation(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You've chosen the VIP sauna. Your HP and MP will recover even faster than that of the regular sauna and you can even find a special item in there. Are you sure you want to go in?")
	return SendYesNo(l, c, m.String(), r.Validate(999, _map.VIPSauna), r.ThinkCarefully)
}

func (r HotelReceptionist) Validate(cost uint32, mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, cost) {
			return r.NotEnoughMeso(cost)(l, c)
		}
		return r.ProcessPurchase(cost, mapId)(l, c)
	}
}

func (r HotelReceptionist) NotEnoughMeso(cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("I'm sorry. It looks like you don't have enough mesos. It will cost you at least ").
			AddText(fmt.Sprintf("%d", cost)).
			AddText("mesos to stay at our hotel.")
		return SendNext(l, c, m.String(), Exit())
	}
}

func (r HotelReceptionist) ProcessPurchase(cost uint32, mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -int32(cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process purchase for character %d.", c.CharacterId)
		}
		err = npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return nil
	}
}
