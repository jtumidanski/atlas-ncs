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

// Chrishrama is located in Dungeon - Sleepywood (105040300)
type Chrishrama struct {
}

func (r Chrishrama) NPCId() uint32 {
	return npc.Chrishrama
}

func (r Chrishrama) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return refine.NewGenericRefine(l, span, c, r.Hello(), r.Categories())
}

func (r Chrishrama) Hello() string {
	return "Hello, I live here, but don't underestimate me. How about I help you by making you a new pair of shoes?"
}

func (r Chrishrama) Categories() []refine.ListItem {
	return []refine.ListItem{
		r.WarriorShoe(),
		r.BowmanShoe(),
		r.MagicianShoe(),
		r.ThiefShoe(),
	}
}

func (r Chrishrama) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Chrishrama) WarriorShoe() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverWarBoots, " - Warrior Lv. 25"), refine.Confirm(item.SilverWarBoots, r.SilverWarBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldWarBoots, " - Warrior Lv. 25"), refine.Confirm(item.GoldWarBoots, r.GoldWarBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkWarBoots, " - Warrior Lv. 25"), refine.Confirm(item.DarkWarBoots, r.DarkWarBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.EmeraldBattleGrieve, " - Warrior Lv. 30"), refine.Confirm(item.EmeraldBattleGrieve, r.EmeraldBattleGrieveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilBattleGrieve, " - Warrior Lv. 30"), refine.Confirm(item.MithrilBattleGrieve, r.MithrilBattleGrieveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverBattleGrieve, " - Warrior Lv. 30"), refine.Confirm(item.SilverBattleGrieve, r.SilverBattleGrieveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodBattleGrieve, " - Warrior Lv. 30"), refine.Confirm(item.BloodBattleGrieve, r.BloodBattleGrieveRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SteelTrigger, " - Warrior Lv. 35"), refine.Confirm(item.SteelTrigger, r.SteelTriggerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilTrigger, " - Warrior Lv. 35"), refine.Confirm(item.MithrilTrigger, r.MithrilTriggerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkTrigger, " - Warrior Lv. 35"), refine.Confirm(item.DarkTrigger, r.DarkTriggerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownJangoonBoots, " - Warrior Lv. 40"), refine.Confirm(item.BrownJangoonBoots, r.BrownJangoonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MaroonJangoonBoots, " - Warrior Lv. 40"), refine.Confirm(item.MaroonJangoonBoots, r.MaroonJangoonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueJangoonBoots, " - Warrior Lv. 40"), refine.Confirm(item.BlueJangoonBoots, r.BlueJangoonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.EmeraldHildonBoots, " - Warrior Lv. 50"), refine.Confirm(item.EmeraldHildonBoots, r.EmeraldHildonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilHildonBoots, " - Warrior Lv. 50"), refine.Confirm(item.MithrilHildonBoots, r.MithrilHildonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrihalconHildonBoots, " - Warrior Lv. 50"), refine.Confirm(item.OrihalconHildonBoots, r.OrihalconHildonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldHildonBoots, " - Warrior Lv. 50"), refine.Confirm(item.GoldHildonBoots, r.GoldHildonBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SapphireCamelBoots, " - Warrior Lv. 60"), refine.Confirm(item.SapphireCamelBoots, r.SapphireCamelBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrihalconCamelBoots, " - Warrior Lv. 60"), refine.Confirm(item.OrihalconCamelBoots, r.OrihalconCamelBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodCamelBoots, " - Warrior Lv. 60"), refine.Confirm(item.BloodCamelBoots, r.BloodCamelBootsRequirements())),
	}
	prompt := refine.PromptCategory("Warrior shoes? Sure thing, which kind?", choices)
	return refine.ListItem{ListText: "Create a Warrior shoe", SelectionState: prompt}
}

func (r Chrishrama) BowmanShoe() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownJackBoots, " - Bowman Lv. 25"), refine.Confirm(item.BrownJackBoots, r.BrownJackBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenJackBoots, " - Bowman Lv. 25"), refine.Confirm(item.GreenJackBoots, r.GreenJackBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedJackBoots, " - Bowman Lv. 25"), refine.Confirm(item.RedJackBoots, r.RedJackBootsRequirements())),
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
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownSteelTippedBoots, " - Bowman Lv.50"), refine.Confirm(item.BrownSteelTippedBoots, r.BrownSteelTippedBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenSteelTippedBoots, " - Bowman Lv. 50"), refine.Confirm(item.GreenSteelTippedBoots, r.GreenSteelTippedBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueSteelTippedBoots, " - Bowman Lv. 50"), refine.Confirm(item.BlueSteelTippedBoots, r.BlueSteelTippedBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleSteelTippedBoots, " - Bowman Lv. 50"), refine.Confirm(item.PurpleSteelTippedBoots, r.PurpleSteelTippedBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedGoreBoots, " - Bowman Lv. 60"), refine.Confirm(item.RedGoreBoots, r.RedGoreBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueGoreBoots, " - Bowman Lv. 60"), refine.Confirm(item.BlueGoreBoots, r.BlueGoreBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenGoreBoots, " - Bowman Lv. 60"), refine.Confirm(item.GreenGoreBoots, r.GreenGoreBootsRequirements())),
	}
	prompt := refine.PromptCategory("Bowman shoes? Sure thing, which kind?", choices)
	return refine.ListItem{ListText: "Create a Bowman shoe", SelectionState: prompt}
}

func (r Chrishrama) MagicianShoe() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueJeweleryShoes, " - Magician Lv. 20"), refine.Confirm(item.BlueJeweleryShoes, r.BlueJeweleryShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleJeweleryShoes, " - Magician Lv. 20"), refine.Confirm(item.PurpleJeweleryShoes, r.PurpleJeweleryShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedJeweleryShoes, " - Magician Lv. 20"), refine.Confirm(item.RedJeweleryShoes, r.RedJeweleryShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverWindshoes, " - Magician Lv. 25"), refine.Confirm(item.SilverWindshoes, r.SilverWindshoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.YellowWindshoes, " - Magician Lv. 25"), refine.Confirm(item.YellowWindshoes, r.YellowWindshoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackWindshoes, " - Magician Lv. 25"), refine.Confirm(item.BlackWindshoes, r.BlackWindshoesRequirements())),
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
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenGoldwindShoes, " - Magician Lv. 50"), refine.Confirm(item.GreenGoldwindShoes, r.GreenGoldwindShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PinkGoldrunners, " - Magician Lv. 60"), refine.Confirm(item.PinkGoldrunners, r.PinkGoldrunnersRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenGoldrunners, " - Magician Lv. 60"), refine.Confirm(item.GreenGoldrunners, r.GreenGoldrunnersRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrangeGoldrunners, " - Magician Lv. 60"), refine.Confirm(item.OrangeGoldrunners, r.OrangeGoldrunnersRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueGoldrunners, " - Magician Lv. 60"), refine.Confirm(item.BlueGoldrunners, r.BlueGoldrunnersRequirements())),
	}
	prompt := refine.PromptCategory("Magician shoes? Sure thing, which kind?", choices)
	return refine.ListItem{ListText: "Create a Magician shoe", SelectionState: prompt}
}

func (r Chrishrama) ThiefShoe() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueLappyShoes, " - Thief Lv. 25"), refine.Confirm(item.BlueLappyShoes, r.BlueLappyShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedLappyShoes, " - Thief Lv. 25"), refine.Confirm(item.RedLappyShoes, r.RedLappyShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenLappyShoes, " - Thief Lv. 25"), refine.Confirm(item.GreenLappyShoes, r.GreenLappyShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlackLappyShoes, " - Thief Lv. 25"), refine.Confirm(item.BlackLappyShoes, r.BlackLappyShoesRequirements())),
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
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodMossBoots, " - Thief Lv. 60"), refine.Confirm(item.BloodMossBoots, r.BloodMossBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldMossBoots, " - Thief Lv. 60"), refine.Confirm(item.GoldMossBoots, r.GoldMossBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkMossBoots, " - Thief Lv. 60"), refine.Confirm(item.DarkMossBoots, r.DarkMossBootsRequirements())),
	}
	prompt := refine.PromptCategory("Thief shoes? Sure thing, which kind?", choices)
	return refine.ListItem{ListText: "Create a Thief shoe", SelectionState: prompt}
}

func (r Chrishrama) SilverWarBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 2}, {ItemId: 4011001, Amount: 1}, {ItemId: 4000021, Amount: 15}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(10000))
}

func (r Chrishrama) GoldWarBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 2}, {ItemId: 4011001, Amount: 1}, {ItemId: 4000021, Amount: 15}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(10000))
}

func (r Chrishrama) DarkWarBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 20}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(12000))
}

func (r Chrishrama) EmeraldBattleGrieveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021003, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 45}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(20000))
}

func (r Chrishrama) MithrilBattleGrieveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011002, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 45}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(20000))
}

func (r Chrishrama) SilverBattleGrieveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 45}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(20000))
}

func (r Chrishrama) BloodBattleGrieveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 4}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 45}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(20000))
}

func (r Chrishrama) SteelTriggerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 20}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(22000))
}

func (r Chrishrama) MithrilTriggerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011002, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 20}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(22000))
}

func (r Chrishrama) DarkTriggerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 2}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 20}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(25000))
}

func (r Chrishrama) BrownJangoonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011003, Amount: 4}, {ItemId: 4000021, Amount: 100}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000033, Amount: 100}}, refine.SetCost(38000))
}

func (r Chrishrama) MaroonJangoonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011005, Amount: 4}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000042, Amount: 250}}, refine.SetCost(38000))
}

func (r Chrishrama) BlueJangoonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011002, Amount: 4}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000041, Amount: 120}}, refine.SetCost(38000))
}

func (r Chrishrama) EmeraldHildonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4021003, Amount: 6}, {ItemId: 4000030, Amount: 65}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(50000))
}

func (r Chrishrama) MithrilHildonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4011002, Amount: 6}, {ItemId: 4000030, Amount: 65}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(50000))
}

func (r Chrishrama) OrihalconHildonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4011005, Amount: 6}, {ItemId: 4000030, Amount: 65}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(50000))
}

func (r Chrishrama) GoldHildonBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011001, Amount: 3}, {ItemId: 4011006, Amount: 6}, {ItemId: 4000030, Amount: 65}, {ItemId: 4003000, Amount: 45}}, refine.SetCost(50000))
}

func (r Chrishrama) SapphireCamelBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4021005, Amount: 8}, {ItemId: 4000030, Amount: 80}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(60000))
}

func (r Chrishrama) OrihalconCamelBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4011005, Amount: 8}, {ItemId: 4000030, Amount: 80}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(60000))
}

func (r Chrishrama) BloodCamelBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4021000, Amount: 8}, {ItemId: 4000030, Amount: 80}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(60000))
}

func (r Chrishrama) BrownJackBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 35}, {ItemId: 4011000, Amount: 3}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(9000))
}

func (r Chrishrama) GreenJackBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 35}, {ItemId: 4021003, Amount: 1}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(9000))
}

func (r Chrishrama) RedJackBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 35}, {ItemId: 4021000, Amount: 1}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(9000))
}

func (r Chrishrama) RedHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021000, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Chrishrama) BlueHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021005, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Chrishrama) GreenHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021003, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Chrishrama) BlackHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021004, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Chrishrama) BrownHunterBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 50}, {ItemId: 4021006, Amount: 2}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Chrishrama) BlueSilkyBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021002, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000030, Amount: 15}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(19000))
}

func (r Chrishrama) GreenSilkyBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021003, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000030, Amount: 15}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Chrishrama) RedSilkyBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000030, Amount: 15}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Chrishrama) RedPierreShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 4}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000024, Amount: 20}}, refine.SetCost(20000))
}

func (r Chrishrama) YellowPierreShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021006, Amount: 4}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000027, Amount: 20}}, refine.SetCost(32000))
}

func (r Chrishrama) BrownPierreShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011003, Amount: 5}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000044, Amount: 40}}, refine.SetCost(32000))
}

func (r Chrishrama) BluePierreShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021002, Amount: 5}, {ItemId: 4003000, Amount: 30}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000009, Amount: 120}}, refine.SetCost(40000))
}

func (r Chrishrama) BrownSteelTippedBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 35}, {ItemId: 4000033, Amount: 80}}, refine.SetCost(40000))
}

func (r Chrishrama) GreenSteelTippedBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 35}, {ItemId: 4000032, Amount: 150}}, refine.SetCost(50000))
}

func (r Chrishrama) BlueSteelTippedBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 35}, {ItemId: 4000041, Amount: 100}}, refine.SetCost(50000))
}

func (r Chrishrama) PurpleSteelTippedBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4021006, Amount: 3}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 35}, {ItemId: 4000042, Amount: 250}}, refine.SetCost(50000))
}

func (r Chrishrama) RedGoreBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 5}, {ItemId: 4021000, Amount: 8}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(50000))
}

func (r Chrishrama) BlueGoreBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 5}, {ItemId: 4021005, Amount: 8}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Chrishrama) GreenGoreBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 5}, {ItemId: 4021003, Amount: 8}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Chrishrama) BlueJeweleryShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021005, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 5}}, refine.SetCost(3000))
}

func (r Chrishrama) PurpleJeweleryShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021001, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 5}}, refine.SetCost(3000))
}

func (r Chrishrama) RedJeweleryShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 5}}, refine.SetCost(3000))
}

func (r Chrishrama) SilverWindshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 1}, {ItemId: 4000021, Amount: 35}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(8000))
}

func (r Chrishrama) YellowWindshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021006, Amount: 1}, {ItemId: 4000021, Amount: 35}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(8000))
}

func (r Chrishrama) BlackWindshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 35}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(8000))
}

func (r Chrishrama) RedMagicshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(18000))
}

func (r Chrishrama) BlueMagicshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021002, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(18000))
}

func (r Chrishrama) WhiteMagicshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(18000))
}

func (r Chrishrama) BlackMagicshoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(18000))
}

func (r Chrishrama) PurpleSaltShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021001, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Chrishrama) RedSaltShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 3}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Chrishrama) BlackSaltShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 2}, {ItemId: 4021006, Amount: 1}, {ItemId: 4000021, Amount: 40}, {ItemId: 4000030, Amount: 25}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(22000))
}

func (r Chrishrama) RedMoonShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 4}, {ItemId: 4000030, Amount: 40}, {ItemId: 4000043, Amount: 35}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(30000))
}

func (r Chrishrama) BlueMoonShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021005, Amount: 4}, {ItemId: 4000030, Amount: 40}, {ItemId: 4000037, Amount: 70}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(30000))
}

func (r Chrishrama) GoldMoonShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 2}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 40}, {ItemId: 4000027, Amount: 20}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(35000))
}

func (r Chrishrama) DarkMoonShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 2}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 40}, {ItemId: 4000014, Amount: 30}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(40000))
}

func (r Chrishrama) PinkGoldwindShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4021000, Amount: 3}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(50000))
}

func (r Chrishrama) BlueGoldwindShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4021005, Amount: 3}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(50000))
}

func (r Chrishrama) PurpleGoldwindShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4021001, Amount: 3}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(50000))
}

func (r Chrishrama) GreenGoldwindShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4021003, Amount: 3}, {ItemId: 4000030, Amount: 60}, {ItemId: 4003000, Amount: 40}}, refine.SetCost(50000))
}

func (r Chrishrama) PinkGoldrunnersRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 4}, {ItemId: 4011005, Amount: 5}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Chrishrama) GreenGoldrunnersRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 4}, {ItemId: 4021003, Amount: 5}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Chrishrama) OrangeGoldrunnersRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 4}, {ItemId: 4011003, Amount: 5}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Chrishrama) BlueGoldrunnersRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 4}, {ItemId: 4021002, Amount: 5}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Chrishrama) BlueLappyShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021005, Amount: 1}, {ItemId: 4000021, Amount: 35}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(9000))
}

func (r Chrishrama) RedLappyShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 1}, {ItemId: 4000021, Amount: 35}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(9000))
}

func (r Chrishrama) GreenLappyShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021003, Amount: 1}, {ItemId: 4000021, Amount: 35}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(9000))
}

func (r Chrishrama) BlackLappyShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 35}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(9000))
}

func (r Chrishrama) BronzeChainBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 3}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Chrishrama) IronChainBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 3}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Chrishrama) SilverChainBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(19000))
}

func (r Chrishrama) GoldChainBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(21000))
}

func (r Chrishrama) RedWhiteLineBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Chrishrama) GreenWhiteLineBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021003, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Chrishrama) BlueWhiteLineBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021002, Amount: 3}, {ItemId: 4021004, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4000030, Amount: 15}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(20000))
}

func (r Chrishrama) BlackRedLinedShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021000, Amount: 5}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000033, Amount: 50}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(40000))
}

func (r Chrishrama) BlackGreenLinedShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021003, Amount: 4}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000032, Amount: 30}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(32000))
}

func (r Chrishrama) BlackYellowLinedShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021006, Amount: 4}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000040, Amount: 3}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(35000))
}

func (r Chrishrama) BlackBlueLinedShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021005, Amount: 4}, {ItemId: 4000030, Amount: 45}, {ItemId: 4000037, Amount: 70}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(35000))
}

func (r Chrishrama) BlueGoniShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011007, Amount: 2}, {ItemId: 4021005, Amount: 3}, {ItemId: 4000030, Amount: 50}, {ItemId: 4000037, Amount: 200}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(50000))
}

func (r Chrishrama) RedGoniShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011007, Amount: 2}, {ItemId: 4021000, Amount: 3}, {ItemId: 4000030, Amount: 50}, {ItemId: 4000043, Amount: 150}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(50000))
}

func (r Chrishrama) GreenGoniShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011007, Amount: 2}, {ItemId: 4021003, Amount: 3}, {ItemId: 4000030, Amount: 50}, {ItemId: 4000045, Amount: 80}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(50000))
}

func (r Chrishrama) PurpleGoniShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011007, Amount: 2}, {ItemId: 4021001, Amount: 3}, {ItemId: 4000030, Amount: 50}, {ItemId: 4000036, Amount: 80}, {ItemId: 4003000, Amount: 35}}, refine.SetCost(50000))
}

func (r Chrishrama) BloodMossBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4021005, Amount: 8}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Chrishrama) GoldMossBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4011005, Amount: 5}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Chrishrama) DarkMossBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4021000, Amount: 1}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Chrishrama) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("There, the shoes are ready. Be careful not to trip!")
	return script.SendOk(l, span, c, m.String())
}

func (r Chrishrama) CannotAfford(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Sorry, I can only accept meso.")
	return script.SendOk(l, span, c, m.String())
}

func (r Chrishrama) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText("Sorry, but I have to have those items to get this exactly right. Perhaps next time.")
		return script.SendOk(l, span, c, m.String())
	}
}

func (r Chrishrama) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, span, c, m.String())
}
