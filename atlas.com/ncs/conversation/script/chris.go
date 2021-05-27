package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Chris is located in Victoria Road - Kerning City (103000000)
type Chris struct {
}

func (r Chris) NPCId() uint32 {
	return npc.Chris
}

func (r Chris) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Chris) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Yes, I do own this forge. If you're willing to pay, I can offer you some of my services.").NewLine().
		OpenItem(0).BlueText().AddText("Refine a mineral ore").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Refine a jewel ore").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("I have Iron Hog's Metal Hoof...").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Upgrade a claw").CloseItem()
	return SendListSelection(l, c, m.String(), r.WhatToDo)
}

func (r Chris) WhatToDo(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.RefineMineral
	case 1:
		return r.RefineJewel
	case 2:
		return r.MetalHoof
	case 3:
		return r.UpgradeClaw
	}
	return nil
}

func (r Chris) RefineMineral(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, what kind of mineral ore would you like to refine?").NewLine().
		OpenItem(0).BlueText().AddText("Bronze").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Steel").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Mithril").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Adamantium").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Silver").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Orihalcon").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Gold").CloseItem()
	return SendListSelection(l, c, m.String(), r.RefineMineralSelection)
}

func (r Chris) RefineJewel(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, what kind of jewel ore would you like to refine?").NewLine().
		OpenItem(0).BlueText().AddText("Garnet").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Amethyst").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Aquamarine").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Emerald").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Opal").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Sapphire").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Topaz").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Diamond").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Black Crystal").CloseItem()
	return SendListSelection(l, c, m.String(), r.RefineJewelSelection)
}

func (r Chris) MetalHoof(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You know about that? Not many people realize the potential in the Iron Hog's Metal Hoof... I can make this into something special, if you want me to.")
	return SendYesNo(l, c, m.String(), r.SpecialRefinement, Exit())
}

func (r Chris) UpgradeClaw(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Ah, you wish to upgrade a claw? Then tell me, which one?").NewLine().
		OpenItem(0).BlueText().AddText("Blood Gigantic#k - Thief Lv. 60#b").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Sapphire Gigantic#k - Thief Lv. 60#b").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Dark Gigantic#k - Thief Lv. 60#b").CloseItem()
	return SendListSelection(l, c, m.String(), r.UpgradeClawSelection)
}

func (r Chris) RefineMineralSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.HowMany(item.BronzePlate, r.BronzeRefineRequirements())
	case 1:
		return r.HowMany(item.SteelPlate, r.SteelRefineRequirements())
	case 2:
		return r.HowMany(item.MithrilPlate, r.MithrilRefineRequirements())
	case 3:
		return r.HowMany(item.AdamantiumPlate, r.AdamantiumRefineRequirements())
	case 4:
		return r.HowMany(item.SilverPlate, r.SilverRefineRequirements())
	case 5:
		return r.HowMany(item.OrihalconPlate, r.OrihalconRefineRequirements())
	case 6:
		return r.HowMany(item.GoldPlate, r.GoldRefineRequirements())
	}
	return nil
}

func (r Chris) RefineJewelSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.HowMany(item.Garnet, r.GarnetRefineRequirements())
	case 1:
		return r.HowMany(item.Amethyst, r.AmethystRefineRequirements())
	case 2:
		return r.HowMany(item.AquaMarine, r.AquamarineRefineRequirements())
	case 3:
		return r.HowMany(item.Emerald, r.EmeraldRefineRequirements())
	case 4:
		return r.HowMany(item.Opal, r.OpalRefineRequirements())
	case 5:
		return r.HowMany(item.Sapphire, r.SapphireRefineRequirements())
	case 6:
		return r.HowMany(item.Topaz, r.TopazRefineRequirements())
	case 7:
		return r.HowMany(item.Diamond, r.DiamondRefineRequirements())
	case 8:
		return r.HowMany(item.BlackCrystal, r.BlackCrystalRefineRequirements())
	}
	return nil
}

func (r Chris) SpecialRefinement(l logrus.FieldLogger, c Context) State {
	return r.HowMany(item.SteelPlate, r.SteelFromHogRequirements())(l, c)
}

func (r Chris) UpgradeClawSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.BloodGigantic, r.BloodGiganticRequirements())
	case 1:
		return r.Confirm(item.SapphireGigantic, r.SapphireGiganticRequirements())
	case 2:
		return r.Confirm(item.DarkGigantic, r.DarkGiganticRequirements())
	}
	return nil
}

func (r Chris) HowMany(itemId uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("So, you want me to make some ").
			ShowItemName1(itemId).
			AddText("s? In that case, how many do you want me to make?")
		return SendGetNumber(l, c, m.String(), r.QuantitySelection(itemId, requirements), 1, 1, 100)
	}
}

func (r Chris) QuantitySelection(itemId uint32, requirements RefinementRequirements) ProcessNumber {
	return func(selection int32) StateProducer {
		return r.ConfirmQuantity(itemId, uint32(selection), requirements)
	}
}

func (r Chris) Confirm(itemId uint32, requirements RefinementRequirements) StateProducer {
	return r.ConfirmQuantity(itemId, 1, requirements)
}

func (r Chris) ConfirmQuantity(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("You want me to make ")
		if amount == 1 {
			m = m.AddText("a ").ShowItemName1(itemId)
		} else {
			m = m.AddText(fmt.Sprintf("%d ", amount)).ShowItemName1(itemId)
		}
		m = m.AddText("? In that case, I'm going to need specific items from you in order to make it. Make sure you have room in your inventory, though!").NewLine()
		for _, req := range requirements.requirements {
			m = m.ShowItemImage2(req.itemId).AddText(fmt.Sprintf(" %d ", req.amount)).ShowItemName1(req.itemId).NewLine()
		}
		if requirements.cost > 0 {
			m = m.ShowItemImage2(item.MoneySack).AddText(fmt.Sprintf(" %d meso", requirements.cost*amount))
		}
		return SendYesNo(l, c, m.String(), r.Validate(itemId, amount, requirements), Exit())
	}
}

func (r Chris) BronzeRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeOre, amount: 10}}, cost: 300}
}

func (r Chris) SteelRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelOre, amount: 10}}, cost: 300}
}

func (r Chris) MithrilRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.MithrilOre, amount: 10}}, cost: 300}
}

func (r Chris) AdamantiumRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.AdamantiumOre, amount: 10}}, cost: 500}
}

func (r Chris) SilverRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SilverOre, amount: 10}}, cost: 500}
}

func (r Chris) OrihalconRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.OrihalconOre, amount: 10}}, cost: 500}
}

func (r Chris) GoldRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.GoldOre, amount: 10}}, cost: 800}
}

func (r Chris) GarnetRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.GarnetOre, amount: 10}}, cost: 500}
}

func (r Chris) AmethystRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.AmethystOre, amount: 10}}, cost: 500}
}

func (r Chris) AquamarineRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.AquaMarineOre, amount: 10}}, cost: 500}
}

func (r Chris) EmeraldRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.EmeraldOre, amount: 10}}, cost: 500}
}

func (r Chris) OpalRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.OpalOre, amount: 10}}, cost: 500}
}

func (r Chris) SapphireRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SapphireOre, amount: 10}}, cost: 500}
}

func (r Chris) TopazRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.TopazOre, amount: 10}}, cost: 500}
}

func (r Chris) DiamondRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.DiamondOre, amount: 10}}, cost: 1000}
}

func (r Chris) BlackCrystalRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BlackCrystalOre, amount: 10}}, cost: 3000}
}

func (r Chris) BloodGiganticRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeGigantic, amount: 1}, {itemId: item.MoonRock, amount: 1}, {itemId: item.Garnet, amount: 8}, {itemId: item.DrakesBlood, amount: 10}}, cost: 80000}
}

func (r Chris) SapphireGiganticRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeGigantic, amount: 1}, {itemId: item.MoonRock, amount: 1}, {itemId: item.Sapphire, amount: 8}, {itemId: item.SapOfAncientTree, amount: 10}}, cost: 80000}
}

func (r Chris) DarkGiganticRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeGigantic, amount: 1}, {itemId: item.MoonRock, amount: 1}, {itemId: item.BlackCrystal, amount: 3}, {itemId: item.TaurospearHorn, amount: 5}}, cost: 100000}
}

func (r Chris) SteelFromHogRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.IronHogsMetalHoof, amount: 100}}}
}

func (r Chris) Validate(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.CanHoldAll(l)(c.CharacterId, itemId, amount) {
			return r.MakeRoom(l, c)
		}
		if !character.HasMeso(l)(c.CharacterId, requirements.cost*amount) {
			return r.CannotAfford(l, c)
		}
		for _, req := range requirements.requirements {
			if !character.HasItems(l)(c.CharacterId, req.itemId, uint32(req.amount)*amount) {
				return r.MissingSomething(req.itemId)(l, c)
			}
		}
		return r.PerformRefine(itemId, amount, requirements)(l, c)
	}
}

func (r Chris) MakeRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Check your inventory for a free slot first.")
	return SendOk(l, c, m.String())
}

func (r Chris) CannotAfford(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Cash only, no credit.")
	return SendOk(l, c, m.String())
}

func (r Chris) MissingSomething(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("I cannot accept substitutes. If you don't have what I need, then I won't be able to help you. Please bring more ").
			ShowItemName1(itemId).
			AddText(".")
		return SendOk(l, c, m.String())
	}
}

func (r Chris) PerformRefine(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -int32(amount*requirements.cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment for refine.")
		}
		for _, req := range requirements.requirements {
			character.GainItem(l)(c.CharacterId, req.itemId, -int32(req.amount)*int32(amount))
		}
		character.GainItem(l)(c.CharacterId, itemId, int32(amount))
		return r.Success(l, c)
	}
}

func (r Chris) Success(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Phew... I almost didn't think that would work for a second... Well, I hope you enjoy it, anyway.")
	return SendOk(l, c, m.String())
}