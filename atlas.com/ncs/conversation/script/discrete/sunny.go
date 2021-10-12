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

// Sunny is located in Orbis - Station <Ludibrium> (200000121)
type Sunny struct {
}

func (r Sunny) NPCId() uint32 {
	return npc.Sunny
}

func (r Sunny) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.TicketToLudibriumRegular) {
		return r.MissingTicket(l, span, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.OrbisStationLudibrium, _map.LudibriumStationOrbis) {
		return r.AlreadyTraveling(l, span, c)
	}

	return r.Confirm(l, span, c)
}

func (r Sunny) MissingTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Make sure you got a Ludibrium ticket to travel in this train. Check your inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r Sunny) AlreadyTraveling(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The train to Ludibrium is already travelling, please be patient for the next one.")
	return script.SendOk(l, span, c, m.String())
}

func (r Sunny) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you want to go to Ludibrium?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, r.ChangeYourMind)
}

func (r Sunny) ChangeYourMind(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return script.SendOk(l, span, c, m.String())
}

func (r Sunny) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.OrbisStationLudibrium, _map.LudibriumStationOrbis) {
		return r.BePatient(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r Sunny) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.TicketToLudibriumRegular, -1)
	return script.WarpById(_map.BeforeTheDepartureLudibrium, 0)(l, span, c)
}

func (r Sunny) BePatient(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The train to Ludibrium is ready to take off, please be patient for the next one.")
	return script.SendOk(l, span, c, m.String())
}
