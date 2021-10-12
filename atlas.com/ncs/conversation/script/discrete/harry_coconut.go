package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/event"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// HarryCoconut is located in Hidden Street - Coconut Harvest (109080000), Hidden Street - Coconut Harvest (109080001), Hidden Street - Coconut Harvest (109080002), and Hidden Street - G? Coconut Season (109080003)
type HarryCoconut struct {
}

func (r HarryCoconut) NPCId() uint32 {
	return npc.HarryCoconut
}

func (r HarryCoconut) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Man... It is hot!!!~ How can I help you?").NewLine().
		OpenItem(0).BlueText().AddText("Leave the event game.").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Buy the weapon (Wooden Club 1 meso)").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r HarryCoconut) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Confirm
	case 1:
		return r.Buy
	}
	return nil
}

func (r HarryCoconut) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If you leave now, you can't participate in this event for the next 24 hours. Are you sure you want to leave?")
	return script.SendYesNo(l, span, c, m.String(), r.Leave, script.Exit())
}

func (r HarryCoconut) Leave(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	event.LeaveEvent(l)(c.WorldId, c.ChannelId, c.CharacterId)
	return script.Exit()(l, span, c)
}

func (r HarryCoconut) Buy(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 1) || !character.CanHold(l)(c.CharacterId, item.WoodenClub) {
		return r.NoMesoOrSpace(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r HarryCoconut) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -1)
	character.GainItem(l, span)(c.CharacterId, item.WoodenClub, 1)
	return script.Exit()(l, span, c)
}

func (r HarryCoconut) NoMesoOrSpace(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't have enough mesos or you don't have any space in your inventory.")
	return script.SendOk(l, span, c, m.String())
}
