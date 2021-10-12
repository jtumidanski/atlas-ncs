package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// SecondMagicPentagram is located in Hidden Street - Black Magician's Lab (261040000)
type SecondMagicPentagram struct {
}

func (r SecondMagicPentagram) NPCId() uint32 {
	return npc.SecondMagicPentagram
}

func (r SecondMagicPentagram) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3345) {
		progress := quest.ProgressInt(l)(c.CharacterId, 3345, 0)

		if progress == 1 {
			quest.SetProgress(l)(c.CharacterId, 3345, 0, 2)
		} else if progress < 4 {
			quest.SetProgress(l)(c.CharacterId, 3345, 0, 0)
		}
	}
	return script.Exit()(l, span, c)
}
