package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// NLCTaxi is located in New Leaf City Town Street - New Leaf City - Town Center (600000000)
type NLCTaxi struct {
}

func (r NLCTaxi) NPCId() uint32 {
	return npc.NLCTaxi
}

func (r NLCTaxi) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.HauntedHouse {
		m := message.NewBuilder().
			AddText("Would you like to return back to ").
			BlueText().AddText("civilization").
			BlackText().AddText(fmt.Sprintf("? The fee is %d mesos.", 15000))
		return script.SendYesNo(l, span, c, m.String(), r.ToNLC, r.NextTime)
	} else {
		m := message.NewBuilder().
			AddText("Would you like to go to the ").
			BlueText().AddText("Haunted Mansion").
			BlackText().AddText(fmt.Sprintf("? The fee is %d mesos.", 15000))
		return script.SendYesNo(l, span, c, m.String(), r.ToHauntedHouse, r.NextTime)
	}
}

func (r NLCTaxi) NextTime(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Alright, see you next time.")
	return script.SendOk(l, span, c, m.String())
}

func (r NLCTaxi) WarpToNLC(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.NewLeafCityTownCenter, 0)(l, span, c)
}

func (r NLCTaxi) ToNLC(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Validate(r.WarpToNLC)(l, span, c)
}

func (r NLCTaxi) ToHauntedHouse(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Validate(r.WarpToHauntedHouse)(l, span, c)
}

func (r NLCTaxi) WarpToHauntedHouse(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.HauntedHouse, 0)(l, span, c)
}

func (r NLCTaxi) Validate(warp func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, 15000) {
			return r.NotEnoughMesos(l, span, c)
		}
		character.GainMeso(l, span)(c.CharacterId, -15000)
		return warp(l, span, c)
	}
}

func (r NLCTaxi) NotEnoughMesos(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hey, what are you trying to pull on? You don't have enough meso to pay the fee.")
	return script.SendOk(l, span, c, m.String())
}
