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

// Rydole is located in Ludibrium - Toy Factory <Aparatus Room> (220020600)
type Rydole struct {
}

func (r Rydole) NPCId() uint32 {
	return npc.Rydole
}

func (r Rydole) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return refine.NewGenericRefine(l, span, c, r.Hello(), r.Categories())
}

func (r Rydole) Hello() string {
	return "Ah, you've found me! I spend most of my time here, working on weapons to make for travellers like yourself. Did you have a request?"
}

func (r Rydole) Categories() []refine.ListItem {
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

func (r Rydole) WhatIsAStimulator() refine.ListItem {
	return refine.ListItem{
		ListText:       "What's a stimulator?",
		SelectionState: r.StimulatorInfo,
	}
}

func (r Rydole) StimulatorInfo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("A stimulator is a special potion that I can add into the process of creating certain items. It gives it stats as though it had dropped from a monster. However, it is possible to have no change, and it is also possible for the item to be below average. There's also a 10% chance of not getting any item when using a stimulator, so please choose wisely.")
	return script.SendOk(l, span, c, m.String())
}

func (r Rydole) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
		StimulatorError:  r.StimulatorError,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Rydole) Warrior() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Gladius, " - Lv. 30 One-Handed Sword"), refine.Confirm(item.Gladius, r.GladiusRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Cutlus, " - Lv. 35 One-Handed Sword"), refine.Confirm(item.Cutlus, r.CutlusRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Traus, " - Lv. 40 One-Handed Sword"), refine.Confirm(item.Traus, r.TrausRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.JeweledKatar, " - Lv. 50 One-Handed Sword"), refine.Confirm(item.JeweledKatar, r.JeweledKatarRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.FiremansAxe, " - Lv. 30 One-Handed Axe"), refine.Confirm(item.FiremansAxe, r.FiremansAxeRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Dankke, " - Lv. 35 One-Handed Axe"), refine.Confirm(item.Dankke, r.DankkeRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueCounter, " - Lv. 40 One-Handed Axe"), refine.Confirm(item.BlueCounter, r.BlueCounterRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Buck, " - Lv. 50 One-Handed Axe"), refine.Confirm(item.Buck, r.BuckRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.WarHammer, " - Lv. 30 One-Handed BW"), refine.Confirm(item.WarHammer, r.WarHammerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.HeavyHammer, " - Lv. 35 One-Handed BW"), refine.Confirm(item.HeavyHammer, r.HeavyHammerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Jacker, " - Lv. 40 One-Handed BW"), refine.Confirm(item.Jacker, r.JackerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.KnuckleMace, " - Lv. 50 One-Handed BW"), refine.Confirm(item.KnuckleMace, r.KnuckleMaceRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Scimitar, " - Lv. 30 Two-Handed Sword"), refine.Confirm(item.Scimitar, r.ScimitarRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Lionheart, " - Lv. 35 Two-Handed Sword"), refine.Confirm(item.Lionheart, r.LionheartRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Zard, " - Lv. 40 Two-Handed Sword"), refine.Confirm(item.Zard, r.ZardRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.LionsFang, " - Lv. 50 Two-Handed Sword"), refine.Confirm(item.LionsFang, r.LionsFangRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueAxe, " - Lv. 30 Two-Handed Axe"), refine.Confirm(item.BlueAxe, r.BlueAxeRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Niam, " - Lv. 35 Two-Handed Axe"), refine.Confirm(item.Niam, r.NiamRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Sabretooth, " - Lv. 40 Two-Handed Axe"), refine.Confirm(item.Sabretooth, r.SabretoothRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.TheRising, " - Lv. 50 Two-Handed Axe"), refine.Confirm(item.TheRising, r.TheRisingRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilMaul, " - Lv. 30 Two-Handed BW"), refine.Confirm(item.MithrilMaul, r.MithrilMaulRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Sledgehammer, " - Lv. 35 Two-Handed BW"), refine.Confirm(item.Sledgehammer, r.SledgehammerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Titan, " - Lv. 40 Two-Handed BW"), refine.Confirm(item.Titan, r.TitanRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldenMole, " - Lv. 50 Two-Handed BW"), refine.Confirm(item.GoldenMole, r.GoldenMoleRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.ForkedSpear, " - Lv. 30 Spear"), refine.Confirm(item.ForkedSpear, r.ForkedSpearRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Nakimaki, " - Lv. 35 Spear"), refine.Confirm(item.Nakimaki, r.NakimakiRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Zeco, " - Lv. 40 Spear"), refine.Confirm(item.Zeco, r.ZecoRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SerpentsTongue, " - Lv. 50 Spear"), refine.Confirm(item.SerpentsTongue, r.SerpentsTongueRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilPolearm, " - Lv. 30 Polearm"), refine.Confirm(item.MithrilPolearm, r.MithrilPolearmRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.AxePolearm, " - Lv. 35 Polearm"), refine.Confirm(item.AxePolearm, r.AxePolearmRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.CrescentPolearm, " - Lv. 40 Polearm"), refine.Confirm(item.CrescentPolearm, r.CrescentPolearmRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.TheNineDragons, " - Lv. 50 Polearm"), refine.Confirm(item.TheNineDragons, r.TheNineDragonsRequirements())),
	}
	prompt := refine.PromptCategory("Very well, then which Warrior weapon shall I work on?", choices)
	return refine.ListItem{
		ListText:       "Create a Warrior weapon",
		SelectionState: prompt,
	}
}

func (r Rydole) Bowman() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Ryden, " - Lv. 30 Bow"), refine.Confirm(item.Ryden, r.RydenRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedViper, " - Lv. 35 Bow"), refine.Confirm(item.RedViper, r.RedViperRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Vaulter2000, " - Lv. 40 Bow"), refine.Confirm(item.Vaulter2000, r.Vaulter2000Requirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Olympus, " - Lv. 50 Bow"), refine.Confirm(item.Olympus, r.OlympusRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.EagleCrow, " - Bowman Lv. 32"), refine.Confirm(item.EagleCrow, r.EagleCrowRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Heckler, " - Bowman Lv. 38"), refine.Confirm(item.Heckler, r.HecklerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverCrow, " - Bowman Lv. 42"), refine.Confirm(item.SilverCrow, r.SilverCrowRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Rower, " - Bowman Lv. 50"), refine.Confirm(item.Rower, r.RowerRequirements())),
	}
	prompt := refine.PromptCategory("Very well, then which Bowman weapon shall I work on?", choices)
	return refine.ListItem{
		ListText:       "Create a Bowman weapon",
		SelectionState: prompt,
	}
}

func (r Rydole) Magician() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilWand, " - Lv. 28 Wand"), refine.Confirm(item.MithrilWand, r.MithrilWandRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.WizardWand, " - Lv. 33 Wand"), refine.Confirm(item.WizardWand, r.WizardWandRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.FairyWand, " - Lv. 38 Wand"), refine.Confirm(item.FairyWand, r.FairyWandRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Cromi, " - Lv. 48 Wand"), refine.Confirm(item.Cromi, r.CromiRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.WizardStaff, " - Lv. 25 Staff"), refine.Confirm(item.WizardStaff, r.WizardStaffRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.ArcStaff, " - Lv. 45 Staff"), refine.Confirm(item.ArcStaff, r.ArcStaffRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Thorns, " - Lv. 55 Staff"), refine.Confirm(item.Thorns, r.ThornsRequirements())),
	}
	prompt := refine.PromptCategory("Very well, then which Magician weapon shall I work on?", choices)
	return refine.ListItem{
		ListText:       "Create a Magician weapon",
		SelectionState: prompt,
	}
}

func (r Rydole) Thief() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.ReefClaw, " - Lv. 30 LUK Dagger"), refine.Confirm(item.ReefClaw, r.ReefClawRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Cass, " - Lv. 30 STR Dagger"), refine.Confirm(item.Cass, r.CassRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Gephart, " - Lv. 35 LUK Dagger"), refine.Confirm(item.Gephart, r.GephartRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Bazlud, " - Lv. 40 STR Dagger"), refine.Confirm(item.Bazlud, r.BazludRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Sai, " - Lv. 50 STR Dagger"), refine.Confirm(item.Sai, r.SaiRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Shinkita, " - Lv. 50 LUK Dagger"), refine.Confirm(item.Shinkita, r.ShinkitaRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelGuards, " - Lv. 30 Claw"), refine.Confirm(item.SteelGuards, r.SteelGuardsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeGuardian, " - Lv. 35 Claw"), refine.Confirm(item.BronzeGuardian, r.BronzeGuardianRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelAvarice, " - Lv. 40 Claw"), refine.Confirm(item.SteelAvarice, r.SteelAvariceRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelSlain, " - Lv. 50 Claw"), refine.Confirm(item.SteelSlain, r.SteelSlainRequirements())),
	}
	prompt := refine.PromptCategory("Very well, then which Thief weapon shall I work on?", choices)
	return refine.ListItem{
		ListText:       "Create a Thief weapon",
		SelectionState: prompt,
	}
}

func (r Rydole) WarriorStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Gladius, " - Lv. 30 One-Handed Sword"), refine.Confirm(item.Gladius, r.GladiusRequirements().AddRequirement(item.OneHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Cutlus, " - Lv. 35 One-Handed Sword"), refine.Confirm(item.Cutlus, r.CutlusRequirements().AddRequirement(item.OneHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Traus, " - Lv. 40 One-Handed Sword"), refine.Confirm(item.Traus, r.TrausRequirements().AddRequirement(item.OneHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.JeweledKatar, " - Lv. 50 One-Handed Sword"), refine.Confirm(item.JeweledKatar, r.JeweledKatarRequirements().AddRequirement(item.OneHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.FiremansAxe, " - Lv. 30 One-Handed Axe"), refine.Confirm(item.FiremansAxe, r.FiremansAxeRequirements().AddRequirement(item.OneHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Dankke, " - Lv. 35 One-Handed Axe"), refine.Confirm(item.Dankke, r.DankkeRequirements().AddRequirement(item.OneHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueCounter, " - Lv. 40 One-Handed Axe"), refine.Confirm(item.BlueCounter, r.BlueCounterRequirements().AddRequirement(item.OneHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Buck, " - Lv. 50 One-Handed Axe"), refine.Confirm(item.Buck, r.BuckRequirements().AddRequirement(item.OneHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.WarHammer, " - Lv. 30 One-Handed BW"), refine.Confirm(item.WarHammer, r.WarHammerRequirements().AddRequirement(item.OneHandedBluntWeaponForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.HeavyHammer, " - Lv. 35 One-Handed BW"), refine.Confirm(item.HeavyHammer, r.HeavyHammerRequirements().AddRequirement(item.OneHandedBluntWeaponForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Jacker, " - Lv. 40 One-Handed BW"), refine.Confirm(item.Jacker, r.JackerRequirements().AddRequirement(item.OneHandedBluntWeaponForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.KnuckleMace, " - Lv. 50 One-Handed BW"), refine.Confirm(item.KnuckleMace, r.KnuckleMaceRequirements().AddRequirement(item.OneHandedBluntWeaponForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Scimitar, " - Lv. 30 Two-Handed Sword"), refine.Confirm(item.Scimitar, r.ScimitarRequirements().AddRequirement(item.TwoHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Lionheart, " - Lv. 35 Two-Handed Sword"), refine.Confirm(item.Lionheart, r.LionheartRequirements().AddRequirement(item.TwoHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Zard, " - Lv. 40 Two-Handed Sword"), refine.Confirm(item.Zard, r.ZardRequirements().AddRequirement(item.TwoHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.LionsFang, " - Lv. 50 Two-Handed Sword"), refine.Confirm(item.LionsFang, r.LionsFangRequirements().AddRequirement(item.TwoHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueAxe, " - Lv. 30 Two-Handed Axe"), refine.Confirm(item.BlueAxe, r.BlueAxeRequirements().AddRequirement(item.TwoHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Niam, " - Lv. 35 Two-Handed Axe"), refine.Confirm(item.Niam, r.NiamRequirements().AddRequirement(item.TwoHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Sabretooth, " - Lv. 40 Two-Handed Axe"), refine.Confirm(item.Sabretooth, r.SabretoothRequirements().AddRequirement(item.TwoHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.TheRising, " - Lv. 50 Two-Handed Axe"), refine.Confirm(item.TheRising, r.TheRisingRequirements().AddRequirement(item.TwoHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilMaul, " - Lv. 30 Two-Handed BW"), refine.Confirm(item.MithrilMaul, r.MithrilMaulRequirements().AddRequirement(item.TwoHandedMaceForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Sledgehammer, " - Lv. 35 Two-Handed BW"), refine.Confirm(item.Sledgehammer, r.SledgehammerRequirements().AddRequirement(item.TwoHandedMaceForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Titan, " - Lv. 40 Two-Handed BW"), refine.Confirm(item.Titan, r.TitanRequirements().AddRequirement(item.TwoHandedMaceForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldenMole, " - Lv. 50 Two-Handed BW"), refine.Confirm(item.GoldenMole, r.GoldenMoleRequirements().AddRequirement(item.TwoHandedMaceForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.ForkedSpear, " - Lv. 30 Spear"), refine.Confirm(item.ForkedSpear, r.ForkedSpearRequirements().AddRequirement(item.SpearForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Nakimaki, " - Lv. 35 Spear"), refine.Confirm(item.Nakimaki, r.NakimakiRequirements().AddRequirement(item.SpearForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Zeco, " - Lv. 40 Spear"), refine.Confirm(item.Zeco, r.ZecoRequirements().AddRequirement(item.SpearForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SerpentsTongue, " - Lv. 50 Spear"), refine.Confirm(item.SerpentsTongue, r.SerpentsTongueRequirements().AddRequirement(item.SpearForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilPolearm, " - Lv. 30 Polearm"), refine.Confirm(item.MithrilPolearm, r.MithrilPolearmRequirements().AddRequirement(item.PoleArmForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.AxePolearm, " - Lv. 35 Polearm"), refine.Confirm(item.AxePolearm, r.AxePolearmRequirements().AddRequirement(item.PoleArmForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.CrescentPolearm, " - Lv. 40 Polearm"), refine.Confirm(item.CrescentPolearm, r.CrescentPolearmRequirements().AddRequirement(item.PoleArmForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.TheNineDragons, " - Lv. 50 Polearm"), refine.Confirm(item.TheNineDragons, r.TheNineDragonsRequirements().AddRequirement(item.PoleArmForgingStimulator, 1))),
	}
	prompt := refine.PromptCategory("Very well, then which Warrior weapon shall I work on?", choices)
	return refine.ListItem{
		ListText:       "Create a Warrior weapon with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Rydole) BowmanStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Ryden, " - Lv. 30 Bow"), refine.Confirm(item.Ryden, r.RydenRequirements().AddRequirement(item.BowProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedViper, " - Lv. 35 Bow"), refine.Confirm(item.RedViper, r.RedViperRequirements().AddRequirement(item.BowProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Vaulter2000, " - Lv. 40 Bow"), refine.Confirm(item.Vaulter2000, r.Vaulter2000Requirements().AddRequirement(item.BowProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Olympus, " - Lv. 50 Bow"), refine.Confirm(item.Olympus, r.OlympusRequirements().AddRequirement(item.BowProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.EagleCrow, " - Bowman Lv. 32"), refine.Confirm(item.EagleCrow, r.EagleCrowRequirements().AddRequirement(item.CrossbowProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Heckler, " - Bowman Lv. 38"), refine.Confirm(item.Heckler, r.HecklerRequirements().AddRequirement(item.CrossbowProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverCrow, " - Bowman Lv. 42"), refine.Confirm(item.SilverCrow, r.SilverCrowRequirements().AddRequirement(item.CrossbowProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Rower, " - Bowman Lv. 50"), refine.Confirm(item.Rower, r.RowerRequirements().AddRequirement(item.CrossbowProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Very well, then which Bowman weapon shall I work on?", choices)
	return refine.ListItem{
		ListText:       "Create a Bowman weapon with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Rydole) MagicianStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilWand, " - Lv. 28 Wand"), refine.Confirm(item.MithrilWand, r.MithrilWandRequirements().AddRequirement(item.WandProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.WizardWand, " - Lv. 33 Wand"), refine.Confirm(item.WizardWand, r.WizardWandRequirements().AddRequirement(item.WandProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.FairyWand, " - Lv. 38 Wand"), refine.Confirm(item.FairyWand, r.FairyWandRequirements().AddRequirement(item.WandProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Cromi, " - Lv. 48 Wand"), refine.Confirm(item.Cromi, r.CromiRequirements().AddRequirement(item.WandProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.WizardStaff, " - Lv. 25 Staff"), refine.Confirm(item.WizardStaff, r.WizardStaffRequirements().AddRequirement(item.StaffProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.ArcStaff, " - Lv. 45 Staff"), refine.Confirm(item.ArcStaff, r.ArcStaffRequirements().AddRequirement(item.StaffProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Thorns, " - Lv. 55 Staff"), refine.Confirm(item.Thorns, r.ThornsRequirements().AddRequirement(item.StaffProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Very well, then which Magician weapon shall I work on?", choices)
	return refine.ListItem{
		ListText:       "Create a Magician weapon with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Rydole) ThiefStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.ReefClaw, " - Lv. 30 LUK Dagger"), refine.Confirm(item.ReefClaw, r.ReefClawRequirements().AddRequirement(item.DaggerForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Cass, " - Lv. 30 STR Dagger"), refine.Confirm(item.Cass, r.CassRequirements().AddRequirement(item.DaggerForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Gephart, " - Lv. 35 LUK Dagger"), refine.Confirm(item.Gephart, r.GephartRequirements().AddRequirement(item.DaggerForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Bazlud, " - Lv. 40 STR Dagger"), refine.Confirm(item.Bazlud, r.BazludRequirements().AddRequirement(item.DaggerForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Sai, " - Lv. 50 STR Dagger"), refine.Confirm(item.Sai, r.SaiRequirements().AddRequirement(item.DaggerForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Shinkita, " - Lv. 50 LUK Dagger"), refine.Confirm(item.Shinkita, r.ShinkitaRequirements().AddRequirement(item.DaggerForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilGuards, " - Lv. 30 Claw"), refine.Confirm(item.MithrilGuards, r.MithrilGuardsRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.AdamantiumGuards, " - Lv. 30 Claw"), refine.Confirm(item.AdamantiumGuards, r.AdamantiumGuardsRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverGuardian, " - Lv. 35 Claw"), refine.Confirm(item.SilverGuardian, r.SilverGuardianRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkGuardian, " - Lv. 35 Claw"), refine.Confirm(item.DarkGuardian, r.DarkGuardianRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodAvarice, " - Lv. 40 Claw"), refine.Confirm(item.BloodAvarice, r.BloodAvariceRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.AdamantiumAvarice, " - Lv. 40 Claw"), refine.Confirm(item.AdamantiumAvarice, r.AdamantiumAvariceRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkAvarice, " - Lv. 40 Claw"), refine.Confirm(item.DarkAvarice, r.DarkAvariceRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodSlain, " - Lv. 50 Claw"), refine.Confirm(item.BloodSlain, r.BloodSlainRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SapphireSlain, " - Lv. 50 Claw"), refine.Confirm(item.SapphireSlain, r.SapphireSlainRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkSlain, " - Lv. 50 Claw"), refine.Confirm(item.DarkSlain, r.DarkSlainRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Very well, then which Thief weapon shall I work on?", choices)
	return refine.ListItem{
		ListText:       "Create a Thief weapon with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Rydole) GladiusRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131000, Amount: 1}, {ItemId: 4011001, Amount: 2}, {ItemId: 4011004, Amount: 2}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(18000))
}

func (r Rydole) CutlusRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131000, Amount: 1}, {ItemId: 4011006, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4021006, Amount: 3}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(35000))
}

func (r Rydole) TrausRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131000, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4011001, Amount: 5}, {ItemId: 4021000, Amount: 5}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(70000))
}

func (r Rydole) JeweledKatarRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131000, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4011006, Amount: 4}, {ItemId: 4021003, Amount: 10}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(200000))
}

func (r Rydole) FiremansAxeRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131001, Amount: 1}, {ItemId: 4011001, Amount: 2}, {ItemId: 4021000, Amount: 2}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(18000))
}

func (r Rydole) DankkeRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131001, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4021000, Amount: 5}, {ItemId: 4011004, Amount: 3}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(35000))
}

func (r Rydole) BlueCounterRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131001, Amount: 1}, {ItemId: 4021005, Amount: 7}, {ItemId: 4011001, Amount: 5}, {ItemId: 4021001, Amount: 5}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(70000))
}

func (r Rydole) BuckRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131001, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4011004, Amount: 8}, {ItemId: 4011001, Amount: 10}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(200000))
}

func (r Rydole) WarHammerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131002, Amount: 1}, {ItemId: 4011001, Amount: 2}, {ItemId: 4011000, Amount: 2}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(18000))
}

func (r Rydole) HeavyHammerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131002, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011000, Amount: 5}, {ItemId: 4011003, Amount: 3}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(35000))
}

func (r Rydole) JackerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131002, Amount: 1}, {ItemId: 4011003, Amount: 7}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011004, Amount: 5}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(70000))
}

func (r Rydole) KnuckleMaceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131002, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4011006, Amount: 4}, {ItemId: 4011001, Amount: 10}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(200000))
}

func (r Rydole) ScimitarRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131003, Amount: 1}, {ItemId: 4011001, Amount: 2}, {ItemId: 4021000, Amount: 1}, {ItemId: 4021004, Amount: 2}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(20000))
}

func (r Rydole) LionheartRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131003, Amount: 1}, {ItemId: 4011006, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4021004, Amount: 5}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(37000))
}

func (r Rydole) ZardRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131003, Amount: 1}, {ItemId: 4021003, Amount: 7}, {ItemId: 4011000, Amount: 5}, {ItemId: 4011001, Amount: 5}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(72000))
}

func (r Rydole) LionsFangRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131003, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021007, Amount: 2}, {ItemId: 4011006, Amount: 4}, {ItemId: 4011001, Amount: 10}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(220000))
}

func (r Rydole) BlueAxeRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131004, Amount: 1}, {ItemId: 4021005, Amount: 2}, {ItemId: 4011001, Amount: 2}, {ItemId: 4003001, Amount: 5}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(20000))
}

func (r Rydole) NiamRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131004, Amount: 1}, {ItemId: 4011004, Amount: 5}, {ItemId: 4011000, Amount: 5}, {ItemId: 4021003, Amount: 3}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(37000))
}

func (r Rydole) SabretoothRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131004, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4011004, Amount: 5}, {ItemId: 4011001, Amount: 5}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(72000))
}

func (r Rydole) TheRisingRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131004, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021007, Amount: 2}, {ItemId: 4011006, Amount: 5}, {ItemId: 4021006, Amount: 7}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(220000))
}

func (r Rydole) MithrilMaulRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131005, Amount: 1}, {ItemId: 4011001, Amount: 2}, {ItemId: 4011004, Amount: 3}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(20000))
}

func (r Rydole) SledgehammerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131005, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011000, Amount: 5}, {ItemId: 4003001, Amount: 10}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(37000))
}

func (r Rydole) TitanRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131005, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011004, Amount: 5}, {ItemId: 4011006, Amount: 3}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(72000))
}

func (r Rydole) GoldenMoleRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131005, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4021006, Amount: 7}, {ItemId: 4011006, Amount: 5}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(220000))
}

func (r Rydole) ForkedSpearRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131006, Amount: 1}, {ItemId: 4011000, Amount: 2}, {ItemId: 4011004, Amount: 3}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(22000))
}

func (r Rydole) NakimakiRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131006, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011002, Amount: 5}, {ItemId: 4021000, Amount: 3}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(39000))
}

func (r Rydole) ZecoRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131006, Amount: 1}, {ItemId: 4011004, Amount: 3}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011000, Amount: 5}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(74000))
}

func (r Rydole) SerpentsTongueRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131006, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4011000, Amount: 7}, {ItemId: 4021000, Amount: 5}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(240000))
}

func (r Rydole) MithrilPolearmRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131007, Amount: 1}, {ItemId: 4011000, Amount: 2}, {ItemId: 4011002, Amount: 3}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(22000))
}

func (r Rydole) AxePolearmRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131007, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011002, Amount: 5}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(39000))
}

func (r Rydole) CrescentPolearmRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131007, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4011002, Amount: 5}, {ItemId: 4011001, Amount: 5}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(74000))
}

func (r Rydole) TheNineDragonsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131007, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021007, Amount: 2}, {ItemId: 4011001, Amount: 7}, {ItemId: 4011002, Amount: 5}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(240000))
}

func (r Rydole) RydenRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131010, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011006, Amount: 5}, {ItemId: 4021003, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(15000))
}

func (r Rydole) RedViperRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131010, Amount: 1}, {ItemId: 4011004, Amount: 7}, {ItemId: 4021000, Amount: 6}, {ItemId: 4021004, Amount: 3}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(20000))
}

func (r Rydole) Vaulter2000Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131010, Amount: 1}, {ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 10}, {ItemId: 4011006, Amount: 3}, {ItemId: 4003000, Amount: 40}, {ItemId: 4000112, Amount: 100}}, refine.SetCost(40000))
}

func (r Rydole) OlympusRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131010, Amount: 1}, {ItemId: 4005002, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4011001, Amount: 10}, {ItemId: 4021005, Amount: 6}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(100000))
}

func (r Rydole) EagleCrowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131011, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011005, Amount: 5}, {ItemId: 4021006, Amount: 3}, {ItemId: 4003001, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(15000))
}

func (r Rydole) HecklerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131011, Amount: 1}, {ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 8}, {ItemId: 4011006, Amount: 4}, {ItemId: 4021006, Amount: 2}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(25000))
}

func (r Rydole) SilverCrowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131011, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4011004, Amount: 6}, {ItemId: 4003001, Amount: 30}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(41000))
}

func (r Rydole) RowerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131011, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4011006, Amount: 5}, {ItemId: 4021006, Amount: 3}, {ItemId: 4003001, Amount: 40}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(100000))
}

func (r Rydole) MithrilWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131008, Amount: 1}, {ItemId: 4011002, Amount: 3}, {ItemId: 4021002, Amount: 1}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(15000))
}

func (r Rydole) WizardWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131008, Amount: 1}, {ItemId: 4021006, Amount: 5}, {ItemId: 4011002, Amount: 3}, {ItemId: 4011001, Amount: 1}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(30000))
}

func (r Rydole) FairyWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131008, Amount: 1}, {ItemId: 4021006, Amount: 5}, {ItemId: 4021005, Amount: 5}, {ItemId: 4021007, Amount: 1}, {ItemId: 4003003, Amount: 1}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(60000))
}

func (r Rydole) CromiRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131008, Amount: 1}, {ItemId: 4011006, Amount: 4}, {ItemId: 4021003, Amount: 3}, {ItemId: 4021007, Amount: 2}, {ItemId: 4021002, Amount: 1}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(100000))
}

func (r Rydole) WizardStaffRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131009, Amount: 1}, {ItemId: 4021006, Amount: 2}, {ItemId: 4021001, Amount: 1}, {ItemId: 4011001, Amount: 1}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(10000))
}

func (r Rydole) ArcStaffRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131009, Amount: 1}, {ItemId: 4011001, Amount: 8}, {ItemId: 4021006, Amount: 5}, {ItemId: 4021001, Amount: 5}, {ItemId: 4021005, Amount: 5}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(80000))
}

func (r Rydole) ThornsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131009, Amount: 1}, {ItemId: 4005001, Amount: 2}, {ItemId: 4021008, Amount: 2}, {ItemId: 4011006, Amount: 5}, {ItemId: 4011004, Amount: 10}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(200000))
}

func (r Rydole) ReefClawRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131012, Amount: 1}, {ItemId: 4011002, Amount: 2}, {ItemId: 4011001, Amount: 3}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(20000))
}

func (r Rydole) CassRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131012, Amount: 1}, {ItemId: 4021005, Amount: 2}, {ItemId: 4011001, Amount: 3}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(20000))
}

func (r Rydole) GephartRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131012, Amount: 1}, {ItemId: 4021005, Amount: 1}, {ItemId: 4011001, Amount: 5}, {ItemId: 4011002, Amount: 3}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(33000))
}

func (r Rydole) BazludRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131012, Amount: 1}, {ItemId: 4011001, Amount: 7}, {ItemId: 4011006, Amount: 3}, {ItemId: 4021006, Amount: 6}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(73000))
}

func (r Rydole) SaiRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131012, Amount: 1}, {ItemId: 4005003, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4011004, Amount: 7}, {ItemId: 4011001, Amount: 10}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(230000))
}

func (r Rydole) ShinkitaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131012, Amount: 1}, {ItemId: 4005003, Amount: 1}, {ItemId: 4021007, Amount: 2}, {ItemId: 4011006, Amount: 5}, {ItemId: 4011001, Amount: 10}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(230000))
}

func (r Rydole) SteelGuardsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 4011000, Amount: 3}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(15000))
}

func (r Rydole) BronzeGuardianRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 4011000, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 80}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(30000))
}

func (r Rydole) SteelAvariceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 4011000, Amount: 3}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 100}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(40000))
}

func (r Rydole) SteelSlainRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 4011000, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(50000))
}

func (r Rydole) MithrilGuardsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472008, Amount: 1}, {ItemId: 4011002, Amount: 3}}, refine.SetCost(10000))
}

func (r Rydole) AdamantiumGuardsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472008, Amount: 1}, {ItemId: 4011003, Amount: 3}}, refine.SetCost(15000))
}

func (r Rydole) SilverGuardianRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472011, Amount: 1}, {ItemId: 4011004, Amount: 4}}, refine.SetCost(20000))
}

func (r Rydole) DarkGuardianRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472011, Amount: 1}, {ItemId: 4021008, Amount: 1}}, refine.SetCost(25000))
}

func (r Rydole) BloodAvariceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472014, Amount: 1}, {ItemId: 4021000, Amount: 5}}, refine.SetCost(30000))
}

func (r Rydole) AdamantiumAvariceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472014, Amount: 1}, {ItemId: 4011003, Amount: 5}}, refine.SetCost(30000))
}

func (r Rydole) DarkAvariceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472014, Amount: 1}, {ItemId: 4021008, Amount: 2}}, refine.SetCost(35000))
}

func (r Rydole) BloodSlainRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472018, Amount: 1}, {ItemId: 4021000, Amount: 6}}, refine.SetCost(40000))
}

func (r Rydole) SapphireSlainRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472018, Amount: 1}, {ItemId: 4021005, Amount: 6}}, refine.SetCost(40000))
}

func (r Rydole) DarkSlainRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4131013, Amount: 1}, {ItemId: 1472018, Amount: 1}, {ItemId: 4005003, Amount: 1}, {ItemId: 4021008, Amount: 3}}, refine.SetCost(50000))
}

func (r Rydole) StimulatorError(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("...ACK! My attention wandered, and before I knew it... Uh, sorry, but there's nothing I can do for you now.")
	return script.SendOk(l, span, c, m.String())
}

func (r Rydole) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Heeere you go! What do you think? Marvellous, isn't it?")
	return script.SendOk(l, span, c, m.String())
}

func (r Rydole) CannotAfford(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm afraid my fees are non-negotiable.")
	return script.SendOk(l, span, c, m.String())
}

func (r Rydole) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText("Sorry, but you're missing a required item. Possibly a manual? Or one of the ores?")
		return script.SendOk(l, span, c, m.String())
	}
}

func (r Rydole) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Verify for a slot in your inventory first.")
	return script.SendOk(l, span, c, m.String())
}