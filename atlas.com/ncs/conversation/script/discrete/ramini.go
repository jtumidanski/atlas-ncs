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

// Ramini is located in  Orbis - Cabin <To Leafre> (200000131)
type Ramini struct {
}

func (r Ramini) NPCId() uint32 {
	return npc.Ramini
}

func (r Ramini) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.TicketToLeafreRegular) {
		return r.NeedTicket(l, span, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.CabinToLeafre, _map.LeafreStationEntrance) {
		return r.NotArrived(l, span, c)
	}

	return r.Confirm(l, span, c)
}

func (r Ramini) NeedTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Make sure you got a Leafre ticket to travel in this flight. Check your inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r Ramini) NotArrived(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The flight has not arrived yet. Come back soon.")
	return script.SendOk(l, span, c, m.String())
}

func (r Ramini) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to board the flight?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, r.ChangeYourMind)
}

func (r Ramini) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.CabinToLeafre, _map.LeafreStationEntrance) {
		return r.NotArrived(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r Ramini) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.TicketToLeafreRegular, -1)
	return script.WarpById(_map.CabinToLeafre2, 0)(l, span, c)
}

func (r Ramini) ChangeYourMind(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return script.SendOk(l, span, c, m.String())
}
