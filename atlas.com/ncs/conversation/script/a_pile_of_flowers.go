package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
	"math"
	"math/rand"
)

// APileOfFlowers is located in Hidden Street - The Forest of Patience <Step 2> (101000101)
type APileOfFlowers struct {
}

func (r APileOfFlowers) NPCId() uint32 {
	return npc.APileOfFlowers
}

func (r APileOfFlowers) Initial(l logrus.FieldLogger, c Context) State {
	return r.AwardPrize(l, c)
}

type FlowerPrizes struct {
	prizes []FlowerPrize
}

type FlowerPrize struct {
	itemId uint32
	chance uint32
}

func (r APileOfFlowers) AwardPrize(l logrus.FieldLogger, c Context) State {
	prizes := r.GetPrizes()
	gender := character.GetGender(l)(c.CharacterId)
	prizes = r.FilterOutGenderedItems(gender, prizes)

	odds := uint32(0)
	for _, p := range prizes.prizes {
		odds += p.chance
	}
	random := int32(math.Floor((rand.Float64() * float64(odds)) + 1))
	var pick int
	for i, p := range prizes.prizes {
		random -= int32(p.chance)
		if random <= 0 {
			pick = i
			break
		}
	}
	if character.QuestStarted(l)(c.CharacterId, 2050) {
		character.GainItem(l)(c.CharacterId, item.PinkAnthurium, 1)
	}
	character.GainItem(l)(c.CharacterId, prizes.prizes[pick].itemId, 1)
	err := npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, _map.Ellinia, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.Ellinia, c.NPCId)
	}
	return nil
}

func (r APileOfFlowers) GetPrizes() FlowerPrizes {
	return FlowerPrizes{prizes: []FlowerPrize{
		{itemId: item.GreenStripedTrainer, chance: 10},
		{itemId: item.GreenDiscoShirt, chance: 10},
		{itemId: item.GreenTieCasualSuit, chance: 10},
		{itemId: item.GreenTieJacket, chance: 15},
		{itemId: item.GreenSnowboardTop, chance: 10},
		{itemId: item.CamouflagedUniform, chance: 10},
		{itemId: item.GreenDoubleCoat, chance: 10},
		{itemId: item.ArmyGeneralHoodie, chance: 10},
		{itemId: item.CamoHoodedJacket, chance: 10},
		{itemId: item.GreenBaseballJacket, chance: 10},
		{itemId: item.WoodenSlingshot, chance: 5},
		{itemId: item.BugNet, chance: 5},
	}}
}

func (r APileOfFlowers) FilterOutGenderedItems(gender byte, prizes FlowerPrizes) FlowerPrizes {
	result := FlowerPrizes{}
	for _, p := range prizes.prizes {
		itemGender := byte(math.Floor(float64((p.itemId / 100) % 10)))
		if gender == itemGender || itemGender == 2 {
			result.prizes = append(result.prizes, p)
		}
	}
	return result
}
