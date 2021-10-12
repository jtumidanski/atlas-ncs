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

// FourthEOSRock is located in Ludibrium - Eos Tower 1st Floor (221020000)
type FourthEOSRock struct {
}

func (r FourthEOSRock) NPCId() uint32 {
	return npc.FourthEOSRock
}

func (r FourthEOSRock) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.EOSRockScroll) {
		return r.ToNext(l, span, c)
	}
	return r.NeedRock(l, span, c)
}

func (r FourthEOSRock) ToNext(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Fourth Eos Rock").
		BlackText().AddText(". Will you head over to ").
		BlueText().AddText("Third Eos Rock").
		BlackText().AddText(" at the 41st floor?")
	return script.SendYesNo(l, span, c, m.String(), r.Process, script.Exit())
}

func (r FourthEOSRock) NeedRock(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("There's a rock that will enable you to teleport to ").
		BlueText().AddText("Third Eos Rock").
		BlackText().AddText(", but it cannot be activated without the scroll.")
	return script.SendOk(l, span, c, m.String())
}

func (r FourthEOSRock) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.EOSRockScroll, -1)
	return script.WarpById(_map.EosTower41stFloor, 3)(l, span, c)
}
