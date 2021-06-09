package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// AriantPrivateHouse1 is located in Town of Ariant - Residential Area 1 (260000202)
type AriantPrivateHouse1 struct {
}

func (r AriantPrivateHouse1) NPCId() uint32 {
	return npc.AriantPrivateHouse1
}

func (r AriantPrivateHouse1) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.QuestStarted(l)(c.CharacterId, 3929) {
		progress := character.QuestProgress(l)(c.CharacterId, 3929)
		slot := 0

		if progress[slot] == '2' {
			next := progress[0:slot] + string('3') + progress[slot+1:]
			character.GainItem(l)(c.CharacterId, item.WrappedFood, -1)
			character.SetQuestProgressString(l)(c.CharacterId, 3929, next)
		}
	}
	return script.Exit()(l, c)
}
