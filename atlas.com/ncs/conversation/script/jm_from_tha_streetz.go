package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// JMFromThaStreetz is located in Victoria Road - Kerning City (103000000)
type JMFromThaStreetz struct {
}

func (r JMFromThaStreetz) NPCId() uint32 {
	return npc.JMFromThaStreetz
}

func (r JMFromThaStreetz) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r JMFromThaStreetz) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Pst... If you have the right goods, I can turn it into something nice...").NewLine().
		OpenItem(0).BlueText().AddText("Create a glove").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Upgrade a glove").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Create a claw").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Upgrade a claw").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Create materials").CloseItem()
	return SendListSelection(l, c, m.String(), r.WhatToDo)
}

func (r JMFromThaStreetz) WhatToDo(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.CreateGlove
	case 1:
		return r.UpgradeGlove
	case 2:
		return r.CreateClaw
	case 3:
		return r.UpgradeClaw
	case 4:
		return r.CreateMaterials
	}
	return nil
}

func (r JMFromThaStreetz) CreateGlove(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, what kind of glove would you like me to make?").NewLine().
		OpenItem(0).BlueText().AddText("Work Gloves").BlackText().AddText(" - Common Lv. 10").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Brown Duo").BlackText().AddText(" - Thief Lv. 15").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Blue Duo").BlackText().AddText(" - Thief Lv. 15").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Black Duo").BlackText().AddText(" - Thief Lv. 15").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Bronze Mischief").BlackText().AddText(" - Thief Lv. 20").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Bronze Wolfskin").BlackText().AddText(" - Thief Lv. 25").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Steel Sylvia").BlackText().AddText(" - Thief Lv. 30").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Steel Arbion").BlackText().AddText(" - Thief Lv. 35").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Red Cleave").BlackText().AddText(" - Thief Lv. 40").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("Blue Moon Glove").BlackText().AddText(" - Thief Lv. 50").CloseItem().NewLine().
		OpenItem(10).BlueText().AddText("Bronze Pow").BlackText().AddText(" - Thief Lv. 60").CloseItem()
	return SendListSelection(l, c, m.String(), r.CreateGloveSelection)
}

func (r JMFromThaStreetz) UpgradeGlove(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("An upgraded glove? Sure thing, but note that upgrades won't carry over to the new item... ").NewLine().
		OpenItem(0).BlueText().AddText("Mithril Mischief").BlackText().AddText(" - Thief Lv. 20").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Dark Mischief").BlackText().AddText(" - Thief Lv. 20").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Mithril Wolfskin").BlackText().AddText(" - Thief Lv. 25").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Dark Wolfskin").BlackText().AddText(" - Thief Lv. 25").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Silver Sylvia").BlackText().AddText(" - Thief Lv. 30").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Gold Sylvia").BlackText().AddText(" - Thief Lv. 30").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Orihalcon Arbion").BlackText().AddText(" - Thief Lv. 35").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Gold Arbion").BlackText().AddText(" - Thief Lv. 35").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Gold Cleave").BlackText().AddText(" - Thief Lv. 40").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("Dark Cleave").BlackText().AddText(" - Thief Lv. 40").CloseItem().NewLine().
		OpenItem(10).BlueText().AddText("Red Moon Glove").BlackText().AddText(" - Thief Lv. 50").CloseItem().NewLine().
		OpenItem(11).BlueText().AddText("Brown Moon Glove").BlackText().AddText(" - Thief Lv. 50").CloseItem().NewLine().
		OpenItem(12).BlueText().AddText("Steal Pow").BlackText().AddText(" - Thief Lv. 60").CloseItem().NewLine().
		OpenItem(13).BlueText().AddText("Gold Pow").BlackText().AddText(" - Thief Lv. 60").CloseItem()
	return SendListSelection(l, c, m.String(), r.UpgradeGloveSelection)
}

func (r JMFromThaStreetz) CreateClaw(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, what kind of claw would you like me to make?").NewLine().
		OpenItem(0).BlueText().AddText("Steel Titans").BlackText().AddText(" - Thief Lv. 15").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Bronze Igor").BlackText().AddText(" - Thief Lv. 20").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Meba").BlackText().AddText(" - Thief Lv. 25").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Steel Guards").BlackText().AddText(" - Thief Lv. 30").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Bronze Guardian").BlackText().AddText(" - Thief Lv. 35").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Steel Avarice").BlackText().AddText(" - Thief Lv. 40").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Steel Slain").BlackText().AddText(" - Thief Lv. 50").CloseItem()
	return SendListSelection(l, c, m.String(), r.CreateClawSelection)
}

func (r JMFromThaStreetz) UpgradeClaw(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("An upgraded claw? Sure thing, but note that upgrades won't carry over to the new item...").NewLine().
		OpenItem(0).BlueText().AddText("Mithril Titans").BlackText().AddText(" - Thief Lv. 15").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Gold Titans").BlackText().AddText(" - Thief Lv. 15").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Steel Igor").BlackText().AddText(" - Thief Lv. 20").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Adamantium Igor").BlackText().AddText(" - Thief Lv. 20").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Mithril Guards").BlackText().AddText(" - Thief Lv. 30").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Adamantium Guards").BlackText().AddText(" - Thief Lv. 3").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Silver Guardian").BlackText().AddText(" - Thief Lv. 35").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Dark Guardian").BlackText().AddText(" - Thief Lv. 35").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Blood Avarice").BlackText().AddText(" - Thief Lv. 40").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("Adamantium Avarice").BlackText().AddText(" - Thief Lv. 40").CloseItem().NewLine().
		OpenItem(10).BlueText().AddText("Dark Avarice").BlackText().AddText(" - Thief Lv. 40").CloseItem().NewLine().
		OpenItem(11).BlueText().AddText("Blood Slain").BlackText().AddText(" - Thief Lv. 50").CloseItem().NewLine().
		OpenItem(12).BlueText().AddText("Sapphire Slain").BlackText().AddText(" - Thief Lv. 5").CloseItem()
	return SendListSelection(l, c, m.String(), r.UpgradeClawSelection)
}

func (r JMFromThaStreetz) CreateMaterials(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Materials? I know of a few materials that I can make for you...").
		NewLine().
		BlueText().OpenItem(0).AddText("Make Processed Wood with Tree Branch").CloseItem().NewLine().
		BlueText().OpenItem(1).AddText("Make Processed Wood with Firewood").CloseItem().NewLine().
		BlueText().OpenItem(2).AddText("Make Screws (packs of 15)").CloseItem().NewLine()
	return SendListSelection(l, c, m.String(), r.MaterialRefineSelection)
}

func (r JMFromThaStreetz) CreateGloveSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.WorkGloves, r.WorkGlovesRequirements())
	case 1:
		return r.Confirm(item.BrownDuo, r.BrownDuoRequirements())
	case 2:
		return r.Confirm(item.BlueDuo, r.BlueDuoRequirements())
	case 3:
		return r.Confirm(item.BlackDuo, r.BlackDuoRequirements())
	case 4:
		return r.Confirm(item.BronzeMischief, r.BronzeMischiefRequirements())
	case 5:
		return r.Confirm(item.BronzeWolfskin, r.BronzeWolfskinRequirements())
	case 6:
		return r.Confirm(item.SteelSylvia, r.SteelSylviaRequirements())
	case 7:
		return r.Confirm(item.SteelArbion, r.SteelArbionRequirements())
	case 8:
		return r.Confirm(item.RedCleave, r.RedCleaveRequirements())
	case 9:
		return r.Confirm(item.BlueMoonGlove, r.BlueMoonGloveRequirements())
	case 10:
		return r.Confirm(item.BronzePow, r.BronzePowRequirements())
	}
	return nil
}

func (r JMFromThaStreetz) UpgradeGloveSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.MithrilMischief, r.MithrilMischiefRequirements())
	case 1:
		return r.Confirm(item.DarkMischief, r.DarkMischiefRequirements())
	case 2:
		return r.Confirm(item.MithrilWolfskin, r.MithrilWolfskinRequirements())
	case 3:
		return r.Confirm(item.DarkWolfskin, r.DarkWolfskinRequirements())
	case 4:
		return r.Confirm(item.SilverSylvia, r.SilverSylviaRequirements())
	case 5:
		return r.Confirm(item.GoldSylvia, r.GoldSylviaRequirements())
	case 6:
		return r.Confirm(item.OrihalconArbion, r.OrihalconArbionRequirements())
	case 7:
		return r.Confirm(item.GoldArbion, r.GoldArbionRequirements())
	case 8:
		return r.Confirm(item.GoldCleave, r.GoldCleaveRequirements())
	case 9:
		return r.Confirm(item.DarkCleave, r.DarkCleaveRequirements())
	case 10:
		return r.Confirm(item.RedMoonGlove, r.RedMoonGloveRequirements())
	case 11:
		return r.Confirm(item.BrownMoonGlove, r.BrownMoonGloveRequirements())
	case 12:
		return r.Confirm(item.StealPow, r.SilverPowRequirements())
	case 13:
		return r.Confirm(item.GoldPow, r.GoldPowRequirements())
	}
	return nil
}

func (r JMFromThaStreetz) CreateClawSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.SteelTitans, r.SteelTitansRequirements())
	case 1:
		return r.Confirm(item.BronzeIgor, r.BronzeIgorRequirements())
	case 2:
		return r.Confirm(item.Meba, r.MebaRequirements())
	case 3:
		return r.Confirm(item.SteelGuards, r.SteelGuardsRequirements())
	case 4:
		return r.Confirm(item.BronzeGuardian, r.BronzeGuardianRequirements())
	case 5:
		return r.Confirm(item.SteelAvarice, r.SteelAvariceRequirements())
	case 6:
		return r.Confirm(item.SteelSlain, r.SteelSlainRequirements())
	}
	return nil
}

func (r JMFromThaStreetz) UpgradeClawSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.MithrilTitans, r.MithrilTitansRequirements())
	case 1:
		return r.Confirm(item.GoldTitans, r.GoldTitansRequirements())
	case 2:
		return r.Confirm(item.SteelIgor, r.SteelIgorRequirements())
	case 3:
		return r.Confirm(item.AdamantiumIgor, r.AdamantiumIgorRequirements())
	case 4:
		return r.Confirm(item.MithrilGuards, r.MithrilGuardsRequirements())
	case 5:
		return r.Confirm(item.AdamantiumGuards, r.AdamantiumGuardsRequirements())
	case 6:
		return r.Confirm(item.SilverGuardian, r.SilverGuardianRequirements())
	case 7:
		return r.Confirm(item.DarkGuardian, r.DarkGuardianRequirements())
	case 8:
		return r.Confirm(item.BloodAvarice, r.BloodAvariceRequirements())
	case 9:
		return r.Confirm(item.AdamantiumAvarice, r.AdamantiumAvariceRequirements())
	case 10:
		return r.Confirm(item.DarkAvarice, r.DarkAvariceRequirements())
	case 11:
		return r.Confirm(item.BloodSlain, r.BloodSlainRequirements())
	case 12:
		return r.Confirm(item.SapphireSlain, r.SapphireSlainRequirements())
	}
	return nil
}

func (r JMFromThaStreetz) MaterialRefineSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.MaterialQuantityPrompt(item.ProcessedWood, r.ProcessedWoodFromTreeBranchRequirements())
	case 1:
		return r.MaterialQuantityPrompt(item.ProcessedWood, r.ProcessedWoodFromFirewoodRequirements())
	case 2:
		return r.MaterialQuantityPrompt(item.Screw, r.ScrewRequirements())
	}
	return nil
}

func (r JMFromThaStreetz) MaterialQuantityPrompt(itemId uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("So, you want me to make some ").
			ShowItemName1(itemId).
			AddText("s? In that case, how many do you want me to make?")
		return SendGetNumber(l, c, m.String(), r.ProcessMaterialQuantity(itemId, requirements), 1, 1, 100)
	}
}

func (r JMFromThaStreetz) ProcessedWoodFromTreeBranchRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.TreeBranch, amount: 10}}, cost: 0}
}

func (r JMFromThaStreetz) ProcessedWoodFromFirewoodRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Firewood, amount: 5}}, cost: 0}
}

func (r JMFromThaStreetz) ScrewRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 1}, {itemId: item.SteelPlate, amount: 1}}, cost: 0}
}

func (r JMFromThaStreetz) WorkGlovesRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 15}}, cost: 1000}
}

func (r JMFromThaStreetz) BrownDuoRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 30}, {itemId: item.Firewood, amount: 20}}, cost: 7000}
}

func (r JMFromThaStreetz) BlueDuoRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 30}, {itemId: item.HornyMushroomCap, amount: 20}}, cost: 7000}
}

func (r JMFromThaStreetz) BlackDuoRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 30}, {itemId: item.WildBoarTooth, amount: 20}}, cost: 7000}
}

func (r JMFromThaStreetz) BronzeMischiefRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 2}, {itemId: item.Leather, amount: 40}}, cost: 10000}
}

func (r JMFromThaStreetz) BronzeWolfskinRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 2}, {itemId: item.SteelPlate, amount: 1}, {itemId: item.Leather, amount: 10}}, cost: 15000}
}

func (r JMFromThaStreetz) SteelSylviaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelPlate, amount: 2}, {itemId: item.Leather, amount: 50}, {itemId: item.Screw, amount: 10}}, cost: 25000}
}

func (r JMFromThaStreetz) SteelArbionRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelPlate, amount: 3}, {itemId: item.BronzePlate, amount: 1}, {itemId: item.Leather, amount: 60}, {itemId: item.Screw, amount: 15}}, cost: 30000}
}

func (r JMFromThaStreetz) RedCleaveRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Garnet, amount: 3}, {itemId: item.DrakeSkull, amount: 200}, {itemId: item.Leather, amount: 80}, {itemId: item.Screw, amount: 30}}, cost: 40000}
}

func (r JMFromThaStreetz) BlueMoonGloveRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Sapphire, amount: 3}, {itemId: item.BlackCrystal, amount: 1}, {itemId: item.DragonSkin, amount: 40}, {itemId: item.Screw, amount: 30}}, cost: 50000}
}

func (r JMFromThaStreetz) BronzePowRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.MoonRock, amount: 1}, {itemId: item.BronzePlate, amount: 8}, {itemId: item.Diamond, amount: 1}, {itemId: item.DragonSkin, amount: 50}, {itemId: item.Screw, amount: 50}}, cost: 70000}
}

func (r JMFromThaStreetz) MithrilMischiefRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeMischief, amount: 1}, {itemId: item.MithrilPlate, amount: 1}}, cost: 5000}
}

func (r JMFromThaStreetz) DarkMischiefRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeMischief, amount: 1}, {itemId: item.Opal, amount: 1}}, cost: 7000}
}

func (r JMFromThaStreetz) MithrilWolfskinRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeWolfskin, amount: 1}, {itemId: item.MithrilPlate, amount: 2}}, cost: 10000}
}

func (r JMFromThaStreetz) DarkWolfskinRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeWolfskin, amount: 1}, {itemId: item.Opal, amount: 2}}, cost: 12000}
}

func (r JMFromThaStreetz) SilverSylviaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelSylvia, amount: 1}, {itemId: item.SilverPlate, amount: 2}}, cost: 15000}
}

func (r JMFromThaStreetz) GoldSylviaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelSylvia, amount: 1}, {itemId: item.GoldPlate, amount: 1}}, cost: 20000}
}

func (r JMFromThaStreetz) OrihalconArbionRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelArbion, amount: 1}, {itemId: item.OrihalconPlate, amount: 3}}, cost: 22000}
}

func (r JMFromThaStreetz) GoldArbionRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelArbion, amount: 1}, {itemId: item.GoldPlate, amount: 2}}, cost: 25000}
}

func (r JMFromThaStreetz) GoldCleaveRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedCleave, amount: 1}, {itemId: item.GoldPlate, amount: 4}}, cost: 40000}
}

func (r JMFromThaStreetz) DarkCleaveRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedCleave, amount: 1}, {itemId: item.BlackCrystal, amount: 2}}, cost: 50000}
}

func (r JMFromThaStreetz) RedMoonGloveRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BlueMoonGlove, amount: 1}, {itemId: item.Garnet, amount: 5}}, cost: 55000}
}

func (r JMFromThaStreetz) BrownMoonGloveRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BlueMoonGlove, amount: 1}, {itemId: item.GoldPlate, amount: 2}, {itemId: item.BlackCrystal, amount: 1}}, cost: 60000}
}

func (r JMFromThaStreetz) SilverPowRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.StealPow, amount: 1}, {itemId: item.SteelPlate, amount: 7}, {itemId: item.DrakeSkull, amount: 200}}, cost: 70000}
}

func (r JMFromThaStreetz) GoldPowRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.StealPow, amount: 1}, {itemId: item.GoldPlate, amount: 7}, {itemId: item.WildKargoEye, amount: 150}}, cost: 80000}
}

func (r JMFromThaStreetz) SteelTitansRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelPlate, amount: 1}, {itemId: item.Leather, amount: 20}, {itemId: item.Screw, amount: 5}}, cost: 2000}
}

func (r JMFromThaStreetz) BronzeIgorRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 2}, {itemId: item.SteelPlate, amount: 1}, {itemId: item.Leather, amount: 30}, {itemId: item.Screw, amount: 10}}, cost: 3000}
}

func (r JMFromThaStreetz) MebaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Garnier, amount: 1}, {itemId: item.SteelPlate, amount: 3}, {itemId: item.Leather, amount: 20}, {itemId: item.ProcessedWood, amount: 30}}, cost: 5000}
}

func (r JMFromThaStreetz) SteelGuardsRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 3}, {itemId: item.SteelPlate, amount: 2}, {itemId: item.Leather, amount: 50}, {itemId: item.Screw, amount: 20}}, cost: 15000}
}

func (r JMFromThaStreetz) BronzeGuardianRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 4}, {itemId: item.SteelPlate, amount: 2}, {itemId: item.Leather, amount: 80}, {itemId: item.Screw, amount: 25}}, cost: 30000}
}

func (r JMFromThaStreetz) SteelAvariceRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 3}, {itemId: item.SteelPlate, amount: 2}, {itemId: item.Leather, amount: 100}, {itemId: item.Screw, amount: 30}}, cost: 40000}
}

func (r JMFromThaStreetz) SteelSlainRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 4}, {itemId: item.SteelPlate, amount: 2}, {itemId: item.DragonSkin, amount: 40}, {itemId: item.Screw, amount: 35}}, cost: 50000}
}

func (r JMFromThaStreetz) MithrilTitansRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelTitans, amount: 1}, {itemId: item.MithrilPlate, amount: 1}}, cost: 1000}
}

func (r JMFromThaStreetz) GoldTitansRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelTitans, amount: 1}, {itemId: item.GoldPlate, amount: 1}}, cost: 2000}
}

func (r JMFromThaStreetz) SteelIgorRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeIgor, amount: 1}, {itemId: item.SteelPlate, amount: 2}}, cost: 3000}
}

func (r JMFromThaStreetz) AdamantiumIgorRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeIgor, amount: 1}, {itemId: item.AdamantiumPlate, amount: 2}}, cost: 5000}
}

func (r JMFromThaStreetz) MithrilGuardsRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelGuards, amount: 1}, {itemId: item.MithrilPlate, amount: 3}}, cost: 10000}
}

func (r JMFromThaStreetz) AdamantiumGuardsRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelGuards, amount: 1}, {itemId: item.AdamantiumPlate, amount: 3}}, cost: 15000}
}

func (r JMFromThaStreetz) SilverGuardianRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeGuardian, amount: 1}, {itemId: item.SilverPlate, amount: 4}}, cost: 20000}
}

func (r JMFromThaStreetz) DarkGuardianRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeGuardian, amount: 1}, {itemId: item.BlackCrystal, amount: 1}}, cost: 25000}
}

func (r JMFromThaStreetz) BloodAvariceRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelAvarice, amount: 1}, {itemId: item.Garnet, amount: 5}}, cost: 30000}
}

func (r JMFromThaStreetz) AdamantiumAvariceRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelAvarice, amount: 1}, {itemId: item.AdamantiumPlate, amount: 5}}, cost: 30000}
}

func (r JMFromThaStreetz) DarkAvariceRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelAvarice, amount: 1}, {itemId: item.BlackCrystal, amount: 2}}, cost: 35000}
}

func (r JMFromThaStreetz) BloodSlainRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelSlain, amount: 1}, {itemId: item.Garnet, amount: 6}}, cost: 40000}
}

func (r JMFromThaStreetz) SapphireSlainRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelSlain, amount: 1}, {itemId: item.Sapphire, amount: 6}}, cost: 40000}
}

func (r JMFromThaStreetz) ProcessMaterialQuantity(itemId uint32, requirements RefinementRequirements) ProcessNumber {
	return func(selection int32) StateProducer {
		return r.ConfirmQuantity(itemId, uint32(selection), requirements)
	}
}

func (r JMFromThaStreetz) Confirm(itemId uint32, requirements RefinementRequirements) StateProducer {
	return r.ConfirmQuantity(itemId, 1, requirements)
}

func (r JMFromThaStreetz) ConfirmQuantity(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
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

func (r JMFromThaStreetz) Validate(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
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

func (r JMFromThaStreetz) MakeRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Check your inventory for a free slot first.")
	return SendOk(l, c, m.String())
}

func (r JMFromThaStreetz) CannotAfford(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I'm afraid you cannot afford my services.")
	return SendOk(l, c, m.String())
}

func (r JMFromThaStreetz) MissingSomething(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("What are you trying to pull? I can't make anything unless you bring me what I ask for. Can you please bring more ").
			ShowItemName1(itemId).
			AddText("?")
		return SendOk(l, c, m.String())
	}
}

func (r JMFromThaStreetz) PerformRefine(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
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

func (r JMFromThaStreetz) Success(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("All done. If you need anything else... Well, I'm not going anywhere.")
	return SendOk(l, c, m.String())
}
