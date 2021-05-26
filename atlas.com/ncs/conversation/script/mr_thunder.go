package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// MrThunder is located in Victoria Road - Perion (102000000)
type MrThunder struct {
}

func (r MrThunder) NPCId() uint32 {
	return npc.MrThunder
}

func (r MrThunder) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r MrThunder) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hm? Who might you be? Oh, you've heard about my forging skills? In that case, I'd be glad to process some of your ores... for a fee.").NewLine().
		OpenItem(0).BlueText().AddText("Refine a mineral ore").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Refine a jewel ore").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Upgrade a helmet").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Upgrade a shield").CloseItem()
	return SendListSelection(l, c, m.String(), r.WhatToDo)
}

func (r MrThunder) WhatToDo(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.RefineMineral
	case 1:
		return r.RefineJewel
	case 2:
		return r.UpgradeHelmet
	case 3:
		return r.UpgradeShield
	}
	return nil
}

func (r MrThunder) RefineMineral(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, what kind of mineral ore would you like to refine?").NewLine().
		OpenItem(0).BlueText().AddText("Bronze").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Steel").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Mithril").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Adamantium").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Silver").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Orihalcon").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Gold").CloseItem()
	return SendListSelection(l, c, m.String(), r.MineralSelection)
}

func (r MrThunder) MineralSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.HowMany(item.BronzePlate, r.BronzeRefineRequirements())
	case 1:
		return r.HowMany(item.SteelPlate, r.SteelRefineRequirements())
	case 2:
		return r.HowMany(item.MithrilPlate, r.MithrilRefineRequirements())
	case 3:
		return r.HowMany(item.AdamantiumPlate, r.AdamantiumRefineRequirements())
	case 4:
		return r.HowMany(item.SilverPlate, r.SilverRefineRequirements())
	case 5:
		return r.HowMany(item.OrihalconPlate, r.OrihalconRefineRequirements())
	case 6:
		return r.HowMany(item.GoldPlate, r.GoldRefineRequirements())
	}
	return nil
}

func (r MrThunder) RefineJewel(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, what kind of jewel ore would you like to refine?").NewLine().
		OpenItem(0).BlueText().AddText("Garnet").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Amethyst").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("AquaMarine").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Emerald").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Opal").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Sapphire").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Topaz").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Diamond").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Black Crystal").CloseItem()
	return SendListSelection(l, c, m.String(), r.JewelSelection)
}

func (r MrThunder) JewelSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.HowMany(item.Garnet, r.GarnetRefineRequirements())
	case 1:
		return r.HowMany(item.Amethyst, r.AmethystRefineRequirements())
	case 2:
		return r.HowMany(item.AquaMarine, r.AquamarineRefineRequirements())
	case 3:
		return r.HowMany(item.Emerald, r.EmeraldRefineRequirements())
	case 4:
		return r.HowMany(item.Opal, r.OpalRefineRequirements())
	case 5:
		return r.HowMany(item.Sapphire, r.SapphireRefineRequirements())
	case 6:
		return r.HowMany(item.Topaz, r.TopazRefineRequirements())
	case 7:
		return r.HowMany(item.Diamond, r.DiamondRefineRequirements())
	case 8:
		return r.HowMany(item.BlackCrystal, r.BlackCrystalRefineRequirements())
	}
	return nil
}

func (r MrThunder) UpgradeHelmet(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Ah, you wish to upgrade a helmet? Then tell me, which one?").NewLine().
		OpenItem(0).BlueText().AddText("Blue Metal Gear").BlackText().AddText(" - Common Lv. 15").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Yellow Metal Gear").BlackText().AddText(" - Common Lv. 15").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Metal Koif").BlackText().AddText(" - Warrior Lv. 10").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Mithril Koif").BlackText().AddText(" - Warrior Lv. 10").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Steel Helmet").BlackText().AddText(" - Warrior Lv. 12").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Mithril Helmet").BlackText().AddText(" - Warrior Lv. 12").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Steel Full Helm").BlackText().AddText(" - Warrior Lv. 15").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Mithril Full Helm").BlackText().AddText(" - Warrior Lv. 15").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Iron Viking Helm").BlackText().AddText(" - Warrior Lv. 20").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("Mithril Viking Helm").BlackText().AddText(" - Warrior Lv. 20").CloseItem().NewLine().
		OpenItem(10).BlueText().AddText("Steel Football Helmet").BlackText().AddText(" - Warrior Lv. 20").CloseItem().NewLine().
		OpenItem(11).BlueText().AddText("Mithril Football Helmet").BlackText().AddText(" - Warrior Lv. 20").CloseItem().NewLine().
		OpenItem(12).BlueText().AddText("Mithril Sharp Helm").BlackText().AddText(" - Warrior Lv. 22").CloseItem().NewLine().
		OpenItem(13).BlueText().AddText("Gold Sharp Helm").BlackText().AddText(" - Warrior Lv. 22").CloseItem().NewLine().
		OpenItem(14).BlueText().AddText("Orihalcon Burgernet Helm").BlackText().AddText(" - Warrior Lv. 25").CloseItem().NewLine().
		OpenItem(15).BlueText().AddText("Gold Burgernet Helm").BlackText().AddText(" - Warrior Lv. 25").CloseItem().NewLine().
		OpenItem(16).BlueText().AddText("Great Red Helmet").BlackText().AddText(" - Warrior Lv. 35").CloseItem().NewLine().
		OpenItem(17).BlueText().AddText("Great Blue Helmet").BlackText().AddText(" - Warrior Lv. 35").CloseItem().NewLine().
		OpenItem(18).BlueText().AddText("Mithril Nordic Helm").BlackText().AddText(" - Warrior Lv. 40").CloseItem().NewLine().
		OpenItem(19).BlueText().AddText("Gold Nordic Helm").BlackText().AddText(" - Warrior Lv. 40").CloseItem().NewLine().
		OpenItem(20).BlueText().AddText("Mithril Crusader Helm").BlackText().AddText(" - Warrior Lv. 50").CloseItem().NewLine().
		OpenItem(21).BlueText().AddText("Silver Crusader Helm").BlackText().AddText(" - Warrior Lv. 50").CloseItem().NewLine().
		OpenItem(22).BlueText().AddText("Old Steel Nordic Helm").BlackText().AddText(" - Warrior Lv. 55").CloseItem().NewLine().
		OpenItem(23).BlueText().AddText("Old Mithril Nordic Helm").BlackText().AddText(" - Warrior Lv. 55").CloseItem()
	return SendListSelection(l, c, m.String(), r.HelmetSelection)
}

func (r MrThunder) HelmetSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.BlueMetalGear, r.BlueMetalGearRequirements())
	case 1:
		return r.Confirm(item.YellowMetalGear, r.YellowMetalGearRequirements())
	case 2:
		return r.Confirm(item.MetalKoif, r.MetalKoifRequirements())
	case 3:
		return r.Confirm(item.MithrilKoif, r.MithrilKoifRequirements())
	case 4:
		return r.Confirm(item.SteelHelmet, r.SteelHelmetRequirements())
	case 5:
		return r.Confirm(item.MithrilHelmet, r.MithrilHelmetRequirements())
	case 6:
		return r.Confirm(item.SteelFullHelm, r.SteelFullHelmRequirements())
	case 7:
		return r.Confirm(item.MithrilFullHelm, r.MithrilFullHelmRequirements())
	case 8:
		return r.Confirm(item.IronVikingHelm, r.IronVikingHelmRequirements())
	case 9:
		return r.Confirm(item.MithrilVikingHelm, r.MithrilVikingHelmRequirements())
	case 10:
		return r.Confirm(item.SteelFootballHelmet, r.SteelFootballHelmetRequirements())
	case 11:
		return r.Confirm(item.MithrilFootballHelmet, r.MithrilFootballHelmetRequirements())
	case 12:
		return r.Confirm(item.MithrilSharpHelm, r.MithrilSharpHelmRequirements())
	case 13:
		return r.Confirm(item.GoldSharpHelm, r.GoldSharpHelmRequirements())
	case 14:
		return r.Confirm(item.OrihalconBurgernetHelm, r.OrihalconBurgernetHelmRequirements())
	case 15:
		return r.Confirm(item.GoldBurgernetHelm, r.GoldBurgernetHelmRequirements())
	case 16:
		return r.Confirm(item.GreatRedHelmet, r.GreatRedHelmetRequirements())
	case 17:
		return r.Confirm(item.GreatBlueHelmet, r.GreatBlueHelmetRequirements())
	case 18:
		return r.Confirm(item.MithrilNordicHelm, r.MithrilNordicHelmRequirements())
	case 19:
		return r.Confirm(item.GoldNordicHelm, r.GoldNordicHelmRequirements())
	case 20:
		return r.Confirm(item.MithrilCrusaderHelm, r.MithrilCrusaderHelmRequirements())
	case 21:
		return r.Confirm(item.SilverCrusaderHelm, r.SilverCrusaderHelmRequirements())
	case 22:
		return r.Confirm(item.OldSteelNordicHelm, r.OldSteelNordicHelmRequirements())
	case 23:
		return r.Confirm(item.OldMithril, r.OldMithrilRequirements())
	}
	return nil
}

func (r MrThunder) UpgradeShield(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Ah, you wish to upgrade a shield? Then tell me, which one?").NewLine().
		OpenItem(0).BlueText().AddText("Adamantium Tower Shield").BlackText().AddText(" - Warrior Lv. 40").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Mithril Tower Shield").BlackText().AddText(" - Warrior Lv. 40").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Silver Legend Shield").BlackText().AddText(" - Warrior Lv. 60").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Adamantium Legend Shield").BlackText().AddText(" - Warrior Lv. 60").CloseItem()
	return SendListSelection(l, c, m.String(), r.ShieldSelection)
}

func (r MrThunder) ShieldSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.AdamantiumTowerShield, r.AdamantiumTowerShieldRequirements())
	case 1:
		return r.Confirm(item.MithrilTowerShield, r.MithrilTowerShieldRequirements())
	case 2:
		return r.Confirm(item.SilverLegendShield, r.SilverLegendShieldRequirements())
	case 3:
		return r.Confirm(item.AdamantiumLegendShield, r.AdamantiumLegendShieldRequirements())
	}
	return nil
}

func (r MrThunder) BronzeRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeOre, amount: 10}}, cost: 300}
}

func (r MrThunder) SteelRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelOre, amount: 10}}, cost: 300}
}

func (r MrThunder) MithrilRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.MithrilOre, amount: 10}}, cost: 300}
}

func (r MrThunder) AdamantiumRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.AdamantiumOre, amount: 10}}, cost: 500}
}

func (r MrThunder) SilverRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SilverOre, amount: 10}}, cost: 500}
}

func (r MrThunder) OrihalconRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.OrihalconOre, amount: 10}}, cost: 500}
}

func (r MrThunder) GoldRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.GoldOre, amount: 10}}, cost: 800}
}

func (r MrThunder) GarnetRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.GarnetOre, amount: 10}}, cost: 500}
}

func (r MrThunder) AmethystRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.AmethystOre, amount: 10}}, cost: 500}
}

func (r MrThunder) AquamarineRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.AquaMarineOre, amount: 10}}, cost: 500}
}

func (r MrThunder) EmeraldRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.EmeraldOre, amount: 10}}, cost: 500}
}

func (r MrThunder) OpalRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.OpalOre, amount: 10}}, cost: 500}
}

func (r MrThunder) SapphireRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SapphireOre, amount: 10}}, cost: 500}
}

func (r MrThunder) TopazRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.TopazOre, amount: 10}}, cost: 500}
}

func (r MrThunder) DiamondRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.DiamondOre, amount: 10}}, cost: 1000}
}

func (r MrThunder) BlackCrystalRefineRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BlackCrystalOre, amount: 10}}, cost: 3000}
}

func (r MrThunder) BlueMetalGearRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.MetalGear, amount: 1}, {itemId: item.MithrilPlate, amount: 1}}, cost: 500}
}
func (r MrThunder) YellowMetalGearRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.MetalGear, amount: 1}, {itemId: item.Topaz, amount: 1}}, cost: 300}
}
func (r MrThunder) MetalKoifRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeKoif, amount: 1}, {itemId: item.SteelPlate, amount: 1}}, cost: 500}
}
func (r MrThunder) MithrilKoifRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeKoif, amount: 1}, {itemId: item.MithrilPlate, amount: 1}}, cost: 800}
}
func (r MrThunder) SteelHelmetRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeHelmet, amount: 1}, {itemId: item.SteelPlate, amount: 1}}, cost: 500}
}
func (r MrThunder) MithrilHelmetRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeHelmet, amount: 1}, {itemId: item.MithrilPlate, amount: 1}}, cost: 800}
}
func (r MrThunder) SteelFullHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeFullHelm, amount: 1}, {itemId: item.SteelPlate, amount: 2}}, cost: 1000}
}
func (r MrThunder) MithrilFullHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeFullHelm, amount: 1}, {itemId: item.MithrilPlate, amount: 2}}, cost: 1500}
}
func (r MrThunder) IronVikingHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeVikingHelm, amount: 1}, {itemId: item.SteelPlate, amount: 3}}, cost: 1500}
}
func (r MrThunder) MithrilVikingHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeVikingHelm, amount: 1}, {itemId: item.MithrilPlate, amount: 3}}, cost: 2000}
}
func (r MrThunder) SteelFootballHelmetRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeFootballHelmet, amount: 1}, {itemId: item.SteelPlate, amount: 3}}, cost: 1500}
}
func (r MrThunder) MithrilFootballHelmetRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeFootballHelmet, amount: 1}, {itemId: item.MithrilPlate, amount: 3}}, cost: 2000}
}
func (r MrThunder) MithrilSharpHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelSharpHelm, amount: 1}, {itemId: item.MithrilPlate, amount: 4}}, cost: 2000}
}
func (r MrThunder) GoldSharpHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelSharpHelm, amount: 1}, {itemId: item.GoldPlate, amount: 4}}, cost: 4000}
}
func (r MrThunder) OrihalconBurgernetHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.IronBurgernetHelm, amount: 1}, {itemId: item.OrihalconPlate, amount: 5}}, cost: 4000}
}
func (r MrThunder) GoldBurgernetHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.IronBurgernetHelm, amount: 1}, {itemId: item.GoldPlate, amount: 5}}, cost: 5000}
}
func (r MrThunder) GreatRedHelmetRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.GreatBrownHelmet, amount: 1}, {itemId: item.Garnet, amount: 3}}, cost: 8000}
}
func (r MrThunder) GreatBlueHelmetRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.GreatBrownHelmet, amount: 1}, {itemId: item.Sapphire, amount: 3}}, cost: 10000}
}
func (r MrThunder) MithrilNordicHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelNordicHelm, amount: 1}, {itemId: item.MithrilPlate, amount: 5}}, cost: 12000}
}
func (r MrThunder) GoldNordicHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelNordicHelm, amount: 1}, {itemId: item.GoldPlate, amount: 6}}, cost: 15000}
}
func (r MrThunder) MithrilCrusaderHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeCrusaderHelm, amount: 1}, {itemId: item.MithrilPlate, amount: 5}}, cost: 20000}
}
func (r MrThunder) SilverCrusaderHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeCrusaderHelm, amount: 1}, {itemId: item.SilverPlate, amount: 4}}, cost: 25000}
}
func (r MrThunder) OldSteelNordicHelmRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.OldBronzeNordicHelm, amount: 1}, {itemId: item.MoonRock, amount: 1}, {itemId: item.SteelPlate, amount: 7}}, cost: 30000}
}
func (r MrThunder) OldMithrilRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.OldBronzeNordicHelm, amount: 1}, {itemId: item.MoonRock, amount: 1}, {itemId: item.MithrilPlate, amount: 7}}, cost: 30000}
}

func (r MrThunder) AdamantiumTowerShieldRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelTowerShield, amount: 1}, {itemId: item.AdamantiumPlate, amount: 10}}, cost: 100000}
}

func (r MrThunder) MithrilTowerShieldRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelTowerShield, amount: 1}, {itemId: item.MithrilPlate, amount: 10}}, cost: 100000}
}

func (r MrThunder) SilverLegendShieldRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.WoodenLegendShield, amount: 1}, {itemId: item.MoonRock, amount: 1}, {itemId: item.SilverPlate, amount: 15}}, cost: 120000}
}

func (r MrThunder) AdamantiumLegendShieldRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.WoodenLegendShield, amount: 1}, {itemId: item.MoonRock, amount: 1}, {itemId: item.AdamantiumPlate, amount: 15}}, cost: 120000}
}

func (r MrThunder) HowMany(itemId uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("So, you want me to make some ").
			ShowItemName1(itemId).
			AddText("s? In that case, how many do you want me to make?")
		return SendGetNumber(l, c, m.String(), r.QuantitySelection(itemId, requirements), 1, 1, 100)
	}
}

func (r MrThunder) QuantitySelection(itemId uint32, requirements RefinementRequirements) ProcessNumber {
	return func(selection int32) StateProducer {
		return r.ConfirmQuantity(itemId, uint32(selection), requirements)
	}
}

func (r MrThunder) Confirm(itemId uint32, requirements RefinementRequirements) StateProducer {
	return r.ConfirmQuantity(itemId, 1, requirements)
}

func (r MrThunder) ConfirmQuantity(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("You want me to make ")
		if amount == 1 {
			m = m.AddText("a ").ShowItemName1(itemId)
		} else {
			m = m.AddText(fmt.Sprintf("%d ", amount)).ShowItemName1(itemId)
		}
		m = m.AddText("? In that case, I'm going to need specific items from you in order to make it. Make sure you have room in your inventory, though!").NewLine()
		for _, req := range requirements.requirements {
			m = m.ShowItemImage2(req.itemId).AddText(fmt.Sprintf(" %d ", req.amount)).ShowItemName1(req.itemId).NewLine()
		}
		if requirements.cost > 0 {
			m = m.ShowItemImage2(item.MoneySack).AddText(fmt.Sprintf(" %d meso", requirements.cost*amount))
		}
		return SendYesNo(l, c, m.String(), r.Validate(itemId, amount, requirements), Exit())
	}
}

func (r MrThunder) Validate(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.CanHoldAll(l)(c.CharacterId, itemId, amount) {
			return r.MakeRoom(l, c)
		}
		if !character.HasMeso(l)(c.CharacterId, requirements.cost*amount) {
			return r.CannotAfford(l, c)
		}
		for _, req := range requirements.requirements {
			if !character.HasItems(l)(c.CharacterId, req.itemId, uint32(req.amount)*amount) {
				return r.MissingSomething(req.itemId)(l, c)
			}
		}
		return r.PerformRefine(itemId, amount, requirements)(l, c)
	}
}

func (r MrThunder) MakeRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Check your inventory for a free slot first.")
	return SendOk(l, c, m.String())
}

func (r MrThunder) CannotAfford(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I'm afraid you cannot afford my services.")
	return SendOk(l, c, m.String())
}

func (r MrThunder) MissingSomething(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("I'm afraid you're missing some ").
			ShowItemName1(itemId).
			AddText(". See you another time, yes?")
		return SendOk(l, c, m.String())
	}
}

func (r MrThunder) PerformRefine(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, -int32(amount*requirements.cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment for refine.")
		}
		for _, req := range requirements.requirements {
			character.GainItem(l)(c.CharacterId, req.itemId, -int32(req.amount)*int32(amount))
		}
		character.GainItem(l)(c.CharacterId, itemId, int32(amount))
		return r.Success(l, c)
	}
}

func (r MrThunder) Success(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("There, finished. What do you think, a piece of art, isn't it? Well, if you need anything else, you know where to find me.")
	return SendOk(l, c, m.String())
}
