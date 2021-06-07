package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SnowmanExit is located in 
type SnowmanExit struct {
}

func (r SnowmanExit) NPCId() uint32 {
	return npc.SnowmanExit
}

func (r SnowmanExit) Initial(l logrus.FieldLogger, c Context) State {
	area := c.MapId % 10
	if area > 0 {
		m := message.NewBuilder().
			AddText("Do you wish to leave this place?")
		return SendYesNo(l, c, m.String(), r.WarpById(c.MapId+1, 0), Exit())
	} else {
		m := message.NewBuilder().
			AddText("Do you wish to return to ").
			BlueText().AddText("Happyville").
			BlackText().AddText("?")
		return SendYesNo(l, c, m.String(), r.WarpToHappyville, Exit())
	}
}

func (r SnowmanExit) WarpById(mapId uint32, portalId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, portalId)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}

func (r SnowmanExit) WarpToHappyville(l logrus.FieldLogger, c Context) State {
	err := npc.Warp(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.Happyville)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.Happyville, c.NPCId)
	}
	return Exit()(l, c)
}
