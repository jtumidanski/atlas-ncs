package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"math"
)

// Natalie is located in Victoria Road - Henesys Hair Salon (100000104)
type Natalie struct {
}

func (r Natalie) NPCId() uint32 {
	return npc.Natalie
}

func (r Natalie) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Natalie) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I'm the head of this hair salon. If you have a ").
		BlueText().ShowItemName1(item.HenesysHairStyleCouponVIP).
		BlackText().AddText(" or a ").
		BlueText().ShowItemName1(item.HenesysHairColorCouponVIP).
		BlackText().AddText(" allow me to take care of your hairdo. Please choose the one you want.").AddNewLine().
		OpenItem(0).AddText("Haircut: ").ShowItemImage2(item.HenesysHairStyleCouponVIP).ShowItemName1(item.HenesysHairStyleCouponVIP).CloseItem().AddNewLine().
		OpenItem(1).AddText("Dye your hair: ").ShowItemImage2(item.HenesysHairColorCouponVIP).ShowItemName1(item.HenesysHairColorCouponVIP).CloseItem().AddNewLine()
	return SendListSelection(l, c, m.String(), r.Choose)
}

func (r Natalie) Choose(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.ICanTotallyStyle
	case 1:
		return r.ICanTotallyColor
	}
	return nil
}

func (r Natalie) ICanTotallyStyle(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I can totally change up your hairstyle and make it look so good. Why don't you change it up a bit? If you have ").
		BlueText().ShowItemName1(item.HenesysHairStyleCouponVIP).
		BlackText().AddText(" I'll change it for you. Choose the one to your liking~.")

	hair := make([]uint32, 0)
	gender := character.GetGender(l)(c.CharacterId)
	if gender == character.GenderMale {
		hair = r.BaseMaleHair()
	} else if gender == character.GenderFemale {
		hair = r.BaseFemaleHair()
	}
	hair = ApplyCharacterColor(l)(c.CharacterId, hair)
	hair = FilterCurrentHair(l)(c.CharacterId, hair)

	return SendStyle(l, c, m.String(), r.ProcessHairChange, hair)
}

func (r Natalie) BaseMaleHair() []uint32 {
	return []uint32{character.HairBlackCatalyst, character.HairBlackTopknot, character.HairBlackWind, character.HairBlackShaggyWax, character.HairBlackAcorn, character.HairBlackAranCut, character.HairBlackTheCoco}
}

func (r Natalie) BaseFemaleHair() []uint32 {
	return []uint32{character.HairBlackAngelica, character.HairBlackChantelle, character.HairBlackFourtailBraids, character.HairBlackCrazyMedusa, character.HairBlackFrizzleDizzle, character.HairBlackAranHair, character.HairBlackFullBangs}
}

func ApplyCharacterColor(l logrus.FieldLogger) func(characterId uint32, hair []uint32) []uint32 {
	return func(characterId uint32, hair []uint32) []uint32 {
		color := character.GetHair(l)(characterId) % 10
		results := make([]uint32, 0)
		for _, h := range hair {
			results = append(results, h+color)
		}
		return results
	}
}

func FilterCurrentHair(l logrus.FieldLogger) func(characterId uint32, hair []uint32) []uint32 {
	return func(characterId uint32, hair []uint32) []uint32 {
		current := character.GetHair(l)(characterId)
		results := make([]uint32, 0)
		for _, h := range hair {
			if h != current {
				results = append(results, h)
			}
		}
		return results
	}
}

func (r Natalie) ICanTotallyColor(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I can totally change your hair color and make it look so good. Why don't you change it up a bit? With ").
		BlueText().ShowItemName1(item.HenesysHairColorCouponVIP).
		BlackText().AddText(" I'll change it for you. Choose the one to your liking.")

	hair := ProduceColorOptions(l, c)

	return SendStyle(l, c, m.String(), r.ProcessColorChange, hair)
}

func ProduceColorOptions(l logrus.FieldLogger, c Context) []uint32 {
	hair := make([]uint32, 0)
	currentHair := character.GetHair(l)(c.CharacterId)
	baseStyle := uint32(math.Floor(float64(currentHair/10)) * 10)
	for i := uint32(0); i < 8; i++ {
		newColor := baseStyle + i
		if newColor != currentHair {
			hair = append(hair, newColor)
		}
	}
	return hair
}

func (r Natalie) ProcessHairChange(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if character.HasItem(l)(c.CharacterId, item.HenesysHairMembershipCoupon) {
			character.SetHair(l)(c.CharacterId, uint32(selection))
			return r.EnjoyHair(l, c)
		} else if character.HasItem(l)(c.CharacterId, item.HenesysHairStyleCouponVIP) {
			character.GainItem(l)(c.CharacterId, item.HenesysHairStyleCouponVIP, -1)
			character.SetHair(l)(c.CharacterId, uint32(selection))
			return r.EnjoyHair(l, c)
		}
		return r.MissingStyleCoupon(l, c)
	}
}

func (r Natalie) EnjoyHair(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Enjoy your new and improved hairstyle!")
	return SendOk(l, c, m.String())
}

func (r Natalie) MissingStyleCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't give you a haircut without it. I'm sorry...")
	return SendOk(l, c, m.String())
}

func (r Natalie) ProcessColorChange(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if character.HasItem(l)(c.CharacterId, item.HenesysHairColorCouponVIP) {
			character.GainItem(l)(c.CharacterId, item.HenesysHairColorCouponVIP, -1)
			character.SetHair(l)(c.CharacterId, uint32(selection))
			return r.EnjoyColor(l, c)
		}
		return r.MissingColorCoupon(l, c)
	}
}

func (r Natalie) EnjoyColor(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Enjoy your new and improved hair color!")
	return SendOk(l, c, m.String())
}

func (r Natalie) MissingColorCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't dye your hair without it. I'm sorry...")
	return SendOk(l, c, m.String())
}
