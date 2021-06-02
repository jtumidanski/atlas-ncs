package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Nara is located in Ludibrium - Ludibrium (220000000)
type Nara struct {
}

func (r Nara) NPCId() uint32 {
	return npc.Nara
}

func (r Nara) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Have you heard of the beach with a spectacular view of the ocean called ").
		BlueText().ShowMap(_map.FlorinaBeach).
		BlackText().AddText(", located a little far from ").
		ShowMap(_map.Ludibrium).
		AddText("? I can take you there right now for either ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 1800)).
		BlackText().AddText(", or if you have ").
		BlueText().ShowItemName1(item.VIPTicketToFlorinaBeach).
		BlackText().AddText(" with you, in which case you'll be in for free.").NewLine().
		OpenItem(0).BlueText().AddText(fmt.Sprintf("I'll pay %d mesos.", 1800)).CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("I have ").ShowItemName1(item.VIPTicketToFlorinaBeach).CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("What is ").ShowItemName1(item.VIPTicketToFlorinaBeach).AddText("?").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r Nara) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.MesoPayment
	case 1:
		return r.Ticket
	case 2:
		return r.TicketInfo
	}
	return nil
}

func (r Nara) TicketInfo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You must be curious about ").
		BlueText().ShowItemName1(item.VIPTicketToFlorinaBeach).
		BlackText().AddText(". Yeah, I can see that. ").
		ShowItemName1(item.VIPTicketToFlorinaBeach).
		AddText(" is an item where as long as you have in possession, you may make your way to ").
		ShowMap(_map.FlorinaBeach).
		AddText(" for free. It's such a rare item that even we had to buy those, but unfortunately I lost mine a few weeks ago during a long weekend.")
	return SendNext(l, c, m.String(), r.CameBackWithoutIt)
}

func (r Nara) CameBackWithoutIt(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I came back without it, and it just feels awful not having it. Hopefully someone picked it up and put it somewhere safe. Anyway this is my story and who knows, you may be able to pick it up and put it to good use. If you have any questions, feel free to ask")
	return SendPrevious(l, c, m.String(), r.TicketInfo)
}

func (r Nara) MesoPayment(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You want to pay ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 1800)).
		BlackText().AddText(" and leave for ").
		ShowMap(_map.FlorinaBeach).
		AddText("? Okay!! Please beware that you may be running into some monsters around there though, so make sure not to get caught off-guard. Okay, would you like to head over to ").
		ShowMap(_map.FlorinaBeach).AddText(" right now?")
	return SendYesNo(l, c, m.String(), r.ValidateMeso, r.MustHaveBusiness)
}

func (r Nara) MustHaveBusiness(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You must have some business to take care of here. You must be tired from all that traveling and hunting. Go take some rest, and if you feel like changing your mind, then come talk to me.")
	return SendOk(l, c, m.String())
}

func (r Nara) ValidateMeso(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 1800) {
		return r.NotEnoughMeso(l, c)
	}
	err := character.GainMeso(l)(c.CharacterId, -1800)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
	}
	return r.Warp(l, c)
}

func (r Nara) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I think you're lacking mesos. There are many ways to gather up some money, you know, like ... selling your armor ... defeating the monsters ... doing quests ... you know what I'm talking about.")
	return SendOk(l, c, m.String())
}

func (r Nara) Warp(l logrus.FieldLogger, c Context) State {
	character.SaveLocation(l)(c.CharacterId, "FLORINA")
	err := npc.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.FlorinaBeach, "st00")
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.FlorinaBeach, c.NPCId)
	}
	return Exit()(l, c)
}

func (r Nara) Ticket(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.VIPTicketToFlorinaBeach) {
		return r.MissingTicket(l, c)
	}
	return r.Warp(l, c)
}

func (r Nara) MissingTicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmmm, so where exactly is ").
		BlueText().ShowItemName1(item.VIPTicketToFlorinaBeach).
		BlackText().AddText("?? Are you sure you have them? Please double-check.")
	return SendOk(l, c, m.String())
}
