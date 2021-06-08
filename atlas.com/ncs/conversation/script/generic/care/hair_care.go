package care

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"math"
	"math/rand"
)

func SetHair(l logrus.FieldLogger, c script.Context, choice uint32) {
	character.SetHair(l)(c.CharacterId, choice)
}

func HairStyleChoices(maleHair []uint32, femaleHair []uint32) ChoicesSupplier {
	return func(l logrus.FieldLogger, c script.Context) []uint32 {
		hair := make([]uint32, 0)
		gender := character.GetGender(l)(c.CharacterId)
		if gender == character.GenderMale {
			hair = maleHair
		} else if gender == character.GenderFemale {
			hair = femaleHair
		}
		hair = ApplyCurrentColor(l)(c.CharacterId, hair)
		hair = FilterCurrent(l)(c.CharacterId, hair)
		return hair
	}
}

func HairColorChoices(l logrus.FieldLogger, c script.Context) []uint32 {
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

func GetRandomHair(l logrus.FieldLogger, c script.Context, maleHair []uint32, femaleHair []uint32) ChoiceSupplier {
	return func() uint32 {
		hair := make([]uint32, 0)
		gender := character.GetGender(l)(c.CharacterId)
		if gender == character.GenderMale {
			hair = maleHair
		} else if gender == character.GenderFemale {
			hair = femaleHair
		}
		hair = ApplyCurrentColor(l)(c.CharacterId, hair)
		return hair[rand.Intn(len(hair))]
	}
}

func WarnRandomStyle(prompt string, coupon uint32, maleHair []uint32, femaleHair []uint32, choice ChoiceConsumer, no script.StateProducer) ChoiceStateProducer {
	return func(config ChoiceConfig) script.StateProducer {
		return func(l logrus.FieldLogger, c script.Context) script.State {
			randomSupplier := GetRandomHair(l, c, maleHair, femaleHair)
			couponProcessor := ProcessCoupon(coupon, choice, SetSingleUse(true))
			choiceProcessor := couponProcessor(randomSupplier)
			return script.SendYesNo(l, c, prompt, choiceProcessor(config), no)
		}
	}
}

func ApplyCurrentColor(l logrus.FieldLogger) func(characterId uint32, hair []uint32) []uint32 {
	return func(characterId uint32, hair []uint32) []uint32 {
		color := character.GetHair(l)(characterId) % 10
		results := make([]uint32, 0)
		for _, h := range hair {
			results = append(results, h+color)
		}
		return results
	}
}

func FilterCurrent(l logrus.FieldLogger) func(characterId uint32, hair []uint32) []uint32 {
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

func HairStyleCouponListText(coupon uint32) ChoiceConfigurator {
	return SetListText(message.NewBuilder().AddText("Haircut: ").ShowItemImage2(coupon).ShowItemName1(coupon).String())
}

func HairStyleEnjoy() ChoiceConfigurator {
	return SetEnjoy("Enjoy your new and improved hairstyle!")
}

func HairStyleCouponMissing() ChoiceConfigurator {
	return SetMissingCoupon("Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't give you a haircut without it. I'm sorry...")
}

func StyleListEntry(coupon uint32) string {
	return message.NewBuilder().AddText("Haircut: ").ShowItemImage2(coupon).ShowItemName1(coupon).String()
}

func StylePrompt(coupon uint32) string {
	return message.NewBuilder().
		AddText("I can totally change up your hairstyle and make it look so good. Why don't you change it up a bit? If you have ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" I'll change it for you. Choose the one to your liking~.").
		String()
}

func ColorCareChoice(coupon uint32) ChoiceConfig {
	listText := DyeListEntry(coupon)
	hairColor := ColorPrompt(coupon)
	vip := ProcessCoupon(coupon, SetHair, SetSingleUse(true))
	return ChoiceConfig{ListText: listText, NextState: ShowChoices(hairColor, HairColorChoices, vip)}
}

func DyeListEntry(coupon uint32) string {
	return message.NewBuilder().AddText("Dye your hair: ").ShowItemImage2(coupon).ShowItemName1(coupon).String()
}

func ColorPrompt(coupon uint32) string {
	return message.NewBuilder().
		AddText("I can totally change your hair color and make it look so good. Why don't you change it up a bit? With ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" I'll change it for you. Choose the one to your liking.").
		String()
}
