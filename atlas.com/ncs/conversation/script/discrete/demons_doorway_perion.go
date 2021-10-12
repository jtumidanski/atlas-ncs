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

// DemonsDoorwayPerion is located in Victoria Road - West Rocky Mountain IV (102020300)
type DemonsDoorwayPerion struct {
}

func (r DemonsDoorwayPerion) NPCId() uint32 {
	return npc.DemonsDoorwayPerion
}

func (r DemonsDoorwayPerion) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 28179) {
		m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
		return script.SendOk(l, span, c, m.String())
	}

	if !character.HasItem(l, span)(c.CharacterId, item.AndrasEmblem) {
		m := message.NewBuilder().AddText("he entrance is blocked by a force that can only be lifted by those holding an emblem.")
		return script.SendOk(l, span, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Would you like to move to ").
		BlueText().ShowMap(_map.AndrasStrollingPath).
		BlackText().AddText("?")
	return script.SendYesNo(l, span, c, m.String(), r.Process, script.Exit())
}

func (r DemonsDoorwayPerion) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.BrokenIronFragment) {
		character.GainItem(l, span)(c.CharacterId, item.BrokenIronFragment, -1)
	}
	if character.HasItem(l, span)(c.CharacterId, item.OrangeMushroomWine) {
		character.GainItem(l, span)(c.CharacterId, item.OrangeMushroomWine, -1)
	}
	return script.WarpById(_map.AndrasStrollingPath, 0)(l, span, c)
}
