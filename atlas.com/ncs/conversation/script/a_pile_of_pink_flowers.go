package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// APileOfPinkFlowers is located in Hidden Street - The Deep Forest of Patience <Step 2> (105040311)
type APileOfPinkFlowers struct {
}

func (r APileOfPinkFlowers) NPCId() uint32 {
	return npc.APileOfPinkFlowers
}

func (r APileOfPinkFlowers) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 2052) && !character.HasItems(l)(c.CharacterId, item.PinkViola, 10) {
		return r.QuestReward(l, c)
	}
	return r.RandomReward(l, c)
}

func (r APileOfPinkFlowers) QuestReward(l logrus.FieldLogger, c Context) State {
	if !character.CanHoldAll(l)(c.CharacterId, item.PinkViola, 10) {
		return r.NoSpace(l, c)
	}
	return r.AwardViola(l, c)
}

func (r APileOfPinkFlowers) NoSpace(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Check for a available slot on your ETC inventory.")
	return SendOk(l, c, m.String())
}

func (r APileOfPinkFlowers) AwardViola(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.PinkViola, 10)
	return r.Warp(l, c)
}

func (r APileOfPinkFlowers) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.Sleepywood, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.Sleepywood, c.NPCId)
	}
	return nil
}

func (r APileOfPinkFlowers) RandomReward(l logrus.FieldLogger, c Context) State {
	prizes := []uint32{item.BronzeOre, item.SteelOre, item.MithrilOre, item.AdamantiumOre, item.SilverOre, item.OrihalconOre}
	character.GainItem(l)(c.CharacterId, prizes[rand.Intn(len(prizes))], 3)
	return Exit()(l, c)
}
