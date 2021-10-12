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

// Dolphin2 is located in Herb Town - Pier on the Beach (251000100)
type Dolphin2 struct {
}

func (r Dolphin2) NPCId() uint32 {
	return npc.Dolphin2
}

func (r Dolphin2) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Will you move to ").
		BlueText().ShowMap(_map.Aquarium).
		BlackText().AddText(" now? The price is ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 1000)).
		BlackText().AddText(".")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, r.TooBusy)
}

func (r Dolphin2) TooBusy(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hmmm ... too busy to do it right now? If you feel like doing it, though, come back and find me.")
	return script.SendOk(l, span, c, m.String())
}

func (r Dolphin2) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 1000) {
		return r.NotEnoughMeso(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r Dolphin2) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -1000)
	return script.WarpById(_map.Aquarium, 0)(l, span, c)
}

func (r Dolphin2) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I don't think you have enough money...")
	return script.SendOk(l, span, c, m.String())
}
