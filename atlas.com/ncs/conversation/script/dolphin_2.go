package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Dolphin2 is located in Herb Town - Pier on the Beach (251000100)
type Dolphin2 struct {
}

func (r Dolphin2) NPCId() uint32 {
	return npc.Dolphin2
}

func (r Dolphin2) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Will you move to ").
		BlueText().ShowMap(_map.Aquarium).
		BlackText().AddText(" now? The price is ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 1000)).
		BlackText().AddText(".")
	return SendYesNo(l, c, m.String(), r.Validate, r.TooBusy)
}

func (r Dolphin2) TooBusy(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hmmm ... too busy to do it right now? If you feel like doing it, though, come back and find me.")
	return SendOk(l, c, m.String())
}

func (r Dolphin2) Validate(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 1000) {
		return r.NotEnoughMeso(l, c)
	}
	return r.Process(l, c)
}

func (r Dolphin2) Process(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -1000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
	}
	return WarpById(_map.Aquarium, 0)(l, c)
}

func (r Dolphin2) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I don't think you have enough money...")
	return SendOk(l, c, m.String())
}
