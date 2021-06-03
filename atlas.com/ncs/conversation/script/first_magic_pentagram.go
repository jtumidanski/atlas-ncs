package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// FirstMagicPentagram is located in Hidden Street - Black Magician's Lab (261040000)
type FirstMagicPentagram struct {
}

func (r FirstMagicPentagram) NPCId() uint32 {
	return npc.FirstMagicPentagram
}

func (r FirstMagicPentagram) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3345) {
		progress := character.QuestProgressInt(l)(c.CharacterId, 3345, 0)

		if progress == 0 {
			character.SetQuestProgress(l)(c.CharacterId, 3345, 0, 1)
		} else if progress < 4 {
			character.SetQuestProgress(l)(c.CharacterId, 3345, 0, 0)
		}
	}
	return Exit()(l, c)
}
