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

// Tian is located in Ludibrium - Station<Orbis> (220000110)
type Tian struct {
}

func (r Tian) NPCId() uint32 {
	return npc.Tian
}

func (r Tian) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.NeedTicket(l, span, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.StationOrbis, _map.OrbisStationEntrance) {
		return r.AlreadyTraveling(l, span, c)
	}

	return r.Confirm(l, span, c)
}

func (r Tian) NeedTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Make sure you got a Orbis ticket to travel in this train. Check your inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tian) AlreadyTraveling(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The train to Orbis is already travelling, please be patient for the next one.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tian) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Do you want to go to Orbis?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, r.TalkToMe)
}

func (r Tian) TalkToMe(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return script.SendOk(l, span, c, m.String())
}

func (r Tian) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.StationOrbis, _map.OrbisStationEntrance) {
		return r.BePatient(l, span, c)
	}
	return r.Warp(l, span, c)
}

func (r Tian) BePatient(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The train to Orbis is ready to take off, please be patient for the next one.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tian) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.TicketToOrbisRegular, -1)
	return script.WarpById(_map.BeforeTheDepartureOrbis, 0)(l, span, c)
}
