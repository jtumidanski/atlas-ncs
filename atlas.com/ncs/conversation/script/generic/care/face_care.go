package care

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"math/rand"
)

func SetFace(l logrus.FieldLogger, c script.Context, choice uint32) {
	character.SetFace(l)(c.CharacterId, choice)
}

func FaceCouponListText(coupon uint32) ChoiceConfigurator {
	return SetListText(message.NewBuilder().AddText("Plastic Surgery: ").ShowItemImage2(coupon).ShowItemName1(coupon).String())
}

func LensCouponListText(coupon uint32) ChoiceConfigurator {
	return SetListText(message.NewBuilder().AddText("Cosmetic Lenses: ").ShowItemImage2(coupon).ShowItemName1(coupon).String())
}

func LensCouponOneTimeListText(coupon uint32) ChoiceConfigurator {
	return SetListText(message.NewBuilder().AddText("One-Time Cosmetic Lenses: ").ShowItemImage2(coupon).AddText(" (anycolor)").String())
}

func FaceEnjoy() ChoiceConfigurator {
	return SetEnjoy("Enjoy your new and improved face!")
}

func LensEnjoy() ChoiceConfigurator {
	return SetEnjoy("Enjoy your new and improved cosmetic lenses!")
}

func FaceCouponMissing() ChoiceConfigurator {
	return SetMissingCoupon("Hmm ... it looks like you don't have the coupon specifically for this place. Sorry to say this, but without the coupon, there's no plastic surgery for you...")
}

func LensCouponMissing() ChoiceConfigurator {
	return SetMissingCoupon("I'm sorry, but I don't think you have our cosmetic lens coupon with you right now. Without the coupon, I'm afraid I can't do it for you..")
}

func WarnRandomFace(prompt string, coupon uint32, male []uint32, female []uint32, choice ChoiceConsumer, no script.StateProducer) ChoiceStateProducer {
	return func(config ChoiceConfig) script.StateProducer {
		return func(l logrus.FieldLogger, c script.Context) script.State {
			randomSupplier := GetRandomFace(l, c, male, female)
			couponProcessor := ProcessCoupon(coupon, choice, SetSingleUse(true))
			choiceProcessor := couponProcessor(randomSupplier)
			return script.SendYesNo(l, c, prompt, choiceProcessor(config), no)
		}
	}
}

func GetRandomFace(l logrus.FieldLogger, c script.Context, male []uint32, female []uint32) ChoiceSupplier {
	return func() uint32 {
		choices := make([]uint32, 0)
		gender := character.GetGender(l)(c.CharacterId)
		if gender == character.GenderMale {
			choices = male
		} else if gender == character.GenderFemale {
			choices = female
		}
		choices = ApplyEyeColor(l)(c.CharacterId, choices)
		return choices[rand.Intn(len(choices))]
	}
}

func ApplyEyeColor(l logrus.FieldLogger) func(characterId uint32, choices []uint32) []uint32 {
	return func(characterId uint32, choices []uint32) []uint32 {
		//TODO need to verify color combination exists
		current := character.GetFace(l)(characterId)
		color := (current % 1000) - (current % 100)
		results := make([]uint32, 0)
		for _, h := range choices {
			results = append(results, h+color)
		}
		return results
	}
}

func WarnRandomLensColor(prompt string, coupon uint32, choice ChoiceConsumer, no script.StateProducer) ChoiceStateProducer {
	return func(config ChoiceConfig) script.StateProducer {
		return func(l logrus.FieldLogger, c script.Context) script.State {
			randomSupplier := GetRandomLensColor(l, c)
			couponProcessor := ProcessCoupon(coupon, choice, SetSingleUse(true))
			choiceProcessor := couponProcessor(randomSupplier)
			return script.SendYesNo(l, c, prompt, choiceProcessor(config), no)
		}
	}
}

func GetRandomLensColor(l logrus.FieldLogger, c script.Context) ChoiceSupplier {
	return func() uint32 {
		hair := LensColorChoices(l, c)
		return hair[rand.Intn(len(hair))]
	}
}

func LensColorChoices(l logrus.FieldLogger, c script.Context) []uint32 {
	var current uint32
	gender := character.GetGender(l)(c.CharacterId)

	if gender == character.GenderMale {
		current = character.GetFace(l)(c.CharacterId)%100 + 20000
	} else if gender == character.GenderFemale {
		current = character.GetFace(l)(c.CharacterId)%100 + 21000
	}
	return []uint32{current, current + 100, current + 200, current + 400, current + 600, current + 700}
}

func LensColorOneTimeChoices(l logrus.FieldLogger, c script.Context) []uint32 {
	var current uint32
	gender := character.GetGender(l)(c.CharacterId)

	if gender == character.GenderMale {
		current = character.GetFace(l)(c.CharacterId)%100 + 20000
	} else if gender == character.GenderFemale {
		current = character.GetFace(l)(c.CharacterId)%100 + 21000
	}

	colors := make([]uint32, 0)
	for i := uint32(0); i < 8; i++ {
		if character.HasItem(l)(c.CharacterId, item.OneTimeCosmeticLensBlack+i) {
			colors = append(colors, current+100*i)
		}
	}
	return colors
}
