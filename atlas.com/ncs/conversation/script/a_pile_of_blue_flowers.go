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

// APileOfBlueFlowers is located in Hidden Street - The Deep Forest of Patience <Step 4> (105040313)
type APileOfBlueFlowers struct {
}

func (r APileOfBlueFlowers) NPCId() uint32 {
	return npc.APileOfBlueFlowers
}

func (r APileOfBlueFlowers) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 2053) && !character.HasItems(l)(c.CharacterId, item.BlueViola, 20) {
		return r.QuestReward(l, c)
	}
	return r.RandomReward(l, c)
}

func (r APileOfBlueFlowers) QuestReward(l logrus.FieldLogger, c Context) State {
	if !character.CanHoldAll(l)(c.CharacterId, item.BlueViola, 20) {
		return r.NoSpace(l, c)
	}
	return r.AwardViola(l, c)
}

func (r APileOfBlueFlowers) NoSpace(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Check for a available slot on your ETC inventory.")
	return SendOk(l, c, m.String())
}

func (r APileOfBlueFlowers) AwardViola(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.BlueViola, 20)
	return r.Warp(l, c)
}

func (r APileOfBlueFlowers) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, _map.Sleepywood, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.Sleepywood, c.NPCId)
	}
	return nil
}

func (r APileOfBlueFlowers) RandomReward(l logrus.FieldLogger, c Context) State {
	prizes := []uint32{item.GarnetOre, item.AquaMarineOre, item.TopazOre}
	character.GainItem(l)(c.CharacterId, prizes[rand.Intn(len(prizes))], 4)
	return Exit()(l, c)
}
