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

// AriantPrivateHouse1 is located in Town of Ariant - Residential Area 1 (260000202)
type AriantPrivateHouse1 struct {
}

func (r AriantPrivateHouse1) NPCId() uint32 {
	return npc.AriantPrivateHouse1
}

func (r AriantPrivateHouse1) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3929) {
		progress := quest.Progress(l)(c.CharacterId, 3929)
		slot := 0

		if progress[slot] == '2' {
			next := progress[0:slot] + string('3') + progress[slot+1:]
			character.GainItem(l, span)(c.CharacterId, item.WrappedFood, -1)
			quest.SetProgressString(l)(c.CharacterId, 3929, next)
		}
	}
	return script.Exit()(l, span, c)
}
