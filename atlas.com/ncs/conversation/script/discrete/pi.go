package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Pi is located in Ludibrium - Ludibrium Village (220000300)
type Pi struct {
}

func (r Pi) NPCId() uint32 {
	return npc.Pi
}

func (r Pi) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return refine.NewGenericRefine(l, c, r.Hello(), r.Categories())
}

func (r Pi) Hello() string {
	return "Hm? Who might you be? Oh, you've heard about my forging skills? In that case, I'd be glad to process some of your ores... for a fee."
}

func (r Pi) Categories() []refine.ListItem {
	return []refine.ListItem{
		r.RefineMineralOre(),
		r.RefineJewelOre(),
		r.RareJewel(),
		r.CrystalOre(),
		r.CreateMaterials(),
		r.CreateArrows(),
	}
}

func (r Pi) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Pi) RefineMineralOre() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.SimpleList("Bronze"), refine.HowMany(item.BronzePlate, r.BronzeRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Steel"), refine.HowMany(item.SteelPlate, r.SteelRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Mithril"), refine.HowMany(item.MithrilPlate, r.MithrilRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Adamantium"), refine.HowMany(item.AdamantiumPlate, r.AdamantiumRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Silver"), refine.HowMany(item.SilverPlate, r.SilverRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Orihalcon"), refine.HowMany(item.OrihalconPlate, r.OrihalconRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Gold"), refine.HowMany(item.GoldPlate, r.GoldRefineRequirements())),
	}
	prompt := refine.PromptCategory("So, what kind of mineral ore would you like to refine?", choices)
	return refine.ListItem{ListText: "Refine a mineral ore", SelectionState: prompt}
}

func (r Pi) RefineJewelOre() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.SimpleList("Garnet"), refine.HowMany(item.Garnet, r.GarnetRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Amethyst"), refine.HowMany(item.Amethyst, r.AmethystRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Aquamarine"), refine.HowMany(item.AquaMarine, r.AquamarineRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Emerald"), refine.HowMany(item.Emerald, r.EmeraldRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Opal"), refine.HowMany(item.Opal, r.OpalRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Sapphire"), refine.HowMany(item.Sapphire, r.SapphireRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Topaz"), refine.HowMany(item.Topaz, r.TopazRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Diamond"), refine.HowMany(item.Diamond, r.DiamondRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Black Crystal"), refine.HowMany(item.BlackCrystal, r.BlackCrystalRefineRequirements())),
	}
	prompt := refine.PromptCategory("So, what kind of jewel ore would you like to refine?", choices)
	return refine.ListItem{ListText: "Refine a jewel ore", SelectionState: prompt}
}

func (r Pi) RareJewel() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.MoonRock), refine.HowMany(item.MoonRock, r.MoonRockRequirements())),
		r.CreateChoice(refine.ItemIdList(item.StarRock), refine.HowMany(item.StarRock, r.StarRockRequirements())),
	}
	prompt := refine.PromptCategory("A rare jewel? Which one were you thinking of?", choices)
	return refine.ListItem{ListText: "Refine a rare jewel", SelectionState: prompt}
}

func (r Pi) CrystalOre() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.PowerCrystal), refine.HowMany(item.PowerCrystal, r.PowerCrystalRequirements())),
		r.CreateChoice(refine.ItemIdList(item.WisdomCrystal), refine.HowMany(item.WisdomCrystal, r.WisdomCrystalRequirements())),
		r.CreateChoice(refine.ItemIdList(item.DEXCrystal), refine.HowMany(item.DEXCrystal, r.DEXCrystalRequirements())),
		r.CreateChoice(refine.ItemIdList(item.LUKCrystal), refine.HowMany(item.LUKCrystal, r.LUKCrystalRequirements())),
	}
	prompt := refine.PromptCategory("Crystal ore? It's hard to find those around here...", choices)
	return refine.ListItem{ListText: "Refine a crystal ore", SelectionState: prompt}
}

func (r Pi) CreateMaterials() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.SimpleList("Make Processed Wood with Tree Branch"), refine.HowMany(item.ProcessedWood, r.ProcessedWoodFromTreeBranchRequirements())),
		r.CreateChoice(refine.SimpleList("Make Processed Wood with Firewood"), refine.HowMany(item.ProcessedWood, r.ProcessedWoodFromFirewoodRequirements())),
		r.CreateChoice(refine.SimpleList("Make Screws (packs of 15)"), refine.HowMany(item.Screw, r.ScrewRequirements())),
	}
	prompt := refine.PromptCategory("Materials? I know of a few materials that I can make for you...", choices)
	return refine.ListItem{ListText: "Create materials", SelectionState: prompt}
}

func (r Pi) CreateArrows() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.ArrowForBow), refine.Confirm(item.ArrowForBow, r.ArrowForBowRequirements())),
		r.CreateChoice(refine.ItemIdList(item.ArrowForCrossbow), refine.Confirm(item.ArrowForCrossbow, r.ArrowForCrossbowRequirements())),
		r.CreateChoice(refine.ItemIdList(item.BronzeArrowForBow), refine.Confirm(item.BronzeArrowForBow, r.BronzeArrowForBowRequirements())),
		r.CreateChoice(refine.ItemIdList(item.BronzeArrowForCrossbow), refine.Confirm(item.BronzeArrowForCrossbow, r.BronzeArrowForCrossbowRequirements())),
		r.CreateChoice(refine.ItemIdList(item.SteelArrowForBow), refine.Confirm(item.SteelArrowForBow, r.SteelArrowForBowRequirements())),
		r.CreateChoice(refine.ItemIdList(item.SteelArrowForCrossbow), refine.Confirm(item.SteelArrowForCrossbow, r.SteelArrowForCrossbow())),
	}
	prompt := refine.PromptCategory("Arrows? Not a problem at all.", choices)
	return refine.ListItem{ListText: "Create arrows", SelectionState: prompt}
}

func (r Pi) BronzeRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeOre, Amount: 10}}, refine.SetCost(300))
}

func (r Pi) SteelRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelOre, Amount: 10}}, refine.SetCost(300))
}

func (r Pi) MithrilRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.MithrilOre, Amount: 10}}, refine.SetCost(300))
}

func (r Pi) AdamantiumRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.AdamantiumOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) SilverRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SilverOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) OrihalconRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.OrihalconOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) GoldRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GoldOre, Amount: 10}}, refine.SetCost(800))
}

func (r Pi) GarnetRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GarnetOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) AmethystRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.AmethystOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) AquamarineRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.AquaMarineOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) EmeraldRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.EmeraldOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) OpalRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.OpalOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) SapphireRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SapphireOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) TopazRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.TopazOre, Amount: 10}}, refine.SetCost(500))
}

func (r Pi) DiamondRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.DiamondOre, Amount: 10}}, refine.SetCost(1000))
}

func (r Pi) BlackCrystalRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlackCrystalOre, Amount: 10}}, refine.SetCost(3000))
}

func (r Pi) ProcessedWoodFromTreeBranchRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.TreeBranch, Amount: 10}})
}

func (r Pi) ProcessedWoodFromFirewoodRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.Firewood, Amount: 5}})
}

func (r Pi) ScrewRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}}, refine.SetAwardAmount(15))
}

func (r Pi) ArrowForBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 1}, {ItemId: item.StiffFeather, Amount: 1}}, refine.SetAwardAmount(1000))
}

func (r Pi) ArrowForCrossbowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.ProcessedWood, Amount: 1}, {ItemId: item.StiffFeather, Amount: 1}}, refine.SetAwardAmount(1000))
}

func (r Pi) BronzeArrowForBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 3}, {ItemId: item.StiffFeather, Amount: 10}}, refine.SetAwardAmount(900))
}

func (r Pi) BronzeArrowForCrossbowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzePlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 3}, {ItemId: item.StiffFeather, Amount: 10}}, refine.SetAwardAmount(900))
}

func (r Pi) SteelArrowForBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 5}, {ItemId: item.SoftFeather, Amount: 15}}, refine.SetAwardAmount(800))
}

func (r Pi) SteelArrowForCrossbow() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelPlate, Amount: 1}, {ItemId: item.ProcessedWood, Amount: 5}, {ItemId: item.SoftFeather, Amount: 15}}, refine.SetAwardAmount(800))
}

func (r Pi) MoonRockRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 1}, {ItemId: 4011001, Amount: 1}, {ItemId: 4011002, Amount: 1}, {ItemId: 4011003, Amount: 1}, {ItemId: 4011004, Amount: 1}, {ItemId: 4011005, Amount: 1}, {ItemId: 4011006, Amount: 1}}, refine.SetCost(10000))
}

func (r Pi) StarRockRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 1}, {ItemId: 4021001, Amount: 1}, {ItemId: 4021002, Amount: 1}, {ItemId: 4021003, Amount: 1}, {ItemId: 4021004, Amount: 1}, {ItemId: 4021005, Amount: 1}, {ItemId: 4021006, Amount: 1}, {ItemId: 4021007, Amount: 1}, {ItemId: 4021008, Amount: 1}}, refine.SetCost(15000))
}

func (r Pi) PowerCrystalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4004000, Amount: 10}}, refine.SetCost(5000))
}

func (r Pi) WisdomCrystalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4004001, Amount: 10}}, refine.SetCost(5000))
}

func (r Pi) DEXCrystalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4004002, Amount: 10}}, refine.SetCost(5000))
}

func (r Pi) LUKCrystalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4004003, Amount: 10}}, refine.SetCost(5000))
}

func (r Pi) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("All done. If you need anything else, you know where to find me.")
	return script.SendOk(l, c, m.String())
}

func (r Pi) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm afraid you cannot afford my services.")
	return script.SendOk(l, c, m.String())
}

func (r Pi) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText("Hold it, I can't finish that without all of the proper materials. Bring them first, then we'll talk.")
		return script.SendOk(l, c, m.String())
	}
}

func (r Pi) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm afraid you have no slots available for this transaction.")
	return script.SendOk(l, c, m.String())
}
