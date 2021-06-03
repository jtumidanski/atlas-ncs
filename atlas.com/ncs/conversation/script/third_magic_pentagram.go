package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// ThirdMagicPentagram is located in Hidden Street - Black Magician's Lab (261040000)
type ThirdMagicPentagram struct {
}

func (r ThirdMagicPentagram) NPCId() uint32 {
	return npc.ThirdMagicPentagram
}

func (r ThirdMagicPentagram) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3345) {
		progress := character.QuestProgressInt(l)(c.CharacterId, 3345, 0)

		if progress == 2 {
			character.SetQuestProgress(l)(c.CharacterId, 3345, 0, 3)
		} else if progress < 4 {
			character.SetQuestProgress(l)(c.CharacterId, 3345, 0, 0)
		}
	}
	return Exit()(l, c)
}
