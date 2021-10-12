package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Shuri is located in Orbis - Orbis (200000000)
type Shuri struct {
}

func (r Shuri) NPCId() uint32 {
	return npc.Shuri
}

func (r Shuri) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.HaveYouHeard(l, span, c)
}

func (r Shuri) HaveYouHeard(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Have you heard of the beach with a spectacular view of the ocean called ").
		BlueText().ShowMap(_map.FlorinaBeach).
		BlackText().AddText(", located a little far from ").ShowMap(c.MapId).AddText("? I can take you there right now for either ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 2000)).
		BlackText().AddText(", or if you have ").
		BlueText().ShowItemName1(item.VIPTicketToFlorinaBeach).
		BlackText().AddText(" with you, in which case you'll be in for free.").NewLine().
		OpenItem(0).BlueText().AddText(fmt.Sprintf("I'll pay %d mesos.", 2000)).CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("I have ").ShowItemName1(item.VIPTicketToFlorinaBeach).CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("What is ").ShowItemName1(item.VIPTicketToFlorinaBeach).CloseItem()
	return script.SendListSelectionExit(l, span, c, m.String(), r.Selection, r.OtherBusiness)
}

func (r Shuri) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Payment
	case 1:
		return r.UseVIP
	case 2:
		return r.WhatIsVIP
	}
	return nil
}

func (r Shuri) Payment(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You want to pay ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 2000)).
		BlackText().AddText(" and leave for ").
		ShowMap(_map.FlorinaBeach).
		AddText("? Okay!! Please beware that you may be running into some monsters around there though, so make sure not to get caught off-guard. Okay, would you like to head over to ").
		ShowMap(_map.FlorinaBeach).
		AddText(" right now?")
	return script.SendYesNoExit(l, span, c, m.String(), r.ValidatePayment, r.OtherBusiness, r.OtherBusiness)
}

func (r Shuri) UseVIP(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So you have ").
		BlueText().ShowItemName1(item.VIPTicketToFlorinaBeach).
		BlackText().AddText("? You can always head over to ").
		ShowMap(_map.FlorinaBeach).
		AddText(" with that.").
		AddText(" Okay!! Please beware that you may be running into some monsters around there though, so make sure not to get caught off-guard. Okay, would you like to head over to ").
		ShowMap(_map.FlorinaBeach).
		AddText(" right now?")
	return script.SendYesNoExit(l, span, c, m.String(), r.ValidateVIP, r.OtherBusiness, r.OtherBusiness)
}

func (r Shuri) WhatIsVIP(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You must be curious about ").
		BlueText().ShowItemName1(item.VIPTicketToFlorinaBeach).
		BlackText().AddText(". Yeah, I can see that. ").
		ShowItemName1(item.VIPTicketToFlorinaBeach).
		AddText(" is an item where as long as you have in possession, you may make your way to ").
		ShowMap(_map.FlorinaBeach).
		AddText(" for free. It's such a rare item that even we had to buy those, but unfortunately I lost mine a few weeks ago during a long weekend.")
	return script.SendNext(l, span, c, m.String(), r.CameBackWithoutIt)
}

func (r Shuri) OtherBusiness(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You must have some business to take care of here. You must be tired from all that travelling and hunting. Go take some rest, and if you feel like changing your mind, then come talk to me.")
	return script.SendOk(l, span, c, m.String())
}

func (r Shuri) ValidatePayment(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 2000) {
		return r.NotEnoughMeso(l, span, c)
	}
	character.GainMeso(l, span)(c.CharacterId, -2000)
	return r.Warp(l, span, c)
}

func (r Shuri) ValidateVIP(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.VIPTicketToFlorinaBeach) {
		return r.MissingTicket(l, span, c)
	}
	return r.Warp(l, span, c)
}

func (r Shuri) CameBackWithoutIt(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I came back without it, and it just feels awful not having it. Hopefully someone picked it up and put it somewhere safe. Anyway this is my story and who knows, you may be able to pick it up and put it to good use. If you have any questions, feel free to ask")
	return script.SendPrevious(l, span, c, m.String(), r.WhatIsVIP)
}

func (r Shuri) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.SaveLocation(l)(c.CharacterId, "FLORINA")
	return script.WarpByName(_map.FlorinaBeach, "st00")(l, span, c)
}

func (r Shuri) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I think you're lacking mesos. There are many ways to gather up some money, you know, like ... selling your armor ... defeating the monsters ... doing quests ... you know what I'm talking about.")
	return script.SendOk(l, span, c, m.String())
}

func (r Shuri) MissingTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm, so where exactly is ").
		BlueText().ShowItemName1(item.VIPTicketToFlorinaBeach).
		BlackText().AddText("?? Are you sure you have them? Please double-check.")
	return script.SendOk(l, span, c, m.String())
}
