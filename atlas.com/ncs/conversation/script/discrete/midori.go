package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Midori is located in Zipangu - Hair Salon (801000001)
type Midori struct {
}

func (r Midori) NPCId() uint32 {
	return npc.Midori
}

func (r Midori) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, c)
}

func (r Midori) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.RegularStyleHair(item.ShowaHairStyleCouponRegular),
		care.ColorCareRandom(item.ShowaHairColorCouponRegular, r.Initial),
	}
}

func (r Midori) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I'm the assistant here. Don't worry, I'm plenty good enough for this. If you have ").
		BlueText().ShowItemName1(item.ShowaHairStyleCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.ShowaHairColorCouponRegular).
		BlackText().AddText(" by any chance, then allow me to take care of the rest, alright?").
		String()
}

func (r Midori) RegularStyleHair(coupon uint32) care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use this REGULAR coupon, your hair may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()

	maleHair := []uint32{30260, 30280, 30340, 30360, 30710, 30780, 30790, 30800, 30810, 30820, 30920}
	femaleHair := []uint32{31350, 31410, 31460, 31540, 31550, 31710, 31720, 31770, 31790, 31800, 31850, 34000}
	next := care.WarnRandomStyle(hairStyle, coupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}