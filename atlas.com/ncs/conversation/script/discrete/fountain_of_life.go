package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// FountainOfLife is located in Cave of Life - Entrance to Horntail's Cave (240050400)
type FountainOfLife struct {
}

func (r FountainOfLife) NPCId() uint32 {
	return npc.FountainOfLife
}

func (r FountainOfLife) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 6280) {
		if character.HasItem(l, span)(c.CharacterId, item.HolyCup) {
			character.GainItem(l, span)(c.CharacterId, item.HolyCup, -1)
			character.GainItem(l, span)(c.CharacterId, item.HolyWaterOfLife, 1)
			m := message.NewBuilder().AddText("(You poured some water from the fountain into the cup.)")
			return script.SendOk(l, span, c, m.String())
		}
	}
	return script.Exit()(l, span, c)
}
