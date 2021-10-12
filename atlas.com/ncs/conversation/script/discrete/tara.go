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

// Tara is located in Ludibrium - Tara and Sarah's House (220000303)
type Tara struct {
}

func (r Tara) NPCId() uint32 {
	return npc.Tara
}

func (r Tara) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return refine.NewGenericRefine(l, span, c, r.Hello(), r.Categories())
}

func (r Tara) Hello() string {
	return "Hello, and welcome to the Ludibrium Shoe Store. How can I help you today??"
}

func (r Tara) Categories() []refine.ListItem {
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

func (r Tara) WhatIsAStimulator() refine.ListItem {
	return refine.ListItem{
		ListText:       "What's a stimulator?",
		SelectionState: r.StimulatorInfo,
	}
}

func (r Tara) StimulatorInfo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("A stimulator is a special potion that I can add into the process of creating certain items. It gives it stats as though it had dropped from a monster. However, it is possible to have no change, and it is also possible for the item to be below average. There's also a 10% chance of not getting any item when using a stimulator, so please choose wisely.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tara) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
		StimulatorError:  r.StimulatorError,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Tara) Warrior() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.EmeraldBattleGrieve, " - Warrior Lv. 30#b"), refine.Confirm(item.EmeraldBattleGrieve, r.EmeraldBattleGrieveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilBattleGrieve, " - Warrior Lv. 30#b"), refine.Confirm(item.MithrilBattleGrieve, r.MithrilBattleGrieveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverBattleGrieve, " - Warrior Lv. 30#b"), refine.Confirm(item.SilverBattleGrieve, r.SilverBattleGrieveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodBattleGrieve, " - Warrior Lv. 30#b"), refine.Confirm(item.BloodBattleGrieve, r.BloodBattleGrieveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelTrigger, " - Warrior Lv. 35#b"), refine.Confirm(item.SteelTrigger, r.SteelTriggerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilTrigger, " - Warrior Lv. 35#b"), refine.Confirm(item.MithrilTrigger, r.MithrilTriggerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkTrigger, " - Warrior Lv. 35#b"), refine.Confirm(item.DarkTrigger, r.DarkTriggerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownJangoonBoots, " - Warrior Lv. 40#b"), refine.Confirm(item.BrownJangoonBoots, r.BrownJangoonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MaroonJangoonBoots, " - Warrior Lv. 40#b"), refine.Confirm(item.MaroonJangoonBoots, r.MaroonJangoonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueJangoonBoots, " - Warrior Lv. 40#b"), refine.Confirm(item.BlueJangoonBoots, r.BlueJangoonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.EmeraldHildonBoots, " - Warrior Lv. 50#b"), refine.Confirm(item.EmeraldHildonBoots, r.EmeraldHildonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilHildonBoots, " - Warrior Lv. 50#b"), refine.Confirm(item.MithrilHildonBoots, r.MithrilHildonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrihalconHildonBoots, " - Warrior Lv. 50#b"), refine.Confirm(item.OrihalconHildonBoots, r.OrihalconHildonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldHildonBoots, " - Warrior Lv. 50#b"), refine.Confirm(item.GoldHildonBoots, r.GoldHildonBootsRequirements())),
	}
	prompt := refine.PromptCategory("Warrior shoes? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create Warrior shoes",
		SelectionState: prompt,
	}
}

func (r Tara) Bowman() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.RedHunterBoots, r.RedHunterBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.BlueHunterBoots, r.BlueHunterBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.GreenHunterBoots, r.GreenHunterBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.BlackHunterBoots, r.BlackHunterBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.BrownHunterBoots, r.BrownHunterBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueSilkyBoots, " - Bowman Lv. 35"), refine.Confirm(item.BlueSilkyBoots, r.BlueSilkyBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenSilkyBoots, " - Bowman Lv. 35"), refine.Confirm(item.GreenSilkyBoots, r.GreenSilkyBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedSilkyBoots, " - Bowman Lv. 35"), refine.Confirm(item.RedSilkyBoots, r.RedSilkyBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedPierreShoes, " - Bowman Lv. 40"), refine.Confirm(item.RedPierreShoes, r.RedPierreShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.YellowPierreShoes, " - Bowman Lv. 40"), refine.Confirm(item.YellowPierreShoes, r.YellowPierreShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownPierreShoes, " - Bowman Lv. 40"), refine.Confirm(item.BrownPierreShoes, r.BrownPierreShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BluePierreShoes, " - Bowman Lv. 40"), refine.Confirm(item.BluePierreShoes, r.BluePierreShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownSteelTippedBoots, " - Bowman Lv. 50"), refine.Confirm(item.BrownSteelTippedBoots, r.BrownSteelTippedBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenSteelTippedBoots, " - Bowman Lv. 50"), refine.Confirm(item.GreenSteelTippedBoots, r.GreenSteelTippedBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueSteelTippedBoots, " - Bowman Lv. 50"), refine.Confirm(item.BlueSteelTippedBoots, r.BlueSteelTippedBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleSteelTippedBoots, " - Bowman Lv. 50#b"), refine.Confirm(item.PurpleSteelTippedBoots, r.PurpleSteelTippedBootsRequirements())),
	}
	prompt := refine.PromptCategory("Bowman shoes? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create Bowman shoes",
		SelectionState: prompt,
	}
}

func (r Tara) Magician() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedMagicshoes, " - Magician Lv. 30"), refine.Confirm(item.RedMagicshoes, r.RedMagicshoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueMagicshoes, " - Magician Lv. 30"), refine.Confirm(item.BlueMagicshoes, r.BlueMagicshoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.WhiteMagicshoes, " - Magician Lv. 30"), refine.Confirm(item.WhiteMagicshoes, r.WhiteMagicshoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackMagicshoes, " - Magician Lv. 30"), refine.Confirm(item.BlackMagicshoes, r.BlackMagicshoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleSaltShoes, " - Magician Lv. 35"), refine.Confirm(item.PurpleSaltShoes, r.PurpleSaltShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedSaltShoes, " - Magician Lv. 35"), refine.Confirm(item.RedSaltShoes, r.RedSaltShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackSaltShoes, " - Magician Lv. 35"), refine.Confirm(item.BlackSaltShoes, r.BlackSaltShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedMoonShoes, " - Magician Lv. 40"), refine.Confirm(item.RedMoonShoes, r.RedMoonShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueMoonShoes, " - Magician Lv. 40"), refine.Confirm(item.BlueMoonShoes, r.BlueMoonShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldMoonShoes, " - Magician Lv. 40"), refine.Confirm(item.GoldMoonShoes, r.GoldMoonShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkMoonShoes, " - Magician Lv. 40"), refine.Confirm(item.DarkMoonShoes, r.DarkMoonShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PinkGoldwindShoes, " - Magician Lv. 50"), refine.Confirm(item.PinkGoldwindShoes, r.PinkGoldwindShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueGoldwindShoes, " - Magician Lv. 50"), refine.Confirm(item.BlueGoldwindShoes, r.BlueGoldwindShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleGoldwindShoes, " - Magician Lv. 50"), refine.Confirm(item.PurpleGoldwindShoes, r.PurpleGoldwindShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenGoldwindShoes, " - Magician Lv. 50#b"), refine.Confirm(item.GreenGoldwindShoes, r.GreenGoldwindShoesRequirements())),
	}
	prompt := refine.PromptCategory("Magician shoes? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create Magician shoes",
		SelectionState: prompt,
	}
}

func (r Tara) Thief() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeChainBoots, " - Thief Lv. 30"), refine.Confirm(item.BronzeChainBoots, r.BronzeChainBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.IronChainBoots, " - Thief Lv. 30"), refine.Confirm(item.IronChainBoots, r.IronChainBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverChainBoots, " - Thief Lv. 30"), refine.Confirm(item.SilverChainBoots, r.SilverChainBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldChainBoots, " - Thief Lv. 30"), refine.Confirm(item.GoldChainBoots, r.GoldChainBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedWhiteLineBoots, " - Thief Lv. 35"), refine.Confirm(item.RedWhiteLineBoots, r.RedWhiteLineBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenWhiteLineBoots, " - Thief Lv. 35"), refine.Confirm(item.GreenWhiteLineBoots, r.GreenWhiteLineBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueWhiteLineBoots, " - Thief Lv. 35"), refine.Confirm(item.BlueWhiteLineBoots, r.BlueWhiteLineBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackRedLinedShoes, " - Thief Lv. 40"), refine.Confirm(item.BlackRedLinedShoes, r.BlackRedLinedShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackGreenLinedShoes, " - Thief Lv. 40"), refine.Confirm(item.BlackGreenLinedShoes, r.BlackGreenLinedShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackYellowLinedShoes, " - Thief Lv. 40"), refine.Confirm(item.BlackYellowLinedShoes, r.BlackYellowLinedShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackBlueLinedShoes, " - Thief Lv. 40"), refine.Confirm(item.BlackBlueLinedShoes, r.BlackBlueLinedShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueGoniShoes, " - Thief Lv. 50"), refine.Confirm(item.BlueGoniShoes, r.BlueGoniShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedGoniShoes, " - Thief Lv. 50"), refine.Confirm(item.RedGoniShoes, r.RedGoniShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenGoniShoes, " - Thief Lv. 50"), refine.Confirm(item.GreenGoniShoes, r.GreenGoniShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleGoniShoes, " - Thief Lv. 50"), refine.Confirm(item.PurpleGoniShoes, r.PurpleGoniShoesRequirements())),
	}
	prompt := refine.PromptCategory("Thief shoes? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create Thief shoes",
		SelectionState: prompt,
	}
}

func (r Tara) WarriorStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.EmeraldBattleGrieve, " - Warrior Lv. 30#b"), refine.Confirm(item.EmeraldBattleGrieve, r.EmeraldBattleGrieveRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilBattleGrieve, " - Warrior Lv. 30#b"), refine.Confirm(item.MithrilBattleGrieve, r.MithrilBattleGrieveRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverBattleGrieve, " - Warrior Lv. 30#b"), refine.Confirm(item.SilverBattleGrieve, r.SilverBattleGrieveRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodBattleGrieve, " - Warrior Lv. 30#b"), refine.Confirm(item.BloodBattleGrieve, r.BloodBattleGrieveRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelTrigger, " - Warrior Lv. 35#b"), refine.Confirm(item.SteelTrigger, r.SteelTriggerRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilTrigger, " - Warrior Lv. 35#b"), refine.Confirm(item.MithrilTrigger, r.MithrilTriggerRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkTrigger, " - Warrior Lv. 35#b"), refine.Confirm(item.DarkTrigger, r.DarkTriggerRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownJangoonBoots, " - Warrior Lv. 40#b"), refine.Confirm(item.BrownJangoonBoots, r.BrownJangoonBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MaroonJangoonBoots, " - Warrior Lv. 40#b"), refine.Confirm(item.MaroonJangoonBoots, r.MaroonJangoonBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueJangoonBoots, " - Warrior Lv. 40#b"), refine.Confirm(item.BlueJangoonBoots, r.BlueJangoonBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.EmeraldHildonBoots, " - Warrior Lv. 50#b"), refine.Confirm(item.EmeraldHildonBoots, r.EmeraldHildonBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilHildonBoots, " - Warrior Lv. 50#b"), refine.Confirm(item.MithrilHildonBoots, r.MithrilHildonBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrihalconHildonBoots, " - Warrior Lv. 50#b"), refine.Confirm(item.OrihalconHildonBoots, r.OrihalconHildonBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldHildonBoots, " - Warrior Lv. 50#b"), refine.Confirm(item.GoldHildonBoots, r.GoldHildonBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Warrior shoes? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create Warrior shoes with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Tara) BowmanStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.RedHunterBoots, r.RedHunterBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.BlueHunterBoots, r.BlueHunterBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.GreenHunterBoots, r.GreenHunterBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.BlackHunterBoots, r.BlackHunterBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownHunterBoots, " - Bowman Lv. 30"), refine.Confirm(item.BrownHunterBoots, r.BrownHunterBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueSilkyBoots, " - Bowman Lv. 35"), refine.Confirm(item.BlueSilkyBoots, r.BlueSilkyBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenSilkyBoots, " - Bowman Lv. 35"), refine.Confirm(item.GreenSilkyBoots, r.GreenSilkyBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedSilkyBoots, " - Bowman Lv. 35"), refine.Confirm(item.RedSilkyBoots, r.RedSilkyBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedPierreShoes, " - Bowman Lv. 40"), refine.Confirm(item.RedPierreShoes, r.RedPierreShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.YellowPierreShoes, " - Bowman Lv. 40"), refine.Confirm(item.YellowPierreShoes, r.YellowPierreShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownPierreShoes, " - Bowman Lv. 40"), refine.Confirm(item.BrownPierreShoes, r.BrownPierreShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BluePierreShoes, " - Bowman Lv. 40"), refine.Confirm(item.BluePierreShoes, r.BluePierreShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownSteelTippedBoots, " - Bowman Lv. 50"), refine.Confirm(item.BrownSteelTippedBoots, r.BrownSteelTippedBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenSteelTippedBoots, " - Bowman Lv. 50"), refine.Confirm(item.GreenSteelTippedBoots, r.GreenSteelTippedBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueSteelTippedBoots, " - Bowman Lv. 50"), refine.Confirm(item.BlueSteelTippedBoots, r.BlueSteelTippedBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleSteelTippedBoots, " - Bowman Lv. 50#b"), refine.Confirm(item.PurpleSteelTippedBoots, r.PurpleSteelTippedBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Bowman shoes? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create Bowman shoes with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Tara) MagicianStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedMagicshoes, " - Magician Lv. 30"), refine.Confirm(item.RedMagicshoes, r.RedMagicshoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueMagicshoes, " - Magician Lv. 30"), refine.Confirm(item.BlueMagicshoes, r.BlueMagicshoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.WhiteMagicshoes, " - Magician Lv. 30"), refine.Confirm(item.WhiteMagicshoes, r.WhiteMagicshoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackMagicshoes, " - Magician Lv. 30"), refine.Confirm(item.BlackMagicshoes, r.BlackMagicshoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleSaltShoes, " - Magician Lv. 35"), refine.Confirm(item.PurpleSaltShoes, r.PurpleSaltShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedSaltShoes, " - Magician Lv. 35"), refine.Confirm(item.RedSaltShoes, r.RedSaltShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackSaltShoes, " - Magician Lv. 35"), refine.Confirm(item.BlackSaltShoes, r.BlackSaltShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedMoonShoes, " - Magician Lv. 40"), refine.Confirm(item.RedMoonShoes, r.RedMoonShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueMoonShoes, " - Magician Lv. 40"), refine.Confirm(item.BlueMoonShoes, r.BlueMoonShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldMoonShoes, " - Magician Lv. 40"), refine.Confirm(item.GoldMoonShoes, r.GoldMoonShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkMoonShoes, " - Magician Lv. 40"), refine.Confirm(item.DarkMoonShoes, r.DarkMoonShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PinkGoldwindShoes, " - Magician Lv. 50"), refine.Confirm(item.PinkGoldwindShoes, r.PinkGoldwindShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueGoldwindShoes, " - Magician Lv. 50"), refine.Confirm(item.BlueGoldwindShoes, r.BlueGoldwindShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleGoldwindShoes, " - Magician Lv. 50"), refine.Confirm(item.PurpleGoldwindShoes, r.PurpleGoldwindShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenGoldwindShoes, " - Magician Lv. 50#b"), refine.Confirm(item.GreenGoldwindShoes, r.GreenGoldwindShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Magician shoes? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create Magician shoes with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Tara) ThiefStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeChainBoots, " - Thief Lv. 30"), refine.Confirm(item.BronzeChainBoots, r.BronzeChainBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.IronChainBoots, " - Thief Lv. 30"), refine.Confirm(item.IronChainBoots, r.IronChainBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverChainBoots, " - Thief Lv. 30"), refine.Confirm(item.SilverChainBoots, r.SilverChainBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldChainBoots, " - Thief Lv. 30"), refine.Confirm(item.GoldChainBoots, r.GoldChainBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedWhiteLineBoots, " - Thief Lv. 35"), refine.Confirm(item.RedWhiteLineBoots, r.RedWhiteLineBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenWhiteLineBoots, " - Thief Lv. 35"), refine.Confirm(item.GreenWhiteLineBoots, r.GreenWhiteLineBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueWhiteLineBoots, " - Thief Lv. 35"), refine.Confirm(item.BlueWhiteLineBoots, r.BlueWhiteLineBootsRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackRedLinedShoes, " - Thief Lv. 40"), refine.Confirm(item.BlackRedLinedShoes, r.BlackRedLinedShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackGreenLinedShoes, " - Thief Lv. 40"), refine.Confirm(item.BlackGreenLinedShoes, r.BlackGreenLinedShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackYellowLinedShoes, " - Thief Lv. 40"), refine.Confirm(item.BlackYellowLinedShoes, r.BlackYellowLinedShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackBlueLinedShoes, " - Thief Lv. 40"), refine.Confirm(item.BlackBlueLinedShoes, r.BlackBlueLinedShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueGoniShoes, " - Thief Lv. 50"), refine.Confirm(item.BlueGoniShoes, r.BlueGoniShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedGoniShoes, " - Thief Lv. 50"), refine.Confirm(item.RedGoniShoes, r.RedGoniShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenGoniShoes, " - Thief Lv. 50"), refine.Confirm(item.GreenGoniShoes, r.GreenGoniShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleGoniShoes, " - Thief Lv. 50"), refine.Confirm(item.PurpleGoniShoes, r.PurpleGoniShoesRequirements().AddRequirement(item.ShoesProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Thief shoes? Sure thing, which kind?", choices)
	return refine.ListItem{
		ListText:       "Create Thief shoes with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Tara) EmeraldBattleGrieveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021003, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 45}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(20000))
}

func (r Tara) MithrilBattleGrieveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011002, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 45}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(20000))
}

func (r Tara) SilverBattleGrieveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 45}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(20000))
}

func (r Tara) BloodBattleGrieveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 45}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(20000))
}

func (r Tara) SteelTriggerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 20}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(22000))
}

func (r Tara) MithrilTriggerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011002, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 20}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(22000))
}

func (r Tara) DarkTriggerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 2}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 20}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(25000))
}

func (r Tara) BrownJangoonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011003, Amount: 4}, {ItemId: 4000021, Amount: 100}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000103, Amount: 100}}, refine.SetCost(38000))
}

func (r Tara) MaroonJangoonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011005, Amount: 4}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000104, Amount: 100}}, refine.SetCost(38000))
}

func (r Tara) BlueJangoonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011002, Amount: 4}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000105, Amount: 100}}, refine.SetCost(38000))
}

func (r Tara) EmeraldHildonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4021003, Amount: 6}, {ItemId: 4000030, Amount: 65}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(50000))
}

func (r Tara) MithrilHildonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4011002, Amount: 6}, {ItemId: 4000030, Amount: 65}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(50000))
}

func (r Tara) OrihalconHildonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4011005, Amount: 6}, {ItemId: 4000030, Amount: 65}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(50000))
}

func (r Tara) GoldHildonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4011006, Amount: 6}, {ItemId: 4000030, Amount: 65}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(50000))
}

func (r Tara) RedHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021000, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Tara) BlueHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021005, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Tara) GreenHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021003, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Tara) BlackHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021004, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Tara) BrownHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021006, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Tara) BlueSilkyBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021002, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000030, Amount: 15}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(19000))
}

func (r Tara) GreenSilkyBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021003, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000030, Amount: 15}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Tara) RedSilkyBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000030, Amount: 15}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Tara) RedPierreShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 4}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000106, Amount: 100}}, refine.SetCost(20000))
}

func (r Tara) YellowPierreShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021006, Amount: 4}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000107, Amount: 100}}, refine.SetCost(32000))
}

func (r Tara) BrownPierreShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011003, Amount: 5}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000108, Amount: 100}}, refine.SetCost(32000))
}

func (r Tara) BluePierreShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021002, Amount: 5}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000099, Amount: 100}}, refine.SetCost(40000))
}

func (r Tara) BrownSteelTippedBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 35}, {ItemId: 4000033, Amount: 80}}, refine.SetCost(40000))
}

func (r Tara) GreenSteelTippedBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 35}, {ItemId: 4000032, Amount: 150}}, refine.SetCost(50000))
}

func (r Tara) BlueSteelTippedBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 35}, {ItemId: 4000041, Amount: 100}}, refine.SetCost(50000))
}

func (r Tara) PurpleSteelTippedBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 35}, {ItemId: 4000042, Amount: 250}}, refine.SetCost(50000))
}

func (r Tara) RedMagicshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(18000))
}

func (r Tara) BlueMagicshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021002, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(18000))
}

func (r Tara) WhiteMagicshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(18000))
}

func (r Tara) BlackMagicshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(18000))
}

func (r Tara) PurpleSaltShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021001, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Tara) RedSaltShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Tara) BlackSaltShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 2}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000021, Amount: 40}, {ItemId: 4000030, Amount: 25}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(22000))
}

func (r Tara) RedMoonShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 4}, {ItemId: 4000030, Amount: 40}, {ItemId: 4000110, Amount: 100}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(30000))
}

func (r Tara) BlueMoonShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021005, Amount: 4}, {ItemId: 4000030, Amount: 40}, {ItemId: 4000111, Amount: 100}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(30000))
}

func (r Tara) GoldMoonShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 2}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 40}, {ItemId: 4000100, Amount: 100}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(35000))
}

func (r Tara) DarkMoonShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 2}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 40}, {ItemId: 4000112, Amount: 100}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(40000))
}

func (r Tara) PinkGoldwindShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4021000, Amount: 3}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(50000))
}

func (r Tara) BlueGoldwindShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4021005, Amount: 3}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(50000))
}

func (r Tara) PurpleGoldwindShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4021001, Amount: 3}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(50000))
}

func (r Tara) GreenGoldwindShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4021003, Amount: 3}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(50000))
}

func (r Tara) BronzeChainBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 3}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Tara) IronChainBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Tara) SilverChainBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Tara) GoldChainBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(21000))
}

func (r Tara) RedWhiteLineBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Tara) GreenWhiteLineBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021003, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Tara) BlueWhiteLineBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021002, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Tara) BlackRedLinedShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 5}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000113, Amount: 100}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(40000))
}

func (r Tara) BlackGreenLinedShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021003, Amount: 4}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000095, Amount: 100}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(32000))
}

func (r Tara) BlackYellowLinedShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021006, Amount: 4}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000096, Amount: 100}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(35000))
}

func (r Tara) BlackBlueLinedShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021005, Amount: 4}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000097, Amount: 100}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(35000))
}

func (r Tara) BlueGoniShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011007, Amount: 2}, {ItemId: 4021005, Amount: 3}, {ItemId: 4000030, Amount: 50}, {ItemId: 4000114, Amount: 100}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(50000))
}

func (r Tara) RedGoniShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011007, Amount: 2}, {ItemId: 4021000, Amount: 3}, {ItemId: 4000030, Amount: 50}, {ItemId: 4000115, Amount: 100}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(50000))
}

func (r Tara) GreenGoniShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011007, Amount: 2}, {ItemId: 4021003, Amount: 3}, {ItemId: 4000030, Amount: 50}, {ItemId: 4000109, Amount: 100}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(50000))
}

func (r Tara) PurpleGoniShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011007, Amount: 2}, {ItemId: 4021001, Amount: 3}, {ItemId: 4000030, Amount: 50}, {ItemId: 4000036, Amount: 80}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(50000))
}

func (r Tara) StimulatorError(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Eek! I think I accidentally added too much stimulator and, well, the whole thing is unusable now... Sorry, but I can't offer a refund.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tara) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("There, the shoes are ready. Be careful, they're still hot.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tara) CannotAfford(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Sorry, we only accept meso.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tara) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText("Sorry, but I have to have those items to get this exactly right. Perhaps next time.")
		return script.SendOk(l, span, c, m.String())
	}
}

func (r Tara) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, span, c, m.String())
}
