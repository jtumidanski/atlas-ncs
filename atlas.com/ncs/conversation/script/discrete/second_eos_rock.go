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

// SecondEOSRock is located in Ludibrium - Eos Tower 71st Floor (221022900)
type SecondEOSRock struct {
}

func (r SecondEOSRock) NPCId() uint32 {
	return npc.SecondEOSRock
}

func (r SecondEOSRock) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.EOSRockScroll) {
		return r.ToNext(l, span, c)
	}
	return r.NeedRock(l, span, c)
}

func (r SecondEOSRock) ToNext(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(". Which of these rocks would you like to teleport to?").NewLine().
		OpenItem(0).BlueText().AddText("First Eos Rock (100th Floor)").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Third Eos Rock (41st Floor)").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r SecondEOSRock) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Confirm100
	case 1:
		return r.Confirm41
	}
	return nil
}

func (r SecondEOSRock) NeedRock(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("There's a rock that will enable you to teleport to ").
		BlueText().AddText("First Eos Rock or Third Eos Rock").
		BlackText().AddText(", but it cannot be activated without the scroll.")
	return script.SendOk(l, span, c, m.String())

}

func (r SecondEOSRock) Confirm100(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(". Will you teleport to ").
		BlueText().AddText("First Eos Rock").
		BlackText().AddText(" at the 100th Floor?")
	return script.SendYesNo(l, span, c, m.String(), r.Process(_map.EosTower100thFloor), script.Exit())
}

func (r SecondEOSRock) Confirm41(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(". Will you teleport to ").
		BlueText().AddText("Third Eos Rock").
		BlackText().AddText(" at the 41st Floor?")
	return script.SendYesNo(l, span, c, m.String(), r.Process(_map.EosTower41stFloor), script.Exit())
}

func (r SecondEOSRock) Process(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainItem(l, span)(c.CharacterId, item.EOSRockScroll, -1)
		return script.WarpById(mapId, 3)(l, span, c)
	}
}
