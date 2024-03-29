package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Alcaster is located in El Nath - El Nath Market (211000100)
type Alcaster struct {
}

func (r Alcaster) NPCId() uint32 {
	return npc.Alcaster
}

func (r Alcaster) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsCompleted(l)(c.CharacterId, 3035) {
		return r.HelpMeOut(l, span, c)
	}
	return r.ThanksToYou(l, span, c)
}

func (r Alcaster) HelpMeOut(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you decide to help me out, then in return, I'll make the item available for sale.")
	return script.SendOk(l, span, c, m.String())
}

func (r Alcaster) TakeALookAround(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I see. Understand that I have many different items here. Take a look around. I'm only selling these items to you, so I won't be ripping you off in any way shape or form.")
	return script.SendOk(l, span, c, m.String())
}

func (r Alcaster) ThanksToYou(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Thanks to you ").
		BlueText().ShowItemName1(item.TheBookOfAncient).
		BlackText().AddText(" is safely sealed. Of course, also as a result, I used up about half of the power I have accumulated over the last 800 years or so...but now I can die in peace. Oh, by the way... are you looking for rare items by any chance? As a sign of appreciation for your hard work, I'll sell some items I have to you, and ONLY you. Pick out the one you want!").NewLine().
		OpenItem(0).BlueText().ShowItemName1(item.HolyWater).AddText(fmt.Sprintf(" (Price: %d mesos", 300)).CloseItem().NewLine().
		OpenItem(1).BlueText().ShowItemName1(item.AllCurePotion).AddText(fmt.Sprintf(" (Price: %d mesos", 400)).CloseItem().NewLine().
		OpenItem(2).BlueText().ShowItemName1(item.TheMagicRock).AddText(fmt.Sprintf(" (Price: %d mesos", 5000)).CloseItem().NewLine().
		OpenItem(3).BlueText().ShowItemName1(item.TheSummoningRock).AddText(fmt.Sprintf(" (Price: %d mesos", 5000)).CloseItem()
	return script.SendListSelectionExit(l, span, c, m.String(), r.Selection, r.TakeALookAround)
}

func (r Alcaster) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.HolyWater, 300, "that cures the state of being sealed and cursed")
	case 1:
		return r.Confirm(item.AllCurePotion, 400, "that cures all")
	case 2:
		return r.Confirm(item.TheMagicRock, 5000, ", possessing magical power, that is used for high-quality skills")
	case 3:
		return r.Confirm(item.TheSummoningRock, 5000, ", possessing the power of summoning that is used for high-quality skills")
	}
	return nil
}

func (r Alcaster) Confirm(itemId uint32, cost uint32, use string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Is ").
			BlueText().ShowItemName1(itemId).
			BlackText().AddText(" really the item that you need? It's the item ").AddText(use).AddText(". It may not be the easiest item to acquire, but I'll give you a good deal on it. It'll cost you ").
			BlueText().AddText(fmt.Sprintf("%d mesos", cost)).
			BlackText().AddText(" per item. How many would you like to purchase?")
		return script.SendGetNumber(l, span, c, m.String(), r.ConfirmQuantity(itemId, cost), 0, 1, 100)
	}
}

func (r Alcaster) ConfirmQuantity(itemId uint32, cost uint32) script.ProcessNumber {
	return func(selection int32) script.StateProducer {
		return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
			if selection <= 0 {
				return r.ICannotSell(l, span, c)
			}

			m := message.NewBuilder().AddText("Are you sure you want to buy ").
				RedText().AddText(fmt.Sprintf("%d ", selection)).ShowItemName1(itemId).AddText("(s)").
				BlackText().AddText("? It 'll cost you ").AddText(fmt.Sprintf("%d", cost)).AddText(" mesos per ").
				ShowItemName1(itemId).AddText(", which will cost you ").
				RedText().AddText(fmt.Sprintf("%d mesos", cost*uint32(selection))).
				BlackText().AddText(" in total.")
			return script.SendYesNo(l, span, c, m.String(), r.Validate(itemId, cost, uint32(selection)), r.TakeALookAround)
		}
	}
}

func (r Alcaster) ICannotSell(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you're not going to buy anything, then I've got nothing to sell neither.")
	return script.SendOk(l, span, c, m.String())
}

func (r Alcaster) Validate(itemId uint32, cost uint32, quantity uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, cost*quantity) || !character.CanHoldAll(l)(c.CharacterId, itemId, quantity) {
			m := message.NewBuilder().
				AddText("Are you sure you have enough mesos? Please check and see if your etc. or use inventory is full, or if you have at least ").
				RedText().AddText(fmt.Sprintf("%d", quantity*cost)).
				BlackText().AddText(" mesos.")
			return script.SendOk(l, span, c, m.String())
		}

		character.GainMeso(l, span)(c.CharacterId, -int32(quantity*cost))
		character.GainItem(l, span)(c.CharacterId, itemId, int32(quantity))
		return r.Success(l, span, c)
	}
}

func (r Alcaster) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Thank you. If you ever find yourself needing items down the road, make sure to drop by here. I may have gotten old over the years, but I can still make magic items with ease.")
	return script.SendOk(l, span, c, m.String())
}
