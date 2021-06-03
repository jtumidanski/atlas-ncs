package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type CamelCab struct {
}

func (r CamelCab) NPCId() uint32 {
	return npc.CamelCab
}

func (r CamelCab) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.OutsideNorthEntranceOfAriant {
		return r.ToMagatia(l, c)
	} else {
		return r.ToAriant(l, c)
	}
}

func (r CamelCab) ToAriant(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Would you like to take the ").
		BlueText().AddText("Camel Cab").
		BlackText().AddText(" to ").
		BlueText().AddText("Ariant").
		BlackText().AddText(", the town of Burning Roads? The fare is ").
		BlueText().AddText("1500 mesos").
		BlackText().AddText(".")
	return SendYesNo(l, c, m.String(), r.Validate(_map.Ariant), r.TooBusy)
}

func (r CamelCab) ToMagatia(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Would you like to take the ").
		BlueText().AddText("Camel Cab").
		BlackText().AddText(" to ").
		BlueText().AddText("Magatia").
		BlackText().AddText(", the town of Alchemy? The fare is ").
		BlueText().AddText("1500 mesos").
		BlackText().AddText(".")
	return SendYesNo(l, c, m.String(), r.Validate(_map.Magatia), r.TooBusy)
}

func (r CamelCab) Validate(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, 1500) {
			return r.ShortOnMesos(l, c)
		}
		return r.Process(mapId)(l, c)
	}
}

func (r CamelCab) ShortOnMesos(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I am sorry, but I think you are short on mesos. I am afraid I can't let you ride this if you do not have enough money to do so. Please come back when you have enough money to use this.")
	return SendOk(l, c, m.String())
}

func (r CamelCab) Process(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -1500)
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
		}
		err = npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return nil
	}
}

func (r CamelCab) TooBusy(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hmmm... too busy to do it right now? If you feel like doing it, though, come back and find me.")
	return SendOk(l, c, m.String())
}
