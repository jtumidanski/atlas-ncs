package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// MrSmith is located in Victoria Road - Perion (102000000)
type MrSmith struct {
}

func (r MrSmith) NPCId() uint32 {
	return npc.MrSmith
}

func (r MrSmith) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r MrSmith) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Um... Hi, I'm Mr. Thunder's apprentice. He's getting up there in age, so he handles most of the heavy-duty work while I handle some of the lighter jobs. What can I do for you?").NewLine().
		OpenItem(0).BlueText().AddText("Make a glove").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Upgrade a glove").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Create materials").CloseItem()
	return SendListSelection(l, c, m.String(), r.WhatToDo)
}

func (r MrSmith) WhatToDo(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.MakeGlove
	case 1:
		return r.UpgradeGlove
	case 2:
		return r.CreateMaterials
	}
	return nil
}

func (r MrSmith) MakeGlove(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Okay, so which glove do you want me to make?").NewLine().
		OpenItem(0).BlueText().AddText("Juno").BlackText().AddText(" - Warrior Lv. 10").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Steel Fingerless Gloves").BlackText().AddText(" - Warrior Lv. 15").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Venon").BlackText().AddText(" - Warrior Lv. 20").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("White Fingerless Gloves").BlackText().AddText(" - Warrior Lv. 25").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Bronze Missel").BlackText().AddText(" - Warrior Lv. 30").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Steel Briggon").BlackText().AddText(" - Warrior Lv. 35").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Iron Knuckle").BlackText().AddText(" - Warrior Lv. 40").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Steel Brist").BlackText().AddText(" - Warrior Lv. 50").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Bronze Clench").BlackText().AddText(" - Warrior Lv. 60").CloseItem()
	return SendListSelection(l, c, m.String(), r.GloveSelection)
}

func (r MrSmith) UpgradeGlove(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Upgrade a glove? That shouldn't be too difficult. Which did you have in mind?").NewLine().
		OpenItem(0).BlueText().AddText("Steel Missel").BlackText().AddText(" - Warrior Lv. 30").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Orihalcon Missel").BlackText().AddText(" - Warrior Lv. 30").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Yellow Briggon").BlackText().AddText(" - Warrior Lv. 35").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Dark Briggon").BlackText().AddText(" - Warrior Lv. 35").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Adamantium Knuckle").BlackText().AddText(" - Warrior Lv. 40").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Dark Knuckle").BlackText().AddText(" - Warrior Lv. 40").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("Mithril Brist").BlackText().AddText(" - Warrior Lv. 50").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Gold Brist").BlackText().AddText(" - Warrior Lv. 50").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Sapphire Clench").BlackText().AddText(" - Warrior Lv. 60").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("Dark Clench").BlackText().AddText(" - Warrior Lv. 60").CloseItem()
	return SendListSelection(l, c, m.String(), r.UpgradeSelection)
}

func (r MrSmith) CreateMaterials(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Materials? I know of a few materials that I can make for you...").
		NewLine().
		BlueText().OpenItem(0).AddText("Make Processed Wood with Tree Branch").CloseItem().NewLine().
		BlueText().OpenItem(1).AddText("Make Processed Wood with Firewood").CloseItem().NewLine().
		BlueText().OpenItem(2).AddText("Make Screws (packs of 15)").CloseItem().NewLine()
	return SendListSelection(l, c, m.String(), r.MaterialRefineSelection)
}

func (r MrSmith) GloveSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.Juno, r.JunoRequirements())
	case 1:
		return r.Confirm(item.SteelFingerlessGloves, r.SteelFingerlessGlovesRequirements())
	case 2:
		return r.Confirm(item.Venon, r.VenonRequirements())
	case 3:
		return r.Confirm(item.WhiteFingerlessGloves, r.WhiteFingerlessGlovesRequirements())
	case 4:
		return r.Confirm(item.BronzeMissel, r.BronzeMisselRequirements())
	case 5:
		return r.Confirm(item.SteelBriggon, r.SteelBriggonRequirements())
	case 6:
		return r.Confirm(item.IronKnuckle, r.IronKnuckleRequirements())
	case 7:
		return r.Confirm(item.SteelBrist, r.SteelBristRequirements())
	case 8:
		return r.Confirm(item.BronzeClench, r.BronzeClenchRequirements())
	}
	return nil
}

func (r MrSmith) UpgradeSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm(item.SteelMissel, r.SteelMisselRequirements())
	case 1:
		return r.Confirm(item.OrihalconMissel, r.OrihalconMisselRequirements())
	case 2:
		return r.Confirm(item.YellowBriggon, r.YellowBriggonRequirements())
	case 3:
		return r.Confirm(item.DarkBriggon, r.DarkBriggonRequirements())
	case 4:
		return r.Confirm(item.AdamantiumKnuckle, r.AdamantiumKnuckleRequirements())
	case 5:
		return r.Confirm(item.DarkKnuckle, r.DarkKnuckleRequirements())
	case 6:
		return r.Confirm(item.MithrilBrist, r.MithrilBristRequirements())
	case 7:
		return r.Confirm(item.GoldBrist, r.GoldBristRequirements())
	case 8:
		return r.Confirm(item.SapphireClench, r.SapphireClenchRequirements())
	case 9:
		return r.Confirm(item.DarkClench, r.DarkClenchRequirements())
	}
	return nil
}

func (r MrSmith) MaterialRefineSelection(selection int32) StateProducer {
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

func (r MrSmith) MaterialQuantityPrompt(itemId uint32, requirements RefinementRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("So, you want me to make some ").
			ShowItemName1(itemId).
			AddText("s? In that case, how many do you want me to make?")
		return SendGetNumber(l, c, m.String(), r.ProcessMaterialQuantity(itemId, requirements), 1, 1, 100)
	}
}

func (r MrSmith) ProcessedWoodFromTreeBranchRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.TreeBranch, amount: 10}}, cost: 0}
}

func (r MrSmith) ProcessedWoodFromFirewoodRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Firewood, amount: 5}}, cost: 0}
}

func (r MrSmith) ScrewRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 1}, {itemId: item.SteelPlate, amount: 1}}, cost: 0}
}

func (r MrSmith) JunoRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 15}, {itemId: item.SteelPlate, amount: 1}}, cost: 1000}
}

func (r MrSmith) SteelFingerlessGlovesRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelPlate, amount: 2}}, cost: 2000}
}

func (r MrSmith) VenonRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 40}, {itemId: item.BronzePlate, amount: 2}}, cost: 5000}
}

func (r MrSmith) WhiteFingerlessGlovesRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelPlate, amount: 2}}, cost: 10000}
}

func (r MrSmith) BronzeMisselRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzePlate, amount: 3}, {itemId: item.SteelPlate, amount: 2}, {itemId: item.Screw, amount: 15}}, cost: 20000}
}

func (r MrSmith) SteelBriggonRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 30}, {itemId: item.SteelPlate, amount: 4}, {itemId: item.Screw, amount: 15}}, cost: 30000}
}

func (r MrSmith) IronKnuckleRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.Leather, amount: 50}, {itemId: item.SteelPlate, amount: 5}, {itemId: item.Screw, amount: 40}}, cost: 40000}
}

func (r MrSmith) SteelBristRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelPlate, amount: 3}, {itemId: item.Diamond, amount: 2}, {itemId: item.DragonSkin, amount: 30}, {itemId: item.Screw, amount: 45}}, cost: 50000}
}

func (r MrSmith) BronzeClenchRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.MoonRock, amount: 1}, {itemId: item.BronzePlate, amount: 8}, {itemId: item.GoldPlate, amount: 2}, {itemId: item.DragonSkin, amount: 50}, {itemId: item.Screw, amount: 50}}, cost: 70000}
}

func (r MrSmith) SteelMisselRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeMissel, amount: 1}, {itemId: item.SteelPlate, amount: 1}}, cost: 20000}
}

func (r MrSmith) OrihalconMisselRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeMissel, amount: 1}, {itemId: item.OrihalconPlate, amount: 2}}, cost: 25000}
}

func (r MrSmith) YellowBriggonRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelBriggon, amount: 1}, {itemId: item.Topaz, amount: 3}}, cost: 30000}
}

func (r MrSmith) DarkBriggonRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelBriggon, amount: 1}, {itemId: item.BlackCrystal, amount: 1}}, cost: 40000}
}

func (r MrSmith) AdamantiumKnuckleRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.IronKnuckle, amount: 1}, {itemId: item.AdamantiumPlate, amount: 4}}, cost: 45000}
}

func (r MrSmith) DarkKnuckleRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.IronKnuckle, amount: 1}, {itemId: item.BlackCrystal, amount: 2}}, cost: 50000}
}

func (r MrSmith) MithrilBristRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelBrist, amount: 1}, {itemId: item.MithrilPlate, amount: 5}}, cost: 55000}
}

func (r MrSmith) GoldBristRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.SteelBrist, amount: 1}, {itemId: item.GoldPlate, amount: 4}}, cost: 60000}
}

func (r MrSmith) SapphireClenchRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeClench, amount: 1}, {itemId: item.MithrilPlate, amount: 3}, {itemId: item.Sapphire, amount: 5}}, cost: 70000}
}

func (r MrSmith) DarkClenchRequirements() RefinementRequirements {
	return RefinementRequirements{requirements: []Requirement{{itemId: item.BronzeClench, amount: 1}, {itemId: item.Diamond, amount: 2}, {itemId: item.BlackCrystal, amount: 2}}, cost: 80000}
}

func (r MrSmith) ProcessMaterialQuantity(itemId uint32, requirements RefinementRequirements) ProcessNumber {
	return func(selection int32) StateProducer {
		return r.ConfirmQuantity(itemId, uint32(selection), requirements)
	}
}

func (r MrSmith) Confirm(itemId uint32, requirements RefinementRequirements) StateProducer {
	return r.ConfirmQuantity(itemId, 1, requirements)
}

func (r MrSmith) ConfirmQuantity(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
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

func (r MrSmith) Validate(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
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

func (r MrSmith) MakeRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Check your inventory for a free slot first.")
	return SendOk(l, c, m.String())
}

func (r MrSmith) CannotAfford(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I may still be an apprentice, but I do need to earn a living.")
	return SendOk(l, c, m.String())
}

func (r MrSmith) MissingSomething(itemId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("I'm still an apprentice, I don't know if I can substitute other items in yet... Can you please bring ").
			ShowItemName1(itemId).
			AddText("?")
		return SendOk(l, c, m.String())
	}
}

func (r MrSmith) PerformRefine(itemId uint32, amount uint32, requirements RefinementRequirements) StateProducer {
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

func (r MrSmith) Success(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Did that come out right? Come by me again if you have anything for me to practice on.")
	return SendOk(l, c, m.String())
}
