package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

type RefinementRequirements struct {
	requirements []Requirement
	cost         uint32
}

type Requirement struct {
	itemId uint32
	amount uint8
}

// Vicious is located in Victoria Road - Henesys Market (100000100)
type Vicious struct {
}

func (r Vicious) NPCId() uint32 {
	return npc.Vicious
}

func (r Vicious) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.Hello(l, c)
}

func (r Vicious) Hello(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello. I am Vicious, retired Sniper. However, I used to be the top student of Athena Pierce. Though I no longer hunt, I can make some archer items that will be useful for you...").
		BlueText().NewLine().
		OpenItem(0).AddText("Create a bow").CloseItem().NewLine().
		OpenItem(1).AddText("Create a crossbow").CloseItem().NewLine().
		OpenItem(2).AddText("Make a glove").CloseItem().NewLine().
		OpenItem(3).AddText("Upgrade a glove").CloseItem().NewLine().
		OpenItem(4).AddText("Create materials").CloseItem().NewLine().
		OpenItem(5).AddText("Create arrows").CloseItem().NewLine()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Vicious) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.BowRefine
	case 1:
		return r.CrossbowRefine
	case 2:
		return r.GloveRefine
	case 3:
		return r.GloveUpgrade
	case 4:
		return r.MaterialRefine
	case 5:
		return r.ArrowRefine
	}
	return nil
}

func (r Vicious) SingleItemConfirm(itemId uint32, requirements RefinementRequirements) script.StateProducer {
	quantityPrompt := func() string {
		return "You want me to make a "
	}
	return r.RefineItemConfirm(itemId, 1, quantityPrompt, requirements)
}

func (r Vicious) MultipleItemConfirm(itemId uint32, amount int32, requirements RefinementRequirements) script.StateProducer {
	quantityPrompt := func() string {
		return fmt.Sprintf("You want me to make %d ", amount)
	}
	return r.RefineItemConfirm(itemId, amount, quantityPrompt, requirements)
}

func (r Vicious) RefineItemConfirm(itemId uint32, amount int32, quantityPrompt func() string, requirements RefinementRequirements) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText(quantityPrompt()).ShowItemName1(itemId).
			AddText("? In that case, I'm going to need specific items from you in order to make it. Make sure you have room in your inventory, though!").NewLine().
			BlueText()
		for _, req := range requirements.requirements {
			m = m.ShowItemImage2(req.itemId).AddText(fmt.Sprintf(" %d ", uint32(req.amount)*uint32(amount))).ShowItemName1(req.itemId).NewLine()
		}
		if requirements.cost > 0 {
			m = m.ShowItemImage2(item.MoneySack).AddText(fmt.Sprintf(" %d meso", requirements.cost*uint32(amount))).NewLine()
		}
		return script.SendYesNo(l, c, m.String(), r.ProcessPurchase(itemId, 1, requirements), script.Exit())
	}
}

func (r Vicious) ProcessPurchase(itemId uint32, amount int8, requirements RefinementRequirements) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		totalCost := uint32(amount) * requirements.cost
		if !character.HasMeso(l)(c.CharacterId, totalCost) {
			return r.Sorry(l, c)
		}

		for _, req := range requirements.requirements {
			if !character.HasItems(l)(c.CharacterId, req.itemId, uint32(req.amount)*uint32(amount)) {
				return r.NeedMoreItems(l, c)
			}
		}

		var toReceive uint32
		switch itemId {
		case item.ArrowForBow:
			toReceive = 1000
		case item.BronzeArrowForBow:
			toReceive = 900
		case item.SteelArrowForBow:
			toReceive = 800
		case item.ArrowForCrossbow:
			toReceive = 1000
		case item.BronzeArrowForCrossbow:
			toReceive = 900
		case item.SteelArrowForCrossbow:
			toReceive = 800
		case item.Screw:
			toReceive = uint32(amount) * 15
		default:
			toReceive = 1
		}

		if !character.CanHoldAll(l)(c.CharacterId, itemId, toReceive) {
			return r.NotEnoughInventorySpace(l, c)
		}

		for _, req := range requirements.requirements {
			character.GainItem(l)(c.CharacterId, req.itemId, -int32(req.amount)*int32(amount))
		}

		err := character.GainMeso(l)(c.CharacterId, -int32(totalCost))
		if err != nil {
			l.WithError(err).Errorf("Unable to receive payment from character %d for %d.", c.CharacterId, itemId)
		}
		character.GainItem(l)(c.CharacterId, itemId, int32(toReceive))

		m := message.NewBuilder().
			AddText("A perfect item, as usual. Come and see me if you need anything else.")
		return script.SendOk(l, c, m.String())
	}
}

func (r Vicious) BowRefine(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I may have been a Sniper, but bows and crossbows aren't too much different. Anyway, which would you like to make?").
		NewLine().
		BlueText().OpenItem(0).ShowItemName1(item.WarBow).BlackText().AddText(" - Bowman Lv. 10").CloseItem().NewLine().
		BlueText().OpenItem(1).ShowItemName1(item.CompositeBow).BlackText().AddText(" - Bowman Lv. 15").CloseItem().NewLine().
		BlueText().OpenItem(2).ShowItemName1(item.HuntersBow).BlackText().AddText(" - Bowman Lv. 20").CloseItem().NewLine().
		BlueText().OpenItem(3).ShowItemName1(item.BattleBow).BlackText().AddText(" - Bowman Lv. 25").CloseItem().NewLine().
		BlueText().OpenItem(4).ShowItemName1(item.Ryden).BlackText().AddText(" - Bowman Lv. 30").CloseItem().NewLine().
		BlueText().OpenItem(5).ShowItemName1(item.RedViper).BlackText().AddText(" - Bowman Lv. 35").CloseItem().NewLine().
		BlueText().OpenItem(6).ShowItemName1(item.Vaulter2000).BlackText().AddText(" - Bowman Lv. 40").CloseItem().NewLine()
	return script.SendListSelection(l, c, m.String(), r.BowRefineSelection)
}

func (r Vicious) BowRefineSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.SingleItemConfirm(item.WarBow, r.WarBowRequirements())
	case 1:
		return r.SingleItemConfirm(item.CompositeBow, r.CompositeBowRequirements())
	case 2:
		return r.SingleItemConfirm(item.HuntersBow, r.HuntersBowRequirements())
	case 3:
		return r.SingleItemConfirm(item.BattleBow, r.BattleBowRequirements())
	case 4:
		return r.SingleItemConfirm(item.Ryden, r.RydenRequirements())
	case 5:
		return r.SingleItemConfirm(item.RedViper, r.RedViperRequirements())
	case 6:
		return r.SingleItemConfirm(item.Vaulter2000, r.Vaulter2000Requirements())
	}
	return nil
}

func (r Vicious) WarBowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.ProcessedWood, amount: 5}, {itemId: item.BlueSnailShell, amount: 30}},
		cost:         800,
	}
}

func (r Vicious) CompositeBowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 1}, {itemId: item.Screw, amount: 3}},
		cost:         2000,
	}
}

func (r Vicious) HuntersBowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.ProcessedWood, amount: 30}, {itemId: item.RedSnailShell, amount: 50}},
		cost:         3000,
	}
}

func (r Vicious) BattleBowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 2}, {itemId: item.Topaz, amount: 2}, {itemId: item.Screw, amount: 8}},
		cost:         5000,
	}
}

func (r Vicious) RydenRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 5}, {itemId: item.GoldPlate, amount: 5}, {itemId: item.Emerald, amount: 3}, {itemId: item.Topaz, amount: 3}, {itemId: item.Screw, amount: 30}},
		cost:         30000,
	}
}

func (r Vicious) RedViperRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SilverPlate, amount: 7}, {itemId: item.Garnet, amount: 6}, {itemId: item.Opal, amount: 3}, {itemId: item.Screw, amount: 35}},
		cost:         40000,
	}
}

func (r Vicious) Vaulter2000Requirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlackCrystal, amount: 1}, {itemId: item.SteelPlate, amount: 10}, {itemId: item.GoldPlate, amount: 3}, {itemId: item.Screw, amount: 40}, {itemId: item.DrakeSkull, amount: 50}},
		cost:         80000,
	}
}

func (r Vicious) CrossbowRefine(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I was a Sniper. Crossbows are my specialty. Which would you like me to make for you?").
		NewLine().
		BlueText().OpenItem(0).ShowItemName1(item.Crossbow).BlackText().AddText(" - Bowman Lv. 10").CloseItem().NewLine().
		BlueText().OpenItem(1).ShowItemName1(item.BattleCrossbow).BlackText().AddText(" - Bowman Lv. 15").CloseItem().NewLine().
		BlueText().OpenItem(2).ShowItemName1(item.Balanche).BlackText().AddText(" - Bowman Lv. 20").CloseItem().NewLine().
		BlueText().OpenItem(3).ShowItemName1(item.MountainCrossbow).BlackText().AddText(" - Bowman Lv. 25").CloseItem().NewLine().
		BlueText().OpenItem(4).ShowItemName1(item.EagleCrow).BlackText().AddText(" - Bowman Lv. 30").CloseItem().NewLine().
		BlueText().OpenItem(5).ShowItemName1(item.Heckler).BlackText().AddText(" - Bowman Lv. 35").CloseItem().NewLine().
		BlueText().OpenItem(6).ShowItemName1(item.SilverCrow).BlackText().AddText(" - Bowman Lv. 40").CloseItem().NewLine().
		BlueText().OpenItem(7).ShowItemName1(item.Rower).BlackText().AddText(" - Bowman Lv. 45").CloseItem().NewLine()
	return script.SendListSelection(l, c, m.String(), r.CrossbowRefineSelection)
}

func (r Vicious) CrossbowRefineSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.SingleItemConfirm(item.Crossbow, r.CrossbowRequirements())
	case 1:
		return r.SingleItemConfirm(item.BattleCrossbow, r.BattleCrossbowRequirements())
	case 2:
		return r.SingleItemConfirm(item.Balanche, r.BalancheRequirements())
	case 3:
		return r.SingleItemConfirm(item.MountainCrossbow, r.MountainCrossbowRequirements())
	case 4:
		return r.SingleItemConfirm(item.EagleCrow, r.EagleCrowRequirements())
	case 5:
		return r.SingleItemConfirm(item.Heckler, r.HecklerRequirements())
	case 6:
		return r.SingleItemConfirm(item.SilverCrow, r.SilverCrowRequirements())
	case 7:
		return r.SingleItemConfirm(item.Rower, r.RowerRequirements())
	}
	return nil
}

func (r Vicious) CrossbowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.ProcessedWood, amount: 7}, {itemId: item.Screw, amount: 2}},
		cost:         1000,
	}
}

func (r Vicious) BattleCrossbowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 1}, {itemId: item.ProcessedWood, amount: 20}, {itemId: item.Screw, amount: 5}},
		cost:         2000,
	}
}

func (r Vicious) BalancheRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 1}, {itemId: item.ProcessedWood, amount: 50}, {itemId: item.Screw, amount: 8}},
		cost:         3000,
	}
}

func (r Vicious) MountainCrossbowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 2}, {itemId: item.Topaz, amount: 1}, {itemId: item.AquaMarine, amount: 1}, {itemId: item.Screw, amount: 10}},
		cost:         10000,
	}
}

func (r Vicious) EagleCrowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 5}, {itemId: item.OrihalconPlate, amount: 5}, {itemId: item.Topaz, amount: 3}, {itemId: item.ProcessedWood, amount: 50}, {itemId: item.Screw, amount: 15}},
		cost:         30000,
	}
}

func (r Vicious) HecklerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlackCrystal, amount: 1}, {itemId: item.SteelPlate, amount: 8}, {itemId: item.GoldPlate, amount: 4}, {itemId: item.Topaz, amount: 2}, {itemId: item.Screw, amount: 30}},
		cost:         50000,
	}
}

func (r Vicious) SilverCrowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlackCrystal, amount: 2}, {itemId: item.SilverPlate, amount: 6}, {itemId: item.ProcessedWood, amount: 30}, {itemId: item.Screw, amount: 30}},
		cost:         80000,
	}
}

func (r Vicious) RowerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlackCrystal, amount: 2}, {itemId: item.GoldPlate, amount: 5}, {itemId: item.Topaz, amount: 3}, {itemId: item.ProcessedWood, amount: 40}, {itemId: item.Screw, amount: 40}},
		cost:         200000,
	}
}

func (r Vicious) GloveRefine(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Okay, so which glove do you want me to make?").
		NewLine().
		BlueText().OpenItem(0).ShowItemName1(item.BasicArcherGloves).BlackText().AddText(" - Bowman Lv. 10").CloseItem().NewLine().
		BlueText().OpenItem(1).ShowItemName1(item.BlueDiros).BlackText().AddText(" - Bowman Lv. 15").CloseItem().NewLine().
		BlueText().OpenItem(2).ShowItemName1(item.BlueSavata).BlackText().AddText(" - Bowman Lv. 25").CloseItem().NewLine().
		BlueText().OpenItem(3).ShowItemName1(item.BrownMarker).BlackText().AddText(" - Bowman Lv. 30").CloseItem().NewLine().
		BlueText().OpenItem(4).ShowItemName1(item.BronzeScaler).BlackText().AddText(" - Bowman Lv. 35").CloseItem().NewLine().
		BlueText().OpenItem(5).ShowItemName1(item.AquaBrace).BlackText().AddText(" - Bowman Lv. 40").CloseItem().NewLine().
		BlueText().OpenItem(6).ShowItemName1(item.BlueWillow).BlackText().AddText(" - Bowman Lv. 50").CloseItem().NewLine().
		BlueText().OpenItem(7).ShowItemName1(item.OakerGarner).BlackText().AddText(" - Bowman Lv. 60").CloseItem().NewLine()
	return script.SendListSelection(l, c, m.String(), r.GloveRefineSelection)
}

func (r Vicious) GloveRefineSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.SingleItemConfirm(item.BasicArcherGloves, r.BasicArcherGlovesRequirements())
	case 1:
		return r.SingleItemConfirm(item.BlueDiros, r.BlueDirosRequirements())
	case 2:
		return r.SingleItemConfirm(item.BlueSavata, r.BlueSavataRequirements())
	case 3:
		return r.SingleItemConfirm(item.BrownMarker, r.BrownMarkerRequirements())
	case 4:
		return r.SingleItemConfirm(item.BronzeScaler, r.BronzeScalerRequirements())
	case 5:
		return r.SingleItemConfirm(item.AquaBrace, r.AquaBraceRequirements())
	case 6:
		return r.SingleItemConfirm(item.BlueWillow, r.BlueWillowRequirements())
	case 7:
		return r.SingleItemConfirm(item.OakerGarner, r.OakerGarnerRequirements())
	}
	return nil
}

func (r Vicious) BasicArcherGlovesRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.Leather, amount: 15}, {itemId: item.BlueMushroomCap, amount: 20}},
		cost:         5000,
	}
}

func (r Vicious) BlueDirosRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.Leather, amount: 20}, {itemId: item.BlueMushroomCap, amount: 20}, {itemId: item.SteelPlate, amount: 2}},
		cost:         10000,
	}
}

func (r Vicious) BlueSavataRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.Leather, amount: 40}, {itemId: item.BlueMushroomCap, amount: 50}, {itemId: item.GoldPlate, amount: 2}},
		cost:         15000,
	}
}

func (r Vicious) BrownMarkerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.Leather, amount: 50}, {itemId: item.GoldPlate, amount: 2}, {itemId: item.Amethyst, amount: 1}},
		cost:         20000,
	}
}

func (r Vicious) BronzeScalerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BronzePlate, amount: 1}, {itemId: item.SteelPlate, amount: 3}, {itemId: item.Leather, amount: 60}, {itemId: item.Screw, amount: 15}},
		cost:         30000,
	}
}

func (r Vicious) AquaBraceRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 3}, {itemId: item.Garnet, amount: 1}, {itemId: item.AquaMarine, amount: 3}, {itemId: item.Leather, amount: 80}, {itemId: item.Screw, amount: 25}},
		cost:         40000,
	}
}

func (r Vicious) BlueWillowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SilverPlate, amount: 3}, {itemId: item.GoldPlate, amount: 1}, {itemId: item.AquaMarine, amount: 2}, {itemId: item.DragonSkin, amount: 40}, {itemId: item.Screw, amount: 35}},
		cost:         50000,
	}
}

func (r Vicious) OakerGarnerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.GoldPlate, amount: 2}, {itemId: item.MoonRock, amount: 1}, {itemId: item.Topaz, amount: 8}, {itemId: item.DragonSkin, amount: 50}, {itemId: item.Screw, amount: 50}},
		cost:         70000,
	}
}

func (r Vicious) GloveUpgrade(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Upgrade a glove? That shouldn't be too difficult. Which did you have in mind?").
		NewLine().
		BlueText().OpenItem(0).ShowItemName1(item.GreenDiros).BlackText().AddText(" - Bowman Lv. 15").CloseItem().NewLine().
		BlueText().OpenItem(1).ShowItemName1(item.RedDiros).BlackText().AddText(" - Bowman Lv. 15").CloseItem().NewLine().
		BlueText().OpenItem(2).ShowItemName1(item.RedSavata).BlackText().AddText(" - Bowman Lv. 25").CloseItem().NewLine().
		BlueText().OpenItem(3).ShowItemName1(item.DarkSavata).BlackText().AddText(" - Bowman Lv. 25").CloseItem().NewLine().
		BlueText().OpenItem(4).ShowItemName1(item.GreenMarker).BlackText().AddText(" - Bowman Lv. 30").CloseItem().NewLine().
		BlueText().OpenItem(5).ShowItemName1(item.BlackMarker).BlackText().AddText(" - Bowman Lv. 30").CloseItem().NewLine().
		BlueText().OpenItem(6).ShowItemName1(item.MithrilScaler).BlackText().AddText(" - Bowman Lv. 35").CloseItem().NewLine().
		BlueText().OpenItem(7).ShowItemName1(item.GoldScaler).BlackText().AddText(" - Bowman Lv. 35").CloseItem().NewLine().
		BlueText().OpenItem(8).ShowItemName1(item.GoldBrace).BlackText().AddText(" - Bowman Lv. 40").CloseItem().NewLine().
		BlueText().OpenItem(9).ShowItemName1(item.DarkBrace).BlackText().AddText(" - Bowman Lv. 40").CloseItem().NewLine().
		BlueText().OpenItem(10).ShowItemName1(item.RedWillow).BlackText().AddText(" - Bowman Lv. 50").CloseItem().NewLine().
		BlueText().OpenItem(11).ShowItemName1(item.DarkWillow).BlackText().AddText(" - Bowman Lv. 50").CloseItem().NewLine().
		BlueText().OpenItem(12).ShowItemName1(item.SephiaGarner).BlackText().AddText(" - Bowman Lv. 60").CloseItem().NewLine().
		BlueText().OpenItem(13).ShowItemName1(item.DarkGarner).BlackText().AddText(" - Bowman Lv. 60").CloseItem().NewLine()
	return script.SendListSelection(l, c, m.String(), r.GloveUpgradeSelection)
}

func (r Vicious) GloveUpgradeSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.SingleItemConfirm(item.GreenDiros, r.GreenDirosRequirements())
	case 1:
		return r.SingleItemConfirm(item.RedDiros, r.RedDirosRequirements())
	case 2:
		return r.SingleItemConfirm(item.RedSavata, r.RedSavataRequirements())
	case 3:
		return r.SingleItemConfirm(item.DarkSavata, r.DarkSavataRequirements())
	case 4:
		return r.SingleItemConfirm(item.GreenMarker, r.GreenMarkerRequirements())
	case 5:
		return r.SingleItemConfirm(item.BlackMarker, r.BlackMarkerRequirements())
	case 6:
		return r.SingleItemConfirm(item.MithrilScaler, r.MithrilScalerRequirements())
	case 7:
		return r.SingleItemConfirm(item.GoldScaler, r.GoldScalerRequirements())
	case 8:
		return r.SingleItemConfirm(item.GoldBrace, r.GoldBraceRequirements())
	case 9:
		return r.SingleItemConfirm(item.DarkBrace, r.DarkBraceRequirements())
	case 10:
		return r.SingleItemConfirm(item.RedWillow, r.RedWillowRequirements())
	case 11:
		return r.SingleItemConfirm(item.DarkWillow, r.DarkWillowRequirements())
	case 12:
		return r.SingleItemConfirm(item.SephiaGarner, r.SephiaGarnerRequirements())
	case 13:
		return r.SingleItemConfirm(item.DarkGarner, r.DarkGarnerRequirements())
	}
	return nil
}

func (r Vicious) GreenDirosRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlueDiros, amount: 1}, {itemId: item.Emerald, amount: 2}},
		cost:         7000,
	}
}

func (r Vicious) RedDirosRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlueDiros, amount: 1}, {itemId: item.Garnet, amount: 1}},
		cost:         7000,
	}
}

func (r Vicious) RedSavataRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlueSavata, amount: 1}, {itemId: item.Garnet, amount: 3}},
		cost:         10000,
	}
}

func (r Vicious) DarkSavataRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlueSavata, amount: 1}, {itemId: item.BlackCrystal, amount: 1}},
		cost:         12000,
	}
}

func (r Vicious) GreenMarkerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BrownMarker, amount: 1}, {itemId: item.Emerald, amount: 3}},
		cost:         15000,
	}
}

func (r Vicious) BlackMarkerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BrownMarker, amount: 1}, {itemId: item.BlackCrystal, amount: 1}},
		cost:         20000,
	}
}

func (r Vicious) MithrilScalerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BronzeScaler, amount: 1}, {itemId: item.MithrilPlate, amount: 4}},
		cost:         22000,
	}
}

func (r Vicious) GoldScalerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BronzeScaler, amount: 1}, {itemId: item.GoldPlate, amount: 2}},
		cost:         25000,
	}
}

func (r Vicious) GoldBraceRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.AquaBrace, amount: 1}, {itemId: item.GoldPlate, amount: 4}},
		cost:         30000,
	}
}

func (r Vicious) DarkBraceRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.AquaBrace, amount: 1}, {itemId: item.BlackCrystal, amount: 2}},
		cost:         40000,
	}
}

func (r Vicious) RedWillowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlueWillow, amount: 1}, {itemId: item.BronzePlate, amount: 1}, {itemId: item.Garnet, amount: 5}},
		cost:         55000,
	}
}

func (r Vicious) DarkWillowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BlueWillow, amount: 1}, {itemId: item.GoldPlate, amount: 2}, {itemId: item.BlackCrystal, amount: 2}},
		cost:         60000,
	}
}

func (r Vicious) SephiaGarnerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.OakerGarner, amount: 1}, {itemId: item.Garnet, amount: 5}, {itemId: item.Diamond, amount: 1}},
		cost:         70000,
	}
}

func (r Vicious) DarkGarnerRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.OakerGarner, amount: 1}, {itemId: item.Diamond, amount: 2}, {itemId: item.BlackCrystal, amount: 2}},
		cost:         80000,
	}
}

func (r Vicious) MaterialRefine(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Materials? I know of a few materials that I can make for you...").
		NewLine().
		BlueText().OpenItem(0).AddText("Make Processed Wood with Tree Branch").CloseItem().NewLine().
		BlueText().OpenItem(1).AddText("Make Processed Wood with Firewood").CloseItem().NewLine().
		BlueText().OpenItem(2).AddText("Make Screws (packs of 15)").CloseItem().NewLine()
	return script.SendListSelection(l, c, m.String(), r.MaterialRefineSelection)
}

func (r Vicious) MaterialRefineSelection(selection int32) script.StateProducer {
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

func (r Vicious) MaterialQuantityPrompt(itemId uint32, requirements RefinementRequirements) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("So, you want me to make some ").
			ShowItemName1(itemId).
			AddText("s? In that case, how many do you want me to make?")
		return script.SendGetNumber(l, c, m.String(), r.ProcessMaterialQuantity(itemId, requirements), 1, 1, 100)
	}
}

func (r Vicious) ProcessedWoodFromTreeBranchRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.TreeBranch, amount: 10}},
		cost:         0,
	}
}

func (r Vicious) ProcessedWoodFromFirewoodRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.Firewood, amount: 5}},
		cost:         0,
	}
}

func (r Vicious) ScrewRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BronzePlate, amount: 1}, {itemId: item.SteelPlate, amount: 1}},
		cost:         0,
	}
}

func (r Vicious) ProcessMaterialQuantity(itemId uint32, requirements RefinementRequirements) script.ProcessNumber {
	return func(selection int32) script.StateProducer {
		return r.MultipleItemConfirm(itemId, selection, requirements)
	}
}

func (r Vicious) ArrowRefine(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Arrows? Not a problem at all.").
		NewLine().
		BlueText().OpenItem(0).ShowItemName1(item.ArrowForBow).CloseItem().NewLine().
		BlueText().OpenItem(1).ShowItemName1(item.ArrowForCrossbow).CloseItem().NewLine().
		BlueText().OpenItem(2).ShowItemName1(item.BronzeArrowForBow).CloseItem().NewLine().
		BlueText().OpenItem(3).ShowItemName1(item.BronzeArrowForCrossbow).CloseItem().NewLine().
		BlueText().OpenItem(4).ShowItemName1(item.SteelArrowForBow).CloseItem().NewLine().
		BlueText().OpenItem(5).ShowItemName1(item.SteelArrowForCrossbow).CloseItem().NewLine()
	return script.SendListSelection(l, c, m.String(), r.ArrowRefineSelection)
}

func (r Vicious) ArrowRefineSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.SingleItemConfirm(item.ArrowForBow, r.ArrowForBowRequirements())
	case 1:
		return r.SingleItemConfirm(item.ArrowForCrossbow, r.ArrowForCrossbowRequirements())
	case 2:
		return r.SingleItemConfirm(item.BronzeArrowForBow, r.BronzeArrowForBowRequirements())
	case 3:
		return r.SingleItemConfirm(item.BronzeArrowForCrossbow, r.BronzeArrowForCrossbowRequirements())
	case 4:
		return r.SingleItemConfirm(item.SteelArrowForBow, r.SteelArrowForBowRequirements())
	case 5:
		return r.SingleItemConfirm(item.SteelArrowForCrossbow, r.SteelArrowForCrossbow())
	}
	return nil
}

func (r Vicious) ArrowForBowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.ProcessedWood, amount: 1}, {itemId: item.StiffFeather, amount: 1}},
		cost:         0,
	}
}

func (r Vicious) ArrowForCrossbowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.ProcessedWood, amount: 1}, {itemId: item.StiffFeather, amount: 1}},
		cost:         0,
	}
}

func (r Vicious) BronzeArrowForBowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BronzePlate, amount: 1}, {itemId: item.ProcessedWood, amount: 3}, {itemId: item.StiffFeather, amount: 10}},
		cost:         0,
	}
}

func (r Vicious) BronzeArrowForCrossbowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.BronzePlate, amount: 1}, {itemId: item.ProcessedWood, amount: 3}, {itemId: item.StiffFeather, amount: 10}},
		cost:         0,
	}
}

func (r Vicious) SteelArrowForBowRequirements() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 1}, {itemId: item.ProcessedWood, amount: 5}, {itemId: item.SoftFeather, amount: 15}},
		cost:         0,
	}
}

func (r Vicious) SteelArrowForCrossbow() RefinementRequirements {
	return RefinementRequirements{
		requirements: []Requirement{{itemId: item.SteelPlate, amount: 1}, {itemId: item.ProcessedWood, amount: 5}, {itemId: item.SoftFeather, amount: 15}},
		cost:         0,
	}
}

func (r Vicious) Sorry(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Sorry, but this is how I make my living. No meso, no item.")
	return script.SendOk(l, c, m.String())
}

func (r Vicious) NeedMoreItems(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Surely you, of all people, would understand the value of having quality items? I can't do that without the items I require.")
	return script.SendOk(l, c, m.String())
}

func (r Vicious) NotEnoughInventorySpace(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please make sure you have room in your inventory, and talk to me again.")
	return script.SendOk(l, c, m.String())
}
