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

// Jane is located in Victoria Road - Lith Harbor (104000000)
type Jane struct {
}

func (r Jane) NPCId() uint32 {
	return npc.Jane
}

func (r Jane) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsCompleted(l)(c.CharacterId, 2013) {
		return r.ItsYou(l, span, c)
	} else if quest.IsCompleted(l)(c.CharacterId, 2010) {
		return r.NotStrongEnough(l, span, c)
	} else {
		return r.MyDream(l, span, c)
	}
}

func (r Jane) MyDream(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("My dream is to travel everywhere, much like you. My father, however, does not allow me to do it, because he thinks it's very dangerous. He may say yes, though, if I show him some sort of a proof that I'm not the weak girl that he thinks I am ...")
	return script.SendOk(l, span, c, m.String())
}

func (r Jane) NotStrongEnough(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You don't seem strong enough to be able to purchase my potion ...")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r Jane) ItsYou(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It's you ... thanks to you I was able to get a lot done. Nowadays I've been making a bunch of items. If you need anything let me know.")
	return script.SendNextExit(l, span, c, m.String(), r.WhatToBuy, r.StillHaveAFew)
}

func (r Jane) StillHaveAFew(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I still have quite a few of the materials you got me before. The items are all there so take your time choosing.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r Jane) WhatToBuy(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Which item would you like to buy?").NewLine().
		BlueText().
		OpenItem(0).ShowItemImage2(item.WhitePotion).AddText(fmt.Sprintf(" (Price : %d mesos)", 310)).CloseItem().NewLine().
		OpenItem(1).ShowItemImage2(item.Unagi).AddText(fmt.Sprintf(" (Price : %d mesos)", 1060)).CloseItem().NewLine().
		OpenItem(2).ShowItemImage2(item.PureWater).AddText(fmt.Sprintf(" (Price : %d mesos)", 1600)).CloseItem().NewLine().
		OpenItem(3).ShowItemImage2(item.Watermelon).AddText(fmt.Sprintf(" (Price : %d mesos)", 3120)).CloseItem().NewLine()
	return script.SendListSelectionExit(l, span, c, m.String(), r.SelectItem, r.StillHaveAFew)
}

func (r Jane) SelectItem(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.HowMany(item.WhitePotion, 310, "300 HP")
	case 1:
		return r.HowMany(item.Unagi, 1060, "1000 HP")
	case 2:
		return r.HowMany(item.PureWater, 1600, "800 MP")
	case 3:
		return r.HowMany(item.Watermelon, 3120, "1000 HP and MP")
	}
	return nil
}

func (r Jane) HowMany(itemId uint32, cost uint32, recover string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("You want ").
			BlueText().ShowItemName1(itemId).
			BlackText().AddText("? ").ShowItemName1(itemId).AddText(" allows you to recover ").AddText(recover).AddText(" How many would you like to buy?")
		return script.SendGetNumberExit(l, span, c, m.String(), r.Confirmation(itemId, cost), r.StillHaveAFew, 1, 1, 100)
	}
}

func (r Jane) Confirmation(itemId uint32, cost uint32) func(selection int32) script.StateProducer {
	return func(selection int32) script.StateProducer {
		return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
			m := message.NewBuilder().
				AddText("Will you purchase ").
				RedText().AddText(fmt.Sprintf("%d", selection)).
				BlackText().AddText(" ").
				BlueText().ShowItemName1(itemId).AddText("(s)").
				BlackText().AddText("? ").ShowItemName1(itemId).AddText(fmt.Sprintf(" costs %d mesos for one, so the total comes out to be ", cost)).
				RedText().AddText(fmt.Sprintf("%d", uint32(selection)*cost)).
				BlackText().AddText(" mesos.")
			return script.SendYesNoExit(l, span, c, m.String(), r.ValidateTransaction(itemId, uint32(selection)*cost, selection), r.StillHaveAFew, r.StillHaveAFew)
		}
	}
}

func (r Jane) ValidateTransaction(itemId uint32, cost uint32, amount int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, cost) {
			return r.LackingMesos(cost)(l, span, c)
		}
		if !character.CanHold(l)(c.CharacterId, itemId) {
			return r.LackingInventory(l, span, c)
		}
		return r.PerformTransaction(itemId, cost, amount)(l, span, c)
	}
}

func (r Jane) LackingMesos(cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Are you lacking mesos by any chance? Please check and see if you have an empty slot available at your etc. inventory, and if you have at least ").
			RedText().AddText(fmt.Sprintf("%d", cost)).
			BlackText().AddText(" mesos with you.")
		return script.SendNext(l, span, c, m.String(), script.Exit())
	}
}

func (r Jane) LackingInventory(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please check and see if you have an empty slot available at your etc. inventory.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r Jane) PerformTransaction(itemId uint32, cost uint32, amount int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainMeso(l, span)(c.CharacterId, -int32(cost))
		character.GainItem(l, span)(c.CharacterId, itemId, amount)
		return r.Success(l, span, c)
	}
}

func (r Jane) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Thank you for coming. Stuff here can always be made so if you need something, please come again.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}
