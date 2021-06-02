package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// Meteorite1 is located in Omega Sector - Kulan Field I (221040000)
type Meteorite1 struct {
}

func (r Meteorite1) NPCId() uint32 {
	return npc.Meteorite1
}

func (r Meteorite1) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3421) {
		id := c.NPCId - 2050014

		progress := character.QuestProgressInt(l)(c.CharacterId, 3421, 1)
		if (progress>>id)%2 == 0 || (progress == 63 && !character.HasItems(l)(c.CharacterId, item.MeteoriteSample, 6)) {
			if character.CanHold(l)(c.CharacterId, item.MeteoriteSample) {
				progress |= 1 << id
				character.GainItem(l)(c.CharacterId, item.MeteoriteSample, 1)
				character.SetQuestProgress(l)(c.CharacterId, 3421, 1, uint32(progress))
			} else {
				character.SendNotice(l)(c.CharacterId, "POP_UP", "Have a ETC slot available for this item.")
			}
		}
	}
	return Exit()(l, c)
}
