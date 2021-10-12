package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CamelCab struct {
}

func (r CamelCab) NPCId() uint32 {
	return npc.CamelCab
}

func (r CamelCab) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.OutsideNorthEntranceOfAriant {
		return r.ToMagatia(l, span, c)
	} else {
		return r.ToAriant(l, span, c)
	}
}

func (r CamelCab) ToAriant(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Would you like to take the ").
		BlueText().AddText("Camel Cab").
		BlackText().AddText(" to ").
		BlueText().AddText("Ariant").
		BlackText().AddText(", the town of Burning Roads? The fare is ").
		BlueText().AddText("1500 mesos").
		BlackText().AddText(".")
	return script.SendYesNo(l, span, c, m.String(), r.Validate(_map.Ariant), r.TooBusy)
}

func (r CamelCab) ToMagatia(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Would you like to take the ").
		BlueText().AddText("Camel Cab").
		BlackText().AddText(" to ").
		BlueText().AddText("Magatia").
		BlackText().AddText(", the town of Alchemy? The fare is ").
		BlueText().AddText("1500 mesos").
		BlackText().AddText(".")
	return script.SendYesNo(l, span, c, m.String(), r.Validate(_map.Magatia), r.TooBusy)
}

func (r CamelCab) Validate(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, 1500) {
			return r.ShortOnMesos(l, span, c)
		}
		return r.Process(mapId)(l, span, c)
	}
}

func (r CamelCab) ShortOnMesos(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I am sorry, but I think you are short on mesos. I am afraid I can't let you ride this if you do not have enough money to do so. Please come back when you have enough money to use this.")
	return script.SendOk(l, span, c, m.String())
}

func (r CamelCab) Process(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainMeso(l, span)(c.CharacterId, -1500)
		return script.WarpById(mapId, 0)(l, span, c)
	}
}

func (r CamelCab) TooBusy(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hmmm... too busy to do it right now? If you feel like doing it, though, come back and find me.")
	return script.SendOk(l, span, c, m.String())
}
