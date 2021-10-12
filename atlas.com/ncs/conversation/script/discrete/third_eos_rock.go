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

// ThirdEOSRock is located in Ludibrium - Eos Tower 41st Floor (221021700)
type ThirdEOSRock struct {
}

func (r ThirdEOSRock) NPCId() uint32 {
	return npc.ThirdEOSRock
}

func (r ThirdEOSRock) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.EOSRockScroll) {
		return r.ToNext(l, span, c)
	}
	return r.NeedRock(l, span, c)
}

func (r ThirdEOSRock) ToNext(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Third Eos Rock").
		BlackText().AddText(". Which of these rocks would you like to teleport to?").NewLine().
		OpenItem(0).BlueText().AddText("Second Eos Rock (71st Floor)").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Fourth Eos Rock (1st Floor)").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r ThirdEOSRock) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Confirm71
	case 1:
		return r.Confirm1
	}
	return nil
}

func (r ThirdEOSRock) NeedRock(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("There's a rock that will enable you to teleport to ").
		BlueText().AddText("Second Eos Rock or Fourth Eos Rock").
		BlackText().AddText(", but it cannot be activated without the scroll.")
	return script.SendOk(l, span, c, m.String())

}

func (r ThirdEOSRock) Confirm71(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Third Eos Rock").
		BlackText().AddText(". Will you teleport to ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(" at the 71st Floor?")
	return script.SendYesNo(l, span, c, m.String(), r.Process(_map.EosTower71stFloor), script.Exit())
}

func (r ThirdEOSRock) Confirm1(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Third Eos Rock").
		BlackText().AddText(". Will you teleport to ").
		BlueText().AddText("Fourth Eos Rock").
		BlackText().AddText(" at the 1st Floor?")
	return script.SendYesNo(l, span, c, m.String(), r.Process(_map.EosTower1stFloor), script.Exit())
}

func (r ThirdEOSRock) Process(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainItem(l, span)(c.CharacterId, item.EOSRockScroll, -1)
		return script.WarpById(mapId, 3)(l, span, c)
	}
}
