package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// ChiefTatamo is located in Leafre - Leafre (240000000)
type ChiefTatamo struct {
}

func (r ChiefTatamo) NPCId() uint32 {
	return npc.ChiefTatamo
}

func (r ChiefTatamo) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("..Can I help you?").NewLine().
		OpenItem(0).BlueText().AddText("Buy the Magic Seed").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Do something for Leafre").CloseItem()
	return script.SendListSelectionExit(l, c, m.String(), r.Selection, r.ThinkCarefully)
}

func (r ChiefTatamo) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.MagicSeed
	case 1:
		return r.DoSomething
	}
	return nil
}

func (r ChiefTatamo) MagicSeed(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You don't seem to be from out town. How can I help you?").NewLine().
		OpenItem(0).BlueText().AddText("I would like some ").ShowItemImage1(item.MagicSeed).BlackText().AddText(".").CloseItem()
	return script.SendListSelectionExit(l, c, m.String(), r.HelpSelection, r.ThinkCarefully)
}

func (r ChiefTatamo) DoSomething(l logrus.FieldLogger, c script.Context) script.State {
	//TODO
	m := message.NewBuilder().AddText("Under development...")
	return script.SendOk(l, c, m.String())
}

func (r ChiefTatamo) HelpSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.HowMany
	}
	return nil
}

func (r ChiefTatamo) ThinkCarefully(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please think carefully. Once you have made your decision, let me know.")
	return script.SendOk(l, c, m.String())
}

func (r ChiefTatamo) HowMany(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		BlueText().ShowItemName1(item.MagicSeed).
		BlackText().AddText(" is a precious item. I cannot give it to you just like that. How about doing me a little favor? Then I'll give it to you. I'll sell the ").
		BlueText().ShowItemName1(item.MagicSeed).
		BlackText().AddText(" to you for ").
		BlueText().AddText("30,000 mesos").
		BlackText().AddText(" each. Are you willing to make the purchase? How many would you like, then?")
	return script.SendGetNumberExit(l, c, m.String(), r.Confirm, r.ThinkCarefully, 0, 0, 100)
}

func (r ChiefTatamo) Confirm(selection int32) script.StateProducer {
	if selection == 0 {
		return r.CannotSellZero
	}
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Buying ").
			BlueText().AddText(fmt.Sprintf("%d ", selection)).ShowItemName1(item.MagicSeed).AddText("(s)").
			BlackText().AddText(" will cost you ").
			BlueText().AddText(fmt.Sprintf("%d mesos", selection*30000)).
			BlackText().AddText(". Are you sure you want to make the purchase?")
		return script.SendYesNo(l, c, m.String(), r.Validate(selection), r.ThinkCarefully)
	}
}

func (r ChiefTatamo) CannotSellZero(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I can't sell you 0.")
	return script.SendOk(l, c, m.String())
}

func (r ChiefTatamo) Validate(amount int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if !character.HasMeso(l)(c.CharacterId, uint32(amount)) || !character.CanHoldAll(l)(c.CharacterId, item.MagicSeed, uint32(amount)) {
			return r.PleaseCheck(l, c)
		}
		return r.Process(amount)(l, c)
	}
}

func (r ChiefTatamo) PleaseCheck(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please check and see if you have enough mesos to make the purchase. Also, I suggest you check the etc. inventory and see if you have enough space available to make the purchase.")
	return script.SendOk(l, c, m.String())
}

func (r ChiefTatamo) Process(amount int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		err := character.GainMeso(l)(c.CharacterId, -amount)
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment from character %d.", c.CharacterId)
		}
		character.GainItem(l)(c.CharacterId, item.MagicSeed, amount)
		return r.Success(l, c)
	}
}

func (r ChiefTatamo) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("See you again~")
	return script.SendOk(l, c, m.String())
}
