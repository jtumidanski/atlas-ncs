package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MinoTheOwner is located in Orbis Park - Orbis Hair Salon (200000202)
type MinoTheOwner struct {
}

func (r MinoTheOwner) NPCId() uint32 {
	return npc.MinoTheOwner
}

func (r MinoTheOwner) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r MinoTheOwner) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hello I'm Mino. If you have either a ").
		BlueText().ShowItemName1(item.OrbisHairStyleCouponVIP).
		BlackText().AddText(", then please let me take care of your hair. Choose what you want to do with it.").NewLine().
		OpenItem(0).AddText("Haircut: ").ShowItemImage2(item.OrbisHairStyleCouponVIP).ShowItemName1(item.OrbisHairStyleCouponVIP).CloseItem().NewLine().
		OpenItem(1).AddText("Dye your hair: ").ShowItemImage2(item.OrbisHairColorCouponVIP).ShowItemName1(item.OrbisHairColorCouponVIP).CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r MinoTheOwner) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.HairStyle
	case 1:
		return r.HairColor
	}
	return nil
}

func (r MinoTheOwner) HairStyle(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I can totally change up your hairstyle and make it look so good. Why don't you change it up a bit? With ").
		BlueText().ShowItemName1(item.OrbisHairStyleCouponVIP).
		BlackText().AddText(", I'll take care of the rest for you. Choose the style of your liking!")

	hair := make([]uint32, 0)
	gender := character.GetGender(l)(c.CharacterId)
	if gender == character.GenderMale {
		hair = r.BaseMaleHair()
	} else if gender == character.GenderFemale {
		hair = r.BaseFemaleHair()
	}
	hair = ApplyHairColor(l)(c.CharacterId, hair)
	hair = FilterCurrentHair(l)(c.CharacterId, hair)

	return SendStyle(l, c, m.String(), r.ProcessHairChange(hair), hair)
}

func (r MinoTheOwner) HairColor(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I can totally change your hair color and make it look so good. Why don't you change it up a bit? With ").
		BlueText().ShowItemName1(item.OrbisHairColorCouponVIP).
		BlackText().AddText(", I'll take care of the rest. Choose the color of your liking!")

	hair := ProduceColorOptions(l, c)

	return SendStyle(l, c, m.String(), r.ProcessColorChange, hair)
}

func (r MinoTheOwner) BaseMaleHair() []uint32 {
	return []uint32{character.HairBlackFoilPerm, character.HairBlackMetrosexual, character.HairBlackMohecanShaggyDo, character.HairBlackTristan, character.HairBlackMessySpike}
}

func (r MinoTheOwner) BaseFemaleHair() []uint32 {
	return []uint32{character.HairBlackMonica, character.HairBlackCaspia, character.HairBlackRose, character.HairBlackTheHoneybun, character.HairBlackPrincessa}
}

func (r MinoTheOwner) ProcessHairChange(choices []uint32) ProcessSelection {
	return func(selection int32) StateProducer {
		choice := choices[selection]
		return func(l logrus.FieldLogger, c Context) State {
			if character.HasItem(l)(c.CharacterId, item.OrbisHairMembershipCoupon) {
				character.SetHair(l)(c.CharacterId, choice)
				return r.EnjoyHair(l, c)
			} else if character.HasItem(l)(c.CharacterId, item.OrbisHairStyleCouponVIP) {
				character.GainItem(l)(c.CharacterId, item.OrbisHairStyleCouponVIP, -1)
				character.SetHair(l)(c.CharacterId, choice)
				return r.EnjoyHair(l, c)
			}
			return r.MissingCoupon(l, c)
		}
	}
}

func (r MinoTheOwner) EnjoyHair(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Enjoy your new and improved hairstyle!")
	return SendOk(l, c, m.String())
}

func (r MinoTheOwner) MissingCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't give you a haircut without it. I'm sorry...")
	return SendOk(l, c, m.String())
}

func (r MinoTheOwner) ProcessColorChange(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if character.HasItem(l)(c.CharacterId, item.OrbisHairColorCouponVIP) {
			character.GainItem(l)(c.CharacterId, item.OrbisHairColorCouponVIP, -1)
			character.SetHair(l)(c.CharacterId, uint32(selection))
			return r.EnjoyColor(l, c)
		}
		return r.MissingColorCoupon(l, c)
	}
}

func (r MinoTheOwner) MissingColorCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't dye your hair without it. I'm sorry...")
	return SendOk(l, c, m.String())
}

func (r MinoTheOwner) EnjoyColor(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Enjoy your new and improved hair color!")
	return SendOk(l, c, m.String())
}
