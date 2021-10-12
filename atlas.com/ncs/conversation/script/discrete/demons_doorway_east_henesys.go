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

// DemonsDoorwayEastHenesys is located in Victoria Road - The Forest East of Henesys (100030000)
type DemonsDoorwayEastHenesys struct {
}

func (r DemonsDoorwayEastHenesys) NPCId() uint32 {
	return npc.DemonsDoorwayEastHenesys
}

func (r DemonsDoorwayEastHenesys) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 28256) {
		m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
		return script.SendOk(l, span, c, m.String())
	}

	if !character.HasItem(l, span)(c.CharacterId, item.CrocellsEmblem) {
		m := message.NewBuilder().AddText("he entrance is blocked by a force that can only be lifted by those holding an emblem.")
		return script.SendOk(l, span, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Would you like to move to ").
		BlueText().ShowMap(_map.CrocellStrollingPath).
		BlackText().AddText("?")
	return script.SendYesNo(l, span, c, m.String(), r.Process, script.Exit())
}

func (r DemonsDoorwayEastHenesys) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.OldWornOutPaper) {
		character.GainItem(l, span)(c.CharacterId, item.OldWornOutPaper, -1)
	}
	if character.HasItem(l, span)(c.CharacterId, item.OldKey) {
		character.GainItem(l, span)(c.CharacterId, item.OldKey, -1)
	}
	if character.HasItem(l, span)(c.CharacterId, item.GreenSlimeEraser) {
		character.GainItem(l, span)(c.CharacterId, item.GreenSlimeEraser, -1)
	}
	if character.HasItem(l, span)(c.CharacterId, item.AyanMercurysMicrophone) {
		character.GainItem(l, span)(c.CharacterId, item.AyanMercurysMicrophone, -1)
	}
	if character.HasItem(l, span)(c.CharacterId, item.InkSack) {
		character.GainItem(l, span)(c.CharacterId, item.InkSack, -1)
	}
	if character.HasItem(l, span)(c.CharacterId, item.DirtyTreasureMap) {
		character.GainItem(l, span)(c.CharacterId, item.DirtyTreasureMap, -1)
	}
	return script.WarpById(_map.CrocellStrollingPath, 0)(l, span, c)
}