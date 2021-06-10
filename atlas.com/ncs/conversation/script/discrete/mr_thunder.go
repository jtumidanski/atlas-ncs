package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MrThunder is located in Victoria Road - Perion (102000000)
type MrThunder struct {
}

func (r MrThunder) NPCId() uint32 {
	return npc.MrThunder
}

func (r MrThunder) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := "Hm? Who might you be? Oh, you've heard about my forging skills? In that case, I'd be glad to process some of your ores... for a fee."
	categories := r.CreateCategories()
	return refine.NewGenericRefine(l, c, hello, categories)
}

func (r MrThunder) CreateCategories() []refine.RefinementCategory {
	return []refine.RefinementCategory{
		r.RefineMineralOre(),
		r.RefineJewelOre(),
		r.UpgradeAHelmet(),
		r.UpgradeAShield(),
	}
}

func (r MrThunder) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r MrThunder) RefineMineralOre() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Refine a mineral ore",
		Prompt:          "So, what kind of mineral ore would you like to refine?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.SimpleList("Bronze"), refine.HowMany(item.BronzePlate, r.BronzeRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Steel"), refine.HowMany(item.SteelPlate, r.SteelRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Mithril"), refine.HowMany(item.MithrilPlate, r.MithrilRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Adamantium"), refine.HowMany(item.AdamantiumPlate, r.AdamantiumRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Silver"), refine.HowMany(item.SilverPlate, r.SilverRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Orihalcon"), refine.HowMany(item.OrihalconPlate, r.OrihalconRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Gold"), refine.HowMany(item.GoldPlate, r.GoldRefineRequirements())),
		},
	}
}

func (r MrThunder) RefineJewelOre() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Refine a jewel ore",
		Prompt:          "So, what kind of jewel ore would you like to refine?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.SimpleList("Garnet"), refine.HowMany(item.Garnet, r.GarnetRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Amethyst"), refine.HowMany(item.Amethyst, r.AmethystRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Aquamarine"), refine.HowMany(item.AquaMarine, r.AquamarineRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Emerald"), refine.HowMany(item.Emerald, r.EmeraldRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Opal"), refine.HowMany(item.Opal, r.OpalRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Sapphire"), refine.HowMany(item.Sapphire, r.SapphireRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Topaz"), refine.HowMany(item.Topaz, r.TopazRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Diamond"), refine.HowMany(item.Diamond, r.DiamondRefineRequirements())),
			r.CreateChoice(refine.SimpleList("Black Crystal"), refine.HowMany(item.BlackCrystal, r.BlackCrystalRefineRequirements())),
		},
	}
}

func (r MrThunder) UpgradeAHelmet() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a helmet",
		Prompt:          "Ah, you wish to upgrade a helmet? Then tell me, which one?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemNameList("Blue Metal Gear", " - Common Lv. 15"), refine.Confirm(item.BlueMetalGear, r.BlueMetalGearRequirements())),
			r.CreateChoice(refine.ItemNameList("Yellow Metal Gear", " - Common Lv. 15"), refine.Confirm(item.YellowMetalGear, r.YellowMetalGearRequirements())),
			r.CreateChoice(refine.ItemNameList("Metal Koif", " - Warrior Lv. 10"), refine.Confirm(item.MetalKoif, r.MetalKoifRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Koif", " - Warrior Lv. 10"), refine.Confirm(item.MithrilKoif, r.MithrilKoifRequirements())),
			r.CreateChoice(refine.ItemNameList("Steel Helmet", " - Warrior Lv. 12"), refine.Confirm(item.SteelHelmet, r.SteelHelmetRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Helmet", " - Warrior Lv. 12"), refine.Confirm(item.MithrilHelmet, r.MithrilHelmetRequirements())),
			r.CreateChoice(refine.ItemNameList("Steel Full Helm", " - Warrior Lv. 15"), refine.Confirm(item.SteelFullHelm, r.SteelFullHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Full Helm", " - Warrior Lv. 15"), refine.Confirm(item.MithrilFullHelm, r.MithrilFullHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Iron Viking Helm", " - Warrior Lv. 20"), refine.Confirm(item.IronVikingHelm, r.IronVikingHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Viking Helm", " - Warrior Lv. 20"), refine.Confirm(item.MithrilVikingHelm, r.MithrilVikingHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Steel Football Helmet", " - Warrior Lv. 20"), refine.Confirm(item.SteelFootballHelmet, r.SteelFootballHelmetRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Football Helmet", " - Warrior Lv. 20"), refine.Confirm(item.MithrilFootballHelmet, r.MithrilFootballHelmetRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Sharp Helm", " - Warrior Lv. 22"), refine.Confirm(item.MithrilSharpHelm, r.MithrilSharpHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Gold Sharp Helm", " - Warrior Lv. 22"), refine.Confirm(item.GoldSharpHelm, r.GoldSharpHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Orihalcon Burgernet Helm", " - Warrior Lv. 25"), refine.Confirm(item.OrihalconBurgernetHelm, r.OrihalconBurgernetHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Gold Burgernet Helm", " - Warrior Lv. 25"), refine.Confirm(item.GoldBurgernetHelm, r.GoldBurgernetHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Great Red Helmet", " - Warrior Lv. 35"), refine.Confirm(item.GreatRedHelmet, r.GreatRedHelmetRequirements())),
			r.CreateChoice(refine.ItemNameList("Great Blue Helmet", " - Warrior Lv. 35"), refine.Confirm(item.GreatBlueHelmet, r.GreatBlueHelmetRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Nordic Helm", " - Warrior Lv. 40"), refine.Confirm(item.MithrilNordicHelm, r.MithrilNordicHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Gold Nordic Helm", " - Warrior Lv. 40"), refine.Confirm(item.GoldNordicHelm, r.GoldNordicHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Crusader Helm", " - Warrior Lv. 50"), refine.Confirm(item.MithrilCrusaderHelm, r.MithrilCrusaderHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Silver Crusader Helm", " - Warrior Lv. 50"), refine.Confirm(item.SilverCrusaderHelm, r.SilverCrusaderHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Old Steel Nordic Helm", " - Warrior Lv. 55"), refine.Confirm(item.OldSteelNordicHelm, r.OldSteelNordicHelmRequirements())),
			r.CreateChoice(refine.ItemNameList("Old Mithril Nordic Helm", " - Warrior Lv. 55"), refine.Confirm(item.OldMithril, r.OldMithrilRequirements())),
		},
	}
}

func (r MrThunder) UpgradeAShield() refine.RefinementCategory {
	return refine.RefinementCategory{
		ListText:        "Upgrade a shield",
		Prompt:          "Ah, you wish to upgrade a shield? Then tell me, which one?",
		SelectionPrompt: refine.PromptCategory,
		Choices: []refine.RefinementChoice{
			r.CreateChoice(refine.ItemNameList("Adamantium Tower Shield", " - Warrior Lv. 40"), refine.Confirm(item.AdamantiumTowerShield, r.AdamantiumTowerShieldRequirements())),
			r.CreateChoice(refine.ItemNameList("Mithril Tower Shield", " - Warrior Lv. 40"), refine.Confirm(item.MithrilTowerShield, r.MithrilTowerShieldRequirements())),
			r.CreateChoice(refine.ItemNameList("Silver Legend Shield", " - Warrior Lv. 60"), refine.Confirm(item.SilverLegendShield, r.SilverLegendShieldRequirements())),
			r.CreateChoice(refine.ItemNameList("Adamantium Legend Shield", " - Warrior Lv. 60"), refine.Confirm(item.AdamantiumLegendShield, r.AdamantiumLegendShieldRequirements())),
		},
	}
}

func (r MrThunder) BronzeRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeOre, Amount: 10}}, Cost: 300}
}

func (r MrThunder) SteelRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelOre, Amount: 10}}, Cost: 300}
}

func (r MrThunder) MithrilRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.MithrilOre, Amount: 10}}, Cost: 300}
}

func (r MrThunder) AdamantiumRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.AdamantiumOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) SilverRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SilverOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) OrihalconRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.OrihalconOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) GoldRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.GoldOre, Amount: 10}}, Cost: 800}
}

func (r MrThunder) GarnetRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.GarnetOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) AmethystRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.AmethystOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) AquamarineRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.AquaMarineOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) EmeraldRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.EmeraldOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) OpalRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.OpalOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) SapphireRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SapphireOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) TopazRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.TopazOre, Amount: 10}}, Cost: 500}
}

func (r MrThunder) DiamondRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.DiamondOre, Amount: 10}}, Cost: 1000}
}

func (r MrThunder) BlackCrystalRefineRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BlackCrystalOre, Amount: 10}}, Cost: 3000}
}

func (r MrThunder) BlueMetalGearRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.MetalGear, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 1}}, Cost: 500}
}
func (r MrThunder) YellowMetalGearRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.MetalGear, Amount: 1}, {ItemId: item.Topaz, Amount: 1}}, Cost: 300}
}
func (r MrThunder) MetalKoifRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeKoif, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}}, Cost: 500}
}
func (r MrThunder) MithrilKoifRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeKoif, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 1}}, Cost: 800}
}
func (r MrThunder) SteelHelmetRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeHelmet, Amount: 1}, {ItemId: item.SteelPlate, Amount: 1}}, Cost: 500}
}
func (r MrThunder) MithrilHelmetRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeHelmet, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 1}}, Cost: 800}
}
func (r MrThunder) SteelFullHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeFullHelm, Amount: 1}, {ItemId: item.SteelPlate, Amount: 2}}, Cost: 1000}
}
func (r MrThunder) MithrilFullHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeFullHelm, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 2}}, Cost: 1500}
}
func (r MrThunder) IronVikingHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeVikingHelm, Amount: 1}, {ItemId: item.SteelPlate, Amount: 3}}, Cost: 1500}
}
func (r MrThunder) MithrilVikingHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeVikingHelm, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 3}}, Cost: 2000}
}
func (r MrThunder) SteelFootballHelmetRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeFootballHelmet, Amount: 1}, {ItemId: item.SteelPlate, Amount: 3}}, Cost: 1500}
}
func (r MrThunder) MithrilFootballHelmetRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeFootballHelmet, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 3}}, Cost: 2000}
}
func (r MrThunder) MithrilSharpHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelSharpHelm, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 4}}, Cost: 2000}
}
func (r MrThunder) GoldSharpHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelSharpHelm, Amount: 1}, {ItemId: item.GoldPlate, Amount: 4}}, Cost: 4000}
}
func (r MrThunder) OrihalconBurgernetHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.IronBurgernetHelm, Amount: 1}, {ItemId: item.OrihalconPlate, Amount: 5}}, Cost: 4000}
}
func (r MrThunder) GoldBurgernetHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.IronBurgernetHelm, Amount: 1}, {ItemId: item.GoldPlate, Amount: 5}}, Cost: 5000}
}
func (r MrThunder) GreatRedHelmetRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.GreatBrownHelmet, Amount: 1}, {ItemId: item.Garnet, Amount: 3}}, Cost: 8000}
}
func (r MrThunder) GreatBlueHelmetRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.GreatBrownHelmet, Amount: 1}, {ItemId: item.Sapphire, Amount: 3}}, Cost: 10000}
}
func (r MrThunder) MithrilNordicHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelNordicHelm, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 5}}, Cost: 12000}
}
func (r MrThunder) GoldNordicHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelNordicHelm, Amount: 1}, {ItemId: item.GoldPlate, Amount: 6}}, Cost: 15000}
}
func (r MrThunder) MithrilCrusaderHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeCrusaderHelm, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 5}}, Cost: 20000}
}
func (r MrThunder) SilverCrusaderHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.BronzeCrusaderHelm, Amount: 1}, {ItemId: item.SilverPlate, Amount: 4}}, Cost: 25000}
}
func (r MrThunder) OldSteelNordicHelmRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.OldBronzeNordicHelm, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.SteelPlate, Amount: 7}}, Cost: 30000}
}
func (r MrThunder) OldMithrilRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.OldBronzeNordicHelm, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 7}}, Cost: 30000}
}

func (r MrThunder) AdamantiumTowerShieldRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelTowerShield, Amount: 1}, {ItemId: item.AdamantiumPlate, Amount: 10}}, Cost: 100000}
}

func (r MrThunder) MithrilTowerShieldRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.SteelTowerShield, Amount: 1}, {ItemId: item.MithrilPlate, Amount: 10}}, Cost: 100000}
}

func (r MrThunder) SilverLegendShieldRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.WoodenLegendShield, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.SilverPlate, Amount: 15}}, Cost: 120000}
}

func (r MrThunder) AdamantiumLegendShieldRequirements() refine.RefinementRequirements {
	return refine.RefinementRequirements{Requirements: []refine.Requirement{{ItemId: item.WoodenLegendShield, Amount: 1}, {ItemId: item.MoonRock, Amount: 1}, {ItemId: item.AdamantiumPlate, Amount: 15}}, Cost: 120000}
}

func (r MrThunder) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, c, m.String())
}

func (r MrThunder) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I'm afraid you cannot afford my services.")
	return script.SendOk(l, c, m.String())
}

func (r MrThunder) MissingSomething(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("I'm afraid you're missing some ").
			ShowItemName1(itemId).
			AddText(". See you another time, yes?")
		return script.SendOk(l, c, m.String())
	}
}

func (r MrThunder) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("There, finished. What do you think, a piece of art, isn't it? Well, if you need anything else, you know where to find me.")
	return script.SendOk(l, c, m.String())
}
