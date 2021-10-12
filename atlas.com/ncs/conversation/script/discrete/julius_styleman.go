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

// JuliusStyleman is located in Amoria - Amoria Hair Salon (680000002)
type JuliusStyleman struct {
}

func (r JuliusStyleman) NPCId() uint32 {
	return npc.JuliusStyleman
}

func (r JuliusStyleman) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, span, c)
}

func (r JuliusStyleman) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.Style(item.AmoriaHairStyleCouponVIP, item.AmoriaHairMembershipCoupon),
		care.ColorCareChoice(item.AmoriaHairColorCouponVIP),
	}
}

func (r JuliusStyleman) Hello() string {
	return message.NewBuilder().
		AddText("Welcome to the Amoria hair shop. If you have a ").
		BlueText().ShowItemName1(item.AmoriaHairStyleCouponVIP).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.AmoriaHairColorCouponVIP).
		BlackText().AddText(", allow me to take care of your hairdo. Please choose the one you want.").
		String()
}

func (r JuliusStyleman) Style(coupon uint32, membershipCoupon uint32) care.ChoiceConfig {
	maleHair := []uint32{30050, 30300, 30410, 30450, 30510, 30570, 30580, 30590, 30660, 30910}
	femaleHair := []uint32{31150, 31220, 31260, 31310, 31420, 31480, 31490, 31580, 31590, 31610, 31630}
	return care.VIPHairCareWithMembership(coupon, membershipCoupon, maleHair, femaleHair)
}
