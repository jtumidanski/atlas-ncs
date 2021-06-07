package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// FirstEOSRock is located in Ludibrium - Eos Tower 100th Floor (221024400)
type FirstEOSRock struct {
}

func (r FirstEOSRock) NPCId() uint32 {
	return npc.FirstEOSRock
}

func (r FirstEOSRock) Initial(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.EOSRockScroll) {
		return r.ToNext(l, c)
	}
	return r.NeedRock(l, c)
}

func (r FirstEOSRock) ToNext(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("First Eos Rock").
		BlackText().AddText(". Will you teleport to ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(" at the 71st floor?")
	return SendYesNo(l, c, m.String(), r.Process, Exit())
}

func (r FirstEOSRock) NeedRock(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("There's a rock that will enable you to teleport to ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(", but it cannot be activated without the scroll.")
	return SendOk(l, c, m.String())
}

func (r FirstEOSRock) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.EOSRockScroll, -1)
	return WarpById(_map.EosTower71stFloor, 3)(l, c)
}
