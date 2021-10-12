package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// TreasureChestB3 is located in Line 3 Construction Site - B3 <Subway Depot> (103000909)
type TreasureChestB3 struct {
}

func (r TreasureChestB3) NPCId() uint32 {
	return npc.TreasureChestB3
}

func (r TreasureChestB3) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 2057) {
		character.GainItem(l, span)(c.CharacterId, item.ShumisSackOfCash, 1)
	} else {
		prizes := []uint32{item.SteelOre, item.MithrilOre, item.AdamantiumOre, item.SilverOre, item.OrihalconOre, item.GoldOre, item.LidiumOre}
		prize := prizes[rand.Intn(len(prizes))]
		character.GainItem(l, span)(c.CharacterId, prize, 1)
	}
	return script.WarpById(_map.SubwayTicketingBooth, 0)(l, span, c)
}
