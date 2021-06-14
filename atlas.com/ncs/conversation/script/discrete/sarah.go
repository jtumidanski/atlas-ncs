package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Sarah is located in Ludibrium - Tara and Sarah's House (220000303)
type Sarah struct {
}

func (r Sarah) NPCId() uint32 {
	return npc.Sarah
}

func (r Sarah) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return refine.NewGenericRefine(l, c, r.Hello(), r.Categories())
}

func (r Sarah) Hello() string {
	return "Hello, and welcome to the Ludibrium Glove Store. How can I help you today?"
}

func (r Sarah) Categories() []refine.ListItem {
	return []refine.ListItem{
		r.WhatIsAStimulator(),
		r.Warrior(),
		r.Bowman(),
		r.Magician(),
		r.Thief(),
		r.WarriorStimulator(),
		r.BowmanStimulator(),
		r.MagicianStimulator(),
		r.ThiefStimulator(),
	}
}

func (r Sarah) WhatIsAStimulator() refine.ListItem {
	return refine.ListItem{
		ListText:       "What's a stimulator?",
		SelectionState: r.StimulatorInfo,
	}
}

func (r Sarah) StimulatorInfo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("A stimulator is a special potion that I can add into the process of creating certain items. It gives it stats as though it had dropped from a monster. However, it is possible to have no change, and it is also possible for the item to be below average. There's also a 10% chance of not getting any item when using a stimulator, so please choose wisely.")
	return script.SendOk(l, c, m.String())
}

func (r Sarah) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
		StimulatorError:  r.StimulatorError,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Sarah) Warrior() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeMissel, " - Warrior Lv. 30"), refine.Confirm(item.BronzeMissel, r.BronzeMisselRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelBriggon, " - Warrior Lv. 35"), refine.Confirm(item.SteelBriggon, r.SteelBriggonRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.IronKnuckle, " - Warrior Lv. 40"), refine.Confirm(item.IronKnuckle, r.IronKnuckleRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelBrist, " - Warrior Lv. 50"), refine.Confirm(item.SteelBrist, r.SteelBristRequirements())),
	}
	prompt := refine.PromptCategory("Warrior glove? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create a Warrior glove",
		SelectionState: prompt,
	}
}

func (r Sarah) Bowman() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownMarker, " - Bowman Lv. 30"), refine.Confirm(item.BrownMarker, r.BrownMarkerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeScaler, " - Bowman Lv. 35"), refine.Confirm(item.BronzeScaler, r.BronzeScalerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.AquaBrace, " - Bowman Lv. 40"), refine.Confirm(item.AquaBrace, r.AquaBraceRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueWillow, " - Bowman Lv. 50"), refine.Confirm(item.BlueWillow, r.BlueWillowRequirements())),
	}
	prompt := refine.PromptCategory("Bowman glove? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create a Bowman glove",
		SelectionState: prompt,
	}
}

func (r Sarah) Magician() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedLutia, " - Magician Lv. 30"), refine.Confirm(item.RedLutia, r.RedLutiaRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedNoel, " - Magician Lv. 35"), refine.Confirm(item.RedNoel, r.RedNoelRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedArten, " - Magician Lv. 40"), refine.Confirm(item.RedArten, r.RedArtenRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedPennance, " - Magician Lv. 50"), refine.Confirm(item.RedPennance, r.RedPennanceRequirements())),
	}
	prompt := refine.PromptCategory("Magician glove? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create a Magician glove",
		SelectionState: prompt,
	}
}

func (r Sarah) Thief() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelSylvia, " - Thief Lv. 30"), refine.Confirm(item.SteelSylvia, r.SteelSylviaRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelArbion, " - Thief Lv. 35"), refine.Confirm(item.SteelArbion, r.SteelArbionRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedCleave, " - Thief Lv. 40"), refine.Confirm(item.RedCleave, r.RedCleaveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueMoonGlove, " - Thief Lv. 50"), refine.Confirm(item.BlueMoonGlove, r.BlueMoonGloveRequirements())),
	}
	prompt := refine.PromptCategory("Thief glove? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create a Thief glove",
		SelectionState: prompt,
	}
}

func (r Sarah) WarriorStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelMissel, " - Warrior Lv. 30"), refine.Confirm(item.SteelMissel, r.SteelMisselRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrihalconMissel, " - Warrior Lv. 30"), refine.Confirm(item.OrihalconMissel, r.OrihalconMisselRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.YellowBriggon, " - Warrior Lv. 35"), refine.Confirm(item.YellowBriggon, r.YellowBriggonRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkBriggon, " - Warrior Lv. 35"), refine.Confirm(item.DarkBriggon, r.DarkBriggonRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.AdamantiumKnuckle, " - Warrior Lv. 40"), refine.Confirm(item.AdamantiumKnuckle, r.AdamantiumKnuckleRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkKnuckle, " - Warrior Lv. 40"), refine.Confirm(item.DarkKnuckle, r.DarkKnuckleRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilBrist, " - Warrior Lv. 50"), refine.Confirm(item.MithrilBrist, r.MithrilBristRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldBrist, " - Warrior Lv. 50"), refine.Confirm(item.GoldBrist, r.GoldBristRequirements())),
	}
	prompt := refine.PromptCategory("Warrior glove with a stimulator? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create a Warrior glove with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Sarah) BowmanStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenMarker, " - Bowman Lv. 30"), refine.Confirm(item.GreenMarker, r.GreenMarkerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackMarker, " - Bowman Lv. 30"), refine.Confirm(item.BlackMarker, r.BlackMarkerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilScaler, " - Bowman Lv. 35"), refine.Confirm(item.MithrilScaler, r.MithrilScalerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldScaler, " - Bowman Lv. 35"), refine.Confirm(item.GoldScaler, r.GoldScalerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldBrace, " - Bowman Lv. 40"), refine.Confirm(item.GoldBrace, r.GoldBraceRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkBrace, " - Bowman Lv. 40"), refine.Confirm(item.DarkBrace, r.DarkBraceRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedWillow, " - Bowman Lv. 50"), refine.Confirm(item.RedWillow, r.RedWillowRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkWillow, " - Bowman Lv. 50"), refine.Confirm(item.DarkWillow, r.DarkWillowRequirements())),
	}
	prompt := refine.PromptCategory("Bowman glove with a stimulator? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create a Bowman glove with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Sarah) MagicianStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueLutia, " - Magician Lv. 30"), refine.Confirm(item.BlueLutia, r.BlueLutiaRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackLutia, " - Magician Lv. 30"), refine.Confirm(item.BlackLutia, r.BlackLutiaRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueNoel, " - Magician Lv. 35"), refine.Confirm(item.BlueNoel, r.BlueNoelRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkNoel, " - Magician Lv. 35"), refine.Confirm(item.DarkNoel, r.DarkNoelRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueArten, " - Magician Lv. 40"), refine.Confirm(item.BlueArten, r.BlueArtenRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkArten, " - Magician Lv. 40"), refine.Confirm(item.DarkArten, r.DarkArtenRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BluePennance, " - Magician Lv. 50"), refine.Confirm(item.BluePennance, r.BluePennanceRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkPennance, " - Magician Lv. 50"), refine.Confirm(item.DarkPennance, r.DarkPennanceRequirements())),
	}
	prompt := refine.PromptCategory("Magician glove with a stimulator? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create a Magician glove with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Sarah) ThiefStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverSylvia, " - Thief Lv. 30"), refine.Confirm(item.SilverSylvia, r.SilverSylviaRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldSylvia, " - Thief Lv. 30"), refine.Confirm(item.GoldSylvia, r.GoldSylviaRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrihalconArbion, " - Thief Lv. 35"), refine.Confirm(item.OrihalconArbion, r.OrihalconArbionRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldArbion, " - Thief Lv. 35"), refine.Confirm(item.GoldArbion, r.GoldArbionRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldCleave, " - Thief Lv. 40"), refine.Confirm(item.GoldCleave, r.GoldCleaveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkCleave, " - Thief Lv. 40"), refine.Confirm(item.DarkCleave, r.DarkCleaveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedMoonGlove, " - Thief Lv. 50"), refine.Confirm(item.RedMoonGlove, r.RedMoonGloveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownMoonGlove, " - Thief Lv. 50"), refine.Confirm(item.BrownMoonGlove, r.BrownMoonGloveRequirements())),
	}
	prompt := refine.PromptCategory("Thief glove with a stimulator? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create a Thief glove with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Sarah) BronzeMisselRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 3}, {ItemId: 4011001, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(18000))
}

func (r Sarah) SteelBriggonRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 30}, {ItemId: 4011001, Amount: 4}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(27000))
}

func (r Sarah) IronKnuckleRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4011001, Amount: 5}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(36000))
}

func (r Sarah) SteelBristRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021007, Amount: 2}, {ItemId: 4000030, Amount: 30}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(45000))
}

func (r Sarah) BrownMarkerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4011006, Amount: 2}, {ItemId: 4021001, Amount: 1}}, refine.SetCost(18000))
}

func (r Sarah) BronzeScalerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4000021, Amount: 60}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(27000))
}

func (r Sarah) AquaBraceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021000, Amount: 1}, {ItemId: 4021002, Amount: 3}, {ItemId: 4000021, Amount: 80}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(36000))
}

func (r Sarah) BlueWillowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 3}, {ItemId: 4011006, Amount: 1}, {ItemId: 4021002, Amount: 2}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(45000))
}

func (r Sarah) RedLutiaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 60}, {ItemId: 4021006, Amount: 1}, {ItemId: 4021000, Amount: 2}}, refine.SetCost(22500))
}

func (r Sarah) RedNoelRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 70}, {ItemId: 4011006, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4021000, Amount: 2}}, refine.SetCost(27000))
}

func (r Sarah) RedArtenRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 80}, {ItemId: 4021000, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(36000))
}

func (r Sarah) RedPennanceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 3}, {ItemId: 4011006, Amount: 2}, {ItemId: 4000030, Amount: 35}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(45000))
}

func (r Sarah) SteelSylviaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(22500))
}

func (r Sarah) SteelArbionRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4011000, Amount: 1}, {ItemId: 4000021, Amount: 60}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(27000))
}

func (r Sarah) RedCleaveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 3}, {ItemId: 4000101, Amount: 100}, {ItemId: 4000021, Amount: 80}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(36000))
}

func (r Sarah) BlueMoonGloveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021005, Amount: 3}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(45000))
}

func (r Sarah) SteelMisselRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082007, Amount: 1}, {ItemId: 4011001, Amount: 1}}, refine.SetCost(18000))
}

func (r Sarah) OrihalconMisselRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082007, Amount: 1}, {ItemId: 4011005, Amount: 2}}, refine.SetCost(22500))
}

func (r Sarah) YellowBriggonRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082008, Amount: 1}, {ItemId: 4021006, Amount: 3}}, refine.SetCost(27000))
}

func (r Sarah) DarkBriggonRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082008, Amount: 1}, {ItemId: 4021008, Amount: 1}}, refine.SetCost(36000))
}

func (r Sarah) AdamantiumKnuckleRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082023, Amount: 1}, {ItemId: 4011003, Amount: 4}}, refine.SetCost(40500))
}

func (r Sarah) DarkKnuckleRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082023, Amount: 1}, {ItemId: 4021008, Amount: 2}}, refine.SetCost(45000))
}

func (r Sarah) MithrilBristRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082009, Amount: 1}, {ItemId: 4011002, Amount: 5}}, refine.SetCost(49500))
}

func (r Sarah) GoldBristRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082009, Amount: 1}, {ItemId: 4011006, Amount: 4}}, refine.SetCost(54000))
}

func (r Sarah) GreenMarkerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082048, Amount: 1}, {ItemId: 4021003, Amount: 3}}, refine.SetCost(13500))
}

func (r Sarah) BlackMarkerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082048, Amount: 1}, {ItemId: 4021008, Amount: 1}}, refine.SetCost(18000))
}

func (r Sarah) MithrilScalerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082068, Amount: 1}, {ItemId: 4011002, Amount: 4}}, refine.SetCost(19800))
}

func (r Sarah) GoldScalerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082068, Amount: 1}, {ItemId: 4011006, Amount: 2}}, refine.SetCost(22500))
}

func (r Sarah) GoldBraceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082071, Amount: 1}, {ItemId: 4011006, Amount: 4}}, refine.SetCost(27000))
}

func (r Sarah) DarkBraceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082071, Amount: 1}, {ItemId: 4021008, Amount: 2}}, refine.SetCost(36000))
}

func (r Sarah) RedWillowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082084, Amount: 1}, {ItemId: 4011000, Amount: 1}, {ItemId: 4021000, Amount: 5}}, refine.SetCost(49500))
}

func (r Sarah) DarkWillowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082084, Amount: 1}, {ItemId: 4011006, Amount: 2}, {ItemId: 4021008, Amount: 2}}, refine.SetCost(54000))
}

func (r Sarah) BlueLutiaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082051, Amount: 1}, {ItemId: 4021005, Amount: 3}}, refine.SetCost(31500))
}

func (r Sarah) BlackLutiaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082051, Amount: 1}, {ItemId: 4021008, Amount: 1}}, refine.SetCost(36000))
}

func (r Sarah) BlueNoelRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082054, Amount: 1}, {ItemId: 4021005, Amount: 3}}, refine.SetCost(36000))
}

func (r Sarah) DarkNoelRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082054, Amount: 1}, {ItemId: 4021008, Amount: 1}}, refine.SetCost(40500))
}

func (r Sarah) BlueArtenRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082062, Amount: 1}, {ItemId: 4021002, Amount: 4}}, refine.SetCost(40500))
}

func (r Sarah) DarkArtenRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082062, Amount: 1}, {ItemId: 4021008, Amount: 2}}, refine.SetCost(45000))
}

func (r Sarah) BluePennanceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082081, Amount: 1}, {ItemId: 4021002, Amount: 5}}, refine.SetCost(49500))
}

func (r Sarah) DarkPennanceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082081, Amount: 1}, {ItemId: 4021008, Amount: 3}}, refine.SetCost(54000))
}

func (r Sarah) SilverSylviaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082042, Amount: 1}, {ItemId: 4011004, Amount: 2}}, refine.SetCost(13500))
}

func (r Sarah) GoldSylviaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082042, Amount: 1}, {ItemId: 4011006, Amount: 1}}, refine.SetCost(18000))
}

func (r Sarah) OrihalconArbionRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082046, Amount: 1}, {ItemId: 4011005, Amount: 3}}, refine.SetCost(19800))
}

func (r Sarah) GoldArbionRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082046, Amount: 1}, {ItemId: 4011006, Amount: 2}}, refine.SetCost(22500))
}

func (r Sarah) GoldCleaveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082075, Amount: 1}, {ItemId: 4011006, Amount: 4}}, refine.SetCost(36000))
}

func (r Sarah) DarkCleaveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082075, Amount: 1}, {ItemId: 4021008, Amount: 2}}, refine.SetCost(45000))
}

func (r Sarah) RedMoonGloveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082065, Amount: 1}, {ItemId: 4021000, Amount: 5}}, refine.SetCost(49500))
}

func (r Sarah) BrownMoonGloveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: item.GlovesProductionStimulator, Amount: 1}, {ItemId: 1082065, Amount: 1}, {ItemId: 4011006, Amount: 2}, {ItemId: 4021008, Amount: 1}}, refine.SetCost(54000))
}

func (r Sarah) StimulatorError(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Eek! I think I accidentally added too much stimulator and, well, the whole thing is unusable now... Sorry, but I can't offer a refund.")
	return script.SendOk(l, c, m.String())
}

func (r Sarah) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("There, the gloves are ready. Be careful, they're still hot.")
	return script.SendOk(l, c, m.String())
}

func (r Sarah) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Sorry, we only accept meso.")
	return script.SendOk(l, c, m.String())
}

func (r Sarah) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText("Sorry, but I have to have those items to get this exactly right. Perhaps next time.")
		return script.SendOk(l, c, m.String())
	}
}

func (r Sarah) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, c, m.String())
}
