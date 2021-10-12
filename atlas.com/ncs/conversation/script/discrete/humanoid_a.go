package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// HumanoidA is located in Sunset Road - Magatia (261000000)
type HumanoidA struct {
}

func (r HumanoidA) NPCId() uint32 {
	return npc.HumanoidA
}

func (r HumanoidA) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3335) && !character.HasItem(l, span)(c.CharacterId, item.SnowRose) {
		return script.WarpByName(_map.WhereSnowRoseGrows, "out00")(l, span, c)
	}
	m := message.NewBuilder().AddText("Emotion that I feel is real? Or just illusion coming from mechanical error?")
	return script.SendOk(l, span, c, m.String())
}
