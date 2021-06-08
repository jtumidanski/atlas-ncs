package script

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// Brittany is located in Victoria Road - Henesys Hair Salon (100000104)
type Brittany struct {
}

func (r Brittany) NPCId() uint32 {
	return npc.Brittany
}

func (r Brittany) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Brittany) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I'm Brittany the assistant. If you have ").
		BlueText().ShowItemName1(item.HenesysHairStyleCouponREG).
		BlackText().AddText(", ").
		BlueText().ShowItemName1(item.HenesysHairStyleCouponEXP).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.HenesysHairColorCouponREG).
		BlackText().AddText(" by any chance, then how about letting me change your hairdo?").NewLine().
		OpenItem(0).AddText("Haircut: ").ShowItemImage2(item.HenesysHairStyleCouponREG).ShowItemName1(item.HenesysHairStyleCouponREG).CloseItem().NewLine().
		OpenItem(1).AddText("Haircut: ").ShowItemImage2(item.HenesysHairStyleCouponEXP).ShowItemName1(item.HenesysHairStyleCouponEXP).CloseItem().NewLine().
		OpenItem(2).AddText("Dye your hair: ").ShowItemImage2(item.HenesysHairColorCouponREG).ShowItemName1(item.HenesysHairColorCouponREG).CloseItem().NewLine()
	return SendListSelection(l, c, m.String(), r.Choose)
}

func (r Brittany) Choose(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.StyleRegular
	case 1:
		return r.StyleExperimental
	case 2:
		return r.ColorRegular
	}
	return nil
}

func (r Brittany) StyleRegular(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you use the REG coupon your hair will change RANDOMLY with a chance to obtain a new experimental style that even you didn't think was possible. Are you going to use ").
		BlueText().ShowItemName1(item.HenesysHairStyleCouponREG).
		BlackText().AddText(" and really change your hairstyle?")
	return SendYesNo(l, c, m.String(), r.PerformRegularStyle, r.Hello)
}

func (r Brittany) StyleExperimental(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you use the EXP coupon your hair will change RANDOMLY with a chance to obtain a new experimental style that even you didn't think was possible. Are you going to use ").
		BlueText().ShowItemName1(item.HenesysHairStyleCouponEXP).
		BlackText().AddText(" and really change your hairstyle?")
	return SendYesNo(l, c, m.String(), r.PerformExperimentalStyle, r.Hello)
}

func (r Brittany) ColorRegular(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you use a regular coupon your hair will change RANDOMLY. Do you still want to use ").
		BlueText().ShowItemName1(item.HenesysHairColorCouponREG).
		BlackText().AddText(" and change it up?")
	return SendYesNo(l, c, m.String(), r.PerformRegularColor, r.Hello)
}

func (r Brittany) PerformRegularStyle(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.HenesysHairStyleCouponREG) {
		return r.MissingStyleCoupon(l, c)
	}

	hair := make([]uint32, 0)
	gender := character.GetGender(l)(c.CharacterId)
	if gender == character.GenderMale {
		hair = r.RegularMaleHair()
	} else if gender == character.GenderFemale {
		hair = r.RegularFemaleHair()
	}
	hair = care.ApplyCurrentColor(l)(c.CharacterId, hair)
	hair = care.FilterCurrent(l)(c.CharacterId, hair)
	randomHair := hair[rand.Intn(len(hair))]

	character.GainItem(l)(c.CharacterId, item.HenesysHairStyleCouponREG, -1)
	character.SetHair(l)(c.CharacterId, randomHair)
	return r.EnjoyStyle(l, c)
}

func (r Brittany) RegularMaleHair() []uint32 {
	return []uint32{character.HairBlackCatalyst, character.HairBlackTopknot, character.HairBlackWind, character.HairBlackShaggyWax, character.HairBlackAcorn, character.HairBlackTheMoRawk, character.HairBlackAranCut, character.HairBlackTheCoco}
}

func (r Brittany) RegularFemaleHair() []uint32 {
	return []uint32{character.HairBlackStella, character.HairBlackRainbow, character.HairBlackAngelica, character.HairBlackChantelle, character.HairBlackFourtailBraids, character.HairBlackCrazyMedusa, character.HairBlackAranHair, character.HairBlackFullBangs}
}

func (r Brittany) MissingStyleCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't give you a haircut without it. I'm sorry...")
	return SendOk(l, c, m.String())
}

func (r Brittany) EnjoyStyle(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Enjoy your new and improved hairstyle!")
	return SendOk(l, c, m.String())
}

func (r Brittany) PerformExperimentalStyle(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.HenesysHairStyleCouponEXP) {
		return r.MissingStyleCoupon(l, c)
	}

	hair := make([]uint32, 0)
	gender := character.GetGender(l)(c.CharacterId)
	if gender == character.GenderMale {
		hair = r.ExperimentalMaleHair()
	} else if gender == character.GenderFemale {
		hair = r.ExperimentalFemaleHair()
	}
	hair = care.ApplyCurrentColor(l)(c.CharacterId, hair)
	hair = care.FilterCurrent(l)(c.CharacterId, hair)
	randomHair := hair[rand.Intn(len(hair))]

	character.GainItem(l)(c.CharacterId, item.HenesysHairStyleCouponEXP, -1)
	character.SetHair(l)(c.CharacterId, randomHair)
	return r.EnjoyStyle(l, c)
}

func (r Brittany) ExperimentalMaleHair() []uint32 {
	return []uint32{character.HairBlackBuzz, character.HairBlackTopknot, character.HairBlackWind, character.HairBlackShaggyWax, character.HairBlackAcorn, character.HairBlackTheMoRawk, character.HairBlackAranCut, character.HairBlackTheCoco}
}

func (r Brittany) ExperimentalFemaleHair() []uint32 {
	return []uint32{character.HairBlackStella, character.HairBlackAngelica, character.HairBlackChantelle, character.HairBlackFourtailBraids, character.HairSkinHead, character.HairBlackCrazyMedusa, character.HairBlackAranHair, character.HairBlackFullBangs}
}

func (r Brittany) PerformRegularColor(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.HenesysHairColorCouponREG) {
		return r.MissingColorCoupon(l, c)
	}
	hair := care.ProduceColorOptions(l, c)
	randomHair := hair[rand.Intn(len(hair))]
	character.GainItem(l)(c.CharacterId, item.HenesysHairColorCouponREG, -1)
	character.SetHair(l)(c.CharacterId, randomHair)
	return r.EnjoyColor(l, c)
}

func (r Brittany) MissingColorCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't dye your hair without it. I'm sorry...")
	return SendOk(l, c, m.String())
}

func (r Brittany) EnjoyColor(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Enjoy your new and improved hair color!")
	return SendOk(l, c, m.String())
}
