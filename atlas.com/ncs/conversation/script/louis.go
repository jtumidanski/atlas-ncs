package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Louis is located in Hidden Street - The Forest of Patience (101000100, 101000101, 101000102, 101000103, and 101000104)
type Louis struct {
}

func (r Louis) NPCId() uint32 {
	return npc.Louis
}

func (r Louis) Initial(l logrus.FieldLogger, c Context) State {
	return r.Return(l, c)
}

func (r Louis) Return(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Would you like to return to Ellinia?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r Louis) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, _map.Ellinia, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.Ellinia, c.NPCId)
	}
	return nil
}
