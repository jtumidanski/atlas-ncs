package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// ShinyStone is located in The Nautilus - Generator Room (120000301)
type ShinyStone struct {
}

func (r ShinyStone) NPCId() uint32 {
	return npc.ShinyStone
}

func (r ShinyStone) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 2166) {
		return r.BeautifulRock(l, c)
	}
	return r.MysteriousPower(l, c)
}

func (r ShinyStone) BeautifulRock(l logrus.FieldLogger, c script.Context) script.State {
	quest.Complete(l)(c.CharacterId, 2166)
	m := message.NewBuilder().
		AddText("It's a beautiful, shiny rock. I can feel the mysterious power surrounding it.")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r ShinyStone) MysteriousPower(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I touched the shiny rock with my hand, and I felt a mysterious power flowing into my body.")
	return script.SendNext(l, c, m.String(), script.Exit())
}
