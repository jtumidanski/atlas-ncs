package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Cherry is located in Victoria Road - Ellinia Station (101000300)
type Cherry struct {
}

func (r Cherry) NPCId() uint32 {
	return npc.Cherry
}

func (r Cherry) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.CheckInventory(l, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.BeforeTakeoffToOrbis, _map.OrbisStationEnterance) {
		return r.AlreadyTraveling(l, c)
	}

	return r.WantToGo(l, c)
}

func (r Cherry) CheckInventory(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Make sure you got a Orbis ticket to travel in this boat. Check your inventory.")
	return SendOk(l, c, m.String())
}

func (r Cherry) AlreadyTraveling(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The boat to Orbis is already travelling, please be patient for the next one.")
	return SendOk(l, c, m.String())
}

func (r Cherry) WantToGo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you want to go to Orbis?")
	return SendYesNo(l, c, m.String(), r.ValidateTravel, r.TalkToMeAgain)
}

func (r Cherry) TalkToMeAgain(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Okay, talk to me if you change your mind!")
	return SendOk(l, c, m.String())
}

func (r Cherry) ValidateTravel(l logrus.FieldLogger, c Context) State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.BeforeTakeoffToOrbis, _map.OrbisStationEnterance) {
		return r.BePatient(l, c)
	}

	character.GainItem(l)(c.CharacterId, item.TicketToOrbisRegular, -1)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.BeforeTakeoffToOrbis, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.BeforeTakeoffToOrbis, c.NPCId)
	}
	return nil
}

func (r Cherry) BePatient(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The boat to Orbis is ready to take off, please be patient for the next one.")
	return SendOk(l, c, m.String())
}
