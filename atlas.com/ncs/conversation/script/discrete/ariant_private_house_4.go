package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// AriantPrivateHouse4 is located in Town of Ariant - Residential Area 4 (260000205)
type AriantPrivateHouse4 struct {
}

func (r AriantPrivateHouse4) NPCId() uint32 {
	return npc.AriantPrivateHouse4
}

func (r AriantPrivateHouse4) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.QuestStarted(l)(c.CharacterId, 3929) {
		progress := character.QuestProgress(l)(c.CharacterId, 3929)
		slot := 1

		if progress[slot] == '2' {
			next := progress[0:slot] + string('3') + progress[slot+1:]
			character.GainItem(l)(c.CharacterId, item.WrappedFood, -1)
			character.SetQuestProgressString(l)(c.CharacterId, 3929, next)
		}
	}
	return script.Exit()(l, c)
}
