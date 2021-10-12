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

// Ari is located in New Leaf City Town Street - NLC Mall (600000001)
type Ari struct {
}

func (r Ari) NPCId() uint32 {
	return npc.Ari
}

func (r Ari) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, span, c)
}

func (r Ari) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.ExperimentalStyleHair(item.NLCHairStyleCouponExperimental),
		care.ColorCareRandom(item.NLCHairColorCouponRegular, r.Initial),
	}
}

func (r Ari) Hello() string {
	return message.NewBuilder().
		AddText("I'm Ari the assistant. If you have ").
		BlueText().ShowItemName1(item.NLCHairStyleCouponExperimental).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.NLCHairColorCouponRegular).
		BlackText().AddText(" by any chance, then how about letting me change your hairdo?").
		String()
}

func (r Ari) ExperimentalStyleHair(coupon uint32) care.ChoiceConfig {
	maleHair := []uint32{30250, 30400, 30430, 30440, 30490, 30730, 30830, 30870, 30880, 33100}
	femaleHair := []uint32{31320, 31450, 31560, 31570, 31690, 31720, 31730, 31830, 34010}
	return care.ExperimentalHairCare(coupon, maleHair, femaleHair, r.Initial)
}