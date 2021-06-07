package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// FourthEOSRock is located in Ludibrium - Eos Tower 1st Floor (221020000)
type FourthEOSRock struct {
}

func (r FourthEOSRock) NPCId() uint32 {
	return npc.FourthEOSRock
}

func (r FourthEOSRock) Initial(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.EOSRockScroll) {
		return r.ToNext(l, c)
	}
	return r.NeedRock(l, c)
}

func (r FourthEOSRock) ToNext(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Fourth Eos Rock").
		BlackText().AddText(". Will you head over to ").
		BlueText().AddText("Third Eos Rock").
		BlackText().AddText(" at the 41st floor?")
	return SendYesNo(l, c, m.String(), r.Process, Exit())
}

func (r FourthEOSRock) NeedRock(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("There's a rock that will enable you to teleport to ").
		BlueText().AddText("Third Eos Rock").
		BlackText().AddText(", but it cannot be activated without the scroll.")
	return SendOk(l, c, m.String())
}

func (r FourthEOSRock) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.EOSRockScroll, -1)
	return WarpById(_map.EosTower41stFloor, 3)(l, c)
}
