package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// WolfGuard is located in Snow Island - Dangerous Forest (140010200)
type WolfGuard struct {
}

func (r WolfGuard) NPCId() uint32 {
	return npc.WolfGuard
}

func (r WolfGuard) Initial(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.Werewolf) {
		return r.Warp(l, c)
	}
	return r.GetLost(l, c)
}

func (r WolfGuard) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.FieldOfWolves, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.FieldOfWolves, c.NPCId)
	}
	return Exit()(l, c)
}

func (r WolfGuard) GetLost(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("What is it? If you you're here to waste my time, get lost!")
	return SendOk(l, c, m.String())
}