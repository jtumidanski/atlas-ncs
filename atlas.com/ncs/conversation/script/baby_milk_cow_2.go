package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// BabyMilkCow2 is located in Hidden Chamber - The Nautilus - Stable (912000100)
type BabyMilkCow2 struct {
}

func (r BabyMilkCow2) NPCId() uint32 {
	return npc.BabyMilkCow2
}

func (r BabyMilkCow2) Initial(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.MilkJug) {
		return r.RemainsEmpty(l, c)
	}
	if character.HasItem(l)(c.CharacterId, item.MilkJugOneThird) {
		return r.DrinkBottle(item.MilkJugOneThird)(l, c)
	}
	if character.HasItem(l)(c.CharacterId, item.MilkJugTwoThird) {
		return r.DrinkBottle(item.MilkJugTwoThird)(l, c)
	}
	return r.DrinkBottle(item.MilkJugFull)(l, c)
}

func (r BabyMilkCow2) RemainsEmpty(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The hungry calf is drinking all the milk! The bottle remains empty...")
	return SendNext(l, c, m.String(), r.EmptyBottle)
}

func (r BabyMilkCow2) EmptyBottle(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The hungry calf isn't interested in the empty bottle.")
	return SendNextPrevious(l, c, m.String(), Exit(), r.RemainsEmpty)
}

func (r BabyMilkCow2) DrinkBottle(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.GainItem(l)(c.CharacterId, itemId, -1)
		character.GainItem(l)(c.CharacterId, item.MilkJug, 1)
		m := message.NewBuilder().
			AddText("The hungry calf is drinking all the milk! The bottle is now empty.")
		return SendNext(l, c, m.String(), Exit())
	}
}
