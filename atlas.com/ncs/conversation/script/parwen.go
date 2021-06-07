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
		return WarpById(_map.DransLab, 1)(l, c)
	}
	m := message.NewBuilder().AddText("uuuuhuk...Why only Ghost are around here?...")
	return SendOk(l, c, m.String())
}
