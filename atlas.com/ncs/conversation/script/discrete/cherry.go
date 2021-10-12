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

// Cherry is located in Victoria Road - Ellinia Station (101000300)
type Cherry struct {
}

func (r Cherry) NPCId() uint32 {
	return npc.Cherry
}

func (r Cherry) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.CheckInventory(l, span, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.BeforeTakeoffToOrbisFromEllinia, _map.OrbisStationEntrance) {
		return r.AlreadyTraveling(l, span, c)
	}

	return r.WantToGo(l, span, c)
}

func (r Cherry) CheckInventory(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Make sure you got a Orbis ticket to travel in this boat. Check your inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r Cherry) AlreadyTraveling(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The boat to Orbis is already travelling, please be patient for the next one.")
	return script.SendOk(l, span, c, m.String())
}

func (r Cherry) WantToGo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Do you want to go to Orbis?")
	return script.SendYesNo(l, span, c, m.String(), r.ValidateTravel, r.TalkToMeAgain)
}

func (r Cherry) TalkToMeAgain(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Okay, talk to me if you change your mind!")
	return script.SendOk(l, span, c, m.String())
}

func (r Cherry) ValidateTravel(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.BeforeTakeoffToOrbisFromEllinia, _map.OrbisStationEntrance) {
		return r.BePatient(l, span, c)
	}

	character.GainItem(l, span)(c.CharacterId, item.TicketToOrbisRegular, -1)
	return script.WarpById(_map.BeforeTakeoffToOrbisFromEllinia, 0)(l, span, c)
}

func (r Cherry) BePatient(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The boat to Orbis is ready to take off, please be patient for the next one.")
	return script.SendOk(l, span, c, m.String())
}
