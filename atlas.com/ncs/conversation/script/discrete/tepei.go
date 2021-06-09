package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Tepei is located in Zipangu - Hair Salon (801000001)
type Tepei struct {
}

func (r Tepei) NPCId() uint32 {
	return npc.Tepei
}

func (r Tepei) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), []care.ChoiceConfig{r.StyleHair(item.ShowaHairColorCouponVIP),
		care.ColorCareChoice(item.OrbisHairColorCouponVIP)})(l, c)
}

func (r Tepei) Hello() string {
	return message.NewBuilder().
		AddText("Welcome to the Showa hair shop. If you have a ").
		BlueText().ShowItemName1(item.ShowaHairStyleCouponVIP).
		BlackText().AddText(", or a ").
		BlueText().ShowItemName1(item.ShowaHairColorCouponVIP).
		BlackText().AddText(", allow me to take care of your hairdo. Please choose the one you want.").
		String()
}

func (r Tepei) StyleHair(coupon uint32) care.ChoiceConfig {
	maleHair := []uint32{30260, 30280, 30340, 30710, 30780, 30800, 30810, 30820, 30920}
	femaleHair := []uint32{31000, 31030, 31100, 31350, 31460, 31550, 31770, 31790, 31850}
	choiceSupplier := care.HairStyleChoices(maleHair, femaleHair)

	vip := care.ProcessCoupon(coupon, care.SetHair, care.SetSingleUse(true))
	hairStyle := care.StylePrompt(coupon)
	next := care.ShowChoices(hairStyle, choiceSupplier, vip)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}
