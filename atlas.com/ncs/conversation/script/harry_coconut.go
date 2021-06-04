package script

import (
	"atlas-ncs/character"
	"atlas-ncs/event"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// HarryCoconut is located in Hidden Street - Coconut Harvest (109080000), Hidden Street - Coconut Harvest (109080001), Hidden Street - Coconut Harvest (109080002), and Hidden Street - G? Coconut Season (109080003)
type HarryCoconut struct {
}

func (r HarryCoconut) NPCId() uint32 {
	return npc.HarryCoconut
}

func (r HarryCoconut) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Man... It is hot!!!~ How can I help you?").NewLine().
		OpenItem(0).BlueText().AddText("Leave the event game.").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Buy the weapon (Wooden Club 1 meso)").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r HarryCoconut) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm
	case 1:
		return r.Buy
	}
	return nil
}

func (r HarryCoconut) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you leave now, you can't participate in this event for the next 24 hours. Are you sure you want to leave?")
	return SendYesNo(l, c, m.String(), r.Leave, Exit())
}

func (r HarryCoconut) Leave(l logrus.FieldLogger, c Context) State {
	event.LeaveEvent(l)(c.WorldId, c.ChannelId, c.CharacterId)
	return Exit()(l, c)
}

func (r HarryCoconut) Buy(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 1) || !character.CanHold(l)(c.CharacterId, item.WoodenClub) {
		return r.NoMesoOrSpace(l, c)
	}
	return r.Process(l, c)
}

func (r HarryCoconut) Process(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -1)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment from character %d.", c.CharacterId)
	}
	character.GainItem(l)(c.CharacterId, item.WoodenClub, 1)
	return Exit()(l, c)
}

func (r HarryCoconut) NoMesoOrSpace(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You don't have enough mesos or you don't have any space in your inventory.")
	return SendOk(l, c, m.String())
}
