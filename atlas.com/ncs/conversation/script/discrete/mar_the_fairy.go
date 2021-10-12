package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/character/inventory"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/pet"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// MarTheFairy is located in Victoria Road - Ellinia (101000000)
type MarTheFairy struct {
}

func (r MarTheFairy) NPCId() uint32 {
	return npc.MarTheFairy
}

func (r MarTheFairy) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r MarTheFairy) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I am Mar the Fairy. If you have a dragon at level 15 or higher and a rock of evolution, I can evolve your dragon. If you are lucky, you may even get a black one! Would you like me to do so?")
	return script.SendYesNo(l, span, c, m.String(), r.ValidateEvolve, r.SeeYouNextTime)
}

func (r MarTheFairy) SeeYouNextTime(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright, see you next time.")
	return script.SendOk(l, span, c, m.String())
}

func (r MarTheFairy) ValidateEvolve(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, pet.DragonEgg) {
		character.GainItem(l, span)(c.CharacterId, pet.DragonEgg, -1)
		character.GainItem(l, span)(c.CharacterId, pet.BabyDragon, 1)
		return r.Hatched(l, span, c)
	} else if !pet.HasPetInSlot(l)(c.CharacterId, 0) {
		return r.PetInSlotOne(l, span, c)
	} else if !character.HasItem(l, span)(c.CharacterId, item.TheRockOfEvolution) || !pet.Is(l)(c.CharacterId, 0, pet.BabyDragon, pet.GreenDragon, pet.RedDragon, pet.BlueDragon, pet.BlackDragon) {
		return r.DoesNotMeetRequirements(l, span, c)
	} else if !pet.IsLevel(l)(c.CharacterId, 0, 15) {
		return r.NotLeveledEnough(l, span, c)
	} else if character.HasItems(l, span)(c.CharacterId, pet.BabyDragon, 2) ||
		character.HasItems(l, span)(c.CharacterId, pet.GreenDragon, 2) ||
		character.HasItems(l, span)(c.CharacterId, pet.RedDragon, 2) ||
		character.HasItems(l, span)(c.CharacterId, pet.BlueDragon, 2) ||
		character.HasItems(l, span)(c.CharacterId, pet.BlackDragon, 2) {
		return r.MultipleItem(l, span, c)
	} else {
		return r.Evolve(l, span, c)
	}
}

func (r MarTheFairy) Hatched(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I don't know how you got that egg, but it has hatched, apparently!")
	return script.SendOk(l, span, c, m.String())
}

func (r MarTheFairy) PetInSlotOne(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Make sure your pet is equipped on slot 1.")
	return script.SendOk(l, span, c, m.String())
}

func (r MarTheFairy) DoesNotMeetRequirements(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You do not meet the requirements. You need ").
		ShowItemImage2(item.TheRockOfEvolution).ShowItemName1(item.TheRockOfEvolution).
		AddText(", as well as either one of ").
		PurpleText().ShowItemImage2(pet.BabyDragon).ShowItemName1(pet.BabyDragon).
		BlackText().AddText(", ").
		GreenText().ShowItemImage2(pet.GreenDragon).ShowItemName1(pet.GreenDragon).
		BlackText().AddText(", ").
		RedText().ShowItemImage2(pet.RedDragon).ShowItemName1(pet.RedDragon).
		BlackText().AddText(", ").
		BlueText().ShowItemImage2(pet.BlueDragon).ShowItemName1(pet.BlueDragon).
		BlackText().AddText(", or ").
		BoldText().ShowItemImage2(pet.BlackDragon).ShowItemName1(pet.BlackDragon).
		NormalText().AddText(" equipped on slot 1. Please come back when you do.")
	return script.SendOk(l, span, c, m.String())
}

func (r MarTheFairy) NotLeveledEnough(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Your pet must be level 15 or above to evolve.")
	return script.SendOk(l, span, c, m.String())
}

func (r MarTheFairy) MultipleItem(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You have a dragon which isn't out, and as well as a dragon which is out. I can remove one for you. Remember that the data for the dragon I am removing will be lost.").NewLine().
		OpenItem(0).RedText().AddText("Remove my CASH first slot.").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Remove the first dragon in my inventory.").CloseItem().NewLine().
		OpenItem(2).GreenText().AddText("No thanks.").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.ProcessRemoval)
}

func (r MarTheFairy) ProcessRemoval(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.RemoveFirstCashSlot
	case 1:
		return r.RemoveFirstDragon
	case 2:
		return r.NoThanks
	}
	return nil
}

func (r MarTheFairy) NoThanks(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Okay, come back next time.")
	return script.SendOk(l, span, c, m.String())
}

func (r MarTheFairy) RemoveFirstDragon(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItems(l, span)(c.CharacterId, pet.BabyDragon, 2) {
		character.GainItem(l, span)(c.CharacterId, pet.BabyDragon, -1)
	} else if character.HasItems(l, span)(c.CharacterId, pet.GreenDragon, 2) {
		character.GainItem(l, span)(c.CharacterId, pet.GreenDragon, -1)
	} else if character.HasItems(l, span)(c.CharacterId, pet.RedDragon, 2) {
		character.GainItem(l, span)(c.CharacterId, pet.RedDragon, -1)
	} else if character.HasItems(l, span)(c.CharacterId, pet.BlueDragon, 2) {
		character.GainItem(l, span)(c.CharacterId, pet.BlueDragon, -1)
	} else if character.HasItems(l, span)(c.CharacterId, pet.BlackDragon, 2) {
		character.GainItem(l, span)(c.CharacterId, pet.BlackDragon, -1)
	}
	m := message.NewBuilder().
		AddText("The first dragon in your inventory is removed.")
	return script.SendOk(l, span, c, m.String())
}

func (r MarTheFairy) RemoveFirstCashSlot(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.RemoveFromSlot(l)(c.CharacterId, inventory.TypeCash, 1, 1)
	m := message.NewBuilder().
		AddText("Your cash first slot is removed.")
	return script.SendOk(l, span, c, m.String())
}

func (r MarTheFairy) Evolve(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	babyIndex := int16(-1)
	for i := int16(0); i < 3; i++ {
		if pet.Is(l)(c.CharacterId, i, pet.BabyDragon) {
			babyIndex = i
			break
		}
	}

	if babyIndex == -1 {
		return r.NotReady(l, span, c)
	}

	random := 1 + rand.Intn(9)
	var dragonId uint32
	if random >= 1 && random <= 3 {
		dragonId = pet.GreenDragon
	} else if random >= 4 && random <= 6 {
		dragonId = pet.RedDragon
	} else if random >= 7 && random <= 9 {
		dragonId = pet.BlueDragon
	} else {
		dragonId = pet.BlackDragon
	}

	character.GainItem(l, span)(c.CharacterId, item.TheRockOfEvolution, -1)
	pet.Evolve(l)(c.CharacterId, babyIndex, dragonId)

	m := message.NewBuilder().
		AddText("Your dragon has now evolved!! It used to be a ").
		ShowItemImage2(pet.BabyDragon).AddText(" ").ShowItemName1(pet.BabyDragon).
		AddText(", and now it's a ").
		ShowItemImage2(dragonId).AddText(" ").ShowItemName1(dragonId).
		AddText("!")
	return script.SendOk(l, span, c, m.String())
}

func (r MarTheFairy) NotReady(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You either don't have a pet dragon ready to evolve or you lack ").
		BlueText().ShowItemName1(item.TheRockOfEvolution).
		BlackText().AddText(".")
	return script.SendOk(l, span, c, m.String())
}
