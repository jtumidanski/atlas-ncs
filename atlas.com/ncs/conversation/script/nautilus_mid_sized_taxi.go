package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// NautilusMidSizedTaxi is located in Victoria Road - Nautilus Harbor (120000000)
type NautilusMidSizedTaxi struct {
}

func (r NautilusMidSizedTaxi) NPCId() uint32 {
	return npc.NautilusMidSizedTaxi
}

func (r NautilusMidSizedTaxi) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r NautilusMidSizedTaxi) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hello, I drive the Nautilus' Mid-Sized Taxi. If you want to go from town to town safely and fast, then ride our cab. We'll gladly take you to your destination with an affordable price.")
	return SendNextExit(l, c, m.String(), r.ChooseDestination, r.MoreToSee)
}

func (r NautilusMidSizedTaxi) MoreToSee(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("There's a lot to see in this town, too. Come back and find us when you need to go to a different town.")
	return SendOk(l, c, m.String())
}

func (r NautilusMidSizedTaxi) ChooseDestination(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder()
	multiplier := 1.0
	if character.IsBeginnerTree(l)(c.CharacterId) {
		m = m.AddText("We have a special 90% discount for beginners. ")
		multiplier = 0.1
	}
	m = m.AddText("Choose your destination, for fees will change from place to place.").NewLine().
		OpenItem(0).BlueText().ShowMap(_map.LithHarbor).AddText(fmt.Sprintf(" (%d mesos)", uint32(1000*multiplier))).CloseItem().NewLine().
		OpenItem(1).BlueText().ShowMap(_map.Perion).AddText(fmt.Sprintf(" (%d mesos)", uint32(1000*multiplier))).CloseItem().NewLine().
		OpenItem(2).BlueText().ShowMap(_map.Henesys).AddText(fmt.Sprintf(" (%d mesos)", uint32(1000*multiplier))).CloseItem().NewLine().
		OpenItem(3).BlueText().ShowMap(_map.Ellinia).AddText(fmt.Sprintf(" (%d mesos)", uint32(800*multiplier))).CloseItem().NewLine().
		OpenItem(4).BlueText().ShowMap(_map.KerningCity).AddText(fmt.Sprintf(" (%d mesos)", uint32(1000*multiplier))).CloseItem()
	return SendListSelectionExit(l, c, m.String(), r.DestinationSelection(multiplier), r.MoreToSee)
}

func (r NautilusMidSizedTaxi) DestinationSelection(multiplier float64) func(selection int32) StateProducer {
	return func(selection int32) StateProducer {
		switch selection {
		case 0:
			return r.NothingMoreToDo(_map.LithHarbor, 1000, multiplier)
		case 1:
			return r.NothingMoreToDo(_map.Perion, 1000, multiplier)
		case 2:
			return r.NothingMoreToDo(_map.Henesys, 1000, multiplier)
		case 3:
			return r.NothingMoreToDo(_map.Ellinia, 800, multiplier)
		case 4:
			return r.NothingMoreToDo(_map.KerningCity, 1000, multiplier)
		}
		return nil
	}
}

func (r NautilusMidSizedTaxi) NothingMoreToDo(mapId uint32, cost uint32, multiplier float64) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		finalCost := uint32(float64(cost) * multiplier)
		m := message.NewBuilder().
			AddText("You don't have anything else to do here, huh? Do you really want to go to ").
			BlueText().ShowMap(mapId).
			BlackText().AddText("? It'll cost you ").
			BlueText().AddText(fmt.Sprintf("%d", finalCost)).AddText(" mesos").
			BlackText().AddText(".")
		return SendYesNoExit(l, c, m.String(), r.Validate(mapId, finalCost), r.MoreToSee, r.MoreToSee)
	}
}

func (r NautilusMidSizedTaxi) Validate(mapId uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, cost) {
			return r.NotEnoughMeso(l, c)
		}
		return r.Process(mapId, cost)(l, c)
	}
}

func (r NautilusMidSizedTaxi) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You don't have enough mesos. Sorry to say this, but without them, you won't be able to ride the cab.")
	return SendOk(l, c, m.String())
}

func (r NautilusMidSizedTaxi) Process(mapId uint32, cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -int32(cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process purchase for character %d.", c.CharacterId)
		}
		return WarpById(mapId, 0)(l, c)
	}
}
