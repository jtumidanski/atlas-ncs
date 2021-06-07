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
		return WarpByName(_map.WhereSnowRoseGrows, "out00")(l, c)
	}
	m := message.NewBuilder().AddText("Emotion that I feel is real? Or just illusion coming from mechanical error?")
	return SendOk(l, c, m.String())
}
