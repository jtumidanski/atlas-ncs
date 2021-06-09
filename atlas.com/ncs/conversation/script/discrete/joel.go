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

// Joel is located in Victoria Road - Ellinia Station (101000300)
type Joel struct {
}

func (r Joel) NPCId() uint32 {
	return npc.Joel
}

func (r Joel) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.Hello(l, c)
}

func (r Joel) Hello(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello, I'm in charge of selling tickets for the ship ride to Orbis Station of Ossyria. The ride to Orbis takes off every 15 minutes, beginning on the hour, and it'll cost you ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 5000)).
		BlackText().AddText(". Are you sure you want to purchase ").
		BlueText().ShowItemName1(item.TicketToOrbisRegular).
		BlackText().AddText("?")
	return script.SendYesNo(l, c, m.String(), r.ValidatePurchase, r.MoreToDo)
}

func (r Joel) ValidatePurchase(l logrus.FieldLogger, c script.Context) script.State {
	if !character.HasMeso(l)(c.CharacterId, 5000) || !character.CanHold(l)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.PurchaseFailed(l, c)
	}

	return r.ProcessPurchase(l, c)
}

func (r Joel) ProcessPurchase(l logrus.FieldLogger, c script.Context) script.State {
	err := character.GainMeso(l)(c.CharacterId, -5000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process purchase for character %d.", c.CharacterId)
	}
	character.GainItem(l)(c.CharacterId, item.TicketToOrbisRegular, 1)
	return script.Exit()(l, c)
}

func (r Joel) PurchaseFailed(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Are you sure you have ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 5000)).
		BlackText().AddText("? If so, then I urge you to check your etc. inventory, and see if it's full or not.")
	return script.SendOk(l, c, m.String())
}

func (r Joel) MoreToDo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You must have some business to take care of here, right?")
	return script.SendOk(l, c, m.String())
}

