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

// APileOfWhiteFlowers is located in 
type APileOfWhiteFlowers struct {
}

func (r APileOfWhiteFlowers) NPCId() uint32 {
	return npc.APileOfWhiteFlowers
}

func (r APileOfWhiteFlowers) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 2054) && !character.HasItems(l, span)(c.CharacterId, item.WhiteViola, 30) {
		return r.QuestReward(l, span, c)
	}
	return r.RandomReward(l, span, c)
}

func (r APileOfWhiteFlowers) QuestReward(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHoldAll(l)(c.CharacterId, item.WhiteViola, 30) {
		return r.NoSpace(l, span, c)
	}
	return r.AwardViola(l, span, c)
}

func (r APileOfWhiteFlowers) NoSpace(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check for a available slot on your ETC inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r APileOfWhiteFlowers) AwardViola(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.WhiteViola, 30)
	return script.WarpById(_map.Sleepywood, 0)(l, span, c)
}

func (r APileOfWhiteFlowers) RandomReward(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	prizes := []uint32{item.GoldOre, item.LidiumOre, item.DiamondOre}
	character.GainItem(l, span)(c.CharacterId, prizes[rand.Intn(len(prizes))], 4)
	return script.Exit()(l, span, c)
}
