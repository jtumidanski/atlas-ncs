package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// BigHeadward is located in Victoria Road - Henesys Hair Salon (100000104)
type BigHeadward struct {
}

func (r BigHeadward) NPCId() uint32 {
	return npc.BigHeadward
}

func (r BigHeadward) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r BigHeadward) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hi, I'm ").
		ShowNPC(npc.BigHeadward).
		AddText(", the most charming and stylish stylist around. If you're looking for the best looking hairdos around, look no further!").NewLine().
		OpenItem(0).ShowItemImage2(item.RoyalHairCoupon).ShowItemName1(item.RoyalHairCoupon).CloseItem().NewLine().
		OpenItem(1).ShowItemImage2(item.SpecialRoyalHairCoupon).ShowItemName1(item.SpecialRoyalHairCoupon).CloseItem()
	return SendListSelection(l, c, m.String(), r.Choice)
}

func (r BigHeadward) Choice(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.RoyalHair
	case 1:
		return r.SpecialRoyalHair
	}
	return nil
}

func (r BigHeadward) RoyalHair(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you use this REGULAR coupon, your hair may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(item.RoyalHairCoupon).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!")
	return SendYesNo(l, c, m.String(), r.ValidateRoyalHair, r.Hello)
}

func (r BigHeadward) ValidateRoyalHair(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.RoyalHairCoupon) {
		return r.MissingCoupon(l, c)
	}
	return r.GiveRoyalHair(l, c)
}

func (r BigHeadward) MissingCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't give you a haircut without it. I'm sorry...")
	return SendOk(l, c, m.String())
}

func (r BigHeadward) GiveRoyalHair(l logrus.FieldLogger, c Context) State {
	hair := make([]uint32, 0)
	gender := character.GetGender(l)(c.CharacterId)
	if gender == character.GenderMale {
		hair = r.RoyalMaleHair()
	} else if gender == character.GenderFemale {
		hair = r.RoyalFemaleHair()
	}
	hair = ApplyHairColor(l)(c.CharacterId, hair)

	character.GainItem(l)(c.CharacterId, item.RoyalHairCoupon, -1)
	random := rand.Intn(len(hair))
	character.SetHair(l)(c.CharacterId, hair[random])

	return r.Success(l, c)
}

func (r BigHeadward) Success(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Enjoy your new and improved hairstyle!")
	return SendOk(l, c, m.String())
}

func (r BigHeadward) SpecialRoyalHair(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Using the SPECIAL coupon you can choose the style your hair will become. Pick the style that best provides you delight...")

	hair := make([]uint32, 0)
	gender := character.GetGender(l)(c.CharacterId)
	if gender == character.GenderMale {
		hair = r.SpecialMaleHair()
	} else if gender == character.GenderFemale {
		hair = r.SpecialFemaleHair()
	}
	hair = ApplyHairColor(l)(c.CharacterId, hair)

	return SendStyle(l, c, m.String(), r.ValidateSpecialRoyal(hair), hair)
}

func (r BigHeadward) ValidateSpecialRoyal(options []uint32) ProcessSelection {
	return func(selection int32) StateProducer {
		if selection < 0 || int(selection) >= len(options) {
			return Exit()
		}

		return func(l logrus.FieldLogger, c Context) State {
			if !character.HasItem(l)(c.CharacterId, item.SpecialRoyalHairCoupon) {
				return r.MissingCoupon(l, c)
			}
			return r.DoSpecialRoyal(options[int(selection)])(l, c)
		}
	}
}

func (r BigHeadward) RoyalMaleHair() []uint32 {
	return []uint32{character.HairZeta, character.HairAllBack, character.HairMilitaryBuzzcut, character.HairMohawk, character.HairBlueFantasy, character.HairBlackMetroMan, character.HairBlackBowlingBall, character.HairBlackPrinceCut}
}

func (r BigHeadward) RoyalFemaleHair() []uint32 {
	return []uint32{character.HairBlackJolie, character.HairBlackZessica, character.HairBlackGrace, character.HairBlackCLHair, character.HairBlackSpunkyDo, character.HairBlackPalmTreeHair, character.HairBlackDesignerHair}
}

func (r BigHeadward) SpecialMaleHair() []uint32 {
	return []uint32{character.HairZeta, character.HairAllBack, character.HairMilitaryBuzzcut, character.HairMohawk, character.HairBlueFantasy, character.HairBlackBabbyCut, character.HairBlackGrandLionman, character.HairBlackMetroMan, character.HairBlackBowlingBall, character.HairBlackCornrow, character.HairBlackEasternMystery, character.HairBlackBoyBandCut, character.HairBlackVolumeCut}
}

func (r BigHeadward) SpecialFemaleHair() []uint32 {
	return []uint32{character.HairBlackFrancesca, character.HairBlackJolie, character.HairBlackMinnie, character.HairBlackZessica, character.HairBlackGrace, character.HairBlackLagunaBeach, character.HairBlackShortTwinTails, character.HairBlackCLHair, character.HairBlackSpunkyDo, character.HairBlackVintageFlip, character.HairBlackPalmTreeHair}
}

func (r BigHeadward) DoSpecialRoyal(hairId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.GainItem(l)(c.CharacterId, item.SpecialRoyalHairCoupon, -1)
		character.SetHair(l)(c.CharacterId, hairId)
		return r.Success(l, c)
	}
}
