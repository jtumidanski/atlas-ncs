package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// AriantPrivateHouse2 is located in Town of Ariant - Residential Area 2 (260000203)
type AriantPrivateHouse2 struct {
}

func (r AriantPrivateHouse2) NPCId() uint32 {
	return npc.AriantPrivateHouse2
}

func (r AriantPrivateHouse2) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3929) {
		progress := quest.Progress(l)(c.CharacterId, 3929)
		slot := 2

		if progress[slot] == '2' {
			next := progress[0:slot] + string('3') + progress[slot+1:]
			character.GainItem(l)(c.CharacterId, item.WrappedFood, -1)
			quest.SetProgressString(l)(c.CharacterId, 3929, next)
		}
	}
	return script.Exit()(l, c)
}
