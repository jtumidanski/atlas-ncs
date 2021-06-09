package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// GrandpaLuo is located in Mu Lung - Mu Lung Hair Salon (250000003)
type GrandpaLuo struct {
}

func (r GrandpaLuo) NPCId() uint32 {
	return npc.GrandpaLuo
}

func (r GrandpaLuo) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), []care.ChoiceConfig{r.StyleHair(item.MuLungHairStyleCouponVIP, item.MuLungHairMembershipCoupon),
		care.ColorCareChoice(item.MuLungHairColorCouponVIP)})(l, c)
}

func (r GrandpaLuo) Hello() string {
	return message.NewBuilder().
		AddText("Welcome to the Mu Lung hair shop. If you have a ").
		BlueText().ShowItemName1(item.MuLungHairStyleCouponVIP).
		BlackText().AddText(", or a ").
		BlueText().ShowItemName1(item.MuLungHairColorCouponVIP).
		BlackText().AddText(", allow me to take care of your hairdo. Please choose the one you want.").
		String()
}

func (r GrandpaLuo) StyleHair(coupon uint32, membershipCoupon uint32) care.ChoiceConfig {
	maleHair := []uint32{30150, 30240, 30370, 30420, 30640, 30710, 30750, 30810}
	femaleHair := []uint32{31140, 31160, 31180, 31300, 31460, 31470, 31660, 31910}
	choiceSupplier := care.HairStyleChoices(maleHair, femaleHair)

	vip := care.ProcessCoupon(coupon, care.SetHair, care.SetSingleUse(true))
	membership := care.ProcessCoupon(membershipCoupon, care.SetHair, care.SetSingleUse(false), care.SetFailFunction(vip))

	hairStyle := care.StylePrompt(coupon)
	next := care.ShowChoices(hairStyle, choiceSupplier, membership)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}
