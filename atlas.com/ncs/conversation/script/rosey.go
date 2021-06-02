package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Rosey is located in Orbis - Before the Departure <Ludibrium> (200000122)
type Rosey struct {
}

func (r Rosey) NPCId() uint32 {
	return npc.Rosey
}

func (r Rosey) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to leave the train?")
	return SendYesNo(l, c, m.String(), r.SeeYouNextTime, Exit())
}

func (r Rosey) SeeYouNextTime(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r Rosey) Warp(l logrus.FieldLogger, c Context) State {
	var mapId uint32
	if c.MapId == _map.BeforeTheDepartureLudibrium {
		mapId = _map.StationLudibrium
	} else {
		mapId = _map.StationOrbis
	}
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
	}
	return Exit()(l, c)
}
