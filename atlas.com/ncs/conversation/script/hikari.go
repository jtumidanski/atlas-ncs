package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Hikari is located in Zipangu - Showa Town (801000000)
type Hikari struct {
}

func (r Hikari) NPCId() uint32 {
	return npc.Hikari
}

func (r Hikari) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText(fmt.Sprintf("Would you like to enter the bathhouse? That'll be %d mesos for you.", 300))
	return SendYesNo(l, c, m.String(), r.Validate, r.AnotherTime)
}

func (r Hikari) AnotherTime(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Please come back some other time.")
	return SendOk(l, c, m.String())
}

func (r Hikari) Validate(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 300) {
		return r.NotEnoughMesos(l, c)
	}
	return r.Process(l, c)
}

func (r Hikari) NotEnoughMesos(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText(fmt.Sprintf("Please check and see if you have %d mesos to enter this place.", 300))
	return SendOk(l, c, m.String())
}

func (r Hikari) Process(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -300)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
	}
	gender := character.GetGender(l)(c.CharacterId)
	destination := uint32(0)
	if gender == character.GenderMale {
		destination = _map.LockerRoomM
	} else if gender == character.GenderFemale {
		destination = _map.LockerRoomF
	}
	return WarpByName(destination, "out00")(l, c)
}
