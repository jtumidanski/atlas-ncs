package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Chris is located in Victoria Road - Kerning City (103000000)
type Chris struct {
}

func (r Chris) NPCId() uint32 {
	return npc.Chris
}

func (r Chris) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	hello := "Yes, I do own this forge. If you're willing to pay, I can offer you some of my services."
	categories := r.CreateCategories()
	return refine.NewGenericRefine(l, span, c, hello, categories)
}

func (r Chris) CreateCategories() []refine.ListItem {
	return []refine.ListItem{
		r.RefineMineralOre(),
		r.RefineJewelOre(),
		r.RefineMetalHoof(),
		r.RefineUpgradeClaw(),
	}
}

func (r Chris) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Chris) RefineMineralOre() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.SimpleList("Bronze"), refine.HowMany(item.BronzePlate, r.BronzeRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Steel"), refine.HowMany(item.SteelPlate, r.SteelRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Mithril"), refine.HowMany(item.MithrilPlate, r.MithrilRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Adamantium"), refine.HowMany(item.AdamantiumPlate, r.AdamantiumRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Silver"), refine.HowMany(item.SilverPlate, r.SilverRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Orihalcon"), refine.HowMany(item.OrihalconPlate, r.OrihalconRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Gold"), refine.HowMany(item.GoldPlate, r.GoldRefineRequirements())),
	}
	categoryPrompt := refine.PromptCategory("So, what kind of mineral ore would you like to refine?", choices)
	return refine.ListItem{ListText: "Refine a mineral ore", SelectionState: categoryPrompt}
}

func (r Chris) RefineJewelOre() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.SimpleList("Garnet"), refine.HowMany(item.Garnet, r.GarnetRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Amethyst"), refine.HowMany(item.Amethyst, r.AmethystRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Aquamarine"), refine.HowMany(item.AquaMarine, r.AquamarineRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Emerald"), refine.HowMany(item.Emerald, r.EmeraldRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Opal"), refine.HowMany(item.Opal, r.OpalRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Sapphire"), refine.HowMany(item.Sapphire, r.SapphireRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Topaz"), refine.HowMany(item.Topaz, r.TopazRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Diamond"), refine.HowMany(item.Diamond, r.DiamondRefineRequirements())),
		r.CreateChoice(refine.SimpleList("Black Crystal"), refine.HowMany(item.BlackCrystal, r.BlackCrystalRefineRequirements()))}
	prompt := refine.PromptCategory("So, what kind of jewel ore would you like to refine?", choices)
	return refine.ListItem{
		ListText:       "Refine a jewel ore",
		SelectionState: prompt,
	}
}

func (r Chris) RefineMetalHoof() refine.ListItem {
	return refine.ListItem{
		ListText:       "I have Iron Hog's Metal Hoof...",
		SelectionState: r.MetalHoof("You know about that? Not many people realize the potential in the Iron Hog's Metal Hoof... I can make this into something special, if you want me to.", r.SpecialRefinement),
	}
}

func (r Chris) RefineUpgradeClaw() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodGigantic, " - Thief Lv. 60"), refine.Confirm(item.BloodGigantic, r.BloodGiganticRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SapphireGigantic, " - Thief Lv. 60"), refine.Confirm(item.SapphireGigantic, r.SapphireGiganticRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkGigantic, " - Thief Lv. 60"), refine.Confirm(item.DarkGigantic, r.DarkGiganticRequirements())),
	}
	prompt := refine.PromptCategory("Ah, you wish to upgrade a claw? Then tell me, which one?", choices)
	return refine.ListItem{ListText: "Upgrade a claw", SelectionState: prompt}
}

func (r Chris) BronzeRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeOre, Amount: 10}}, refine.SetCost(300))
}

func (r Chris) SteelRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SteelOre, Amount: 10}}, refine.SetCost(300))
}

func (r Chris) MithrilRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.MithrilOre, Amount: 10}}, refine.SetCost(300))
}

func (r Chris) AdamantiumRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.AdamantiumOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) SilverRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SilverOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) OrihalconRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.OrihalconOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) GoldRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GoldOre, Amount: 10}}, refine.SetCost(800))
}

func (r Chris) GarnetRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GarnetOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) AmethystRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.AmethystOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) AquamarineRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.AquaMarineOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) EmeraldRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.EmeraldOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) OpalRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.OpalOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) SapphireRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.SapphireOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) TopazRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.TopazOre, Amount: 10}}, refine.SetCost(500))
}

func (r Chris) DiamondRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.DiamondOre, Amount: 10}}, refine.SetCost(1000))
}

func (r Chris) BlackCrystalRefineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BlackCrystalOre, Amount: 10}}, refine.SetCost(3000))
}

func (r Chris) BloodGiganticRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeGigantic, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.Garnet, Amount: 8}, {ItemId: item.DrakesBlood, Amount: 10}}, refine.SetCost(80000))
}

func (r Chris) SapphireGiganticRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeGigantic, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.Sapphire, Amount: 8}, {ItemId: item.SapOfAncientTree, Amount: 10}}, refine.SetCost(80000))
}

func (r Chris) DarkGiganticRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.BronzeGigantic, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.BlackCrystal, Amount: 3}, {ItemId: item.TaurospearHorn, Amount: 5}}, refine.SetCost(100000))
}

func (r Chris) SteelFromHogRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.IronHogsMetalHoof, Amount: 100}})
}

func (r Chris) MetalHoof(prompt string, next script.StateProducer) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		return script.SendYesNo(l, span, c, prompt, next, script.Exit())
	}
}

func (r Chris) SpecialRefinement(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.HowMany(item.SteelPlate, r.SteelFromHogRequirements())(config)(l, span, c)
}

func (r Chris) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, span, c, m.String())
}

func (r Chris) CannotAfford(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Cash only, no credit.")
	return script.SendOk(l, span, c, m.String())
}

func (r Chris) MissingSomething(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText("I cannot accept substitutes. If you don't have what I need, then I won't be able to help you. Please bring more ").ShowItemName1(itemId).AddText(".")
		return script.SendOk(l, span, c, m.String())
	}
}

func (r Chris) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Phew... I almost didn't think that would work for a second... Well, I hope you enjoy it, anyway.")
	return script.SendOk(l, span, c, m.String())
}
