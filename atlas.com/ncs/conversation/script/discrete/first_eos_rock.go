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

// FirstEOSRock is located in Ludibrium - Eos Tower 100th Floor (221024400)
type FirstEOSRock struct {
}

func (r FirstEOSRock) NPCId() uint32 {
	return npc.FirstEOSRock
}

func (r FirstEOSRock) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.EOSRockScroll) {
		return r.ToNext(l, span, c)
	}
	return r.NeedRock(l, span, c)
}

func (r FirstEOSRock) ToNext(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("First Eos Rock").
		BlackText().AddText(". Will you teleport to ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(" at the 71st floor?")
	return script.SendYesNo(l, span, c, m.String(), r.Process, script.Exit())
}

func (r FirstEOSRock) NeedRock(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("There's a rock that will enable you to teleport to ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(", but it cannot be activated without the scroll.")
	return script.SendOk(l, span, c, m.String())
}

func (r FirstEOSRock) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.EOSRockScroll, -1)
	return script.WarpById(_map.EosTower71stFloor, 3)(l, span, c)
}
