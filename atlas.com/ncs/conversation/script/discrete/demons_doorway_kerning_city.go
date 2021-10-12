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

// DemonsDoorwayKerningCity is located in Victoria Road - Kerning City Middle Forest III (103030200)
type DemonsDoorwayKerningCity struct {
}

func (r DemonsDoorwayKerningCity) NPCId() uint32 {
	return npc.DemonsDoorwayKerningCity
}

func (r DemonsDoorwayKerningCity) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 28219) {
		m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
		return script.SendOk(l, span, c, m.String())
	}

	if !character.HasItem(l, span)(c.CharacterId, item.ValeforsEmblem) {
		m := message.NewBuilder().AddText("he entrance is blocked by a force that can only be lifted by those holding an emblem.")
		return script.SendOk(l, span, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Would you like to move to ").
		BlueText().ShowMap(_map.ValeforStrollingPath).
		BlackText().AddText("?")
	return script.SendYesNo(l, span, c, m.String(), r.Process, script.Exit())
}

func (r DemonsDoorwayKerningCity) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.LargeModelOfACoin) {
		character.GainItem(l, span)(c.CharacterId, item.LargeModelOfACoin, -1)
	}
	if character.HasItem(l, span)(c.CharacterId, item.GoldenFeather) {
		character.GainItem(l, span)(c.CharacterId, item.GoldenFeather, -1)
	}
	return script.WarpById(_map.ValeforStrollingPath, 0)(l, span, c)
}