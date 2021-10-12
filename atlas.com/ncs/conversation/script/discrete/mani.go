package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Mani is located in New Leaf City Town Street - NLC Mall (600000001)
type Mani struct {
}

func (r Mani) NPCId() uint32 {
	return npc.Mani
}

func (r Mani) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, span, c)
}

func (r Mani) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.StyleHair(item.NLCHairStyleCouponVIP, item.NLCHairMembershipCoupon),
		care.ColorCareChoice(item.NLCHairColorCouponVIP),
	}
}

func (r Mani) Hello() string {
	return message.NewBuilder().
		AddText("I'm the head of this hair salon Mani. If you have a ").
		BlueText().ShowItemName1(item.NLCHairStyleCouponVIP).
		BlackText().AddText(" or a ").
		BlueText().ShowItemName1(item.NLCHairColorCouponVIP).
		BlackText().AddText(", allow me to take care of your hairdo. Please choose the one you want.").
		String()
}

func (r Mani) StyleHair(coupon uint32, membershipCoupon uint32) care.ChoiceConfig {
	maleHair := []uint32{30250, 30490, 30730, 30870, 30880, 33100}
	femaleHair := []uint32{31320, 31450, 31560, 31730, 31830}
	return care.VIPHairCareWithMembership(coupon, membershipCoupon, maleHair, femaleHair)
}
