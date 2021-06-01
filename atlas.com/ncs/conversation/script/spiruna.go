package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Spiruna is located in Orbis - Old Man's House (200050001)
type Spiruna struct {
}

func (r Spiruna) NPCId() uint32 {
	return npc.Spiruna
}

func (r Spiruna) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestCompleted(l)(c.CharacterId, 3034) {
		return r.GoAway(l, c)
	}

	return r.RefineDarkCrystal(l, c)
}

func (r Spiruna) GoAway(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Go away, I'm trying to meditate.")
	return SendOk(l, c, m.String())
}

func (r Spiruna) RefineDarkCrystal(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You've been so much of a help to me... If you have any Dark Crystal Ore, I can refine it for you for only ").
		BlueText().AddText(fmt.Sprintf("%d meso", 500000)).
		BlackText().AddText(" each.")
	return SendYesNo(l, c, m.String(), r.HowMany, Exit())
}

func (r Spiruna) HowMany(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Okay, so how many do you want me to make?")
	return SendGetNumber(l, c, m.String(), r.Validate, 1, 1, 100)
}

func (r Spiruna) Validate(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, uint32(500000*selection)) {
			return r.NotForFree(l, c)
		}

		if !character.HasItems(l)(c.CharacterId, item.DarkCrystalOre, uint32(10*selection)) {
			return r.MoreOre(l, c)
		}

		if !character.CanHoldAll(l)(c.CharacterId, item.DarkCrystal, uint32(selection)) {
			return r.MoreInventorySpace(l, c)
		}

		return r.Process(selection)(l, c)
	}
}

func (r Spiruna) NotForFree(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I'm sorry, but I am NOT doing this for free.")
	return SendOk(l, c, m.String())
}

func (r Spiruna) MoreOre(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I need that ore to refine the Crystal. No exceptions..")
	return SendOk(l, c, m.String())
}

func (r Spiruna) MoreInventorySpace(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Are you having trouble with no empty slots on your inventory? Sort that out first!")
	return SendOk(l, c, m.String())
}

func (r Spiruna) Process(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -selection*500000)
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment from character %d.", c.CharacterId)
		}
		character.GainItem(l)(c.CharacterId, item.DarkCrystalOre, -selection*10)
		character.GainItem(l)(c.CharacterId, item.DarkCrystal, selection)
		return r.Success(l, c)
	}
}

func (r Spiruna) Success(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Use it wisely.")
	return SendOk(l, c, m.String())
}
