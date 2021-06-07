package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Tommie is located in Leafre - Station (240000110)
type Tommie struct {
}

func (r Tommie) NPCId() uint32 {
	return npc.Tommie
}

func (r Tommie) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.NeedTicket(l, c)
	}
	if !character.TransportBoarding(l)(c.CharacterId, _map.Station, _map.OrbisStationEntrance) {
		return r.ComeBackSoon(l, c)
	}
	return r.Confirm(l, c)
}

func (r Tommie) NeedTicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Make sure you got an Orbis ticket to travel in this flight. Check your inventory.")
	return SendOk(l, c, m.String())
}

func (r Tommie) ComeBackSoon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The flight has not arrived yet. Come back soon.")
	return SendOk(l, c, m.String())
}

func (r Tommie) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to board the flight?")
	return SendYesNo(l, c, m.String(), r.Validate, r.TalkToMe)
}

func (r Tommie) TalkToMe(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Okay, talk to me if you change your mind!")
	return SendOk(l, c, m.String())
}

func (r Tommie) Validate(l logrus.FieldLogger, c Context) State {
	if !character.TransportBoarding(l)(c.CharacterId, _map.Station, _map.OrbisStationEntrance) {
		return r.ComeBackSoon(l, c)
	}
	return r.Process(l, c)
}

func (r Tommie) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.TicketToOrbisRegular, -1)
	return WarpById(_map.BeforeTakeoffToOrbisFromLeafre, 0)(l, c)
}
