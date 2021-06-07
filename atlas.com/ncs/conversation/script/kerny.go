package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Kerny is located in 
type Kerny struct {
}

func (r Kerny) NPCId() uint32 {
	return npc.Kerny
}

func (r Kerny) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.BeforeDepartureToKerningCity {
		m := message.NewBuilder().AddText("The plane is taking off soon, are you sure you want to leave now? The ticket is not refundable.")
		return SendYesNo(l, c, m.String(), r.WarpBackToSingapore, Exit())
	}
	if c.MapId == _map.OnTheWayToKerningCity {
		m := message.NewBuilder().AddText("We're reaching Kerning City in a minute, please sit down and wait.")
		return SendOk(l, c, m.String())
	}
	if c.MapId == _map.OnTheWayToCBD {
		m := message.NewBuilder().AddText("We're reaching Singapore in a minute, please sit down and wait.")
		return SendOk(l, c, m.String())
	}
	return Exit()(l, c)
}

func (r Kerny) WarpBackToSingapore(l logrus.FieldLogger, c Context) State {
	err := npc.Warp(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ChangiAirport)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ChangiAirport, c.NPCId)
	}
	return r.SeeYouAgain(l, c)
}

func (r Kerny) SeeYouAgain(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hope to see you again soon!")
	return SendOk(l, c, m.String())
}
