package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Egnet is located in Orbis - Station <To Ariant> (200000152)
type Egnet struct {
}

func (r Egnet) NPCId() uint32 {
	return npc.Egnet
}

func (r Egnet) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to leave the genie?")
	return SendYesNo(l, c, m.String(), r.Alright, Exit())
}

func (r Egnet) Alright(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r Egnet) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.StationToAriant, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.StationToAriant, 0)
	}
	return Exit()(l, c)
}
