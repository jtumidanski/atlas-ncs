package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Vicious is located in Victoria Road - Henesys Market (100000100)
type Vicious struct {
}

func (r Vicious) NPCId() uint32 {
	return npc.Vicious
}

func (r Vicious) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := "Hello. I am Vicious, retired Sniper. However, I used to be the top student of Athena Pierce. Though I no longer hunt, I can make some archer items that will be useful for you..."
	categories := r.CreateCategories()
	return refine.NewGenericRefine(l, c, hello, categories)
}

func (r Vicious) CreateCategories() []refine.RefinementCategory {
	return []refine.RefinementCategory{
		r.CreateABow(),
		r.CreateACrossbow(),
		r.MakeAGlove(),
		r.UpgradeAGlove(),
		r.CreateMaterials(),
		r.CreateArrows(),
	}
}

func (r Vicious) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.Sorry,
		RequirementError: r.NeedMoreItems,
		InventoryError:   r.NotEnoughInventorySpace,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Vicious) CreateABow() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Create a bow",
		Prompt:          "I may have been a Sniper, but bows and crossbows aren't too much different. Anyway, which would you like to make?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdList(item.WarBow, " - Bowman Lv. 10"), refine.Confirm(item.WarBow, r.WarBowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.CompositeBow, " - Bowman Lv. 15"), refine.Confirm(item.CompositeBow, r.CompositeBowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.HuntersBow, " - Bowman Lv. 20"), refine.Confirm(item.HuntersBow, r.HuntersBowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BattleBow, " - Bowman Lv. 25"), refine.Confirm(item.BattleBow, r.BattleBowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.Ryden, " - Bowman Lv. 30"), refine.Confirm(item.Ryden, r.RydenRequirements())),
			r.CreateChoice(refine.ItemIdList(item.RedViper, " - Bowman Lv. 35"), refine.Confirm(item.RedViper, r.RedViperRequirements())),
			r.CreateChoice(refine.ItemIdList(item.Vaulter2000, " - Bowman Lv. 40"), refine.Confirm(item.Vaulter2000, r.Vaulter2000Requirements())),
		},
	}
}

func (r Vicious) CreateACrossbow() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Create a crossbow",
		Prompt:          "I was a Sniper. Crossbows are my specialty. Which would you like me to make for you?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdList(item.Crossbow, " - Bowman Lv. 10"), refine.Confirm(item.Crossbow, r.CrossbowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BattleCrossbow, " - Bowman Lv. 15"), refine.Confirm(item.BattleCrossbow, r.BattleCrossbowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.Balanche, " - Bowman Lv. 20"), refine.Confirm(item.Balanche, r.BalancheRequirements())),
			r.CreateChoice(refine.ItemIdList(item.MountainCrossbow, " - Bowman Lv. 25"), refine.Confirm(item.MountainCrossbow, r.MountainCrossbowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.EagleCrow, " - Bowman Lv. 30"), refine.Confirm(item.EagleCrow, r.EagleCrowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.Heckler, " - Bowman Lv. 35"), refine.Confirm(item.Heckler, r.HecklerRequirements())),
			r.CreateChoice(refine.ItemIdList(item.SilverCrow, " - Bowman Lv. 40"), refine.Confirm(item.SilverCrow, r.SilverCrowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.Rower, " - Bowman Lv. 45"), refine.Confirm(item.Rower, r.RowerRequirements())),
		},
	}
}

func (r Vicious) MakeAGlove() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Make a glove",
		Prompt:          "Okay, so which glove do you want me to make?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdList(item.BasicArcherGloves, " - Bowman Lv. 10"), refine.Confirm(item.BasicArcherGloves, r.BasicArcherGlovesRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BlueDiros, " - Bowman Lv. 15"), refine.Confirm(item.BlueDiros, r.BlueDirosRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BlueSavata, " - Bowman Lv. 25"), refine.Confirm(item.BlueSavata, r.BlueSavataRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BrownMarker, " - Bowman Lv. 30"), refine.Confirm(item.BrownMarker, r.BrownMarkerRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BronzeScaler, " - Bowman Lv. 35"), refine.Confirm(item.BronzeScaler, r.BronzeScalerRequirements())),
			r.CreateChoice(refine.ItemIdList(item.AquaBrace, " - Bowman Lv. 40"), refine.Confirm(item.AquaBrace, r.AquaBraceRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BlueWillow, " - Bowman Lv. 50"), refine.Confirm(item.BlueWillow, r.BlueWillowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.OakerGarner, " - Bowman Lv. 60"), refine.Confirm(item.OakerGarner, r.OakerGarnerRequirements())),
		},
	}
}

func (r Vicious) UpgradeAGlove() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a glove",
		Prompt:          "Upgrade a glove? That shouldn't be too difficult. Which did you have in mind?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdList(item.GreenDiros, " - Bowman Lv. 15"), refine.Confirm(item.GreenDiros, r.GreenDirosRequirements())),
			r.CreateChoice(refine.ItemIdList(item.RedDiros, " - Bowman Lv. 15"), refine.Confirm(item.RedDiros, r.RedDirosRequirements())),
			r.CreateChoice(refine.ItemIdList(item.RedSavata, " - Bowman Lv. 25"), refine.Confirm(item.RedSavata, r.RedSavataRequirements())),
			r.CreateChoice(refine.ItemIdList(item.DarkSavata, " - Bowman Lv. 25"), refine.Confirm(item.DarkSavata, r.DarkSavataRequirements())),
			r.CreateChoice(refine.ItemIdList(item.GreenMarker, " - Bowman Lv. 30"), refine.Confirm(item.GreenMarker, r.GreenMarkerRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BlackMarker, " - Bowman Lv. 30"), refine.Confirm(item.BlackMarker, r.BlackMarkerRequirements())),
			r.CreateChoice(refine.ItemIdList(item.MithrilScaler, " - Bowman Lv. 35"), refine.Confirm(item.MithrilScaler, r.MithrilScalerRequirements())),
			r.CreateChoice(refine.ItemIdList(item.GoldScaler, " - Bowman Lv. 35"), refine.Confirm(item.GoldScaler, r.GoldScalerRequirements())),
			r.CreateChoice(refine.ItemIdList(item.GoldBrace, " - Bowman Lv. 40"), refine.Confirm(item.GoldBrace, r.GoldBraceRequirements())),
			r.CreateChoice(refine.ItemIdList(item.DarkBrace, " - Bowman Lv. 40"), refine.Confirm(item.DarkBrace, r.DarkBraceRequirements())),
			r.CreateChoice(refine.ItemIdList(item.RedWillow, " - Bowman Lv. 50"), refine.Confirm(item.RedWillow, r.RedWillowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.DarkWillow, " - Bowman Lv. 50"), refine.Confirm(item.DarkWillow, r.DarkWillowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.SephiaGarner, " - Bowman Lv. 60"), refine.Confirm(item.SephiaGarner, r.SephiaGarnerRequirements())),
			r.CreateChoice(refine.ItemIdList(item.DarkGarner, " - Bowman Lv. 60"), refine.Confirm(item.DarkGarner, r.DarkGarnerRequirements())),
		},
	}
}

func (r Vicious) CreateMaterials() refine.RefinementCategory {
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

func (r Vicious) CreateArrows() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Create arrows",
		Prompt:          "Arrows? Not a problem at all.",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdList(item.ArrowForBow, ""), refine.Confirm(item.ArrowForBow, r.ArrowForBowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.ArrowForCrossbow, ""), refine.Confirm(item.ArrowForCrossbow, r.ArrowForCrossbowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BronzeArrowForBow, ""), refine.Confirm(item.BronzeArrowForBow, r.BronzeArrowForBowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.BronzeArrowForCrossbow, ""), refine.Confirm(item.BronzeArrowForCrossbow, r.BronzeArrowForCrossbowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.SteelArrowForBow, ""), refine.Confirm(item.SteelArrowForBow, r.SteelArrowForBowRequirements())),
			r.CreateChoice(refine.ItemIdList(item.SteelArrowForCrossbow, ""), refine.Confirm(item.SteelArrowForCrossbow, r.SteelArrowForCrossbow())),
		},
	}
}

func (r Vicious) WarBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 5}, {ItemId: item.BlueSnailShell, Amount: 30}}, refine.SetCost(80))
}

func (r Vicious) CompositeBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Screw, Amount: 3}}, refine.SetCost(200))
}

func (r Vicious) HuntersBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 30}, {ItemId: item.RedSnailShell, Amount: 50}}, refine.SetCost(300))
}

func (r Vicious) BattleBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Topaz, Amount: 2}, {ItemId: item.Screw, Amount: 8}}, refine.SetCost(500))
}

func (r Vicious) RydenRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 5}, {ItemId: item.GoldPlate, Amount: 5}, {ItemId: item.Emerald, Amount: 3}, {ItemId: item.Topaz, Amount: 3}, {ItemId: item.Screw, Amount: 30}}, refine.SetCost(3000))
}

func (r Vicious) RedViperRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SilverPlate, Amount: 7}, {ItemId: item.Garnet, Amount: 6}, {ItemId: item.Opal, Amount: 3}, {ItemId: item.Screw, Amount: 35}}, refine.SetCost(4000))
}

func (r Vicious) Vaulter2000Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlackCrystal, Amount: 1}, {ItemId: item.SteelPlate, Amount: 10}, {ItemId: item.GoldPlate, Amount: 3}, {ItemId: item.Screw, Amount: 40}, {ItemId: item.DrakeSkull, Amount: 50}}, refine.SetCost(8000))
}

func (r Vicious) CrossbowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 7}, {ItemId: item.Screw, Amount: 2}}, refine.SetCost(100))
}

func (r Vicious) BattleCrossbowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 20}, {ItemId: item.Screw, Amount: 5}}, refine.SetCost(200))
}

func (r Vicious) BalancheRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 50}, {ItemId: item.Screw, Amount: 8}}, refine.SetCost(300))
}

func (r Vicious) MountainCrossbowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Topaz, Amount: 1}, {ItemId: item.AquaMarine, Amount: 1}, {ItemId: item.Screw, Amount: 10}}, refine.SetCost(1000))
}

func (r Vicious) EagleCrowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 5}, {ItemId: item.OrihalconPlate, Amount: 5}, {ItemId: item.Topaz, Amount: 3}, {ItemId: item.ProcessedWood, Amount: 50}, {ItemId: item.Screw, Amount: 15}}, refine.SetCost(3000))
}

func (r Vicious) HecklerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlackCrystal, Amount: 1}, {ItemId: item.SteelPlate, Amount: 8}, {ItemId: item.GoldPlate, Amount: 4}, {ItemId: item.Topaz, Amount: 2}, {ItemId: item.Screw, Amount: 30}}, refine.SetCost(5000))
}

func (r Vicious) SilverCrowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlackCrystal, Amount: 2}, {ItemId: item.SilverPlate, Amount: 6}, {ItemId: item.ProcessedWood, Amount: 30}, {ItemId: item.Screw, Amount: 30}}, refine.SetCost(8000))
}

func (r Vicious) RowerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlackCrystal, Amount: 2}, {ItemId: item.GoldPlate, Amount: 5}, {ItemId: item.Topaz, Amount: 3}, {ItemId: item.ProcessedWood, Amount: 40}, {ItemId: item.Screw, Amount: 40}}, refine.SetCost(20000))
}

func (r Vicious) BasicArcherGlovesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 15}, {ItemId: item.BlueMushroomCap, Amount: 20}}, refine.SetCost(500))
}

func (r Vicious) BlueDirosRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 20}, {ItemId: item.BlueMushroomCap, Amount: 20}, {ItemId: item.SteelPlate, Amount: 2}}, refine.SetCost(1000))
}

func (r Vicious) BlueSavataRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 40}, {ItemId: item.BlueMushroomCap, Amount: 50}, {ItemId: item.GoldPlate, Amount: 2}}, refine.SetCost(1500))
}

func (r Vicious) BrownMarkerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 50}, {ItemId: item.GoldPlate, Amount: 2}, {ItemId: item.Amethyst, Amount: 1}}, refine.SetCost(2000))
}

func (r Vicious) BronzeScalerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.SteelPlate, Amount: 3}, {ItemId: item.Leather, Amount: 60}, {ItemId: item.Screw, Amount: 15}}, refine.SetCost(3000))
}

func (r Vicious) AquaBraceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 3}, {ItemId: item.Garnet, Amount: 1}, {ItemId: item.AquaMarine, Amount: 3}, {ItemId: item.Leather, Amount: 80}, {ItemId: item.Screw, Amount: 25}}, refine.SetCost(4000))
}

func (r Vicious) BlueWillowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SilverPlate, Amount: 3}, {ItemId: item.GoldPlate, Amount: 1}, {ItemId: item.AquaMarine, Amount: 2}, {ItemId: item.DragonSkin, Amount: 40}, {ItemId: item.Screw, Amount: 35}}, refine.SetCost(5000))
}

func (r Vicious) OakerGarnerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GoldPlate, Amount: 2}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.Topaz, Amount: 8}, {ItemId: item.DragonSkin, Amount: 50}, {ItemId: item.Screw, Amount: 50}}, refine.SetCost(7000))
}

func (r Vicious) GreenDirosRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueDiros, Amount: 1}, {ItemId: item.Emerald, Amount: 2}}, refine.SetCost(700))
}

func (r Vicious) RedDirosRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueDiros, Amount: 1}, {ItemId: item.Garnet, Amount: 1}}, refine.SetCost(700))
}

func (r Vicious) RedSavataRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueSavata, Amount: 1}, {ItemId: item.Garnet, Amount: 3}}, refine.SetCost(1000))
}

func (r Vicious) DarkSavataRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueSavata, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 1}}, refine.SetCost(1200))
}

func (r Vicious) GreenMarkerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BrownMarker, Amount: 1}, {ItemId: item.Emerald, Amount: 3}}, refine.SetCost(1500))
}

func (r Vicious) BlackMarkerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BrownMarker, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 1}}, refine.SetCost(2000))
}

func (r Vicious) MithrilScalerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeScaler, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 4}}, refine.SetCost(2200))
}

func (r Vicious) GoldScalerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeScaler, Amount: 1}, {ItemId: item.GoldPlate, Amount: 2}}, refine.SetCost(2500))
}

func (r Vicious) GoldBraceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.AquaBrace, Amount: 1}, {ItemId: item.GoldPlate, Amount: 4}}, refine.SetCost(3000))
}

func (r Vicious) DarkBraceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.AquaBrace, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 2}}, refine.SetCost(4000))
}

func (r Vicious) RedWillowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueWillow, Amount: 1}, {ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.Garnet, Amount: 5}}, refine.SetCost(5500))
}

func (r Vicious) DarkWillowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueWillow, Amount: 1}, {ItemId: item.GoldPlate, Amount: 2}, {ItemId: item.BlackCrystal, Amount: 2}}, refine.SetCost(6000))
}

func (r Vicious) SephiaGarnerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.OakerGarner, Amount: 1}, {ItemId: item.Garnet, Amount: 5}, {ItemId: item.Diamond, Amount: 1}}, refine.SetCost(7000))
}

func (r Vicious) DarkGarnerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.OakerGarner, Amount: 1}, {ItemId: item.Diamond, Amount: 2}, {ItemId: item.BlackCrystal, Amount: 2}}, refine.SetCost(8000))
}

func (r Vicious) ProcessedWoodFromTreeBranchRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.TreeBranch, Amount: 10}})
}

func (r Vicious) ProcessedWoodFromFirewoodRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Firewood, Amount: 5}})
}

func (r Vicious) ScrewRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}}, refine.SetAwardAmount(15))
}

func (r Vicious) ArrowForBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 1}, {ItemId: item.StiffFeather, Amount: 1}}, refine.SetAwardAmount(1000))
}

func (r Vicious) ArrowForCrossbowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 1}, {ItemId: item.StiffFeather, Amount: 1}}, refine.SetAwardAmount(1000))
}

func (r Vicious) BronzeArrowForBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 3}, {ItemId: item.StiffFeather, Amount: 10}}, refine.SetAwardAmount(900))
}

func (r Vicious) BronzeArrowForCrossbowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 3}, {ItemId: item.StiffFeather, Amount: 10}}, refine.SetAwardAmount(900))
}

func (r Vicious) SteelArrowForBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 5}, {ItemId: item.SoftFeather, Amount: 15}}, refine.SetAwardAmount(800))
}

func (r Vicious) SteelArrowForCrossbow() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 5}, {ItemId: item.SoftFeather, Amount: 15}}, refine.SetAwardAmount(800))
}

func (r Vicious) Sorry(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Sorry, but this is how I make my living. No meso, no item.")
	return script.SendOk(l, c, m.String())
}

func (r Vicious) NeedMoreItems(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Surely you, of all people, would understand the value of having quality items? I can't do that without the items I require.")
		return script.SendOk(l, c, m.String())
	}
}

func (r Vicious) NotEnoughInventorySpace(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please make sure you have room in your inventory, and talk to me again.")
	return script.SendOk(l, c, m.String())
}

func (r Vicious) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("A perfect item, as usual. Come and see me if you need anything else.")
	return script.SendOk(l, c, m.String())
}
