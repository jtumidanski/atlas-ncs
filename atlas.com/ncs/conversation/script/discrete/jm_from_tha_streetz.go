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
			r.CreateChoice(refine.ItemNameList("Work Gloves", " - Common Lv. 10"), refine.Confirm(item.WorkGloves, r.WorkGlovesRequirements())),
			r.CreateChoice(refine.ItemNameList("Brown Duo", " - Thief Lv. 15"), refine.Confirm(item.BrownDuo, r.BrownDuoRequirements())),
			r.CreateChoice(refine.ItemNameList("Blue Duo", " - Thief Lv. 15"), refine.Confirm(item.BlueDuo, r.BlueDuoRequirements())),
			r.CreateChoice(refine.ItemNameList("Black Duo", " - Thief Lv. 15"), refine.Confirm(item.BlackDuo, r.BlackDuoRequirements())),
			r.CreateChoice(refine.ItemNameList("Bronze Mischief", " - Thief Lv. 20"), refine.Confirm(item.BronzeMischief, r.BronzeMischiefRequirements())),
			r.CreateChoice(refine.ItemNameList("Bronze Wolfskin", " - Thief Lv. 25"), refine.Confirm(item.BronzeWolfskin, r.BronzeWolfskinRequirements())),
			r.CreateChoice(refine.ItemNameList("Steel Sylvia", " - Thief Lv. 30"), refine.Confirm(item.SteelSylvia, r.SteelSylviaRequirements())),
			r.CreateChoice(refine.ItemNameList("Steel Arbion", " - Thief Lv. 35"), refine.Confirm(item.SteelArbion, r.SteelArbionRequirements())),
			r.CreateChoice(refine.ItemNameList("Red Cleave", " - Thief Lv. 40"), refine.Confirm(item.RedCleave, r.RedCleaveRequirements())),
			r.CreateChoice(refine.ItemNameList("Blue Moon Glove", " - Thief Lv. 50"), refine.Confirm(item.BlueMoonGlove, r.BlueMoonGloveRequirements())),
			r.CreateChoice(refine.ItemNameList("Bronze Pow", " - Thief Lv. 60"), refine.Confirm(item.BronzePow, r.BronzePowRequirements())),
		},
	}
}

func (r JMFromThaStreetz) UpgradeAGlove() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a glove",
		Prompt:          "An upgraded glove? Sure thing, but note that upgrades won't carry over to the new item... ",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemNameList("Mithril Mischief", " - Thief Lv. 20"), refine.Confirm(item.MithrilMischief, r.MithrilMischiefRequirements())),
			r.CreateChoice(refine.ItemNameList("Dark Mischief", " - Thief Lv. 20"), refine.Confirm(item.DarkMischief, r.DarkMischiefRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Wolfskin", " - Thief Lv. 25"), refine.Confirm(item.MithrilWolfskin, r.MithrilWolfskinRequirements())),
			r.CreateChoice(refine.ItemNameList("Dark Wolfskin", " - Thief Lv. 25"), refine.Confirm(item.DarkWolfskin, r.DarkWolfskinRequirements())),
			r.CreateChoice(refine.ItemNameList("Silver Sylvia", " - Thief Lv. 30"), refine.Confirm(item.SilverSylvia, r.SilverSylviaRequirements())),
			r.CreateChoice(refine.ItemNameList("Gold Sylvia", " - Thief Lv. 30"), refine.Confirm(item.GoldSylvia, r.GoldSylviaRequirements())),
			r.CreateChoice(refine.ItemNameList("Orihalcon Arbion", " - Thief Lv. 35"), refine.Confirm(item.OrihalconArbion, r.OrihalconArbionRequirements())),
			r.CreateChoice(refine.ItemNameList("Gold Arbion", " - Thief Lv. 35"), refine.Confirm(item.GoldArbion, r.GoldArbionRequirements())),
			r.CreateChoice(refine.ItemNameList("Gold Cleave", " - Thief Lv. 40"), refine.Confirm(item.GoldCleave, r.GoldCleaveRequirements())),
			r.CreateChoice(refine.ItemNameList("Dark Cleave", " - Thief Lv. 40"), refine.Confirm(item.DarkCleave, r.DarkCleaveRequirements())),
			r.CreateChoice(refine.ItemNameList("Red Moon Glove", " - Thief Lv. 50"), refine.Confirm(item.RedMoonGlove, r.RedMoonGloveRequirements())),
			r.CreateChoice(refine.ItemNameList("Brown Moon Glove", " - Thief Lv. 50"), refine.Confirm(item.BrownMoonGlove, r.BrownMoonGloveRequirements())),
			r.CreateChoice(refine.ItemNameList("Steal Pow", " - Thief Lv. 60"), refine.Confirm(item.StealPow, r.SilverPowRequirements())),
			r.CreateChoice(refine.ItemNameList("Gold Pow", " - Thief Lv. 60"), refine.Confirm(item.GoldPow, r.GoldPowRequirements())),
		},
	}
}

func (r JMFromThaStreetz) CreateAClaw() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Create a claw",
		Prompt:          "So, what kind of claw would you like me to make?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemNameList("Steel Titans", " - Thief Lv. 15"), refine.Confirm(item.SteelTitans, r.SteelTitansRequirements())),
			r.CreateChoice(refine.ItemNameList("Bronze Igor", " - Thief Lv. 20"), refine.Confirm(item.BronzeIgor, r.BronzeIgorRequirements())),
			r.CreateChoice(refine.ItemNameList("Meba", " - Thief Lv. 25"), refine.Confirm(item.Meba, r.MebaRequirements())),
			r.CreateChoice(refine.ItemNameList("Steel Guards", " - Thief Lv. 30"), refine.Confirm(item.SteelGuards, r.SteelGuardsRequirements())),
			r.CreateChoice(refine.ItemNameList("Bronze Guardian", " - Thief Lv. 35"), refine.Confirm(item.BronzeGuardian, r.BronzeGuardianRequirements())),
			r.CreateChoice(refine.ItemNameList("Steel Avarice", " - Thief Lv. 40"), refine.Confirm(item.SteelAvarice, r.SteelAvariceRequirements())),
			r.CreateChoice(refine.ItemNameList("Steel Slain", " - Thief Lv. 50"), refine.Confirm(item.SteelSlain, r.SteelSlainRequirements())),
		},
	}
}

func (r JMFromThaStreetz) UpgradeAClaw() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a claw",
		Prompt:          "An upgraded claw? Sure thing, but note that upgrades won't carry over to the new item...",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemNameList("Mithril Titans", " - Thief Lv. 15"), refine.Confirm(item.MithrilTitans, r.MithrilTitansRequirements())),
			r.CreateChoice(refine.ItemNameList("Gold Titans", " - Thief Lv. 15"), refine.Confirm(item.GoldTitans, r.GoldTitansRequirements())),
			r.CreateChoice(refine.ItemNameList("Steel Igor", " - Thief Lv. 20"), refine.Confirm(item.SteelIgor, r.SteelIgorRequirements())),
			r.CreateChoice(refine.ItemNameList("Adamantium Igor", " - Thief Lv. 20"), refine.Confirm(item.AdamantiumIgor, r.AdamantiumIgorRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Guards", " - Thief Lv. 30"), refine.Confirm(item.MithrilGuards, r.MithrilGuardsRequirements())),
			r.CreateChoice(refine.ItemNameList("Adamantium Guards", " - Thief Lv. 3"), refine.Confirm(item.AdamantiumGuards, r.AdamantiumGuardsRequirements())),
			r.CreateChoice(refine.ItemNameList("Silver Guardian", " - Thief Lv. 35"), refine.Confirm(item.SilverGuardian, r.SilverGuardianRequirements())),
			r.CreateChoice(refine.ItemNameList("Dark Guardian", " - Thief Lv. 35"), refine.Confirm(item.DarkGuardian, r.DarkGuardianRequirements())),
			r.CreateChoice(refine.ItemNameList("Blood Avarice", " - Thief Lv. 40"), refine.Confirm(item.BloodAvarice, r.BloodAvariceRequirements())),
			r.CreateChoice(refine.ItemNameList("Adamantium Avarice", " - Thief Lv. 40"), refine.Confirm(item.AdamantiumAvarice, r.AdamantiumAvariceRequirements())),
			r.CreateChoice(refine.ItemNameList("Dark Avarice", " - Thief Lv. 40"), refine.Confirm(item.DarkAvarice, r.DarkAvariceRequirements())),
			r.CreateChoice(refine.ItemNameList("Blood Slain", " - Thief Lv. 50"), refine.Confirm(item.BloodSlain, r.BloodSlainRequirements())),
			r.CreateChoice(refine.ItemNameList("Sapphire Slain", " - Thief Lv. 5"), refine.Confirm(item.SapphireSlain, r.SapphireSlainRequirements())),
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

func (r JMFromThaStreetz) ProcessedWoodFromTreeBranchRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.TreeBranch, Amount: 10}}, Cost: 0}
}

func (r JMFromThaStreetz) ProcessedWoodFromFirewoodRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Firewood, Amount: 5}}, Cost: 0}
}

func (r JMFromThaStreetz) ScrewRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}}, Cost: 0}
}

func (r JMFromThaStreetz) WorkGlovesRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Leather, Amount: 15}}, Cost: 1000}
}

func (r JMFromThaStreetz) BrownDuoRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Leather, Amount: 30}, {ItemId: item.Firewood, Amount: 20}}, Cost: 7000}
}

func (r JMFromThaStreetz) BlueDuoRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Leather, Amount: 30}, {ItemId: item.HornyMushroomCap, Amount: 20}}, Cost: 7000}
}

func (r JMFromThaStreetz) BlackDuoRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Leather, Amount: 30}, {ItemId: item.WildBoarTooth, Amount: 20}}, Cost: 7000}
}

func (r JMFromThaStreetz) BronzeMischiefRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 2}, {ItemId: item.Leather, Amount: 40}}, Cost: 10000}
}

func (r JMFromThaStreetz) BronzeWolfskinRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 2}, {ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Leather, Amount: 10}}, Cost: 15000}
}

func (r JMFromThaStreetz) SteelSylviaRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Leather, Amount: 50}, {ItemId: item.Screw, Amount: 10}}, Cost: 25000}
}

func (r JMFromThaStreetz) SteelArbionRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelPlate, Amount: 3}, {ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.Leather, Amount: 60}, {ItemId: item.Screw, Amount: 15}}, Cost: 30000}
}

func (r JMFromThaStreetz) RedCleaveRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Garnet, Amount: 3}, {ItemId: item.DrakeSkull, Amount: 200}, {ItemId: item.Leather, Amount: 80}, {ItemId: item.Screw, Amount: 30}}, Cost: 40000}
}

func (r JMFromThaStreetz) BlueMoonGloveRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Sapphire, Amount: 3}, {ItemId: item.BlackCrystal, Amount: 1}, {ItemId: item.DragonSkin, Amount: 40}, {ItemId: item.Screw, Amount: 30}}, Cost: 50000}
}

func (r JMFromThaStreetz) BronzePowRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.MoonRock, Amount: 1}, {ItemId: item.BronzePlate, Amount: 8}, {ItemId: item.Diamond, Amount: 1}, {ItemId: item.DragonSkin, Amount: 50}, {ItemId: item.Screw, Amount: 50}}, Cost: 70000}
}

func (r JMFromThaStreetz) MithrilMischiefRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeMischief, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 1}}, Cost: 5000}
}

func (r JMFromThaStreetz) DarkMischiefRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeMischief, Amount: 1}, {ItemId: item.Opal, Amount: 1}}, Cost: 7000}
}

func (r JMFromThaStreetz) MithrilWolfskinRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeWolfskin, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 2}}, Cost: 10000}
}

func (r JMFromThaStreetz) DarkWolfskinRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeWolfskin, Amount: 1}, {ItemId: item.Opal, Amount: 2}}, Cost: 12000}
}

func (r JMFromThaStreetz) SilverSylviaRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelSylvia, Amount: 1}, {ItemId: item.SilverPlate, Amount: 2}}, Cost: 15000}
}

func (r JMFromThaStreetz) GoldSylviaRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelSylvia, Amount: 1}, {ItemId: item.GoldPlate, Amount: 1}}, Cost: 20000}
}

func (r JMFromThaStreetz) OrihalconArbionRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelArbion, Amount: 1}, {ItemId: item.OrihalconPlate, Amount: 3}}, Cost: 22000}
}

func (r JMFromThaStreetz) GoldArbionRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelArbion, Amount: 1}, {ItemId: item.GoldPlate, Amount: 2}}, Cost: 25000}
}

func (r JMFromThaStreetz) GoldCleaveRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.RedCleave, Amount: 1}, {ItemId: item.GoldPlate, Amount: 4}}, Cost: 40000}
}

func (r JMFromThaStreetz) DarkCleaveRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.RedCleave, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 2}}, Cost: 50000}
}

func (r JMFromThaStreetz) RedMoonGloveRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BlueMoonGlove, Amount: 1}, {ItemId: item.Garnet, Amount: 5}}, Cost: 55000}
}

func (r JMFromThaStreetz) BrownMoonGloveRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BlueMoonGlove, Amount: 1}, {ItemId: item.GoldPlate, Amount: 2}, {ItemId: item.BlackCrystal, Amount: 1}}, Cost: 60000}
}

func (r JMFromThaStreetz) SilverPowRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.StealPow, Amount: 1}, {ItemId: item.SteelPlate, Amount: 7}, {ItemId: item.DrakeSkull, Amount: 200}}, Cost: 70000}
}

func (r JMFromThaStreetz) GoldPowRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.StealPow, Amount: 1}, {ItemId: item.GoldPlate, Amount: 7}, {ItemId: item.WildKargoEye, Amount: 150}}, Cost: 80000}
}

func (r JMFromThaStreetz) SteelTitansRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Leather, Amount: 20}, {ItemId: item.Screw, Amount: 5}}, Cost: 2000}
}

func (r JMFromThaStreetz) BronzeIgorRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 2}, {ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Leather, Amount: 30}, {ItemId: item.Screw, Amount: 10}}, Cost: 3000}
}

func (r JMFromThaStreetz) MebaRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Garnier, Amount: 1}, {ItemId: item.SteelPlate, Amount: 3}, {ItemId: item.Leather, Amount: 20}, {ItemId: item.ProcessedWood, Amount: 30}}, Cost: 5000}
}

func (r JMFromThaStreetz) SteelGuardsRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 3}, {ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Leather, Amount: 50}, {ItemId: item.Screw, Amount: 20}}, Cost: 15000}
}

func (r JMFromThaStreetz) BronzeGuardianRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 4}, {ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Leather, Amount: 80}, {ItemId: item.Screw, Amount: 25}}, Cost: 30000}
}

func (r JMFromThaStreetz) SteelAvariceRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 3}, {ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Leather, Amount: 100}, {ItemId: item.Screw, Amount: 30}}, Cost: 40000}
}

func (r JMFromThaStreetz) SteelSlainRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 4}, {ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.DragonSkin, Amount: 40}, {ItemId: item.Screw, Amount: 35}}, Cost: 50000}
}

func (r JMFromThaStreetz) MithrilTitansRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelTitans, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 1}}, Cost: 1000}
}

func (r JMFromThaStreetz) GoldTitansRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelTitans, Amount: 1}, {ItemId: item.GoldPlate, Amount: 1}}, Cost: 2000}
}

func (r JMFromThaStreetz) SteelIgorRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeIgor, Amount: 1}, {ItemId: item.SteelPlate, Amount: 2}}, Cost: 3000}
}

func (r JMFromThaStreetz) AdamantiumIgorRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeIgor, Amount: 1}, {ItemId: item.AdamantiumPlate, Amount: 2}}, Cost: 5000}
}

func (r JMFromThaStreetz) MithrilGuardsRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelGuards, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 3}}, Cost: 10000}
}

func (r JMFromThaStreetz) AdamantiumGuardsRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelGuards, Amount: 1}, {ItemId: item.AdamantiumPlate, Amount: 3}}, Cost: 15000}
}

func (r JMFromThaStreetz) SilverGuardianRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeGuardian, Amount: 1}, {ItemId: item.SilverPlate, Amount: 4}}, Cost: 20000}
}

func (r JMFromThaStreetz) DarkGuardianRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeGuardian, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 1}}, Cost: 25000}
}

func (r JMFromThaStreetz) BloodAvariceRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelAvarice, Amount: 1}, {ItemId: item.Garnet, Amount: 5}}, Cost: 30000}
}

func (r JMFromThaStreetz) AdamantiumAvariceRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelAvarice, Amount: 1}, {ItemId: item.AdamantiumPlate, Amount: 5}}, Cost: 30000}
}

func (r JMFromThaStreetz) DarkAvariceRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelAvarice, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 2}}, Cost: 35000}
}

func (r JMFromThaStreetz) BloodSlainRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelSlain, Amount: 1}, {ItemId: item.Garnet, Amount: 6}}, Cost: 40000}
}

func (r JMFromThaStreetz) SapphireSlainRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelSlain, Amount: 1}, {ItemId: item.Sapphire, Amount: 6}}, Cost: 40000}
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
