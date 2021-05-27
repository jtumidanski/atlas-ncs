package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// TreasureChestB1 is located in Line 3 Construction Site - B1 <Subway Depot> (103000902)
type TreasureChestB1 struct {
}

func (r TreasureChestB1) NPCId() uint32 {
	return npc.TreasureChestB1
}

func (r TreasureChestB1) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 2055) {
		character.GainItem(l)(c.CharacterId, item.ShumisCoin, 1)
	} else {
		prizes := []uint32{item.GarnetOre, item.AmethystOre, item.AquaMarineOre, item.EmeraldOre, item.OpalOre}
		prize := prizes[rand.Intn(len(prizes))]
		character.GainItem(l)(c.CharacterId, prize, 1)
	}
	err := npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, _map.SubwayTicketingBooth, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.SubwayTicketingBooth, c.NPCId)
	}
	return nil
}
