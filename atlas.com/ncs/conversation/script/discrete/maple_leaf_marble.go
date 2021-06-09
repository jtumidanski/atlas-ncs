package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

// MapleLeafMarble is located in Orbis - Top of the Hill (200000300)
type MapleLeafMarble struct {
}

func (r MapleLeafMarble) NPCId() uint32 {
	return npc.MapleLeafMarble
}

func (r MapleLeafMarble) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.HasItem(l)(c.CharacterId, item.GlassMarble) {
		character.GainItem(l)(c.CharacterId, item.GlassMarble, -1)
		if character.CanHold(l)(c.CharacterId, item.MapleMarble) {
			character.GainItem(l)(c.CharacterId, item.MapleMarble, 1)
		}
	}
	return script.Exit()(l, c)
}
