package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Andre is located in Victoria Road - Kerning City Hair Salon (103000005)
type Andre struct {
}

func (r Andre) NPCId() uint32 {
	return npc.Andre
}

func (r Andre) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, c)
}

func (r Andre) Hello() string {
	return message.NewBuilder().
		AddText("I'm Andre, Don's assistant. Everyone calls me Andre, though. If you have a ").
		BlueText().ShowItemName1(item.KerningCityHairStyleCouponRegular).
		BlackText().AddText(", ").
		BlueText().ShowItemName1(item.KerningCityHairStyleCouponExperimental).
		BlackText().AddText(" or a ").
		BlueText().ShowItemName1(item.KerningCityHairColorCouponRegular).
		BlackText().AddText(", please let me change your hairdo!").
		String()
}

func (r Andre) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.RegularStyleHair(item.KerningCityHairStyleCouponRegular), r.ExperimentalStyleHair(item.KerningCityHairStyleCouponExperimental), care.ColorCareRandom(item.KerningCityHairColorCouponRegular, r.Initial)}
}

func (r Andre) RegularStyleHair(coupon uint32) care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use this REGULAR coupon, your hair may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()

	maleHair := []uint32{30040, 30130, 30520, 30770, 30780, 30850, 30920, 33040}
	femaleHair := []uint32{31060, 31140, 31330, 31440, 31520, 31750, 31760, 31880, 34050}
	next := care.WarnRandomStyle(hairStyle, coupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}

func (r Andre) ExperimentalStyleHair(coupon uint32) care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use the EXP coupon your hair will change RANDOMLY with a chance to obtain a new experimental style that even you didn't think was possible. Are you going to use ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and really change your hairstyle?").
		String()

	maleHair := []uint32{30130, 30430, 30520, 30770, 30780, 30850, 30920, 33040}
	femaleHair := []uint32{31060, 31140, 31330, 31520, 31760, 31880, 34010, 34050}
	next := care.WarnRandomStyle(hairStyle, coupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}
