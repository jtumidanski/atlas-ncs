package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// Carson is located in Magatia - Zenumist Society (261000010)
type Carson struct {
}

func (r Carson) NPCId() uint32 {
	return npc.Carson
}

func (r Carson) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3310) && !character.HasItem(l)(c.CharacterId, item.LightlessMagicDevice) {
		return script.WarpByName( _map.ClosedLab, "out00")(l, c)
	}

	m := message.NewBuilder().
		AddText("Alchemy....and Alchemist.....both of them are important. But more importantly, it is the Magatia that tolerate everything. The honor of Magatia should be protected by me.")
	return script.SendOk(l, c, m.String())
}
