package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Asesson is located in Ariant - Ariant Station Platform (260000100)
type Asesson struct {
}

func (r Asesson) NPCId() uint32 {
	return npc.Asesson
}

func (r Asesson) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.MissingTicket(l, c)
	}
	if !character.TransportBoarding(l)(c.CharacterId, _map.AriantStationPlatform, _map.OrbisStationEntrance) {
		return r.ReadyForTakeoff(l, c)
	}
	return r.Confirm(l, c)
}

func (r Asesson) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("This will not be a short flight, so you need to take care of some things, I suggest you do that first before getting on board. Do you still wish to board the genie?")
	return SendYesNo(l, c, m.String(), r.Validate, r.ChangeYourMind)
}

func (r Asesson) ReadyForTakeoff(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("This genie is getting ready for takeoff. I'm sorry, but you'll have to get on the next ride. The ride schedule is available through the guide at the ticketing booth.")
	return SendOk(l, c, m.String())
}

func (r Asesson) MissingTicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Make sure you have a ticket to travel in this genie. Check your inventory.")
	return SendOk(l, c, m.String())
}

func (r Asesson) ChangeYourMind(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return SendOk(l, c, m.String())
}

func (r Asesson) Validate(l logrus.FieldLogger, c Context) State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.AriantStationPlatform, _map.OrbisStationEntrance) {
		return r.ReadyForTakeoff(l, c)
	}
	return r.Process(l, c)
}

func (r Asesson) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.TicketToOrbisRegular, -1)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.BeforeTakeoffToOrbisFromAriant, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.BeforeTakeoffToOrbisFromAriant, c.NPCId)
	}
	return Exit()(l, c)
}
