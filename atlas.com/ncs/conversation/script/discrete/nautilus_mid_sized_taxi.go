package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// NautilusMidSizedTaxi is located in Victoria Road - Nautilus Harbor (120000000)
type NautilusMidSizedTaxi struct {
}

func (r NautilusMidSizedTaxi) NPCId() uint32 {
	return npc.NautilusMidSizedTaxi
}

func (r NautilusMidSizedTaxi) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r NautilusMidSizedTaxi) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello, I drive the Nautilus' Mid-Sized Taxi. If you want to go from town to town safely and fast, then ride our cab. We'll gladly take you to your destination with an affordable price.")
	return script.SendNextExit(l, span, c, m.String(), r.ChooseDestination, r.MoreToSee)
}

func (r NautilusMidSizedTaxi) MoreToSee(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("There's a lot to see in this town, too. Come back and find us when you need to go to a different town.")
	return script.SendOk(l, span, c, m.String())
}

func (r NautilusMidSizedTaxi) ChooseDestination(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder()
	multiplier := 1.0
	if character.IsBeginnerTree(l, span)(c.CharacterId) {
		m = m.AddText("We have a special 90% discount for beginners. ")
		multiplier = 0.1
	}
	m = m.AddText("Choose your destination, for fees will change from place to place.").NewLine().
		OpenItem(0).BlueText().ShowMap(_map.LithHarbor).AddText(fmt.Sprintf(" (%d mesos)", uint32(1000*multiplier))).CloseItem().NewLine().
		OpenItem(1).BlueText().ShowMap(_map.Perion).AddText(fmt.Sprintf(" (%d mesos)", uint32(1000*multiplier))).CloseItem().NewLine().
		OpenItem(2).BlueText().ShowMap(_map.Henesys).AddText(fmt.Sprintf(" (%d mesos)", uint32(1000*multiplier))).CloseItem().NewLine().
		OpenItem(3).BlueText().ShowMap(_map.Ellinia).AddText(fmt.Sprintf(" (%d mesos)", uint32(800*multiplier))).CloseItem().NewLine().
		OpenItem(4).BlueText().ShowMap(_map.KerningCity).AddText(fmt.Sprintf(" (%d mesos)", uint32(1000*multiplier))).CloseItem()
	return script.SendListSelectionExit(l, span, c, m.String(), r.DestinationSelection(multiplier), r.MoreToSee)
}

func (r NautilusMidSizedTaxi) DestinationSelection(multiplier float64) func(selection int32) script.StateProducer {
	return func(selection int32) script.StateProducer {
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

func (r NautilusMidSizedTaxi) NothingMoreToDo(mapId uint32, cost uint32, multiplier float64) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		finalCost := uint32(float64(cost) * multiplier)
		m := message.NewBuilder().
			AddText("You don't have anything else to do here, huh? Do you really want to go to ").
			BlueText().ShowMap(mapId).
			BlackText().AddText("? It'll cost you ").
			BlueText().AddText(fmt.Sprintf("%d", finalCost)).AddText(" mesos").
			BlackText().AddText(".")
		return script.SendYesNoExit(l, span, c, m.String(), r.Validate(mapId, finalCost), r.MoreToSee, r.MoreToSee)
	}
}

func (r NautilusMidSizedTaxi) Validate(mapId uint32, cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, cost) {
			return r.NotEnoughMeso(l, span, c)
		}
		return r.Process(mapId, cost)(l, span, c)
	}
}

func (r NautilusMidSizedTaxi) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You don't have enough mesos. Sorry to say this, but without them, you won't be able to ride the cab.")
	return script.SendOk(l, span, c, m.String())
}

func (r NautilusMidSizedTaxi) Process(mapId uint32, cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainMeso(l, span)(c.CharacterId, -int32(cost))
		return script.WarpById(mapId, 0)(l, span, c)
	}
}
