package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ArwenTheFairy is located in Victoria Road - Ellinia (101000000)
type ArwenTheFairy struct {
}

func (r ArwenTheFairy) NPCId() uint32 {
	return npc.ArwenTheFairy
}

func (r ArwenTheFairy) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.IsLevel(l)(c.CharacterId, 40) {
		return r.HelloMaker(l, c)
	}
	return r.HelloStranger(l, c)
}

func (r ArwenTheFairy) HelloMaker(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Yeah... I am the master alchemist of the fairies. But the fairies are not supposed to be in contact with a human being for a long period of time... A strong person like you will be fine, though. If you get me the materials, I'll make you a special item.")
	return script.SendNext(l, c, m.String(), r.WhatToMake)
}

func (r ArwenTheFairy) HelloStranger(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I can make rare, valuable items but unfortunately I can't make it to a stranger like you.")
	return script.SendOk(l, c, m.String())
}

func (r ArwenTheFairy) WhatToMake(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("What do you want to make?").NewLine().
		OpenItem(0).BlueText().AddText("Moon Rock").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Star Rock").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Black Feather").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.ProcessMake)
}

func (r ArwenTheFairy) ProcessMake(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.ConfirmMoonRock
	case 1:
		return r.ConfirmStarRock
	case 2:
		return r.ConfirmBlackFeather
	}
	return nil
}

type ArwenRequirements struct {
	items []ArwenRequirement
	cost  uint32
}

type ArwenRequirement struct {
	itemId uint32
}

func (r ArwenTheFairy) ConfirmMoonRock(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So you want to make a Moon Rock? To do that you need to refine one of each of these: ").
		BlueText().AddText("Bronze Plate").
		BlackText().AddText(", ").
		BlueText().AddText("Steel Plate").
		BlackText().AddText(",").NewLine().
		BlueText().AddText("Mithril Plate").
		BlackText().AddText(", ").
		BlueText().AddText("Adamantium Plate").
		BlackText().AddText(", ").
		BlueText().AddText("Silver Plate").
		BlackText().AddText(", ").
		BlueText().AddText("Orihalcon Plate").
		BlackText().AddText(" and ").
		BlueText().AddText("Gold Plate").
		BlackText().AddText(". Throw in 10,000 mesos and I'll make it for you.")
	return script.SendYesNoExit(l, c, m.String(), r.ValidateMoonRock, r.WhatToMake, r.GetMaterialsReady(item.MoonRock))
}

func (r ArwenTheFairy) ConfirmStarRock(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So you want to make a Star Rock? To do that you need to refine one of each of these: ").
		BlueText().AddText("Garnet").
		BlackText().AddText(", ").
		BlueText().AddText("Amethyst").
		BlackText().AddText(", ").
		BlueText().AddText("AquaMarine").
		BlackText().AddText(", ").
		BlueText().AddText("Emerald").
		BlackText().AddText(", ").
		BlueText().AddText("Opal").
		BlackText().AddText(", ").
		BlueText().AddText("Sapphire").
		BlackText().AddText(", ").
		BlueText().AddText("Topaz").
		BlackText().AddText(", ").
		BlueText().AddText("Diamond").
		BlackText().AddText(" and ").
		BlueText().AddText("Black Crystal").
		BlackText().AddText(". Throw in 15,000 mesos and I'll make it for you.")
	return script.SendYesNoExit(l, c, m.String(), r.ValidateStarRock, r.WhatToMake, r.GetMaterialsReady(item.StarRock))
}

func (r ArwenTheFairy) ConfirmBlackFeather(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So you want to make a Black Feather? To do that you need ").
		BlueText().AddText("1 Flaming Feather").
		BlackText().AddText(", ").
		BlueText().AddText("1 Moon Rock").
		BlackText().AddText(" and ").
		BlueText().AddText("1 Black Crystal").
		BlackText().AddText(". Throw in 30,000 mesos and I'll make it for you. Oh yeah, this piece of feather is a very special item, so if you drop it by any chance, it'll disappear, as well as you won't be able to give it away to someone else.")
	return script.SendYesNoExit(l, c, m.String(), r.ValidateBlackFeather, r.WhatToMake, r.GetMaterialsReady(item.BlackFeather))
}

func (r ArwenTheFairy) ValidateMoonRock(l logrus.FieldLogger, c script.Context) script.State {
	return r.Validate(item.MoonRock, r.MoonRockRequirements())(l, c)
}

func (r ArwenTheFairy) ValidateStarRock(l logrus.FieldLogger, c script.Context) script.State {
	return r.Validate(item.StarRock, r.StarRockRequirements())(l, c)
}

func (r ArwenTheFairy) ValidateBlackFeather(l logrus.FieldLogger, c script.Context) script.State {
	return r.Validate(item.BlackFeather, r.BlackFeatherRequirements())(l, c)
}

func (r ArwenTheFairy) MoonRockRequirements() ArwenRequirements {
	return ArwenRequirements{
		items: []ArwenRequirement{{itemId: 4011000}, {itemId: 4011001}, {itemId: 4011002}, {itemId: 4011003}, {itemId: 4011004}, {itemId: 4011005}, {itemId: 4011006}},
		cost:  10000,
	}
}

func (r ArwenTheFairy) StarRockRequirements() ArwenRequirements {
	return ArwenRequirements{
		items: []ArwenRequirement{{itemId: 4021000}, {itemId: 4021001}, {itemId: 4021002}, {itemId: 4021003}, {itemId: 4021004}, {itemId: 4021005}, {itemId: 4021006}, {itemId: 4021007}, {itemId: 4021008}},
		cost:  15000,
	}
}

func (r ArwenTheFairy) BlackFeatherRequirements() ArwenRequirements {
	return ArwenRequirements{
		items: []ArwenRequirement{{itemId: 4001006}, {itemId: 4011007}, {itemId: 4021008}},
		cost:  30000,
	}
}

func (r ArwenTheFairy) Validate(itemId uint32, requirements ArwenRequirements) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if !character.HasMeso(l)(c.CharacterId, requirements.cost) {
			return r.MoreMeso(l, c)
		}

		for _, i := range requirements.items {
			if !character.HasItem(l)(c.CharacterId, i.itemId) {
				return r.NeedMoreOfItem(i.itemId)(l, c)
			}
		}

		err := character.GainMeso(l)(c.CharacterId, -int32(requirements.cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process purchase for character %d.", c.CharacterId)
		}
		for _, i := range requirements.items {
			character.GainItem(l)(c.CharacterId, i.itemId, -1)
		}
		character.GainItem(l)(c.CharacterId, itemId, 1)
		return r.Success(itemId)(l, c)
	}
}

func (r ArwenTheFairy) MoreMeso(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Are you sure you have enough mesos?")
	return script.SendOk(l, c, m.String())
}

func (r ArwenTheFairy) NeedMoreOfItem(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Please check and see if you have a ").
			ShowItemName1(itemId)
		return script.SendOk(l, c, m.String())
	}
}

func (r ArwenTheFairy) Success(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Ok here, take ").
			ShowItemName1(itemId).
			AddText(". It's well-made, probably because I'm using good materials. If you need my help down the road, feel free to come back.")
		return script.SendOk(l, c, m.String())
	}
}

func (r ArwenTheFairy) GetMaterialsReady(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("It's not easy making ").
			ShowItemName1(itemId).
			AddText(". Please get the materials ready.")
		return script.SendOk(l, c, m.String())
	}
}
