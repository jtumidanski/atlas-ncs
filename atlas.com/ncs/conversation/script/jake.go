package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Jake is located in Victoria Road - Subway Ticketing Booth (103000100)
type Jake struct {
}

func (r Jake) NPCId() uint32 {
	return npc.Jake
}

func (r Jake) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Jake) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hi, I'm the ticket salesman.")
	return SendNext(l, c, m.String(), r.SelectTicket)
}

func (r Jake) SelectTicket(l logrus.FieldLogger, c Context) State {
	zones := 0
	if character.QuestStarted(l)(c.CharacterId, 2055) || character.QuestCompleted(l)(c.CharacterId, 2055) {
		zones++
	}
	if character.QuestStarted(l)(c.CharacterId, 2056) || character.QuestCompleted(l)(c.CharacterId, 2056) {
		zones++
	}
	if character.QuestStarted(l)(c.CharacterId, 2057) || character.QuestCompleted(l)(c.CharacterId, 2057) {
		zones++
	}
	if zones == 0 {
		return Exit()(l, c)
	}

	m := message.NewBuilder().
		AddText("Which ticket would you like?").NewLine()
	for i := 0; i < zones; i++ {
		m = m.OpenItem(i).BlueText().AddText(fmt.Sprintf("Construction Site B %d (%d mesos)", i, 1000)).CloseItem()
	}
	return SendListSelection(l, c, m.String(), r.SiteSelection)
}

func (r Jake) SiteSelection(selection int32) StateProducer {
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

func (r Jake) Validate(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, 1000) {
			return r.NotEnoughMeso(l, c)
		}
		return r.Process(itemId)(l, c)
	}
}

func (r Jake) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You do not have enough mesos.")
	return SendOk(l, c, m.String())
}

func (r Jake) Process(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -1000)
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
		}
		character.GainItem(l)(c.CharacterId, itemId, 1)
		return Exit()(l, c)
	}
}
