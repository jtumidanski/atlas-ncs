package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Francois is located in Victoria Road - Ellinia (101000000)
type Francois struct {
}

func (r Francois) NPCId() uint32 {
	return npc.Francois
}

func (r Francois) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := "Welcome to my eco-safe refining operation! What would you like today?"
	categories := r.CreateCategories()
	return refine.NewGenericRefine(l, c, hello, categories)
}

func (r Francois) CreateCategories() []refine.RefinementCategory {
	return []refine.RefinementCategory{
		r.MakeAGlove(),
		r.UpgradeAGlove(),
		r.UpgradeAHat(),
		r.MakeAWand(),
		r.MakeAStaff(),
	}
}

func (r Francois) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Francois) MakeAGlove() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Make a glove",
		Prompt:          "So, what kind of mineral ore would you like to refine?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.Lemona, " - Magician Lv. 15"), refine.Confirm(item.Lemona, r.LemonaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueMorrican, " - Magician Lv. 20"), refine.Confirm(item.BlueMorrican, r.BlueMorricanRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.OceanMesana, " - Magician Lv. 25"), refine.Confirm(item.OceanMesana, r.OceanMesanaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedLutia, " - Magician Lv. 30"), refine.Confirm(item.RedLutia, r.RedLutiaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedNoel, " - Magician Lv. 35"), refine.Confirm(item.RedNoel, r.RedNoelRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedArten, " - Magician Lv. 40"), refine.Confirm(item.RedArten, r.RedArtenRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedPennance, " - Magician Lv. 50"), refine.Confirm(item.RedPennance, r.RedPennanceRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelManute, " - Magician Lv. 60"), refine.Confirm(item.SteelManute, r.SteelManuteRequirements())),
		},
	}
}

func (r Francois) UpgradeAGlove() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a glove",
		Prompt:          "So, what kind of mineral ore would you like to refine?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenMorrican, " - Magician Lv. 20"), refine.Confirm(item.GreenMorrican, r.GreenMorricanRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleMorrican, " - Magician Lv. 20"), refine.Confirm(item.PurpleMorrican, r.PurpleMorricanRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodMesana, " - Magician Lv. 25"), refine.Confirm(item.BloodMesana, r.BloodMesanaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkMesana, " - Magician Lv. 25"), refine.Confirm(item.DarkMesana, r.DarkMesanaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueLutia, " - Magician Lv. 30"), refine.Confirm(item.BlueLutia, r.BlueLutiaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackLutia, " - Magician Lv. 30"), refine.Confirm(item.BlackLutia, r.BlackLutiaRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueNoel, " - Magician Lv. 35"), refine.Confirm(item.BlueNoel, r.BlueNoelRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkNoel, " - Magician Lv. 35"), refine.Confirm(item.DarkNoel, r.DarkNoelRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueArten, " - Magician Lv. 40"), refine.Confirm(item.BlueArten, r.BlueArtenRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkArten, " - Magician Lv. 40"), refine.Confirm(item.DarkArten, r.DarkArtenRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.BluePennance, " - Magician Lv. 50"), refine.Confirm(item.BluePennance, r.BluePennanceRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkPennance, " - Magician Lv. 50"), refine.Confirm(item.DarkPennance, r.DarkPennanceRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldManute, " - Magician Lv. 60"), refine.Confirm(item.GoldManute, r.GoldManuteRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkManute, " - Magician Lv. 60"), refine.Confirm(item.DarkManute, r.DarkManuteRequirements())),
		},
	}
}

func (r Francois) UpgradeAHat() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a hat",
		Prompt:          "So, what kind of mineral ore would you like to refine?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelPride, " - Magician Lv. 30"), refine.Confirm(item.SteelPride, r.SteelPrideRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldenPride, " - Magician Lv. 30"), refine.Confirm(item.GoldenPride, r.GoldenPrideRequirements())),
		},
	}
}

func (r Francois) MakeAWand() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Make a wand",
		Prompt:          "So, what kind of mineral ore would you like to refine?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.WoodenWand, " - Common Lv. 8"), refine.Confirm(item.WoodenWand, r.WoodenWandRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.HardwoodWand, " - Common Lv. 13"), refine.Confirm(item.HardwoodWand, r.HardwoodWandRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.MetalWand, " - Common Lv. 18"), refine.Confirm(item.MetalWand, r.MetalWandRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.IceWand, " - Magician Lv. 23"), refine.Confirm(item.IceWand, r.IceWandRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilWand, " - Magician Lv. 28"), refine.Confirm(item.MithrilWand, r.MithrilWandRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.WizardWand, " - Magician Lv. 33"), refine.Confirm(item.WizardWand, r.WizardWandRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.FairyWand, " - Magician Lv. 38"), refine.Confirm(item.FairyWand, r.FairyWandRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.Cromi, " - Magician Lv. 48"), refine.Confirm(item.Cromi, r.CromiRequirements())),
		},
	}
}

func (r Francois) MakeAStaff() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Make a staff",
		Prompt:          "So, what kind of mineral ore would you like to refine?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.WoodenStaff, " - Magician Lv. 10"), refine.Confirm(item.WoodenStaff, r.WoodenStaffRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.SapphireStaff, " - Magician Lv. 15"), refine.Confirm(item.SapphireStaff, r.SapphireStaffRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.EmeraldStaff, " - Magician Lv. 15"), refine.Confirm(item.EmeraldStaff, r.EmeraldStaffRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.OldWoodenStaff, " - Magician Lv. 20"), refine.Confirm(item.OldWoodenStaff, r.OldWoodenStaffRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.WizardStaff, " - Magician Lv. 25"), refine.Confirm(item.WizardStaff, r.WizardStaffRequirements())),
			r.CreateChoice(refine.ItemIdAndDescriptionList(item.ArcStaff, " - Magician Lv. 45"), refine.Confirm(item.ArcStaff, r.ArcStaffRequirements())),
		},
	}
}

func (r Francois) LemonaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 15}}, refine.SetCost(7000))
}

func (r Francois) BlueMorricanRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 30}, {ItemId: item.SteelPlate, Amount: 1}}, refine.SetCost(15000))
}

func (r Francois) OceanMesanaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 50}, {ItemId: item.GoldPlate, Amount: 2}}, refine.SetCost(20000))
}

func (r Francois) RedLutiaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 60}, {ItemId: item.Topaz, Amount: 1}, {ItemId: item.Garnet, Amount: 2}}, refine.SetCost(25000))
}

func (r Francois) RedNoelRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 70}, {ItemId: item.GoldPlate, Amount: 1}, {ItemId: item.SteelPlate, Amount: 3}, {ItemId: item.Garnet, Amount: 2}}, refine.SetCost(30000))
}

func (r Francois) RedArtenRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Leather, Amount: 80}, {ItemId: item.Garnet, Amount: 3}, {ItemId: item.Topaz, Amount: 3}, {ItemId: item.Screw, Amount: 30}}, refine.SetCost(40000))
}

func (r Francois) RedPennanceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Garnet, Amount: 3}, {ItemId: item.GoldPlate, Amount: 2}, {ItemId: item.DragonSkin, Amount: 35}, {ItemId: item.Screw, Amount: 40}}, refine.SetCost(50000))
}

func (r Francois) SteelManuteRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.MoonRock, Amount: 1}, {ItemId: item.SteelPlate, Amount: 8}, {ItemId: item.Diamond, Amount: 1}, {ItemId: item.DragonSkin, Amount: 50}, {ItemId: item.Screw, Amount: 50}}, refine.SetCost(70000))
}

func (r Francois) GreenMorricanRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueMorrican, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}}, refine.SetCost(20000))
}

func (r Francois) PurpleMorricanRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlueMorrican, Amount: 1}, {ItemId: item.Amethyst, Amount: 2}}, refine.SetCost(25000))
}

func (r Francois) BloodMesanaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.OceanMesana, Amount: 1}, {ItemId: item.Garnet, Amount: 3}}, refine.SetCost(30000))
}

func (r Francois) DarkMesanaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.OceanMesana, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 1}}, refine.SetCost(40000))
}

func (r Francois) BlueLutiaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedLutia, Amount: 1}, {ItemId: item.Sapphire, Amount: 3}}, refine.SetCost(35000))
}

func (r Francois) BlackLutiaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedLutia, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 1}}, refine.SetCost(40000))
}

func (r Francois) BlueNoelRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedNoel, Amount: 1}, {ItemId: item.Sapphire, Amount: 3}}, refine.SetCost(40000))
}

func (r Francois) DarkNoelRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedNoel, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 1}}, refine.SetCost(45000))
}

func (r Francois) BlueArtenRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedArten, Amount: 1}, {ItemId: item.AquaMarine, Amount: 4}}, refine.SetCost(45000))
}

func (r Francois) DarkArtenRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedArten, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 2}}, refine.SetCost(50000))
}

func (r Francois) BluePennanceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedPennance, Amount: 1}, {ItemId: item.AquaMarine, Amount: 5}}, refine.SetCost(55000))
}

func (r Francois) DarkPennanceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.RedPennance, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 3}}, refine.SetCost(60000))
}

func (r Francois) GoldManuteRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelManute, Amount: 1}, {ItemId: item.SilverPlate, Amount: 3}, {ItemId: item.GoldPlate, Amount: 5}}, refine.SetCost(70000))
}

func (r Francois) DarkManuteRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelManute, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 2}, {ItemId: item.GoldPlate, Amount: 3}}, refine.SetCost(80000))
}

func (r Francois) SteelPrideRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePride, Amount: 1}, {ItemId: item.SteelPlate, Amount: 3}}, refine.SetCost(40000))
}

func (r Francois) GoldenPrideRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePride, Amount: 1}, {ItemId: item.GoldPlate, Amount: 3}}, refine.SetCost(50000))
}

func (r Francois) WoodenWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 5}}, refine.SetCost(1000))
}

func (r Francois) HardwoodWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 10}, {ItemId: item.OrangeMushroomCap, Amount: 50}}, refine.SetCost(3000))
}

func (r Francois) MetalWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.BlueMushroomCap, Amount: 30}, {ItemId: item.Screw, Amount: 5}}, refine.SetCost(5000))
}

func (r Francois) IceWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.MithrilPlate, Amount: 2}, {ItemId: item.PieceOfIce, Amount: 1}, {ItemId: item.Screw, Amount: 10}}, refine.SetCost(12000))
}

func (r Francois) MithrilWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.MithrilPlate, Amount: 3}, {ItemId: item.AquaMarine, Amount: 1}, {ItemId: item.Screw, Amount: 10}}, refine.SetCost(30000))
}

func (r Francois) WizardWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Topaz, Amount: 5}, {ItemId: item.MithrilPlate, Amount: 3}, {ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Screw, Amount: 15}}, refine.SetCost(60000))
}

func (r Francois) FairyWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Topaz, Amount: 5}, {ItemId: item.Sapphire, Amount: 5}, {ItemId: item.Diamond, Amount: 1}, {ItemId: item.FairyWing, Amount: 1}, {ItemId: item.Screw, Amount: 20}}, refine.SetCost(120000))
}

func (r Francois) CromiRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GoldPlate, Amount: 4}, {ItemId: item.Emerald, Amount: 3}, {ItemId: item.Diamond, Amount: 2}, {ItemId: item.AquaMarine, Amount: 1}, {ItemId: item.PieceOfIce, Amount: 1}, {ItemId: item.Screw, Amount: 30}}, refine.SetCost(200000))
}

func (r Francois) WoodenStaffRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 5}}, refine.SetCost(2000))
}

func (r Francois) SapphireStaffRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Sapphire, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Screw, Amount: 5}}, refine.SetCost(2000))
}

func (r Francois) EmeraldStaffRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Emerald, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Screw, Amount: 5}}, refine.SetCost(2000))
}

func (r Francois) OldWoodenStaffRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 50}, {ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Screw, Amount: 10}}, refine.SetCost(5000))
}

func (r Francois) WizardStaffRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Topaz, Amount: 2}, {ItemId: item.Amethyst, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.Screw, Amount: 15}}, refine.SetCost(12000))
}

func (r Francois) ArcStaffRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 8}, {ItemId: item.Topaz, Amount: 5}, {ItemId: item.Amethyst, Amount: 5}, {ItemId: item.Sapphire, Amount: 5}, {ItemId: item.Screw, Amount: 30}, {ItemId: item.SlimeBubble, Amount: 50}, {ItemId: item.FairyWing, Amount: 1}}, refine.SetCost(180000))
}

func (r Francois) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, c, m.String())
}

func (r Francois) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Sorry, but all of us need money to live. Come back when you can pay my fees, yes?")
	return script.SendOk(l, c, m.String())
}

func (r Francois) MissingSomething(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Uhm... I don't keep extra material on me. Sorry. Can you please bring ").
			ShowItemName1(itemId).
			AddText("?")
		return script.SendOk(l, c, m.String())
	}
}

func (r Francois) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("It's a success! Oh, I've never felt so alive! Please come back again!")
	return script.SendOk(l, c, m.String())
}
