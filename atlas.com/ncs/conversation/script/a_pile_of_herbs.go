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

// APileOfHerbs is located in Hidden Street - The Forest of Patience <Step 5> (101000104)
type APileOfHerbs struct {
}

func (r APileOfHerbs) NPCId() uint32 {
	return npc.APileOfHerbs
}

func (r APileOfHerbs) Initial(l logrus.FieldLogger, c Context) State {
	return r.AwardPrize(l, c)
}

type HerbPrizes struct {
	prizes []HerbPrize
}

type HerbPrize struct {
	itemId uint32
	chance uint32
}

func (r APileOfHerbs) AwardPrize(l logrus.FieldLogger, c Context) State {
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
	if character.QuestStarted(l)(c.CharacterId, 2051) {
		character.GainItem(l)(c.CharacterId, item.DoubleRootedRedGinseng, 1)
	}
	character.GainItem(l)(c.CharacterId, prizes.prizes[pick].itemId, 1)
	err := npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, _map.Ellinia, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.Ellinia, c.NPCId)
	}
	return nil
}

func (r APileOfHerbs) GetPrizes() HerbPrizes {
	return HerbPrizes{prizes: []HerbPrize{
		{itemId: item.GreenTrainerPants, chance: 10},
		{itemId: item.GreenDiscoPants, chance: 10},
		{itemId: item.MilitaryCargoShorts, chance: 10},
		{itemId: item.HawaiianSkirt, chance: 10},
		{itemId: item.GreenLongSkirt, chance: 10},
		{itemId: item.MilitaryCargoShorts2, chance: 10},
		{itemId: item.GreenCampingShorts, chance: 10},
		{itemId: item.GreenSnowboardPants, chance: 10},
		{itemId: item.CamouflagedArmyPants, chance: 10},
		{itemId: item.MilitaryCargoPants, chance: 10},
		{itemId: item.OliveSkinnyJeans, chance: 10},
		{itemId: item.SunflowerStalk, chance: 5},
		{itemId: item.WonkysLeaf, chance: 5},
	}}
}

func (r APileOfHerbs) FilterOutGenderedItems(gender byte, prizes HerbPrizes) HerbPrizes {
	result := HerbPrizes{}
	for _, p := range prizes.prizes {
		itemGender := byte(math.Floor(float64((p.itemId / 100) % 10)))
		if gender == itemGender || itemGender == 2 {
			result.prizes = append(result.prizes, p)
		}
	}
	return result
}