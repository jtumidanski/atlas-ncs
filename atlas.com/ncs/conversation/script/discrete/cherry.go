package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
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

func (r Cherry) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.CheckInventory(l, c)
	}

	if !character.TransportBoarding(l)(c.CharacterId, _map.BeforeTakeoffToOrbisFromEllinia, _map.OrbisStationEntrance) {
		return r.AlreadyTraveling(l, c)
	}

	return r.WantToGo(l, c)
}

func (r Cherry) CheckInventory(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Make sure you got a Orbis ticket to travel in this boat. Check your inventory.")
	return script.SendOk(l, c, m.String())
}

func (r Cherry) AlreadyTraveling(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The boat to Orbis is already travelling, please be patient for the next one.")
	return script.SendOk(l, c, m.String())
}

func (r Cherry) WantToGo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Do you want to go to Orbis?")
	return script.SendYesNo(l, c, m.String(), r.ValidateTravel, r.TalkToMeAgain)
}

func (r Cherry) TalkToMeAgain(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Okay, talk to me if you change your mind!")
	return script.SendOk(l, c, m.String())
}

func (r Cherry) ValidateTravel(l logrus.FieldLogger, c script.Context) script.State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.BeforeTakeoffToOrbisFromEllinia, _map.OrbisStationEntrance) {
		return r.BePatient(l, c)
	}

	character.GainItem(l)(c.CharacterId, item.TicketToOrbisRegular, -1)
	return script.WarpById(_map.BeforeTakeoffToOrbisFromEllinia, 0)(l, c)
}

func (r Cherry) BePatient(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The boat to Orbis is ready to take off, please be patient for the next one.")
	return script.SendOk(l, c, m.String())
}