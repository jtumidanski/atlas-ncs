package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// Bush1 is located in Victoria Road - Nautilus Harbor (120000000)
type Bush1 struct {
}

func (r Bush1) NPCId() uint32 {
	return npc.Bush1
}

func (r Bush1) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 2186) {
		return r.PileOfBoxes(l, c)
	}
	return r.DoYouWant(l, c)
}

func (r Bush1) PileOfBoxes(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Just a pile of boxes, nothing special...")
	return SendOk(l, c, m.String())
}

func (r Bush1) DoYouWant(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you want to obtain a glasses?")
	return SendNext(l, c, m.String(), r.Validate)
}

func (r Bush1) Validate(l logrus.FieldLogger, c Context) State {
	if character.HasAnyItem(l)(c.CharacterId, item.AbelsGlasses, item.MiscellaneousGlasses1, item.MiscellaneousGlasses2) {
		return r.AlreadyHave(l, c)
	}

	if character.CanHold(l)(c.CharacterId, item.AbelsGlasses) {
		return r.NoInventoryRoom(l, c)
	}

	random := rand.Intn(3)
	if random == 0 {
		character.GainItem(l)(c.CharacterId, item.AbelsGlasses, 1)
	} else if random == 1 {
		character.GainItem(l)(c.CharacterId, item.MiscellaneousGlasses1, 1)
	} else {
		character.GainItem(l)(c.CharacterId, item.MiscellaneousGlasses2, 1)
	}
	return Exit()(l, c)
}

func (r Bush1) AlreadyHave(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You ").
		BlueText().AddText("already have").
		BlackText().AddText(" the glasses that was here!")
	return SendOk(l, c, m.String())
}

func (r Bush1) NoInventoryRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Check for a available slot on your ETC inventory.")
	return SendOk(l, c, m.String())
}
