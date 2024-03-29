package care

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math"
	"math/rand"
)

func SetHair(l logrus.FieldLogger, c script.Context, choice uint32) {
	character.SetHair(l)(c.CharacterId, choice)
}

func HairStyleChoices(maleHair []uint32, femaleHair []uint32) ChoicesSupplier {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) []uint32 {
		hair := make([]uint32, 0)
		gender := character.GetGender(l, span)(c.CharacterId)
		if gender == character.GenderMale {
			hair = maleHair
		} else if gender == character.GenderFemale {
			hair = femaleHair
		}
		hair = ApplyCurrentHairColor(l, span)(c.CharacterId, hair)
		hair = FilterCurrentHair(l, span)(c.CharacterId, hair)
		return hair
	}
}

func HairColorChoices(l logrus.FieldLogger, span opentracing.Span, c script.Context) []uint32 {
	hair := make([]uint32, 0)
	currentHair := character.GetHair(l, span)(c.CharacterId)
	baseStyle := uint32(math.Floor(float64(currentHair/10)) * 10)
	for i := uint32(0); i < 8; i++ {
		newColor := baseStyle + i
		if newColor != currentHair {
			hair = append(hair, newColor)
		}
	}
	return hair
}

func GetRandomHair(l logrus.FieldLogger, span opentracing.Span, c script.Context, maleHair []uint32, femaleHair []uint32) ChoiceSupplier {
	return func() uint32 {
		hair := make([]uint32, 0)
		gender := character.GetGender(l, span)(c.CharacterId)
		if gender == character.GenderMale {
			hair = maleHair
		} else if gender == character.GenderFemale {
			hair = femaleHair
		}
		hair = ApplyCurrentHairColor(l, span)(c.CharacterId, hair)
		return hair[rand.Intn(len(hair))]
	}
}

func WarnRandomStyle(prompt string, coupon uint32, maleHair []uint32, femaleHair []uint32, choice ChoiceConsumer, no script.StateProducer) ChoiceStateProducer {
	return func(config ChoiceConfig) script.StateProducer {
		return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
			randomSupplier := GetRandomHair(l, span, c, maleHair, femaleHair)
			couponProcessor := ProcessCoupon(coupon, choice, SetSingleUse(true))
			choiceProcessor := couponProcessor(randomSupplier)
			return script.SendYesNo(l, span, c, prompt, choiceProcessor(config), no)
		}
	}
}

func WarnRandomColor(prompt string, choice ChoiceHandlerProducer, no script.StateProducer) ChoiceStateProducer {
	return func(config ChoiceConfig) script.StateProducer {
		return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
			randomSupplier := GetRandomHairColor(l, span, c)
			choiceProcessor := choice(randomSupplier)
			return script.SendYesNo(l, span, c, prompt, choiceProcessor(config), no)
		}
	}
}

func GetRandomHairColor(l logrus.FieldLogger, span opentracing.Span, c script.Context) ChoiceSupplier {
	return func() uint32 {
		hair := HairColorChoices(l, span, c)
		return hair[rand.Intn(len(hair))]
	}
}

func ApplyCurrentHairColor(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, hair []uint32) []uint32 {
	return func(characterId uint32, hair []uint32) []uint32 {
		//TODO need to verify color combination exists
		color := character.GetHair(l, span)(characterId) % 10
		results := make([]uint32, 0)
		for _, h := range hair {
			results = append(results, h+color)
		}
		return results
	}
}

func FilterCurrentHair(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, hair []uint32) []uint32 {
	return func(characterId uint32, hair []uint32) []uint32 {
		current := character.GetHair(l, span)(characterId)
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

func ColorRandomPrompt(coupon uint32) string {
	return message.NewBuilder().
		AddText("If you use a regular coupon your hair will change RANDOMLY. Do you still want to use ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and change it up?").
		String()
}

func HairColorEnjoy() ChoiceConfigurator {
	return SetEnjoy("Enjoy your new and improved haircolor!")
}

func HairColorCouponMissing() ChoiceConfigurator {
	return SetMissingCoupon("Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't dye your hair without it. I'm sorry...")
}

func ColorCareRandom(coupon uint32, no script.StateProducer) ChoiceConfig {
	hairColor := ColorRandomPrompt(coupon)
	processCoupon := ProcessCoupon(coupon, SetHair, SetSingleUse(true))
	return NewChoiceConfig(WarnRandomColor(hairColor, processCoupon, no), SetListText(DyeListEntry(coupon)), HairColorCouponMissing(), HairColorEnjoy())
}

func ColorCareChoice(coupon uint32) ChoiceConfig {
	hairColor := ColorPrompt(coupon)
	vip := ProcessCoupon(coupon, SetHair, SetSingleUse(true))
	return NewChoiceConfig(ShowChoices(hairColor, HairColorChoices, vip), SetListText(DyeListEntry(coupon)), HairColorCouponMissing(), HairColorEnjoy())
}

func RegularHairCare(coupon uint32, maleHair []uint32, femaleHair []uint32, no script.StateProducer) ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use this REGULAR coupon, your hair may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()

	next := WarnRandomStyle(prompt, coupon, maleHair, femaleHair, SetHair, no)
	return NewChoiceConfig(next, HairStyleCouponListText(coupon), HairStyleCouponMissing(), HairStyleEnjoy())
}

func ExperimentalHairCare(coupon uint32, maleHair []uint32, femaleHair []uint32, no script.StateProducer) ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use the EXP coupon your hair will change RANDOMLY with a chance to obtain a new experimental style that even you didn't think was possible. Are you going to use ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and really change your hairstyle?").
		String()

	next := WarnRandomStyle(prompt, coupon, maleHair, femaleHair, SetHair, no)
	return NewChoiceConfig(next, HairStyleCouponListText(coupon), HairStyleCouponMissing(), HairStyleEnjoy())
}

func VIPHairCareWithMembership(coupon uint32, membershipCoupon uint32, maleHair []uint32, femaleHair []uint32) ChoiceConfig {
	choiceSupplier := HairStyleChoices(maleHair, femaleHair)

	vip := ProcessCoupon(coupon, SetHair, SetSingleUse(true))
	membership := ProcessCoupon(membershipCoupon, SetHair, SetSingleUse(false), SetFailFunction(vip))

	hairStyle := StylePrompt(coupon)
	next := ShowChoices(hairStyle, choiceSupplier, membership)
	return NewChoiceConfig(next, HairStyleCouponListText(coupon), HairStyleCouponMissing(), HairStyleEnjoy())
}
