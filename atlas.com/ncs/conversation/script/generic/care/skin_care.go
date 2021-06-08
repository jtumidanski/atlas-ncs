package care

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

func SkinCareChoice(coupon uint32) ChoiceConfig {
	skinPrompt := Prompt(coupon)
	skinColors := FixedChoices([]uint32{0, 1, 2, 3, 4})
	couponProcessor := ProcessCoupon(coupon, SetSkin, SetSingleUse(true))

	return ChoiceConfig{
		ListText:      ListEntry(coupon),
		NextState:     ShowChoices(skinPrompt, skinColors, couponProcessor),
		MissingCoupon: MissingSkinCoupon(),
		Enjoy:         EnjoyNewSkin(),
	}
}

func EnjoyNewSkin() string {
	return "Enjoy your new and improved skin!"
}

func MissingSkinCoupon() string {
	return "Um... you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you..."
}

func Prompt(_ uint32) string {
	return "With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!"
}

func ListEntry(coupon uint32) string {
	return message.NewBuilder().AddText("Skin Care: ").ShowItemImage2(coupon).ShowItemName1(coupon).String()
}

func SetSkin(l logrus.FieldLogger, c script.Context, choice uint32) {
	character.SetSkin(l)(c.CharacterId, byte(choice))
}
