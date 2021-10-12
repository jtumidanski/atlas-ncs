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

// AriantPrivateHouse2Cupboard is located in Town of Ariant - Residential Area 2 (260000203)
type AriantPrivateHouse2Cupboard struct {
}

func (r AriantPrivateHouse2Cupboard) NPCId() uint32 {
	return npc.AriantPrivateHouse2Cupboard
}

func (r AriantPrivateHouse2Cupboard) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3926) {
		progress := quest.Progress(l)(c.CharacterId, 3926)
		slot := 2

		if progress[slot] == '2' {
			next := progress[0:slot] + string('3') + progress[slot+1:]
			character.GainItem(l, span)(c.CharacterId, item.SmallSackOfJewelry, -1)
			quest.SetProgressString(l)(c.CharacterId, 3926, next)
		}
	}
	return script.Exit()(l, span, c)
}
