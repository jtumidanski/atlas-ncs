package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// AriantPrivateHouse6Cupboard is located in Town of Ariant - Residential Area 6 (260000207)
type AriantPrivateHouse6Cupboard struct {
}

func (r AriantPrivateHouse6Cupboard) NPCId() uint32 {
	return npc.AriantPrivateHouse6Cupboard
}

func (r AriantPrivateHouse6Cupboard) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3926) {
		progress := quest.Progress(l)(c.CharacterId, 3926)
		slot := 3

		if progress[slot] == '2' {
			next := progress[0:slot] + string('3') + progress[slot+1:]
			character.GainItem(l)(c.CharacterId, item.SmallSackOfJewelry, -1)
			quest.SetProgressString(l)(c.CharacterId, 3926, next)
		}
	}
	return script.Exit()(l, c)
}
