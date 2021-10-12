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

// Tommie is located in Leafre - Station (240000110)
type Tommie struct {
}

func (r Tommie) NPCId() uint32 {
	return npc.Tommie
}

func (r Tommie) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.NeedTicket(l, span, c)
	}
	if !character.TransportBoarding(l)(c.CharacterId, _map.Station, _map.OrbisStationEntrance) {
		return r.ComeBackSoon(l, span, c)
	}
	return r.Confirm(l, span, c)
}

func (r Tommie) NeedTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Make sure you got an Orbis ticket to travel in this flight. Check your inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tommie) ComeBackSoon(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The flight has not arrived yet. Come back soon.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tommie) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to board the flight?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, r.TalkToMe)
}

func (r Tommie) TalkToMe(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return script.SendOk(l, span, c, m.String())
}

func (r Tommie) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.Station, _map.OrbisStationEntrance) {
		return r.ComeBackSoon(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r Tommie) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.TicketToOrbisRegular, -1)
	return script.WarpById(_map.BeforeTakeoffToOrbisFromLeafre, 0)(l, span, c)
}
