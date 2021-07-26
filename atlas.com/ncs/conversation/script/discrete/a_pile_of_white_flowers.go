package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// APileOfWhiteFlowers is located in 
type APileOfWhiteFlowers struct {
}

func (r APileOfWhiteFlowers) NPCId() uint32 {
	return npc.APileOfWhiteFlowers
}

func (r APileOfWhiteFlowers) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 2054) && !character.HasItems(l)(c.CharacterId, item.WhiteViola, 30) {
		return r.QuestReward(l, c)
	}
	return r.RandomReward(l, c)
}

func (r APileOfWhiteFlowers) QuestReward(l logrus.FieldLogger, c script.Context) script.State {
	if !character.CanHoldAll(l)(c.CharacterId, item.WhiteViola, 30) {
		return r.NoSpace(l, c)
	}
	return r.AwardViola(l, c)
}

func (r APileOfWhiteFlowers) NoSpace(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check for a available slot on your ETC inventory.")
	return script.SendOk(l, c, m.String())
}

func (r APileOfWhiteFlowers) AwardViola(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.WhiteViola, 30)
	return script.WarpById(_map.Sleepywood, 0)(l, c)
}

func (r APileOfWhiteFlowers) RandomReward(l logrus.FieldLogger, c script.Context) script.State {
	prizes := []uint32{item.GoldOre, item.LidiumOre, item.DiamondOre}
	character.GainItem(l)(c.CharacterId, prizes[rand.Intn(len(prizes))], 4)
	return script.Exit()(l, c)
}
