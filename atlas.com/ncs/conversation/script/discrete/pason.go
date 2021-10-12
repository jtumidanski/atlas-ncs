package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Pason is located in Victoria Road - Lith Harbor (104000000)
type Pason struct {
}

func (r Pason) NPCId() uint32 {
	return npc.Pason
}

func (r Pason) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.HaveYouHeard(l, span, c)
}

func (r Pason) HaveYouHeard(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Have you heard of the beach with a spectacular view of the ocean called ").
		BlueText().AddText("Florina Beach").
		BlackText().AddText(", located near Lith Harbor? I can take you there right now for either ").
		BlueText().AddText("1500 mesos").
		BlackText().AddText(", or if you have a ").
		BlueText().AddText("VIP Ticket to Florina Beach").
		BlackText().AddText(" with you, in which case you'll be there for free.").
		NewLine().NewLine().
		OpenItem(0).BlueText().AddText("I'll pay 1500 mesos.").CloseItem().NewLine().
		OpenItem(1).AddText("I have a VIP Ticket to Florina Beach.").CloseItem().NewLine().
		OpenItem(2).AddText("What is a VIP Ticket to Florina Beach?").CloseItem().BlackText()
	return script.SendListSelectionExit(l, span, c, m.String(), r.Choices, r.UnfinishedBusiness)
}

func (r Pason) Choices(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.MesoPayment
	case 1:
		return r.VIPTicket
	case 2:
		return r.WhatIsVIP
	}
	return nil
}

func (r Pason) MesoPayment(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 1500) {
		return r.LackingMeso(l, span, c)
	}
	return r.ProcessMeso(l, span, c)
}

func (r Pason) UnfinishedBusiness(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You must have some business to take care of here. You must be tired from all that traveling and hunting. Go take some rest, and if you feel like changing your mind, then come talk to me.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r Pason) LackingMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I think you're lacking mesos. There are many ways to gather up some money, you know, like... selling your armor... defeating monsters... doing quests... you know what I'm talking about.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r Pason) ProcessMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -1500)
	return r.Warp(l, span, c)
}

func (r Pason) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.SaveLocation(l)(c.CharacterId, "FLORINA")
	return script.WarpByName(_map.FlorinaBeach, "st00")(l, span, c)
}

func (r Pason) VIPTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So you have a ").
		BlueText().AddText("VIP Ticket to Florina Beach").
		BlackText().AddText("? You can always head over to Florina Beach with that. Alright then, but just be aware that you may be running into some monsters there too. Okay, would you like to head over to Florina Beach right now?")
	return script.SendYesNoExit(l, span, c, m.String(), r.ValidateVIPTicket, r.UnfinishedBusiness, r.UnfinishedBusiness)
}

func (r Pason) LackingVIP(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm, so where exactly is your ").
		BlueText().AddText("VIP Ticket to Florina Beach").
		BlackText().AddText("? Are you sure you have one? Please double-check.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r Pason) WhatIsVIP(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You must be curious about a ").
		BlueText().AddText("VIP Ticket to Florina Beach").
		BlackText().AddText(". Haha, that's very understandable. A VIP Ticket to Florina Beach is an item where as long as you have in possession, you may make your way to Florina Beach for free. It's such a rare item that even we had to buy those, but unfortunately I lost mine a few weeks ago during my precious summer break.")
	return script.SendNext(l, span, c, m.String(), r.MayBeAbleToPickItUp)
}

func (r Pason) ValidateVIPTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.VIPTicketToFlorinaBeach) {
		return r.LackingVIP(l, span, c)
	}
	return r.Warp(l, span, c)
}

func (r Pason) MayBeAbleToPickItUp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I came back without it, and it just feels awful not having it. Hopefully someone picked it up and put it somewhere safe. Anyway, this is my story and who knows, you may be able to pick it up and put it to good use. If you have any questions, feel free to ask.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}
