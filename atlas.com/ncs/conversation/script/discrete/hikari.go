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

// Hikari is located in Zipangu - Showa Town (801000000)
type Hikari struct {
}

func (r Hikari) NPCId() uint32 {
	return npc.Hikari
}

func (r Hikari) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText(fmt.Sprintf("Would you like to enter the bathhouse? That'll be %d mesos for you.", 300))
	return script.SendYesNo(l, span, c, m.String(), r.Validate, r.AnotherTime)
}

func (r Hikari) AnotherTime(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please come back some other time.")
	return script.SendOk(l, span, c, m.String())
}

func (r Hikari) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 300) {
		return r.NotEnoughMesos(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r Hikari) NotEnoughMesos(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText(fmt.Sprintf("Please check and see if you have %d mesos to enter this place.", 300))
	return script.SendOk(l, span, c, m.String())
}

func (r Hikari) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -300)
	gender := character.GetGender(l, span)(c.CharacterId)
	destination := uint32(0)
	if gender == character.GenderMale {
		destination = _map.LockerRoomM
	} else if gender == character.GenderFemale {
		destination = _map.LockerRoomF
	}
	return script.WarpByName(destination, "out00")(l, span, c)
}
