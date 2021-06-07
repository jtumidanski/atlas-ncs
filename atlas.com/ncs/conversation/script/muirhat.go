package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Muirhat is located in The Nautilus - Top Floor - Hallway (120000100)
type Muirhat struct {
}

func (r Muirhat) NPCId() uint32 {
	return npc.Muirhat
}

func (r Muirhat) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 2175) {
		return r.ChasingOneAnother(l, c)
	}
	return r.Validate(l, c)
}

func (r Muirhat) ChasingOneAnother(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The Black Magician and his followers. Kyrin and the Crew of Nautilus. ").NewLine().
		AddText("They'll be chasing one another until one of them doesn't exist, that's for sure.")
	return SendOk(l, c, m.String())
}

func (r Muirhat) Validate(l logrus.FieldLogger, c Context) State {
	if !character.CanHold(l)(c.CharacterId, item.ReturnScrollToNautilus) {
		return r.NoFreeInventory(l, c)
	}
	return r.TakeThis(l, c)
}

func (r Muirhat) NoFreeInventory(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("No free inventory spot available. Please make room in your USE inventory first.")
	return SendOk(l, c, m.String())
}

func (r Muirhat) TakeThis(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Please take this ").
		BlueText().ShowItemName1(item.ReturnScrollToNautilus).
		BlackText().AddText(", it will make your life a lot easier.")
	return SendNext(l, c, m.String(), r.Process)
}

func (r Muirhat) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.ReturnScrollToNautilus, 1)
	return WarpById(_map.TheRestingSpotPigPark, 0)(l, c)
}
