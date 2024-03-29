package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Bell is located in 
type Bell struct {
}

func (r Bell) NPCId() uint32 {
	return npc.Bell
}

func (r Bell) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.SubwayTicketingBooth {
		m := message.NewBuilder().
			AddText("The ride to New Leaf City of Masteria takes off every minute, beginning on the hour, and it'll cost you ").
			BlueText().AddText(fmt.Sprintf("%d mesos", 5000)).
			BlackText().AddText(". Are you sure you want to purchase ").
			BlueText().ShowItemName1(item.SubwayTicketToNLCRegular).
			BlackText().AddText("?")
		return script.SendYesNo(l, span, c, m.String(), r.ValidatePurchase(item.SubwayTicketToNLCRegular), script.Exit())
	} else if c.MapId == _map.NLCSubwayStation {
		m := message.NewBuilder().
			AddText("The ride to Kerning City of Victoria Island takes off every minute, beginning on the hour, and it'll cost you ").
			BlueText().AddText(fmt.Sprintf("%d mesos", 5000)).
			BlackText().AddText(". Are you sure you want to purchase ").
			BlueText().ShowItemName1(item.SubwayTicketToKerningCityRegular).
			BlackText().AddText("?")
		return script.SendYesNo(l, span, c, m.String(), r.ValidatePurchase(item.SubwayTicketToKerningCityRegular), script.Exit())
	} else if c.MapId == _map.WaitingRoomFromNewLeafCityToKerningCity {
		m := message.NewBuilder().
			AddText("Do you want to leave before the train start? There will be no refund.")
		return script.SendYesNo(l, span, c, m.String(), script.Warp(_map.NLCSubwayStation), script.Exit())
	} else if c.MapId == _map.WaitingRoomFromKerningCityToNewLeafCity {
		m := message.NewBuilder().
			AddText("Do you want to leave before the train start? There will be no refund.")
		return script.SendYesNo(l, span, c, m.String(), script.Warp(_map.SubwayTicketingBooth), script.Exit())
	}
	return script.Exit()(l, span, c)
}

func (r Bell) ValidatePurchase(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, 5000) {
			return r.NotEnoughMeso(l, span, c)
		}
		if !character.CanHold(l)(c.CharacterId, itemId) {
			return r.NeedInventoryRoom(l, span, c)
		}
		return r.ProcessPurchase(itemId)(l, span, c)
	}
}

func (r Bell) ProcessPurchase(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainMeso(l, span)(c.CharacterId, -5000)
		character.GainItem(l, span)(c.CharacterId, itemId, 1)
		return r.ThereYouGo(l, span, c)
	}
}

func (r Bell) ThereYouGo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("There you go.")
	return script.SendOk(l, span, c, m.String())
}

func (r Bell) NeedInventoryRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't have a etc. slot available.")
	return script.SendOk(l, span, c, m.String())
}

func (r Bell) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't have enough mesos.")
	return script.SendOk(l, span, c, m.String())
}
