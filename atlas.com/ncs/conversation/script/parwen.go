package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Parwen is located in Hidden Street - Authorized Personnel Only (261020401)
type Parwen struct {
}

func (r Parwen) NPCId() uint32 {
	return npc.Parwen
}

func (r Parwen) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3320) || character.QuestCompleted(l)(c.CharacterId, 3320) {
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.DransLab, 1)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.DransLab, c.NPCId)
		}
		return Exit()(l, c)
	}
	m := message.NewBuilder().AddText("uuuuhuk...Why only Ghost are around here?...")
	return SendOk(l, c, m.String())
}
