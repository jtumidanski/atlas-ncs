package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// OrbisMagicSpot20 is located in Orbis - Orbis Tower <20th Floor> (200080200)
type OrbisMagicSpot20 struct {
}

func (r OrbisMagicSpot20) NPCId() uint32 {
	return npc.OrbisMagicSpot20
}

func (r OrbisMagicSpot20) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.OrbisRockScroll) {
		return r.NeedScroll(l, span, c)
	}
	return r.Confirm(l, span, c)
}

func (r OrbisMagicSpot20) NeedScroll(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("There's a ").
		BlueText().ShowNPC(npc.OrbisMagicSpot20).
		BlackText().AddText(" that'll enable you to teleport to where ").
		BlueText().ShowNPC(npc.OrbisMagicSpot1).
		BlackText().AddText(" is, but you can't activate it without the scroll.")
	return script.SendOk(l, span, c, m.String())
}

func (r OrbisMagicSpot20) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You can use ").
		BlueText().ShowItemName1(item.OrbisRockScroll).
		BlackText().AddText(" to activate ").
		BlueText().ShowNPC(npc.OrbisMagicSpot20).
		BlackText().AddText(". Will you teleport to where ").
		BlueText().ShowNPC(npc.OrbisMagicSpot1).
		BlackText().AddText(" is?")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, script.Exit())
}

func (r OrbisMagicSpot20) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.OrbisRockScroll, -1)
	return script.WarpById(_map.OrbisTower1stFloor, 0)(l, span, c)
}
