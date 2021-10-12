package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// StaffSeargantCharlie is located in Orbis - Orbis (200000000)
type StaffSeargantCharlie struct {
}

func (r StaffSeargantCharlie) NPCId() uint32 {
	return npc.StaffSeargantCharlie
}

func (r StaffSeargantCharlie) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r StaffSeargantCharlie) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey, got a little bit of time? Well, my job is to collect items here and sell them elsewhere, but these days the monsters have become much more hostile so it's been difficult to getting good items ... What do you think? Do you want to do some business with me?")
	return script.SendNext(l, span, c, m.String(), r.DealIsSimple)
}

func (r StaffSeargantCharlie) DealIsSimple(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The deal is simple. You get me something I need, I get you something you need. The problem is, I deal with a whole bunch of people, so the items I have to offer may change every time you see me. What do you think? Still want to do it?")
	return script.SendYesNo(l, span, c, m.String(), r.Choices, r.ShouldNotBeBad)
}

func (r StaffSeargantCharlie) ShouldNotBeBad(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm...it shouldn't be a bad deal for you. Come see me at the right time and you may get a much better item to be offered. Anyway, let me know when you have a change of heart.")
	return script.SendOk(l, span, c, m.String())
}

var CharliesChoices = []uint32{item.SolidHorn, item.StarPixiesStarpiece, item.LunarPixiesMoonpiece, item.LusterPixiesSunpiece, item.NependeathSeed, item.DarkNependeathSeed, item.JrYetiSkin, item.YetiHorn, item.DarkJrYetiSkin, item.DarkYetiHorn, item.HectorTail, item.WhitePangTail, item.PepeBeak, item.DarkPepeBeak, item.WerewolfToenail, item.LycanthropeToenail, item.FlyEyeWing, item.JrCerebesTooth, item.FirebombFlame, item.CellionTail, item.LionerTail, item.GrupinTail, item.ZombiesLostTooth, item.CerebesTooth, item.BainsSpikeyCollar}

var CharliesPrizes = [][]CharliesPrize{
	{{itemId: item.OrangePotion, quantity: 20}, {itemId: item.Lemon, quantity: 10}, {itemId: item.BluePotion, quantity: 15}, {itemId: item.ProcessedWood, quantity: 15}, {itemId: item.FriedChicken, quantity: 15}, {itemId: item.ReturnScrollNearest, quantity: 15}},
	{{itemId: item.BluePotion, quantity: 20}, {itemId: item.OrangePotion, quantity: 30}, {itemId: item.Meat, quantity: 40}, {itemId: item.ProcessedWood, quantity: 20}, {itemId: item.ScrollForHelmutDefense10, quantity: 1}},
	{{itemId: item.WhitePotion, quantity: 25}, {itemId: item.ManaElixer, quantity: 10}, {itemId: item.PureWater, quantity: 5}, {itemId: item.DragonSkin, quantity: 15}, {itemId: item.ScrollForShieldForDefense10, quantity: 1}},
	{{itemId: item.WhitePotion, quantity: 30}, {itemId: item.ManaElixer, quantity: 15}, {itemId: item.Salad, quantity: 20}, {itemId: item.Screw, quantity: 5}, {itemId: item.ScrollForCapeForIntelligence60, quantity: 1}},
	{{itemId: item.WhitePotion, quantity: 15}, {itemId: item.Lemon, quantity: 15}, {itemId: item.BluePotion, quantity: 25}, {itemId: item.ProcessedWood, quantity: 30}, {itemId: item.ScrollForEarringForIntelligence10, quantity: 1}},
	{{itemId: item.WhitePotion, quantity: 30}, {itemId: item.ManaElixer, quantity: 15}, {itemId: item.Salad, quantity: 20}, {itemId: item.Screw, quantity: 5}, {itemId: item.ScrollForTopwearForDefense10, quantity: 1}},
	{{itemId: item.WhitePotion, quantity: 30}, {itemId: item.Salad, quantity: 20}, {itemId: item.ManaElixer, quantity: 15}, {itemId: item.Screw, quantity: 5}, {itemId: item.ScrollForTopwearForDefense10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 25}, {itemId: item.Salad, quantity: 20}, {itemId: item.GarnetOre, quantity: 7}, {itemId: item.AmethystOre, quantity: 7}, {itemId: item.AquaMarineOre, quantity: 3}, {itemId: item.DiamondOre, quantity: 2}, {itemId: item.ScrollForShoesForSpeed10, quantity: 1}},
	{{itemId: item.HotDog, quantity: 30}, {itemId: item.HotDogSupreme, quantity: 15}, {itemId: item.RedBeanPorridge, quantity: 30}, {itemId: item.FairyWing, quantity: 1}, {itemId: item.ScrollForArmorForDefense10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 25}, {itemId: item.SapphireOre, quantity: 7}, {itemId: item.EmeraldOre, quantity: 7}, {itemId: item.OpalOre, quantity: 7}, {itemId: item.BlackCrystalOre, quantity: 2}, {itemId: item.ScrollForGlovesForDexterity10, quantity: 1}},
	{{itemId: item.WarriorPotion, quantity: 15}, {itemId: item.SniperPotion, quantity: 15}, {itemId: item.WizardPotion, quantity: 15}, {itemId: item.AncientScroll, quantity: 1}, {itemId: item.ScrollForOverallArmorForDexterity10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 20}, {itemId: item.SilverOre, quantity: 7}, {itemId: item.AdamantiumOre, quantity: 7}, {itemId: item.OrihalconOre, quantity: 7}, {itemId: item.PieceOfIce, quantity: 1}, {itemId: item.ScrollForBottomwearForDefense10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 20}, {itemId: item.MithrilOre, quantity: 7}, {itemId: item.SteelOre, quantity: 7}, {itemId: item.BronzeOre, quantity: 7}, {itemId: item.GoldOre, quantity: 2}, {itemId: item.Screw, quantity: 5}, {itemId: item.ScrollForShoesForDexterity10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 20}, {itemId: item.SilverOre, quantity: 7}, {itemId: item.OrihalconOre, quantity: 7}, {itemId: item.GoldOre, quantity: 3}, {itemId: item.DiamondOre, quantity: 2}, {itemId: item.BlackCrystalOre, quantity: 2}, {itemId: item.ScrollForShoesForJump10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 30}, {itemId: item.TopazOre, quantity: 7}, {itemId: item.BlackCrystalOre, quantity: 2}, {itemId: item.DiamondOre, quantity: 2}, {itemId: item.Icicle, quantity: 1}, {itemId: item.ScrollForGlovesForAttack10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 30}, {itemId: item.TopazOre, quantity: 7}, {itemId: item.BlackCrystalOre, quantity: 2}, {itemId: item.DiamondOre, quantity: 2}, {itemId: item.ScrollForCapeForDexterity10, quantity: 1}},
	{{itemId: item.OrangePotion, quantity: 30}, {itemId: item.BluePotion, quantity: 20}, {itemId: item.ProcessedWood, quantity: 20}, {itemId: item.Meat, quantity: 40}, {itemId: item.ScrollForHelmutDefense10, quantity: 1}},
	{{itemId: item.WhitePotion, quantity: 15}, {itemId: item.BluePotion, quantity: 25}, {itemId: item.Lemon, quantity: 15}, {itemId: item.AllCurePotion, quantity: 15}, {itemId: item.ProcessedWood, quantity: 30}, {itemId: item.ScrollForEarringForIntelligence10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 25}, {itemId: item.HotDogSupreme, quantity: 25}, {itemId: item.SilverOre, quantity: 8}, {itemId: item.OrihalconOre, quantity: 8}, {itemId: item.GoldOre, quantity: 3}, {itemId: item.DiamondOre, quantity: 2}, {itemId: item.BlackCrystalOre, quantity: 2}, {itemId: item.ScrollForShoesForJump10, quantity: 1}},
	{{itemId: item.WhitePotion, quantity: 30}, {itemId: item.Salad, quantity: 20}, {itemId: item.ManaElixer, quantity: 15}, {itemId: item.Screw, quantity: 5}, {itemId: item.ScrollForCapeForWeaponDefense10, quantity: 1}},
	{{itemId: item.WhitePotion, quantity: 30}, {itemId: item.Salad, quantity: 20}, {itemId: item.ManaElixer, quantity: 15}, {itemId: item.Screw, quantity: 5}, {itemId: item.ScrollForCapeForWeaponDefense10, quantity: 1}},
	{{itemId: item.WhitePotion, quantity: 30}, {itemId: item.Salad, quantity: 20}, {itemId: item.ManaElixer, quantity: 15}, {itemId: item.Screw, quantity: 5}, {itemId: item.ScrollForCapeForWeaponDefense10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 20}, {itemId: item.HotDog, quantity: 30}, {itemId: item.HotDogSupreme, quantity: 15}, {itemId: item.AllCurePotion, quantity: 20}, {itemId: item.FairyWing, quantity: 1}, {itemId: item.ScrollForCapeForMagicDefense10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 25}, {itemId: item.AllCurePotion, quantity: 50}, {itemId: item.RedBeanPorridge, quantity: 35}, {itemId: item.GarnetOre, quantity: 8}, {itemId: item.AmethystOre, quantity: 8}, {itemId: item.AquaMarineOre, quantity: 8}, {itemId: item.DiamondOre, quantity: 2}, {itemId: item.ScrollForCapeForLuck10, quantity: 1}},
	{{itemId: item.ManaElixer, quantity: 35}, {itemId: item.TopazOre, quantity: 9}, {itemId: item.BlackCrystalOre, quantity: 4}, {itemId: item.DiamondOre, quantity: 4}, {itemId: item.ScrollForCapeForHP10, quantity: 1}},
}

type CharliesPrize struct {
	itemId   uint32
	quantity uint32
}

func (r StaffSeargantCharlie) Choices(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ok! First you need to choose the item that you'll trade with. The better the item, the more likely the chance that I'll give you something much nicer in return.").NewLine()
	for i, choice := range CharliesChoices {
		m = m.OpenItem(i).ShowItemImage1(choice).AddText("  ").ShowItemName1(choice).CloseItem().NewLine()

	}
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r StaffSeargantCharlie) Selection(selection int32) script.StateProducer {
	itemId := CharliesChoices[selection]
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Let's see, you want to trade your ").
			BlueText().AddText("100 ").
			ShowItemName1(itemId).
			BlackText().AddText(" with my stuff right? Before trading make sure you have an empty slot available on your use or etc. inventory. Now, do you want to trade with me?")
		return script.SendYesNo(l, span, c, m.String(), r.Validate(selection, itemId), r.ShouldNotBeBad)
	}
}

func (r StaffSeargantCharlie) Validate(selection int32, itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasItems(l, span)(c.CharacterId, itemId, 100) {
			return r.AreYouSure(itemId)(l, span, c)
		}

		random := rand.Intn(len(CharliesPrizes[selection]))
		prize := CharliesPrizes[selection][random]
		if !character.CanHoldAll(l)(c.CharacterId, prize.itemId, prize.quantity) {
			return r.InventoryFull(l, span, c)
		}

		character.GainItem(l, span)(c.CharacterId, itemId, -100)
		character.GainExperience(l)(c.CharacterId, 500)
		character.GainItem(l, span)(c.CharacterId, prize.itemId, int32(prize.quantity))
		return r.WhatDoYouThink(itemId, prize)(l, span, c)
	}
}

func (r StaffSeargantCharlie) AreYouSure(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Hmmm... are you sure you have ").
			BlueText().AddText("100 ").
			ShowItemName1(itemId).
			BlackText().AddText("? If so, then please check and see if your item inventory is full or not.")
		return script.SendOk(l, span, c, m.String())
	}
}

func (r StaffSeargantCharlie) InventoryFull(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Your use and etc. inventory seems to be full. You need the free spaces to trade with me! Make room, and then find me.")
	return script.SendOk(l, span, c, m.String())
}

func (r StaffSeargantCharlie) WhatDoYouThink(itemId uint32, prize CharliesPrize) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("For your ").
			BlueText().AddText("100 ").
			ShowItemName1(itemId).
			BlackText().AddText(", here's my ").
			BlueText().AddText(fmt.Sprintf("%d ", prize.quantity)).
			ShowItemName1(prize.itemId).
			BlackText().AddText(". What do you think? Do you like the items I gave you in return? I plan on being here for a while, so if you gather up more items, I'm always open for a trade ...")
		return script.SendOk(l, span, c, m.String())
	}
}
