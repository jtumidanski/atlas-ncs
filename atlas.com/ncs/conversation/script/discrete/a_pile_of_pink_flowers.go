package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// APileOfPinkFlowers is located in Hidden Street - The Deep Forest of Patience <Step 2> (105040311)
type APileOfPinkFlowers struct {
}

func (r APileOfPinkFlowers) NPCId() uint32 {
	return npc.APileOfPinkFlowers
}

func (r APileOfPinkFlowers) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 2052) && !character.HasItems(l)(c.CharacterId, item.PinkViola, 10) {
		return r.QuestReward(l, span, c)
	}
	return r.RandomReward(l, span, c)
}

func (r APileOfPinkFlowers) QuestReward(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHoldAll(l)(c.CharacterId, item.PinkViola, 10) {
		return r.NoSpace(l, span, c)
	}
	return r.AwardViola(l, span, c)
}

func (r APileOfPinkFlowers) NoSpace(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check for a available slot on your ETC inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r APileOfPinkFlowers) AwardViola(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.PinkViola, 10)
	return script.WarpById(_map.Sleepywood, 0)(l, span, c)
}

func (r APileOfPinkFlowers) RandomReward(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	prizes := []uint32{item.BronzeOre, item.SteelOre, item.MithrilOre, item.AdamantiumOre, item.SilverOre, item.OrihalconOre}
	character.GainItem(l)(c.CharacterId, prizes[rand.Intn(len(prizes))], 3)
	return script.Exit()(l, span, c)
}
