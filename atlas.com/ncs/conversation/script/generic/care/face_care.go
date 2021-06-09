package care

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
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

func FaceEnjoy() ChoiceConfigurator {
	return SetEnjoy("Enjoy your new and improved face!")
}

func FaceCouponMissing() ChoiceConfigurator {
	return SetMissingCoupon("Hmm ... it looks like you don't have the coupon specifically for this place. Sorry to say this, but without the coupon, there's no plastic surgery for you...")
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