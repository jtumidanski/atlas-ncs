package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// AncientIcyStone is located in Hidden Street - Ice Valley (921100100)
type AncientIcyStone struct {
}

func (r AncientIcyStone) NPCId() uint32 {
	return npc.AncientIcyStone
}

func (r AncientIcyStone) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.OrihalconHammer) {
		return Exit()(l, c)
	}

	if !character.CanHold(l)(c.CharacterId, item.AncientIcePowder) {
		return Exit()(l, c)
	}

	character.GainItem(l)(c.CharacterId, item.AncientIcePowder, 1)
	character.GainItem(l)(c.CharacterId, item.OrihalconHammer, -1)
	return Exit()(l, c)
}
