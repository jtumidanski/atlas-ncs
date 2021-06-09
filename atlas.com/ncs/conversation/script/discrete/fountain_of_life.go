package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// FountainOfLife is located in Cave of Life - Entrance to Horntail's Cave (240050400)
type FountainOfLife struct {
}

func (r FountainOfLife) NPCId() uint32 {
	return npc.FountainOfLife
}

func (r FountainOfLife) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.QuestStarted(l)(c.CharacterId, 6280) {
		if character.HasItem(l)(c.CharacterId, item.HolyCup) {
			character.GainItem(l)(c.CharacterId, item.HolyCup, -1)
			character.GainItem(l)(c.CharacterId, item.HolyWaterOfLife, 1)
			m := message.NewBuilder().AddText("(You poured some water from the fountain into the cup.)")
			return script.SendOk(l, c, m.String())
		}
	}
	return script.Exit()(l, c)
}
