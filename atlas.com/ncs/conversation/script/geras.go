package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Geras is located in Orbis - Station <To Ariant> (200000151)
type Geras struct {
}

func (r Geras) NPCId() uint32 {
	return npc.Geras
}

func (r Geras) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToAriantRegular) {
		return r.MissingTicket(l, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.StationToAriant, _map.AriantStationPlatform) {
		return r.BePatient(l, c)
	}

	return r.Confirm(l, c)
}

func (r Geras) MissingTicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Make sure you got an Ariant ticket to travel in this genie. Check your inventory.")
	return SendOk(l, c, m.String())
}

func (r Geras) BePatient(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("This genie is getting ready for takeoff. I'm sorry, but you'll have to get on the next ride. The ride schedule is available through the guide at the ticketing booth.")
	return SendOk(l, c, m.String())
}

func (r Geras) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("This will not be a short flight, so you need to take care of some things, I suggest you do that first before getting on board. Do you still wish to board the genie?")
	return SendYesNo(l, c, m.String(), r.Validate, r.ChangeYourMind)
}

func (r Geras) Validate(l logrus.FieldLogger, c Context) State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.StationToAriant, _map.AriantStationPlatform) {
		return r.BePatient(l, c)
	}
	return r.Process(l, c)
}

func (r Geras) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.TicketToAriantRegular, -1)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.StationToAriant2, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.StationToAriant2, c.NPCId)
	}
	return Exit()(l, c)
}

func (r Geras) ChangeYourMind(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return SendOk(l, c, m.String())
}
