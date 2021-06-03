package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// AriantPrivateHouse6 is located in Town of Ariant - Residential Area 6 (260000207)
type AriantPrivateHouse6 struct {
}

func (r AriantPrivateHouse6) NPCId() uint32 {
	return npc.AriantPrivateHouse6
}

func (r AriantPrivateHouse6) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3929) {
		progress := character.QuestProgress(l)(c.CharacterId, 3929)
		slot := 3

		if progress[slot] == '2' {
			next := progress[0:slot] + string('3') + progress[slot+1:]
			character.GainItem(l)(c.CharacterId, item.WrappedFood, -1)
			character.SetQuestProgressString(l)(c.CharacterId, 3929, next)
		}
	}
	return Exit()(l, c)
}
