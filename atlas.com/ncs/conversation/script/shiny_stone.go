package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ShinyStone is located in The Nautilus - Generator Room (120000301)
type ShinyStone struct {
}

func (r ShinyStone) NPCId() uint32 {
	return npc.ShinyStone
}

func (r ShinyStone) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 2166) {
		return r.BeautifulRock(l, c)
	}
	return r.MysteriousPower(l, c)
}

func (r ShinyStone) BeautifulRock(l logrus.FieldLogger, c Context) State {
	character.CompleteQuest(l)(c.CharacterId, 2166)
	m := message.NewBuilder().
		AddText("It's a beautiful, shiny rock. I can feel the mysterious power surrounding it.")
	return SendNext(l, c, m.String(), Exit())
}

func (r ShinyStone) MysteriousPower(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I touched the shiny rock with my hand, and I felt a mysterious power flowing into my body.")
	return SendNext(l, c, m.String(), Exit())
}
