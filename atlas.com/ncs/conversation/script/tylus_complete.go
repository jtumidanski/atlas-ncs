package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// TylusComplete is located in Hidden Street - Protecting Tylus : Complete (921100301)
type TylusComplete struct {
}

func (r TylusComplete) NPCId() uint32 {
	return npc.TylusComplete
}

func (r TylusComplete) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You did a great job back there, ").
		ShowCharacterName().
		AddText(", well done. Now I will transport you back to El Nath. Have the pendant in your possession and talk to me when you feel ready to receive the new skill.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r TylusComplete) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ElNath, "in01")
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ElNath, c.NPCId)
	}
	return Exit()(l, c)
}
