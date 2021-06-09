package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// RinzTheAssistant is located in Orbis Park - Orbis Hair Salon (200000202)
type RinzTheAssistant struct {
}

func (r RinzTheAssistant) NPCId() uint32 {
	return npc.RinzTheAssistant
}

func (r RinzTheAssistant) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, c)
}

func (r RinzTheAssistant) Hello() string {
	return message.NewBuilder().
		AddText("I'm Rinz, the assistant. Do you have ").
		BlueText().ShowItemName1(item.OrbisDirtyHairCoupon).
		BlackText().AddText(", ").
		BlueText().ShowItemName1(item.OrbisHairStyleCouponRegular).
		BlackText().AddText(", ").
		BlueText().ShowItemName1(item.OrbisHairStyleCouponExperimental).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.OrbisHairColorCouponRegular).
		BlackText().AddText(" with you? If so, what do you think about letting me take care of your hairdo? What do you want to do with your hair?").
		String()
}

func (r RinzTheAssistant) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.DirtyHair(), r.RegularStyleHair(), r.ExperimentalStyleHair(), care.ColorCareRandom(item.OrbisHairColorCouponRegular, r.Initial)}
}

func (r RinzTheAssistant) DirtyHair() care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use the DRT coupon your hair will change RANDOMLY with a chance to obtain the basic styles that I came up with. Are you going to use ").
		BlueText().ShowItemName1(item.OrbisDirtyHairCoupon).
		BlackText().AddText(" and really change your hairstyle?").
		String()

	maleHair := []uint32{30030, 30020, 30000, 30270, 30230}
	femaleHair := []uint32{31040, 31000, 31250, 31220, 31260}
	next := care.WarnRandomStyle(hairStyle, item.OrbisDirtyHairCoupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(item.OrbisDirtyHairCoupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}

func (r RinzTheAssistant) RegularStyleHair() care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use this REGULAR coupon, your hair may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(item.OrbisHairStyleCouponRegular).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()

	maleHair := []uint32{30230, 30260, 30280, 30340, 30490, 30530, 30630, 30740}
	femaleHair := []uint32{31110, 31220, 31230, 31630, 31650, 31710, 31790, 31890, 31930}
	next := care.WarnRandomStyle(hairStyle, item.OrbisHairStyleCouponRegular, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(item.OrbisHairStyleCouponRegular), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}

func (r RinzTheAssistant) ExperimentalStyleHair() care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use the EXP coupon your hair will change RANDOMLY with a chance to obtain a new experimental style that even you didn't think was possible. Are you going to use ").
		BlueText().ShowItemName1(item.OrbisHairStyleCouponExperimental).
		BlackText().AddText(" and really change your hairstyle?").
		String()

	maleHair := []uint32{30230, 30280, 30340, 30490, 30530, 30740}
	femaleHair := []uint32{31110, 31220, 31230, 31710, 31790, 31890, 31930}
	next := care.WarnRandomStyle(hairStyle, item.OrbisHairStyleCouponExperimental, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(item.OrbisHairStyleCouponExperimental), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}

