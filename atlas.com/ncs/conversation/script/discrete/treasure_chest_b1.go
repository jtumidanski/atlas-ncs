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

// TreasureChestB1 is located in Line 3 Construction Site - B1 <Subway Depot> (103000902)
type TreasureChestB1 struct {
}

func (r TreasureChestB1) NPCId() uint32 {
	return npc.TreasureChestB1
}

func (r TreasureChestB1) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 2055) {
		character.GainItem(l, span)(c.CharacterId, item.ShumisCoin, 1)
	} else {
		prizes := []uint32{item.GarnetOre, item.AmethystOre, item.AquaMarineOre, item.EmeraldOre, item.OpalOre}
		prize := prizes[rand.Intn(len(prizes))]
		character.GainItem(l, span)(c.CharacterId, prize, 1)
	}
	return script.WarpById(_map.SubwayTicketingBooth, 0)(l, span, c)
}
