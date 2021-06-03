package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// AriantPrivateHouse2Cupboard is located in Town of Ariant - Residential Area 2 (260000203)
type AriantPrivateHouse2Cupboard struct {
}

func (r AriantPrivateHouse2Cupboard) NPCId() uint32 {
	return npc.AriantPrivateHouse2Cupboard
}

func (r AriantPrivateHouse2Cupboard) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3926) {
		progress := character.QuestProgress(l)(c.CharacterId, 3926)
		slot := 2

		if progress[slot] == '2' {
			next := progress[0:slot] + string('3') + progress[slot+1:]
			character.GainItem(l)(c.CharacterId, item.SmallSackOfJewelry, -1)
			character.SetQuestProgressString(l)(c.CharacterId, 3926, next)
		}
	}
	return Exit()(l, c)
}
