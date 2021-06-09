package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// SecondMagicPentagram is located in Hidden Street - Black Magician's Lab (261040000)
type SecondMagicPentagram struct {
}

func (r SecondMagicPentagram) NPCId() uint32 {
	return npc.SecondMagicPentagram
}

func (r SecondMagicPentagram) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.QuestStarted(l)(c.CharacterId, 3345) {
		progress := character.QuestProgressInt(l)(c.CharacterId, 3345, 0)

		if progress == 1 {
			character.SetQuestProgress(l)(c.CharacterId, 3345, 0, 2)
		} else if progress < 4 {
			character.SetQuestProgress(l)(c.CharacterId, 3345, 0, 0)
		}
	}
	return script.Exit()(l, c)
}
