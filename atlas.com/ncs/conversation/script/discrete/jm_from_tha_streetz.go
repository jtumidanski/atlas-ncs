package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// JMFromThaStreetz is located in Victoria Road - Kerning City (103000000)
type JMFromThaStreetz struct {
}

func (r JMFromThaStreetz) NPCId() uint32 {
	return npc.JMFromThaStreetz
}

func (r JMFromThaStreetz) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := "Pst... If you have the right goods, I can turn it into something nice..."
	categories := r.CreateCategories()
	return refine.NewGenericRefine(l, c, hello, categories)
}

func (r JMFromThaStreetz) CreateCategories() []refine.RefinementCategory {
	return []refine.RefinementCategory{
		r.CreateAGlove(),
		r.UpgradeAGlove(),
		r.CreateAClaw(),
		r.UpgradeAClaw(),
		r.CreateMaterials(),
	}
}

func (r JMFromThaStreetz) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r JMFromThaStreetz) CreateAGlove() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Create a glove",
		Prompt:          "So, what kind of glove would you like me to make?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.WorkGloves, " - Common Lv. 10"), refine.Confirm(item.WorkGloves, r.WorkGlovesRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownDuo, " - Thief Lv. 15"), refine.Confirm(item.BrownDuo, r.BrownDuoRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueDuo, " - Thief Lv. 15"), refine.Confirm(item.BlueDuo, r.BlueDuoRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackDuo, " - Thief Lv. 15"), refine.Confirm(item.BlackDuo, r.BlackDuoRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeMischief, " - Thief Lv. 20"), refine.Confirm(item.BronzeMischief, r.BronzeMischiefRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeWolfskin, " - Thief Lv. 25"), refine.Confirm(item.BronzeWolfskin, r.BronzeWolfskinRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelSylvia, " - Thief Lv. 30"), refine.Confirm(item.SteelSylvia, r.SteelSylviaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelArbion, " - Thief Lv. 35"), refine.Confirm(item.SteelArbion, r.SteelArbionRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedCleave, " - Thief Lv. 40"), refine.Confirm(item.RedCleave, r.RedCleaveRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueMoonGlove, " - Thief Lv. 50"), refine.Confirm(item.BlueMoonGlove, r.BlueMoonGloveRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzePow, " - Thief Lv. 60"), refine.Confirm(item.BronzePow, r.BronzePowRequirements())),
		},
	}
}

func (r JMFromThaStreetz) UpgradeAGlove() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a glove",
		Prompt:          "An upgraded glove? Sure thing, but note that upgrades won't carry over to the new item... ",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilMischief, " - Thief Lv. 20"), refine.Confirm(item.MithrilMischief, r.MithrilMischiefRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkMischief, " - Thief Lv. 20"), refine.Confirm(item.DarkMischief, r.DarkMischiefRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilWolfskin, " - Thief Lv. 25"), refine.Confirm(item.MithrilWolfskin, r.MithrilWolfskinRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkWolfskin, " - Thief Lv. 25"), refine.Confirm(item.DarkWolfskin, r.DarkWolfskinRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverSylvia, " - Thief Lv. 30"), refine.Confirm(item.SilverSylvia, r.SilverSylviaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldSylvia, " - Thief Lv. 30"), refine.Confirm(item.GoldSylvia, r.GoldSylviaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrihalconArbion, " - Thief Lv. 35"), refine.Confirm(item.OrihalconArbion, r.OrihalconArbionRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldArbion, " - Thief Lv. 35"), refine.Confirm(item.GoldArbion, r.GoldArbionRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldCleave, " - Thief Lv. 40"), refine.Confirm(item.GoldCleave, r.GoldCleaveRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkCleave, " - Thief Lv. 40"), refine.Confirm(item.DarkCleave, r.DarkCleaveRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedMoonGlove, " - Thief Lv. 50"), refine.Confirm(item.RedMoonGlove, r.RedMoonGloveRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownMoonGlove, " - Thief Lv. 50"), refine.Confirm(item.BrownMoonGlove, r.BrownMoonGloveRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.StealPow, " - Thief Lv. 60"), refine.Confirm(item.StealPow, r.SilverPowRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldPow, " - Thief Lv. 60"), refine.Confirm(item.GoldPow, r.GoldPowRequirements())),
		},
	}
}

func (r JMFromThaStreetz) CreateAClaw() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Create a claw",
		Prompt:          "So, what kind of claw would you like me to make?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelTitans, " - Thief Lv. 15"), refine.Confirm(item.SteelTitans, r.SteelTitansRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeIgor, " - Thief Lv. 20"), refine.Confirm(item.BronzeIgor, r.BronzeIgorRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.Meba, " - Thief Lv. 25"), refine.Confirm(item.Meba, r.MebaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelGuards, " - Thief Lv. 30"), refine.Confirm(item.SteelGuards, r.SteelGuardsRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeGuardian, " - Thief Lv. 35"), refine.Confirm(item.BronzeGuardian, r.BronzeGuardianRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelAvarice, " - Thief Lv. 40"), refine.Confirm(item.SteelAvarice, r.SteelAvariceRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelSlain, " - Thief Lv. 50"), refine.Confirm(item.SteelSlain, r.SteelSlainRequirements())),
		},
	}
}

func (r JMFromThaStreetz) UpgradeAClaw() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a claw",
		Prompt:          "An upgraded claw? Sure thing, but note that upgrades won't carry over to the new item...",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilTitans, " - Thief Lv. 15"), refine.Confirm(item.MithrilTitans, r.MithrilTitansRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldTitans, " - Thief Lv. 15"), refine.Confirm(item.GoldTitans, r.GoldTitansRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelIgor, " - Thief Lv. 20"), refine.Confirm(item.SteelIgor, r.SteelIgorRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.AdamantiumIgor, " - Thief Lv. 20"), refine.Confirm(item.AdamantiumIgor, r.AdamantiumIgorRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilGuards, " - Thief Lv. 30"), refine.Confirm(item.MithrilGuards, r.MithrilGuardsRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.AdamantiumGuards, " - Thief Lv. 3"), refine.Confirm(item.AdamantiumGuards, r.AdamantiumGuardsRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverGuardian, " - Thief Lv. 35"), refine.Confirm(item.SilverGuardian, r.SilverGuardianRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkGuardian, " - Thief Lv. 35"), refine.Confirm(item.DarkGuardian, r.DarkGuardianRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodAvarice, " - Thief Lv. 40"), refine.Confirm(item.BloodAvarice, r.BloodAvariceRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.AdamantiumAvarice, " - Thief Lv. 40"), refine.Confirm(item.AdamantiumAvarice, r.AdamantiumAvariceRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkAvarice, " - Thief Lv. 40"), refine.Confirm(item.DarkAvarice, r.DarkAvariceRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodSlain, " - Thief Lv. 50"), refine.Confirm(item.BloodSlain, r.BloodSlainRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SapphireSlain, " - Thief Lv. 5"), refine.Confirm(item.SapphireSlain, r.SapphireSlainRequirements())),
		},
	}
}

func (r JMFromThaStreetz) CreateMaterials() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Create materials",
		Prompt:          "Materials? I know of a few materials that I can make for you...",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.SimpleList("Make Processed Wood with Tree Branch"), refine.HowMany(item.ProcessedWood, r.ProcessedWoodFromTreeBranchRequirements())),
			r.CreateChoice(refine.SimpleList("Make Processed Wood with Firewood"), refine.HowMany(item.ProcessedWood, r.ProcessedWoodFromFirewoodRequirements())),
			r.CreateChoice(refine.SimpleList("Make Screws (packs of 15)"), refine.HowMany(item.Screw, r.ScrewRequirements())),
		},
	}
}

func (r JMFromThaStreetz) ProcessedWoodFromTreeBranchRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.TreeBranch, Amount: 10}})
}

func (r JMFromThaStreetz) ProcessedWoodFromFirewoodRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Firewood, Amount: 5}})
}

func (r JMFromThaStreetz) ScrewRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}}, refine.SetAwardAmount(15))
}

func (r JMFromThaStreetz) WorkGlovesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 15}}, refine.SetCost(1000))
}

func (r JMFromThaStreetz) BrownDuoRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 30}, {ItemId: item.Firewood, Amount: 20}}, refine.SetCost(7000))
}

func (r JMFromThaStreetz) BlueDuoRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 30}, {ItemId: item.HornyMushroomCap, Amount: 20}}, refine.SetCost(7000))
}

func (r JMFromThaStreetz) BlackDuoRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 30}, {ItemId: item.WildBoarTooth, Amount: 20}}, refine.SetCost(7000))
}

func (r JMFromThaStreetz) BronzeMischiefRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 2}, {ItemId: item.Leather, Amount: 40}}, refine.SetCost(10000))
}

func (r JMFromThaStreetz) BronzeWolfskinRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 2}, {ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Leather, Amount: 10}}, refine.SetCost(15000))
}

func (r JMFromThaStreetz) SteelSylviaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Leather, Amount: 50}, {ItemId: item.Screw, Amount: 10}}, refine.SetCost(25000))
}

func (r JMFromThaStreetz) SteelArbionRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 3}, {ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.Leather, Amount: 60}, {ItemId: item.Screw, Amount: 15}}, refine.SetCost(30000))
}

func (r JMFromThaStreetz) RedCleaveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Garnet, Amount: 3}, {ItemId: item.DrakeSkull, Amount: 200}, {ItemId: item.Leather, Amount: 80}, {ItemId: item.Screw, Amount: 30}}, refine.SetCost(40000))
}

func (r JMFromThaStreetz) BlueMoonGloveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Sapphire, Amount: 3}, {ItemId: item.BlackCrystal, Amount: 1}, {ItemId: item.DragonSkin, Amount: 40}, {ItemId: item.Screw, Amount: 30}}, refine.SetCost(50000))
}

func (r JMFromThaStreetz) BronzePowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.MoonRock, Amount: 1}, {ItemId: item.BronzePlate, Amount: 8}, {ItemId: item.Diamond, Amount: 1}, {ItemId: item.DragonSkin, Amount: 50}, {ItemId: item.Screw, Amount: 50}}, refine.SetCost(70000))
}

func (r JMFromThaStreetz) MithrilMischiefRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeMischief, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 1}}, refine.SetCost(5000))
}

func (r JMFromThaStreetz) DarkMischiefRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeMischief, Amount: 1}, {ItemId: item.Opal, Amount: 1}}, refine.SetCost(7000))
}

func (r JMFromThaStreetz) MithrilWolfskinRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeWolfskin, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 2}}, refine.SetCost(10000))
}

func (r JMFromThaStreetz) DarkWolfskinRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeWolfskin, Amount: 1}, {ItemId: item.Opal, Amount: 2}}, refine.SetCost(12000))
}

func (r JMFromThaStreetz) SilverSylviaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelSylvia, Amount: 1}, {ItemId: item.SilverPlate, Amount: 2}}, refine.SetCost(15000))
}

func (r JMFromThaStreetz) GoldSylviaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelSylvia, Amount: 1}, {ItemId: item.GoldPlate, Amount: 1}}, refine.SetCost(20000))
}

func (r JMFromThaStreetz) OrihalconArbionRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelArbion, Amount: 1}, {ItemId: item.OrihalconPlate, Amount: 3}}, refine.SetCost(22000))
}

func (r JMFromThaStreetz) GoldArbionRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelArbion, Amount: 1}, {ItemId: item.GoldPlate, Amount: 2}}, refine.SetCost(25000))
}

func (r JMFromThaStreetz) GoldCleaveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedCleave, Amount: 1}, {ItemId: item.GoldPlate, Amount: 4}}, refine.SetCost(40000))
}

func (r JMFromThaStreetz) DarkCleaveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedCleave, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 2}}, refine.SetCost(50000))
}

func (r JMFromThaStreetz) RedMoonGloveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueMoonGlove, Amount: 1}, {ItemId: item.Garnet, Amount: 5}}, refine.SetCost(55000))
}

func (r JMFromThaStreetz) BrownMoonGloveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueMoonGlove, Amount: 1}, {ItemId: item.GoldPlate, Amount: 2}, {ItemId: item.BlackCrystal, Amount: 1}}, refine.SetCost(60000))
}

func (r JMFromThaStreetz) SilverPowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.StealPow, Amount: 1}, {ItemId: item.SteelPlate, Amount: 7}, {ItemId: item.DrakeSkull, Amount: 200}}, refine.SetCost(70000))
}

func (r JMFromThaStreetz) GoldPowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.StealPow, Amount: 1}, {ItemId: item.GoldPlate, Amount: 7}, {ItemId: item.WildKargoEye, Amount: 150}}, refine.SetCost(80000))
}

func (r JMFromThaStreetz) SteelTitansRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Leather, Amount: 20}, {ItemId: item.Screw, Amount: 5}}, refine.SetCost(2000))
}

func (r JMFromThaStreetz) BronzeIgorRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 2}, {ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Leather, Amount: 30}, {ItemId: item.Screw, Amount: 10}}, refine.SetCost(3000))
}

func (r JMFromThaStreetz) MebaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Garnier, Amount: 1}, {ItemId: item.SteelPlate, Amount: 3}, {ItemId: item.Leather, Amount: 20}, {ItemId: item.ProcessedWood, Amount: 30}}, refine.SetCost(5000))
}

func (r JMFromThaStreetz) SteelGuardsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 3}, {ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Leather, Amount: 50}, {ItemId: item.Screw, Amount: 20}}, refine.SetCost(15000))
}

func (r JMFromThaStreetz) BronzeGuardianRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 4}, {ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Leather, Amount: 80}, {ItemId: item.Screw, Amount: 25}}, refine.SetCost(30000))
}

func (r JMFromThaStreetz) SteelAvariceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 3}, {ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Leather, Amount: 100}, {ItemId: item.Screw, Amount: 30}}, refine.SetCost(40000))
}

func (r JMFromThaStreetz) SteelSlainRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 4}, {ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.DragonSkin, Amount: 40}, {ItemId: item.Screw, Amount: 35}}, refine.SetCost(50000))
}

func (r JMFromThaStreetz) MithrilTitansRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelTitans, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 1}}, refine.SetCost(1000))
}

func (r JMFromThaStreetz) GoldTitansRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelTitans, Amount: 1}, {ItemId: item.GoldPlate, Amount: 1}}, refine.SetCost(2000))
}

func (r JMFromThaStreetz) SteelIgorRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeIgor, Amount: 1}, {ItemId: item.SteelPlate, Amount: 2}}, refine.SetCost(3000))
}

func (r JMFromThaStreetz) AdamantiumIgorRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeIgor, Amount: 1}, {ItemId: item.AdamantiumPlate, Amount: 2}}, refine.SetCost(5000))
}

func (r JMFromThaStreetz) MithrilGuardsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelGuards, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 3}}, refine.SetCost(10000))
}

func (r JMFromThaStreetz) AdamantiumGuardsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelGuards, Amount: 1}, {ItemId: item.AdamantiumPlate, Amount: 3}}, refine.SetCost(15000))
}

func (r JMFromThaStreetz) SilverGuardianRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeGuardian, Amount: 1}, {ItemId: item.SilverPlate, Amount: 4}}, refine.SetCost(20000))
}

func (r JMFromThaStreetz) DarkGuardianRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeGuardian, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 1}}, refine.SetCost(25000))
}

func (r JMFromThaStreetz) BloodAvariceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelAvarice, Amount: 1}, {ItemId: item.Garnet, Amount: 5}}, refine.SetCost(30000))
}

func (r JMFromThaStreetz) AdamantiumAvariceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelAvarice, Amount: 1}, {ItemId: item.AdamantiumPlate, Amount: 5}}, refine.SetCost(30000))
}

func (r JMFromThaStreetz) DarkAvariceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelAvarice, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 2}}, refine.SetCost(35000))
}

func (r JMFromThaStreetz) BloodSlainRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelSlain, Amount: 1}, {ItemId: item.Garnet, Amount: 6}}, refine.SetCost(40000))
}

func (r JMFromThaStreetz) SapphireSlainRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelSlain, Amount: 1}, {ItemId: item.Sapphire, Amount: 6}}, refine.SetCost(40000))
}

func (r JMFromThaStreetz) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, c, m.String())
}

func (r JMFromThaStreetz) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm afraid you cannot afford my services.")
	return script.SendOk(l, c, m.String())
}

func (r JMFromThaStreetz) MissingSomething(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("What are you trying to pull? I can't make anything unless you bring me what I ask for. Can you please bring more ").
			ShowItemName1(itemId).
			AddText("?")
		return script.SendOk(l, c, m.String())
	}
}

func (r JMFromThaStreetz) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("All done. If you need anything else... Well, I'm not going anywhere.")
	return script.SendOk(l, c, m.String())
}
