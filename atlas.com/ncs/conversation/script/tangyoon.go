package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Tangyoon is located in The Nautilus - Cafeteria (120000103)
type Tangyoon struct {
}

func (r Tangyoon) NPCId() uint32 {
	return npc.Tangyoon
}

func (r Tangyoon) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 2180) {
		return Exit()(l, c)
	}

	return r.SendToStable(l, c)
}

func (r Tangyoon) SendToStable(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Okay, I'll now send you to the stable where my cows are. Watch out for the calves that drink all the milk. You don't want your effort to go to waste.")
	return SendNext(l, c, m.String(), r.IGetConfused)
}

func (r Tangyoon) IGetConfused(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It won't be easy to tell at a glance between a calf and a cow. Those calves may only be a month or two old, but they have already grown to the size of their mother. They even look alike...even I get confused at times! Good luck!")
	return SendNextPrevious(l, c, m.String(), r.Validate, r.SendToStable)
}

func (r Tangyoon) Validate(l logrus.FieldLogger, c Context) State {
	if !character.CanHold(l)(c.CharacterId, item.MilkJug) {
		return r.FullInventory(l, c)
	}
	return r.Award(l, c)
}

func (r Tangyoon) FullInventory(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I can't give you the empty bottle because your inventory is full. Please make some room in your Etc window.")
	return SendOk(l, c, m.String())
}

func (r Tangyoon) Award(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.MilkJug, 1)
	return WarpById(_map.Stable, 0)(l, c)
}
