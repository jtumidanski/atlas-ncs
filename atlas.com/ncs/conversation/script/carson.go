package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Carson is located in Magatia - Zenumist Society (261000010)
type Carson struct {
}

func (r Carson) NPCId() uint32 {
	return npc.Carson
}

func (r Carson) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3310) && !character.HasItem(l)(c.CharacterId, item.LightlessMagicDevice) {
		err := npc.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ClosedLab, "out00")
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ClosedLab, c.NPCId)
		}
		return Exit()(l, c)
	}

	m := message.NewBuilder().
		AddText("Alchemy....and Alchemist.....both of them are important. But more importantly, it is the Magatia that tolerate everything. The honor of Magatia should be protected by me.")
	return SendOk(l, c, m.String())
}
