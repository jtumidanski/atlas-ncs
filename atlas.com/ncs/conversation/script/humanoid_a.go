package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// HumanoidA is located in Sunset Road - Magatia (261000000)
type HumanoidA struct {
}

func (r HumanoidA) NPCId() uint32 {
	return npc.HumanoidA
}

func (r HumanoidA) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3335) && !character.HasItem(l)(c.CharacterId, item.SnowRose) {
		err := npc.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.WhereSnowRoseGrows, "out00")
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.WhereSnowRoseGrows, c.NPCId)
		}
		return Exit()(l, c)
	}
	m := message.NewBuilder().AddText("Emotion that I feel is real? Or just illusion coming from mechanical error?")
	return SendOk(l, c, m.String())
}
