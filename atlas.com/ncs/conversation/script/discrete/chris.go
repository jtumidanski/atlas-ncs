package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Chris is located in Victoria Road - Kerning City (103000000)
type Chris struct {
}

func (r Chris) NPCId() uint32 {
	return npc.Chris
}

func (r Chris) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := "Yes, I do own this forge. If you're willing to pay, I can offer you some of my services."
	categories := r.CreateCategories()
	return refine.NewGenericRefine(l, c, hello, categories)
}

func (r Chris) BronzeRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeOre, Amount: 10}}, Cost: 300}
}

func (r Chris) SteelRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelOre, Amount: 10}}, Cost: 300}
}

func (r Chris) MithrilRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.MithrilOre, Amount: 10}}, Cost: 300}
}

func (r Chris) AdamantiumRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.AdamantiumOre, Amount: 10}}, Cost: 500}
}

func (r Chris) SilverRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SilverOre, Amount: 10}}, Cost: 500}
}

func (r Chris) OrihalconRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.OrihalconOre, Amount: 10}}, Cost: 500}
}

func (r Chris) GoldRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.GoldOre, Amount: 10}}, Cost: 800}
}

func (r Chris) GarnetRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.GarnetOre, Amount: 10}}, Cost: 500}
}

func (r Chris) AmethystRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.AmethystOre, Amount: 10}}, Cost: 500}
}

func (r Chris) AquamarineRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.AquaMarineOre, Amount: 10}}, Cost: 500}
}

func (r Chris) EmeraldRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.EmeraldOre, Amount: 10}}, Cost: 500}
}

func (r Chris) OpalRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.OpalOre, Amount: 10}}, Cost: 500}
}

func (r Chris) SapphireRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SapphireOre, Amount: 10}}, Cost: 500}
}

func (r Chris) TopazRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.TopazOre, Amount: 10}}, Cost: 500}
}

func (r Chris) DiamondRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.DiamondOre, Amount: 10}}, Cost: 1000}
}

func (r Chris) BlackCrystalRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BlackCrystalOre, Amount: 10}}, Cost: 3000}
}

func (r Chris) BloodGiganticRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeGigantic, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.Garnet, Amount: 8}, {ItemId: item.DrakesBlood, Amount: 10}}, Cost: 80000}
}

func (r Chris) SapphireGiganticRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeGigantic, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.Sapphire, Amount: 8}, {ItemId: item.SapOfAncientTree, Amount: 10}}, Cost: 80000}
}

func (r Chris) DarkGiganticRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeGigantic, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 3}, {ItemId: item.TaurospearHorn, Amount: 5}}, Cost: 100000}
}

func (r Chris) SteelFromHogRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.IronHogsMetalHoof, Amount: 100}}}
}

func (r Chris) CreateCategories() []refine.RefinementCategory {
	return []refine.RefinementCategory{
		r.RefineMineralOre(),
		r.RefineJewelOre(),
		r.RefineMetalHoof(),
		r.RefineUpgradeClaw(),
	}
}

func (r Chris) RefineMineralOre() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Refine a mineral ore",
		Prompt:          "So, what kind of mineral ore would you like to refine?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateRefinementChoice("Bronze", refine.HowMany(item.BronzePlate, r.BronzeRefineRequirements())),
			r.CreateRefinementChoice("Steel", refine.HowMany(item.SteelPlate, r.SteelRefineRequirements())),
			r.CreateRefinementChoice("Mithril", refine.HowMany(item.MithrilPlate, r.MithrilRefineRequirements())),
			r.CreateRefinementChoice("Adamantium", refine.HowMany(item.AdamantiumPlate, r.AdamantiumRefineRequirements())),
			r.CreateRefinementChoice("Silver", refine.HowMany(item.SilverPlate, r.SilverRefineRequirements())),
			r.CreateRefinementChoice("Orihalcon", refine.HowMany(item.OrihalconPlate, r.OrihalconRefineRequirements())),
			r.CreateRefinementChoice("Gold", refine.HowMany(item.GoldPlate, r.GoldRefineRequirements())),
		},
	}
}

func (r Chris) RefineJewelOre() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Refine a jewel ore",
		Prompt:          "So, what kind of jewel ore would you like to refine?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateRefinementChoice("Garnet", refine.HowMany(item.Garnet, r.GarnetRefineRequirements())),
			r.CreateRefinementChoice("Amethyst", refine.HowMany(item.Amethyst, r.AmethystRefineRequirements())),
			r.CreateRefinementChoice("Aquamarine", refine.HowMany(item.AquaMarine, r.AquamarineRefineRequirements())),
			r.CreateRefinementChoice("Emerald", refine.HowMany(item.Emerald, r.EmeraldRefineRequirements())),
			r.CreateRefinementChoice("Opal", refine.HowMany(item.Opal, r.OpalRefineRequirements())),
			r.CreateRefinementChoice("Sapphire", refine.HowMany(item.Sapphire, r.SapphireRefineRequirements())),
			r.CreateRefinementChoice("Topaz", refine.HowMany(item.Topaz, r.TopazRefineRequirements())),
			r.CreateRefinementChoice("Diamond", refine.HowMany(item.Diamond, r.DiamondRefineRequirements())),
			r.CreateRefinementChoice("Black Crystal", refine.HowMany(item.BlackCrystal, r.BlackCrystalRefineRequirements())),
		},
	}
}

func (r Chris) RefineMetalHoof() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "I have Iron Hog's Metal Hoof...",
		Prompt:          "You know about that? Not many people realize the potential in the Iron Hog's Metal Hoof... I can make this into something special, if you want me to.",
		SelectionPrompt: r.MetalHoof,
		Choices:         []refine.RefinementChoice{},
	}
}

func (r Chris) RefineUpgradeClaw() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a claw",
		Prompt:          "Ah, you wish to upgrade a claw? Then tell me, which one?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateRefinementChoice("Blood Gigantic#k - Thief Lv. 60", refine.Confirm(item.BloodGigantic, r.BloodGiganticRequirements())),
			r.CreateRefinementChoice("Sapphire Gigantic#k - Thief Lv. 60", refine.Confirm(item.SapphireGigantic, r.SapphireGiganticRequirements())),
			r.CreateRefinementChoice("Dark Gigantic#k - Thief Lv. 60", refine.Confirm(item.DarkGigantic, r.DarkGiganticRequirements())),
		},
	}
}

func (r Chris) CreateRefinementChoice(listText string, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	return refine.RefinementChoice{
		ListText:        listText,
		SelectionPrompt: selectionPrompt,
		Config: refine.TerminalConfig{
			Success:          r.Success,
			MesoError:        r.CannotAfford,
			RequirementError: r.MissingSomething,
			InventoryError:   r.MakeRoom,
		},
	}
}

func (r Chris) MetalHoof(category refine.RefinementCategory) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		return script.SendYesNo(l, c, category.Prompt, r.SpecialRefinement, script.Exit())
	}
}

func (r Chris) SpecialRefinement(l logrus.FieldLogger, c script.Context) script.State {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.HowMany(item.SteelPlate, r.SteelFromHogRequirements())(config)(l, c)
}

func (r Chris) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, c, m.String())
}

func (r Chris) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Cash only, no credit.")
	return script.SendOk(l, c, m.String())
}

func (r Chris) MissingSomething(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText("I cannot accept substitutes. If you don't have what I need, then I won't be able to help you. Please bring more ").ShowItemName1(itemId).AddText(".")
		return script.SendOk(l, c, m.String())
	}
}

func (r Chris) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Phew... I almost didn't think that would work for a second... Well, I hope you enjoy it, anyway.")
	return script.SendOk(l, c, m.String())
}
