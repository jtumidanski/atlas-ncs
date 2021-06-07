package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// NLCTaxi is located in New Leaf City Town Street - New Leaf City - Town Center (600000000)
type NLCTaxi struct {
}

func (r NLCTaxi) NPCId() uint32 {
	return npc.NLCTaxi
}

func (r NLCTaxi) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.HauntedHouse {
		m := message.NewBuilder().
			AddText("Would you like to return back to ").
			BlueText().AddText("civilization").
			BlackText().AddText(fmt.Sprintf("? The fee is %d mesos.", 15000))
		return SendYesNo(l, c, m.String(), r.ToNLC, r.NextTime)
	} else {
		m := message.NewBuilder().
			AddText("Would you like to go to the ").
			BlueText().AddText("Haunted Mansion").
			BlackText().AddText(fmt.Sprintf("? The fee is %d mesos.", 15000))
		return SendYesNo(l, c, m.String(), r.ToHauntedHouse, r.NextTime)
	}
}

func (r NLCTaxi) NextTime(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Alright, see you next time.")
	return SendOk(l, c, m.String())
}

func (r NLCTaxi) WarpToNLC(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.NewLeafCityTownCenter, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.NewLeafCityTownCenter, c.NPCId)
	}
	return Exit()(l, c)
}

func (r NLCTaxi) ToNLC(l logrus.FieldLogger, c Context) State {
	return r.Validate(r.WarpToNLC)(l, c)
}

func (r NLCTaxi) ToHauntedHouse(l logrus.FieldLogger, c Context) State {
	return r.Validate(r.WarpToHauntedHouse)(l, c)
}

func (r NLCTaxi) WarpToHauntedHouse(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.HauntedHouse, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.HauntedHouse, c.NPCId)
	}
	return Exit()(l, c)
}

func (r NLCTaxi) Validate(warp func(l logrus.FieldLogger, c Context) State) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, 15000) {
			return r.NotEnoughMesos(l, c)
		}
		err := character.GainMeso(l)(c.CharacterId, -15000)
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment from character %d.", c.ChannelId)
		}
		return warp(l, c)
	}
}

func (r NLCTaxi) NotEnoughMesos(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hey, what are you trying to pull on? You don't have enough meso to pay the fee.")
	return SendOk(l, c, m.String())
}
