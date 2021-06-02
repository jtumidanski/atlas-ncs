package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Rupi is located in Hidden Street - Happyville (209000000)
type Rupi struct {
}

func (r Rupi) NPCId() uint32 {
	return npc.Rupi
}

func (r Rupi) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you want to get out of Happyville?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r Rupi) Warp(l logrus.FieldLogger, c Context) State {
	mapId := character.SavedLocation(l)(c.CharacterId, "HAPPYVILLE")
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
	}
	return Exit()(l, c)
}
