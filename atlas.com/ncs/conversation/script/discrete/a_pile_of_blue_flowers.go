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

// APileOfBlueFlowers is located in Hidden Street - The Deep Forest of Patience <Step 4> (105040313)
type APileOfBlueFlowers struct {
}

func (r APileOfBlueFlowers) NPCId() uint32 {
	return npc.APileOfBlueFlowers
}

func (r APileOfBlueFlowers) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 2053) && !character.HasItems(l)(c.CharacterId, item.BlueViola, 20) {
		return r.QuestReward(l, span, c)
	}
	return r.RandomReward(l, span, c)
}

func (r APileOfBlueFlowers) QuestReward(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHoldAll(l)(c.CharacterId, item.BlueViola, 20) {
		return r.NoSpace(l, span, c)
	}
	return r.AwardViola(l, span, c)
}

func (r APileOfBlueFlowers) NoSpace(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check for a available slot on your ETC inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r APileOfBlueFlowers) AwardViola(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.BlueViola, 20)
	return script.WarpById(_map.Sleepywood, 0)(l, span, c)
}

func (r APileOfBlueFlowers) RandomReward(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	prizes := []uint32{item.GarnetOre, item.AquaMarineOre, item.TopazOre}
	character.GainItem(l)(c.CharacterId, prizes[rand.Intn(len(prizes))], 4)
	return script.Exit()(l, span, c)
}
