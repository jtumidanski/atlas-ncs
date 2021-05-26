package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// RegularCabHenesys is located in Victoria Road - Henesys (100000000)
type RegularCabHenesys struct {
}

func (r RegularCabHenesys) NPCId() uint32 {
	return npc.RegularCabHenesys
}

func (r RegularCabHenesys) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r RegularCabHenesys) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hello, I drive the Regular Cab. If you want to go from town to town safely and fast, then ride our cab. We'll gladly take you to your destination with an affordable price.")
	return SendNextExit(l, c, m.String(), r.WhereToGo, r.MoreToSee)
}

func (r RegularCabHenesys) MoreToSee(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("There's a lot to see in this town, too. Come back and find us when you need to go to a different town.")
	return SendNext(l, c, m.String(), Exit())
}

func (r RegularCabHenesys) WhereToGo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder()
	beginner := character.IsBeginnerTree(l)(c.CharacterId)

	if beginner {
		m = m.AddText("We have a special 90% discount for beginners. ")
	}
	m = m.
		AddText("Choose your destination, for fees will change from place to place.").
		BlueText().NewLine().
		OpenItem(0).BlueText().ShowMap(_map.LithHarbor).CloseItem().NewLine().
		OpenItem(1).BlueText().ShowMap(_map.Perion).CloseItem().NewLine().
		OpenItem(2).BlueText().ShowMap(_map.Ellinia).CloseItem().NewLine().
		OpenItem(3).BlueText().ShowMap(_map.KerningCity).CloseItem().NewLine().
		OpenItem(4).BlueText().ShowMap(_map.Nautalis).CloseItem()
	return SendListSelectionExit(l, c, m.String(), r.SelectTownConfirm(beginner), r.MoreToSee)
}

func (r RegularCabHenesys) SelectTownConfirm(beginner bool) ProcessSelection {
	return func(selection int32) StateProducer {
		switch selection {
		case 0:
			return r.ConfirmLithHarbor(r.Cost(selection, beginner))
		case 1:
			return r.ConfirmPerion(r.Cost(selection, beginner))
		case 2:
			return r.ConfirmEllinia(r.Cost(selection, beginner))
		case 3:
			return r.ConfirmKerningCity(r.Cost(selection, beginner))
		case 4:
			return r.ConfirmNautalis(r.Cost(selection, beginner))
		}
		return nil
	}
}

func (r RegularCabHenesys) Cost(index int32, beginner bool) uint32 {
	costDivisor := 1
	if beginner {
		costDivisor = 10
	}

	cost := uint32(0)
	switch index {
	case 0:
		cost = 1000
		break
	case 1:
		cost = 1000
		break
	case 2:
		cost = 800
		break
	case 3:
		cost = 1000
		break
	case 4:
		cost = 800
		break
	}
	return cost / uint32(costDivisor)
}

func (r RegularCabHenesys) ConfirmPerion(cost uint32) StateProducer {
	return r.ConfirmMap(_map.Perion, cost)
}

func (r RegularCabHenesys) ConfirmEllinia(cost uint32) StateProducer {
	return r.ConfirmMap(_map.Ellinia, cost)
}

func (r RegularCabHenesys) ConfirmLithHarbor(cost uint32) StateProducer {
	return r.ConfirmMap(_map.LithHarbor, cost)
}

func (r RegularCabHenesys) ConfirmKerningCity(cost uint32) StateProducer {
	return r.ConfirmMap(_map.KerningCity, cost)
}

func (r RegularCabHenesys) ConfirmNautalis(cost uint32) StateProducer {
	return r.ConfirmMap(_map.Nautalis, cost)
}

func (r RegularCabHenesys) ConfirmMap(mapId uint32, cost uint32) StateProducer {
	m := message.NewBuilder().
		AddText("You don't have anything else to do here, huh? Do you really want to go to ").
		BlueText().ShowMap(mapId).
		BlackText().AddText("? It'll cost you ").
		BlueText().AddText(fmt.Sprintf("%d mesos", cost))
	return func(l logrus.FieldLogger, c Context) State {
		return SendYesNoExit(l, c, m.String(), r.PerformTransaction(mapId, cost), r.MoreToSee, r.MoreToSee)
	}
}

func (r RegularCabHenesys) PerformTransaction(mapId uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, cost) {
			m := message.NewBuilder().
				AddText("You don't have enough mesos. Sorry to say this, but without them, you won't be able to ride the cab.")
			return SendNextExit(l, c, m.String(), Exit(), Exit())
		}

		err := character.GainMeso(l)(c.CharacterId, -int32(cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to complete meso transaction with %d.", c.CharacterId)
			return nil
		}

		err = npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to map %d. Refunding mesos.", c.CharacterId)
			err = character.GainMeso(l)(c.CharacterId, int32(cost))
			if err != nil {
				l.WithError(err).Errorf("Error processing refund, %d has lost %d mesos.", c.CharacterId, cost)
			}
		}
		return nil
	}
}
