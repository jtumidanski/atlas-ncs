package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// ThirdMagicPentagram is located in Hidden Street - Black Magician's Lab (261040000)
type ThirdMagicPentagram struct {
}

func (r ThirdMagicPentagram) NPCId() uint32 {
	return npc.ThirdMagicPentagram
}

func (r ThirdMagicPentagram) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3345) {
		progress := quest.ProgressInt(l)(c.CharacterId, 3345, 0)

		if progress == 2 {
			quest.SetProgress(l)(c.CharacterId, 3345, 0, 3)
		} else if progress < 4 {
			quest.SetProgress(l)(c.CharacterId, 3345, 0, 0)
		}
	}
	return script.Exit()(l, c)
}
