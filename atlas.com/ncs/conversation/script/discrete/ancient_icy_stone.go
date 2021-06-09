package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
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

func (r AncientIcyStone) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.HasItem(l)(c.CharacterId, item.OrihalconHammer) {
		return script.Exit()(l, c)
	}

	if !character.CanHold(l)(c.CharacterId, item.AncientIcePowder) {
		return script.Exit()(l, c)
	}

	character.GainItem(l)(c.CharacterId, item.AncientIcePowder, 1)
	character.GainItem(l)(c.CharacterId, item.OrihalconHammer, -1)
	return script.Exit()(l, c)
}
