package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Mini is located in Ludibrium - Ludibrium Hair Salon (220000004)
type Mini struct {
}

func (r Mini) NPCId() uint32 {
	return npc.Mini
}

func (r Mini) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, c)
}

func (r Mini) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.RegularStyleHair(item.LudibriumHairStyleCouponRegular), r.ExperimentalStyleHair(item.LudibriumHairStyleCouponExperimental), care.ColorCareRandom(item.LudibriumHairColorCouponRegular, r.Initial)}
}

func (r Mini) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I'm the assistant here. Don't worry, I'm plenty good enough for this. If you have ").
		BlueText().ShowItemName1(item.LudibriumHairStyleCouponRegular).
		BlackText().AddText(", ").
		BlueText().ShowItemName1(item.LudibriumHairStyleCouponExperimental).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.LudibriumHairColorCouponRegular).
		BlackText().AddText(" by any chance, then allow me to take care of the rest, alright?").
		String()
}

func (r Mini) RegularStyleHair(coupon uint32) care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use this REGULAR coupon, your hair may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()

	maleHair := []uint32{30190, 30220, 30250, 30540, 30610, 30620, 30640, 30650, 30660, 30840, 30870, 30940, 30990}
	femaleHair := []uint32{31170, 31270, 31290, 31510, 31540, 31550, 31600, 31640, 31680, 31810, 31830, 31840, 31870}
	next := care.WarnRandomStyle(hairStyle, coupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}

func (r Mini) ExperimentalStyleHair(coupon uint32) care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use the EXP coupon your hair will change RANDOMLY with a chance to obtain a new experimental style that even you didn't think was possible. Are you going to use ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and really change your hairstyle?").
		String()

	maleHair := []uint32{30030, 30190, 30220, 30250, 30540, 30610, 30620, 30640, 30650, 30660, 30840, 30990}
	femaleHair := []uint32{31170, 31270, 31430, 31510, 31540, 31550, 31600, 31680, 31810, 31830, 31840, 31870}
	next := care.WarnRandomStyle(hairStyle, coupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}
