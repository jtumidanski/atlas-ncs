package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// TreasureChestB2 is located in Line 3 Construction Site - B2 <Subway Depot> (103000905)
type TreasureChestB2 struct {
}

func (r TreasureChestB2) NPCId() uint32 {
	return npc.TreasureChestB2
}

func (r TreasureChestB2) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 2056) {
		character.GainItem(l)(c.CharacterId, item.ShumisRollOfCash, 1)
	} else {
		prizes := []uint32{item.SapphireOre, item.TopazOre, item.DiamondOre, item.BlackCrystalOre, item.BronzeOre}
		prize := prizes[rand.Intn(len(prizes))]
		character.GainItem(l)(c.CharacterId, prize, 1)
	}
	return script.WarpById(_map.SubwayTicketingBooth, 0)(l, c)
}
