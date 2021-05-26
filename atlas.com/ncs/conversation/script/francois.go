package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Francois is located in Victoria Road - Ellinia (101000000)
type Francois struct {
}

func (r Francois) NPCId() uint32 {
	return npc.Francois
}

func (r Francois) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Francois) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Welcome to my eco-safe refining operation! What would you like today?").NewLine().
		OpenItem(0).BlueText().AddText("Make a glove").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Upgrade a glove").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Upgrade a hat").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Make a wand").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Make a staff").CloseItem()
	return SendListSelection(l, c, m.String(), r.WhatToDo)
}

func (r Francois) WhatToDo(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.MakeGlove
	case 1:
		return r.UpgradeGlove
	case 2:
		return r.UpgradeHat
	case 3:
		return r.MakeWand
	case 4:
		return r.MakeStaff
	}
	return nil
}

func (r Francois) MakeGlove(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, what kind of glove would you like me to make?").NewLine().
		OpenItem(0).BlueText().AddText("Lemona").BlackText().AddText(" - Magician Lv. 15").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Blue Morrican").BlackText().AddText(" - Magician Lv. 20").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Ocean Mesana").BlackText().AddText(" - Magician Lv. 25").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Red Lutia").BlackText().AddText(" - Magician Lv. 30").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Red Noel").BlackText().AddText(" - Magician Lv. 35").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Red Arten").BlackText().AddText(" - Magician Lv. 40").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Red Pennance").BlackText().AddText(" - Magician Lv. 50").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Steel Manute").BlackText().AddText(" - Magician Lv. 60").CloseItem()
	return SendListSelection(l, c, m.String(), r.GloveSelection)
}

func (r Francois) UpgradeGlove(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, what kind of glove are you looking to upgrade to?").NewLine().
		OpenItem(0).BlueText().AddText("Green Morrican").BlackText().AddText(" - Magician Lv. 20").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Purple Morrican").BlackText().AddText(" - Magician Lv. 20").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Blood Mesana").BlackText().AddText(" - Magician Lv. 25").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Dark Mesana").BlackText().AddText(" - Magician Lv. 25").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Blue Lutia").BlackText().AddText(" - Magician Lv. 30").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Black Lutia").BlackText().AddText(" - Magician Lv. 30").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Blue Noel").BlackText().AddText(" - Magician Lv. 35").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Dark Noel").BlackText().AddText(" - Magician Lv. 35").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Blue Arten").BlackText().AddText(" - Magician Lv. 40").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("Dark Arten").BlackText().AddText(" - Magician Lv. 40").CloseItem().NewLine().
		OpenItem(10).BlueText().AddText("Blue Pennance").BlackText().AddText(" - Magician Lv. 50").CloseItem().NewLine().
		OpenItem(11).BlueText().AddText("Dark Pennance").BlackText().AddText(" - Magician Lv. 50").CloseItem().NewLine().
		OpenItem(12).BlueText().AddText("Gold Manute").BlackText().AddText(" - Magician Lv. 60").CloseItem().NewLine().
		OpenItem(13).BlueText().AddText("Dark Manute").BlackText().AddText(" - Magician Lv. 60").CloseItem()
	return SendListSelection(l, c, m.String(), r.UpgradeGloveSelection)
}

func (r Francois) UpgradeHat(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("A hat? Which one were you thinking of?").NewLine().
		OpenItem(0).BlueText().AddText("Steel Pride").BlackText().AddText(" - Magician Lv. 30").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Golden Pride").BlackText().AddText(" - Magician Lv. 30").CloseItem()
	return SendListSelection(l, c, m.String(), r.UpgradeHatSelection)
}

func (r Francois) MakeWand(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("A wand, huh? Prefer the smaller weapon that fits in your pocket? Which type are you seeking?").NewLine().
		OpenItem(0).BlueText().AddText("Wooden Wand").BlackText().AddText(" - Common Lv. 8").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Hardwood Wand").BlackText().AddText(" - Common Lv. 13").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Metal Wand").BlackText().AddText(" - Common Lv. 18").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Ice Wand").BlackText().AddText(" - Magician Lv. 23").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Mithril Wand").BlackText().AddText(" - Magician Lv. 28").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Wizard Wand").BlackText().AddText(" - Magician Lv. 33").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Fairy Wand").BlackText().AddText(" - Magician Lv. 38").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Cromi").BlackText().AddText(" - Magician Lv. 48").CloseItem()
	return SendListSelection(l, c, m.String(), r.MakeWandSelection)
}

func (r Francois) MakeStaff(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Ah, a staff, a great symbol of one's power! Which are you looking to make?").NewLine().
		OpenItem(0).BlueText().AddText("Wooden Staff").BlackText().AddText(" - Magician Lv. 10").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Sapphire Staff").BlackText().AddText(" - Magician Lv. 15").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Emerald Staff").BlackText().AddText(" - Magician Lv. 15").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Old Wooden Staff").BlackText().AddText(" - Magician Lv. 20").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Wizard Staff").BlackText().AddText(" - Magician Lv. 25").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Arc Staff").BlackText().AddText(" - Magician Lv. 45").CloseItem()
	return SendListSelection(l, c, m.String(), r.MakeStaffSelection)
}

func (r Francois) GloveSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.Lemona, r.LemonaRequirements())
	case 1:
		return r.Confirm(item.BlueMorrican, r.BlueMorricanRequirements())
	case 2:
		return r.Confirm(item.OceanMesana, r.OceanMesanaRequirements())
	case 3:
		return r.Confirm(item.RedLutia, r.RedLutiaRequirements())
	case 4:
		return r.Confirm(item.RedNoel, r.RedNoelRequirements())
	case 5:
		return r.Confirm(item.RedArten, r.RedArtenRequirements())
	case 6:
		return r.Confirm(item.RedPennance, r.RedPennanceRequirements())
	case 7:
		return r.Confirm(item.SteelManute, r.SteelManuteRequirements())
	}
	return nil
}

func (r Francois) UpgradeGloveSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.GreenMorrican, r.GreenMorricanRequirements())
	case 1:
		return r.Confirm(item.PurpleMorrican, r.PurpleMorricanRequirements())
	case 2:
		return r.Confirm(item.BloodMesana, r.BloodMesanaRequirements())
	case 3:
		return r.Confirm(item.DarkMesana, r.DarkMesanaRequirements())
	case 4:
		return r.Confirm(item.BlueLutia, r.BlueLutiaRequirements())
	case 5:
		return r.Confirm(item.BlackLutia, r.BlackLutiaRequirements())
	case 6:
		return r.Confirm(item.BlueNoel, r.BlueNoelRequirements())
	case 7:
		return r.Confirm(item.DarkNoel, r.DarkNoelRequirements())
	case 8:
		return r.Confirm(item.BlueArten, r.BlueArtenRequirements())
	case 9:
		return r.Confirm(item.DarkArten, r.DarkArtenRequirements())
	case 10:
		return r.Confirm(item.BluePennance, r.BluePennanceRequirements())
	case 11:
		return r.Confirm(item.DarkPennance, r.DarkPennanceRequirements())
	case 12:
		return r.Confirm(item.GoldManute, r.GoldManuteRequirements())
	case 13:
		return r.Confirm(item.DarkManute, r.DarkManuteRequirements())
	}
	return nil
}

func (r Francois) UpgradeHatSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.SteelPride, r.SteelPrideRequirements())
	case 1:
		return r.Confirm(item.GoldenPride, r.GoldenPrideRequirements())
	}
	return nil
}

func (r Francois) MakeWandSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.WoodenWand, r.WoodenWandRequirements())
	case 1:
		return r.Confirm(item.HardwoodWand, r.HardwoodWandRequirements())
	case 2:
		return r.Confirm(item.MetalWand, r.MetalWandRequirements())
	case 3:
		return r.Confirm(item.IceWand, r.IceWandRequirements())
	case 4:
		return r.Confirm(item.MithrilWand, r.MithrilWandRequirements())
	case 5:
		return r.Confirm(item.WizardWand, r.WizardWandRequirements())
	case 6:
		return r.Confirm(item.FairyWand, r.FairyWandRequirements())
	case 7:
		return r.Confirm(item.Cromi, r.CromiRequirements())
	}
	return nil
}

func (r Francois) MakeStaffSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.WoodenStaff, r.WoodenStaffRequirements())
	case 1:
		return r.Confirm(item.SapphireStaff, r.SapphireStaffRequirements())
	case 2:
		return r.Confirm(item.EmeraldStaff, r.EmeraldStaffRequirements())
	case 3:
		return r.Confirm(item.OldWoodenStaff, r.OldWoodenStaffRequirements())
	case 4:
		return r.Confirm(item.WizardStaff, r.WizardStaffRequirements())
	case 5:
		return r.Confirm(item.ArcStaff, r.ArcStaffRequirements())
	}
	return nil
}

func (r Francois) Confirm(itemId uint32, requirements RefinementRequirements) StateProducer {
	return r.ConfirmQuantity(itemId, 1, requirements)
}

func (r Francois) ConfirmQuantity(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
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

func (r Francois) LemonaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 15}}, cost: 7000}
}

func (r Francois) BlueMorricanRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 30}, {itemId: item.SteelPlate, amount: 1}}, cost: 15000}
}

func (r Francois) OceanMesanaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 50}, {itemId: item.GoldPlate, amount: 2}}, cost: 20000}
}

func (r Francois) RedLutiaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 60}, {itemId: item.Topaz, amount: 1}, {itemId: item.Garnet, amount: 2}}, cost: 25000}
}

func (r Francois) RedNoelRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 70}, {itemId: item.GoldPlate, amount: 1}, {itemId: item.SteelPlate, amount: 3}, {itemId: item.Garnet, amount: 2}}, cost: 30000}
}

func (r Francois) RedArtenRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 80}, {itemId: item.Garnet, amount: 3}, {itemId: item.Topaz, amount: 3}, {itemId: item.Screw, amount: 30}}, cost: 40000}
}

func (r Francois) RedPennanceRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Garnet, amount: 3}, {itemId: item.GoldPlate, amount: 2}, {itemId: item.DragonSkin, amount: 35}, {itemId: item.Screw, amount: 40}}, cost: 50000}
}

func (r Francois) SteelManuteRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.MoonRock, amount: 1}, {itemId: item.SteelPlate, amount: 8}, {itemId: item.Diamond, amount: 1}, {itemId: item.DragonSkin, amount: 50}, {itemId: item.Screw, amount: 50}}, cost: 70000}
}

func (r Francois) GreenMorricanRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BlueMorrican, amount: 1}, {itemId: item.SteelPlate, amount: 1}}, cost: 20000}
}

func (r Francois) PurpleMorricanRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BlueMorrican, amount: 1}, {itemId: item.Amethyst, amount: 2}}, cost: 25000}
}

func (r Francois) BloodMesanaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.OceanMesana, amount: 1}, {itemId: item.Garnet, amount: 3}}, cost: 30000}
}

func (r Francois) DarkMesanaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.OceanMesana, amount: 1}, {itemId: item.BlackCrystal, amount: 1}}, cost: 40000}
}

func (r Francois) BlueLutiaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedLutia, amount: 1}, {itemId: item.Sapphire, amount: 3}}, cost: 35000}
}

func (r Francois) BlackLutiaRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedLutia, amount: 1}, {itemId: item.BlackCrystal, amount: 1}}, cost: 40000}
}

func (r Francois) BlueNoelRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedNoel, amount: 1}, {itemId: item.Sapphire, amount: 3}}, cost: 40000}
}

func (r Francois) DarkNoelRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedNoel, amount: 1}, {itemId: item.BlackCrystal, amount: 1}}, cost: 45000}
}

func (r Francois) BlueArtenRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedArten, amount: 1}, {itemId: item.AquaMarine, amount: 4}}, cost: 45000}
}

func (r Francois) DarkArtenRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedArten, amount: 1}, {itemId: item.BlackCrystal, amount: 2}}, cost: 50000}
}

func (r Francois) BluePennanceRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedPennance, amount: 1}, {itemId: item.AquaMarine, amount: 5}}, cost: 55000}
}

func (r Francois) DarkPennanceRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.RedPennance, amount: 1}, {itemId: item.BlackCrystal, amount: 3}}, cost: 60000}
}

func (r Francois) GoldManuteRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelManute, amount: 1}, {itemId: item.SilverPlate, amount: 3}, {itemId: item.GoldPlate, amount: 5}}, cost: 70000}
}

func (r Francois) DarkManuteRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelManute, amount: 1}, {itemId: item.BlackCrystal, amount: 2}, {itemId: item.GoldPlate, amount: 3}}, cost: 80000}
}

func (r Francois) SteelPrideRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePride, amount: 1}, {itemId: item.SteelPlate, amount: 3}}, cost: 40000}
}

func (r Francois) GoldenPrideRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePride, amount: 1}, {itemId: item.GoldPlate, amount: 3}}, cost: 50000}
}

func (r Francois) WoodenWandRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.ProcessedWood, amount: 5}}, cost: 1000}
}

func (r Francois) HardwoodWandRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.ProcessedWood, amount: 10}, {itemId: item.OrangeMushroomCap, amount: 50}}, cost: 3000}
}

func (r Francois) MetalWandRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelPlate, amount: 1}, {itemId: item.BlueMushroomCap, amount: 30}, {itemId: item.Screw, amount: 5}}, cost: 5000}
}

func (r Francois) IceWandRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.MithrilPlate, amount: 2}, {itemId: item.PieceOfIce, amount: 1}, {itemId: item.Screw, amount: 10}}, cost: 12000}
}

func (r Francois) MithrilWandRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.MithrilPlate, amount: 3}, {itemId: item.AquaMarine, amount: 1}, {itemId: item.Screw, amount: 10}}, cost: 30000}
}

func (r Francois) WizardWandRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Topaz, amount: 5}, {itemId: item.MithrilPlate, amount: 3}, {itemId: item.SteelPlate, amount: 1}, {itemId: item.Screw, amount: 15}}, cost: 60000}
}

func (r Francois) FairyWandRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Topaz, amount: 5}, {itemId: item.Sapphire, amount: 5}, {itemId: item.Diamond, amount: 1}, {itemId: item.FairyWing, amount: 1}, {itemId: item.Screw, amount: 20}}, cost: 120000}
}

func (r Francois) CromiRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.GoldPlate, amount: 4}, {itemId: item.Emerald, amount: 3}, {itemId: item.Diamond, amount: 2}, {itemId: item.AquaMarine, amount: 1}, {itemId: item.PieceOfIce, amount: 1}, {itemId: item.Screw, amount: 30}}, cost: 200000}
}

func (r Francois) WoodenStaffRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.ProcessedWood, amount: 5}}, cost: 2000}
}

func (r Francois) SapphireStaffRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Sapphire, amount: 1}, {itemId: item.SteelPlate, amount: 1}, {itemId: item.Screw, amount: 5}}, cost: 2000}
}

func (r Francois) EmeraldStaffRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Emerald, amount: 1}, {itemId: item.SteelPlate, amount: 1}, {itemId: item.Screw, amount: 5}}, cost: 2000}
}

func (r Francois) OldWoodenStaffRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.ProcessedWood, amount: 50}, {itemId: item.SteelPlate, amount: 1}, {itemId: item.Screw, amount: 10}}, cost: 5000}
}

func (r Francois) WizardStaffRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Topaz, amount: 2}, {itemId: item.Amethyst, amount: 1}, {itemId: item.SteelPlate, amount: 1}, {itemId: item.Screw, amount: 15}}, cost: 12000}
}

func (r Francois) ArcStaffRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelPlate, amount: 8}, {itemId: item.Topaz, amount: 5}, {itemId: item.Amethyst, amount: 5}, {itemId: item.Sapphire, amount: 5}, {itemId: item.Screw, amount: 30}, {itemId: item.SlimeBubble, amount: 50}, {itemId: item.FairyWing, amount: 1}}, cost: 180000}
}

func (r Francois) Validate(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
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

func (r Francois) MakeRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Check your inventory for a free slot first.")
	return SendOk(l, c, m.String())
}

func (r Francois) CannotAfford(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Sorry, but all of us need money to live. Come back when you can pay my fees, yes?")
	return SendOk(l, c, m.String())
}

func (r Francois) MissingSomething(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("Uhm... I don't keep extra material on me. Sorry. Can you please bring ").
			ShowItemName1(itemId).
			AddText("?")
		return SendOk(l, c, m.String())
	}
}

func (r Francois) PerformRefine(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
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

func (r Francois) Success(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It's a success! Oh, I've never felt so alive! Please come back again!")
	return SendOk(l, c, m.String())
}
