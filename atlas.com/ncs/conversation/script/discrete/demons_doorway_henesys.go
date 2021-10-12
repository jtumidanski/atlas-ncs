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

// DemonsDoorwayHenesys is located in Victoria Road - Henesys Hunting Ground III (104040002)
type DemonsDoorwayHenesys struct {
}

func (r DemonsDoorwayHenesys) NPCId() uint32 {
	return npc.DemonsDoorwayHenesys
}

func (r DemonsDoorwayHenesys) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 28238) {
		m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
		return script.SendOk(l, span, c, m.String())
	}

	if !character.HasItem(l, span)(c.CharacterId, item.AmdusiasEmblem) {
		m := message.NewBuilder().AddText("he entrance is blocked by a force that can only be lifted by those holding an emblem.")
		return script.SendOk(l, span, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Would you like to move to ").
		BlueText().ShowMap(_map.AmdusiasStrollingPath).
		BlackText().AddText("?")
	return script.SendYesNo(l, span, c, m.String(), r.Process, script.Exit())
}

func (r DemonsDoorwayHenesys) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.UnseenlyOcarina) {
		character.GainItem(l, span)(c.CharacterId, item.UnseenlyOcarina, -1)
	}
	if character.HasItem(l, span)(c.CharacterId, item.DarkweensMonsterDrum) {
		character.GainItem(l, span)(c.CharacterId, item.DarkweensMonsterDrum, -1)
	}
	if character.HasItem(l, span)(c.CharacterId, item.SolomonsSealedBow) {
		character.GainItem(l, span)(c.CharacterId, item.SolomonsSealedBow, -1)
	}
	return script.WarpById(_map.AmdusiasStrollingPath, 0)(l, span, c)
}
