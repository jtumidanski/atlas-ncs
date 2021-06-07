package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KonpeiExit is located in Zipangu - Near the Hideout (801040000)
type KonpeiExit struct {
}

func (r KonpeiExit) NPCId() uint32 {
	return npc.KonpeiExit
}

func (r KonpeiExit) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Here you are, right in front of the hideout! What? You want to").NewLine().
		AddText("return to ").
		ShowMap(_map.ShowaTown).
		AddText("?")
	return SendYesNo(l, c, m.String(), r.Warp, r.TalkToMe)
}

func (r KonpeiExit) TalkToMe(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you want to return to ").
		ShowMap(_map.ShowaTown).
		AddText(", then talk to me.")
	return SendOk(l, c, m.String())
}

func (r KonpeiExit) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.Warp(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ShowaTown)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ShowaTown, c.NPCId)
	}
	return Exit()(l, c)
}
