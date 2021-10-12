package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Meteorite4 is located in Omega Sector - Kulan Field II (221040100)
type Meteorite4 struct {
}

func (r Meteorite4) NPCId() uint32 {
	return npc.Meteorite4
}

func (r Meteorite4) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3421) {
		id := c.NPCId - 2050014

		progress := quest.ProgressInt(l)(c.CharacterId, 3421, 1)
		if (progress>>id)%2 == 0 || (progress == 63 && !character.HasItems(l, span)(c.CharacterId, item.MeteoriteSample, 6)) {
			if character.CanHold(l)(c.CharacterId, item.MeteoriteSample) {
				progress |= 1 << id
				character.GainItem(l, span)(c.CharacterId, item.MeteoriteSample, 1)
				quest.SetProgress(l)(c.CharacterId, 3421, 1, uint32(progress))
			} else {
				character.SendNotice(l)(c.CharacterId, "POP_UP", "Have a ETC slot available for this item.")
			}
		}
	}
	return script.Exit()(l, span, c)
}
