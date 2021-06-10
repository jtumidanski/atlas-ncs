package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MrSmith is located in Victoria Road - Perion (102000000)
type MrSmith struct {
}

func (r MrSmith) NPCId() uint32 {
	return npc.MrSmith
}

func (r MrSmith) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := "Um... Hi, I'm Mr. Thunder's apprentice. He's getting up there in age, so he handles most of the heavy-duty work while I handle some of the lighter jobs. What can I do for you?"
	categories := r.CreateCategories()
	return refine.NewGenericRefine(l, c, hello, categories)
}

func (r MrSmith) CreateCategories() []refine.RefinementCategory {
	return []refine.RefinementCategory{
		r.MakeAGlove(),
		r.UpgradeAGlove(),
		r.CreateMaterials(),
	}
}

func (r MrSmith) CreateRefinementChoice(itemText string, itemDescription string, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	return refine.RefinementChoice{
		ListText:        message.NewBuilder().BlueText().AddText(itemText).BlackText().AddText(itemDescription).String(),
		SelectionPrompt: selectionPrompt,
		Config: refine.TerminalConfig{
			Success:          r.Success,
			MesoError:        r.CannotAfford,
			RequirementError: r.MissingSomething,
			InventoryError:   r.MakeRoom,
		},
	}
}

func (r MrSmith) MakeAGlove() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Make a glove",
		Prompt:          "Okay, so which glove do you want me to make?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateRefinementChoice("Juno", " - Warrior Lv. 10", refine.Confirm(item.Juno, r.JunoRequirements())),
			r.CreateRefinementChoice("Steel Fingerless Gloves", " - Warrior Lv. 15", refine.Confirm(item.SteelFingerlessGloves, r.SteelFingerlessGlovesRequirements())),
			r.CreateRefinementChoice("Venon", " - Warrior Lv. 20", refine.Confirm(item.Venon, r.VenonRequirements())),
			r.CreateRefinementChoice("White Fingerless Gloves", " - Warrior Lv. 25", refine.Confirm(item.WhiteFingerlessGloves, r.WhiteFingerlessGlovesRequirements())),
			r.CreateRefinementChoice("Bronze Missel", " - Warrior Lv. 30", refine.Confirm(item.BronzeMissel, r.BronzeMisselRequirements())),
			r.CreateRefinementChoice("Steel Briggon", " - Warrior Lv. 35", refine.Confirm(item.SteelBriggon, r.SteelBriggonRequirements())),
			r.CreateRefinementChoice("Iron Knuckle", " - Warrior Lv. 40", refine.Confirm(item.IronKnuckle, r.IronKnuckleRequirements())),
			r.CreateRefinementChoice("Steel Brist", " - Warrior Lv. 50", refine.Confirm(item.SteelBrist, r.SteelBristRequirements())),
			r.CreateRefinementChoice("Bronze Clench", " - Warrior Lv. 60", refine.Confirm(item.BronzeClench, r.BronzeClenchRequirements())),
		},
	}
}

func (r MrSmith) UpgradeAGlove() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a glove",
		Prompt:          "Upgrade a glove? That shouldn't be too difficult. Which did you have in mind?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateRefinementChoice("Steel Missel", " - Warrior Lv. 30", refine.Confirm(item.SteelMissel, r.SteelMisselRequirements())),
			r.CreateRefinementChoice("Orihalcon Missel", " - Warrior Lv. 30", refine.Confirm(item.OrihalconMissel, r.OrihalconMisselRequirements())),
			r.CreateRefinementChoice("Yellow Briggon", " - Warrior Lv. 35", refine.Confirm(item.YellowBriggon, r.YellowBriggonRequirements())),
			r.CreateRefinementChoice("Dark Briggon", " - Warrior Lv. 35", refine.Confirm(item.DarkBriggon, r.DarkBriggonRequirements())),
			r.CreateRefinementChoice("Adamantium Knuckle", " - Warrior Lv. 40", refine.Confirm(item.AdamantiumKnuckle, r.AdamantiumKnuckleRequirements())),
			r.CreateRefinementChoice("Dark Knuckle", " - Warrior Lv. 40", refine.Confirm(item.DarkKnuckle, r.DarkKnuckleRequirements())),
			r.CreateRefinementChoice("Mithril Brist", " - Warrior Lv. 50", refine.Confirm(item.MithrilBrist, r.MithrilBristRequirements())),
			r.CreateRefinementChoice("Gold Brist", " - Warrior Lv. 50", refine.Confirm(item.GoldBrist, r.GoldBristRequirements())),
			r.CreateRefinementChoice("Sapphire Clench", " - Warrior Lv. 60", refine.Confirm(item.SapphireClench, r.SapphireClenchRequirements())),
			r.CreateRefinementChoice("Dark Clench", " - Warrior Lv. 60", refine.Confirm(item.DarkClench, r.DarkClenchRequirements())),
		},
	}
}

func (r MrSmith) CreateMaterials() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Create materials",
		Prompt:          "Materials? I know of a few materials that I can make for you...",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateRefinementChoice("Make Processed Wood with Tree Branch", "", refine.HowMany(item.ProcessedWood, r.ProcessedWoodFromTreeBranchRequirements())),
			r.CreateRefinementChoice("Make Processed Wood with Firewood", "", refine.HowMany(item.ProcessedWood, r.ProcessedWoodFromFirewoodRequirements())),
			r.CreateRefinementChoice("Make Screws (packs of 15)", "", refine.HowMany(item.Screw, r.ScrewRequirements())),
		},
	}
}

func (r MrSmith) ProcessedWoodFromTreeBranchRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.TreeBranch, Amount: 10}}, Cost: 0}
}

func (r MrSmith) ProcessedWoodFromFirewoodRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Firewood, Amount: 5}}, Cost: 0}
}

func (r MrSmith) ScrewRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}}, Cost: 0}
}

func (r MrSmith) JunoRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Leather, Amount: 15}, {ItemId: item.SteelPlate, Amount: 1}}, Cost: 1000}
}

func (r MrSmith) SteelFingerlessGlovesRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelPlate, Amount: 2}}, Cost: 2000}
}

func (r MrSmith) VenonRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Leather, Amount: 40}, {ItemId: item.BronzePlate, Amount: 2}}, Cost: 5000}
}

func (r MrSmith) WhiteFingerlessGlovesRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelPlate, Amount: 2}}, Cost: 10000}
}

func (r MrSmith) BronzeMisselRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzePlate, Amount: 3}, {ItemId: item.SteelPlate, Amount: 2}, {ItemId: item.Screw, Amount: 15}}, Cost: 20000}
}

func (r MrSmith) SteelBriggonRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Leather, Amount: 30}, {ItemId: item.SteelPlate, Amount: 4}, {ItemId: item.Screw, Amount: 15}}, Cost: 30000}
}

func (r MrSmith) IronKnuckleRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.Leather, Amount: 50}, {ItemId: item.SteelPlate, Amount: 5}, {ItemId: item.Screw, Amount: 40}}, Cost: 40000}
}

func (r MrSmith) SteelBristRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelPlate, Amount: 3}, {ItemId: item.Diamond, Amount: 2}, {ItemId: item.DragonSkin, Amount: 30}, {ItemId: item.Screw, Amount: 45}}, Cost: 50000}
}

func (r MrSmith) BronzeClenchRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.MoonRock, Amount: 1}, {ItemId: item.BronzePlate, Amount: 8}, {ItemId: item.GoldPlate, Amount: 2}, {ItemId: item.DragonSkin, Amount: 50}, {ItemId: item.Screw, Amount: 50}}, Cost: 70000}
}

func (r MrSmith) SteelMisselRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeMissel, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}}, Cost: 20000}
}

func (r MrSmith) OrihalconMisselRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeMissel, Amount: 1}, {ItemId: item.OrihalconPlate, Amount: 2}}, Cost: 25000}
}

func (r MrSmith) YellowBriggonRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelBriggon, Amount: 1}, {ItemId: item.Topaz, Amount: 3}}, Cost: 30000}
}

func (r MrSmith) DarkBriggonRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelBriggon, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 1}}, Cost: 40000}
}

func (r MrSmith) AdamantiumKnuckleRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.IronKnuckle, Amount: 1}, {ItemId: item.AdamantiumPlate, Amount: 4}}, Cost: 45000}
}

func (r MrSmith) DarkKnuckleRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.IronKnuckle, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 2}}, Cost: 50000}
}

func (r MrSmith) MithrilBristRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelBrist, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 5}}, Cost: 55000}
}

func (r MrSmith) GoldBristRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelBrist, Amount: 1}, {ItemId: item.GoldPlate, Amount: 4}}, Cost: 60000}
}

func (r MrSmith) SapphireClenchRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeClench, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 3}, {ItemId: item.Sapphire, Amount: 5}}, Cost: 70000}
}

func (r MrSmith) DarkClenchRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeClench, Amount: 1}, {ItemId: item.Diamond, Amount: 2}, {ItemId: item.BlackCrystal, Amount: 2}}, Cost: 80000}
}

func (r MrSmith) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, c, m.String())
}

func (r MrSmith) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I may still be an apprentice, but I do need to earn a living.")
	return script.SendOk(l, c, m.String())
}

func (r MrSmith) MissingSomething(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("I'm still an apprentice, I don't know if I can substitute other items in yet... Can you please bring ").
			ShowItemName1(itemId).
			AddText("?")
		return script.SendOk(l, c, m.String())
	}
}

func (r MrSmith) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Did that come out right? Come by me again if you have anything for me to practice on.")
	return script.SendOk(l, c, m.String())
}
