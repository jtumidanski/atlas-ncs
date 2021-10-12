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

// Jake is located in Victoria Road - Subway Ticketing Booth (103000100)
type Jake struct {
}

func (r Jake) NPCId() uint32 {
	return npc.Jake
}

func (r Jake) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r Jake) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hi, I'm the ticket salesman.")
	return script.SendNext(l, span, c, m.String(), r.SelectTicket)
}

func (r Jake) SelectTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	zones := 0
	if quest.IsStarted(l)(c.CharacterId, 2055) || quest.IsCompleted(l)(c.CharacterId, 2055) {
		zones++
	}
	if quest.IsStarted(l)(c.CharacterId, 2056) || quest.IsCompleted(l)(c.CharacterId, 2056) {
		zones++
	}
	if quest.IsStarted(l)(c.CharacterId, 2057) || quest.IsCompleted(l)(c.CharacterId, 2057) {
		zones++
	}
	if zones == 0 {
		return script.Exit()(l, span, c)
	}

	m := message.NewBuilder().
		AddText("Which ticket would you like?").NewLine()
	for i := 0; i < zones; i++ {
		m = m.OpenItem(i).BlueText().AddText(fmt.Sprintf("Construction Site B %d (%d mesos)", i, 1000)).CloseItem()
	}
	return script.SendListSelection(l, span, c, m.String(), r.SiteSelection)
}

func (r Jake) SiteSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Validate(item.TicketToConstructionSiteB1)
	case 1:
		return r.Validate(item.TicketToConstructionSiteB2)
	case 2:
		return r.Validate(item.TicketToConstructionSiteB3)
	}
	return nil
}

func (r Jake) Validate(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, 1000) {
			return r.NotEnoughMeso(l, span, c)
		}
		return r.Process(itemId)(l, span, c)
	}
}

func (r Jake) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You do not have enough mesos.")
	return script.SendOk(l, span, c, m.String())
}

func (r Jake) Process(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainMeso(l, span)(c.CharacterId, -1000)
		character.GainItem(l, span)(c.CharacterId, itemId, 1)
		return script.Exit()(l, span, c)
	}
}
