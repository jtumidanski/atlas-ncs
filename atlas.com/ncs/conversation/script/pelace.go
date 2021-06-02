package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Pelace is located in  Orbis - Cabin <To Leafre> (200000132)
type Pelace struct {
}

func (r Pelace) NPCId() uint32 {
	return npc.Pelace
}

func (r Pelace) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to leave the flight?")
	return SendYesNo(l, c, m.String(), r.Alright, Exit())
}

func (r Pelace) Alright(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r Pelace) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.CabinToLeafre, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.CabinToLeafre, 0)
	}
	return Exit()(l, c)
}
