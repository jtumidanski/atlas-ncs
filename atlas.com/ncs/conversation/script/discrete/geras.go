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

// Geras is located in Orbis - Station <To Ariant> (200000151)
type Geras struct {
}

func (r Geras) NPCId() uint32 {
	return npc.Geras
}

func (r Geras) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.TicketToAriantRegular) {
		return r.MissingTicket(l, span, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.StationToAriant, _map.AriantStationPlatform) {
		return r.BePatient(l, span, c)
	}

	return r.Confirm(l, span, c)
}

func (r Geras) MissingTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Make sure you got an Ariant ticket to travel in this genie. Check your inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r Geras) BePatient(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("This genie is getting ready for takeoff. I'm sorry, but you'll have to get on the next ride. The ride schedule is available through the guide at the ticketing booth.")
	return script.SendOk(l, span, c, m.String())
}

func (r Geras) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("This will not be a short flight, so you need to take care of some things, I suggest you do that first before getting on board. Do you still wish to board the genie?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, r.ChangeYourMind)
}

func (r Geras) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.StationToAriant, _map.AriantStationPlatform) {
		return r.BePatient(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r Geras) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.TicketToAriantRegular, -1)
	return script.WarpById(_map.StationToAriant2, 0)(l, span, c)
}

func (r Geras) ChangeYourMind(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return script.SendOk(l, span, c, m.String())
}
