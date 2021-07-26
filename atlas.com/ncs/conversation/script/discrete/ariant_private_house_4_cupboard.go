package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// AriantPrivateHouse4Cupboard is located in Town of Ariant - Residential Area 4 (260000205)
type AriantPrivateHouse4Cupboard struct {
}

func (r AriantPrivateHouse4Cupboard) NPCId() uint32 {
	return npc.AriantPrivateHouse4Cupboard
}

func (r AriantPrivateHouse4Cupboard) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3926) {
		progress := quest.Progress(l)(c.CharacterId, 3926)
		slot := 1

		if progress[slot] == '2' {
			next := progress[0:slot] + string('3') + progress[slot+1:]
			character.GainItem(l)(c.CharacterId, item.SmallSackOfJewelry, -1)
			quest.SetProgressString(l)(c.CharacterId, 3926, next)
		}
	}
	return script.Exit()(l, c)
}
