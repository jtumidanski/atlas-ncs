package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Agatha is located in Orbis - Orbis Station Enterence (200000100)
type Agatha struct {
}

func (r Agatha) NPCId() uint32 {
	return npc.Agatha
}

func (r Agatha) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello, I'm in charge of selling tickets for the ship ride for every destination. Which ticket would you like to purchase?").NewLine().
		OpenItem(0).BlueText().AddText("Ellinia").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Ludibrium").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Leafre").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Ariant").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r Agatha) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Explanation("Ellinia", 15, 5000, item.TicketToElliniaRegular)
	case 1:
		return r.Explanation("Ludibrium", 10, 6000, item.TicketToLudibriumRegular)
	case 2:
		return r.Explanation("Leafre", 10, 30000, item.TicketToLeafreRegular)
	case 3:
		return r.Explanation("Ariant", 10, 6000, item.TicketToAriantRegular)
	}
	return nil
}

func (r Agatha) Explanation(townName string, duration uint32, cost uint32, ticket uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("The ride to ").
			AddText(townName).
			AddText(" takes off every ").
			AddText(fmt.Sprintf("%d", duration)).
			AddText(" minutes, beginning on the hour, and it'll cost you ").
			BlueText().AddText(fmt.Sprintf("%d mesos", cost)).
			BlackText().AddText(". Are you sure you want to purchase ").
			BlueText().ShowItemName1(ticket).
			BlackText().AddText("?")
		return script.SendYesNo(l, span, c, m.String(), r.ValidateChoice(ticket, cost), script.Exit())
	}
}

func (r Agatha) ValidateChoice(ticket uint32, cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l)(c.CharacterId, cost) || !character.CanHold(l)(c.CharacterId, ticket) {
			return r.AreYouSure(cost)(l, span, c)
		}
		return r.ProcessPurchase(ticket, cost)(l, span, c)
	}
}

func (r Agatha) AreYouSure(cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText("Are you sure you have ").
			BlueText().AddText(fmt.Sprintf("%d mesos", cost)).
			BlackText().AddText("? If so, then I urge you to check you etc. inventory, and see if it's full or not.")
		return script.SendOk(l, span, c, m.String())
	}
}

func (r Agatha) ProcessPurchase(ticket uint32, cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		err := character.GainMeso(l)(c.CharacterId, -int32(cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process purchase for character %d.", c.CharacterId)
		}
		character.GainItem(l)(c.CharacterId, ticket, 1)
		return script.Exit()(l, span, c)
	}
}
