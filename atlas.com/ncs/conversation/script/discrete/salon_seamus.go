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

// SalonSeamus is located in Amoria - Amoria Hair Salon (680000002)
type SalonSeamus struct {
}

func (r SalonSeamus) NPCId() uint32 {
	return npc.SalonSeamus
}

func (r SalonSeamus) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, span, c)
}

func (r SalonSeamus) Hello() string {
	return message.NewBuilder().
		AddText("I'm Salon Seamus. If you have ").
		BlueText().ShowItemName1(item.AmoriaHairStyleCouponExperimental).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.AmoriaHairColorCouponRegular).
		BlackText().AddText(" by any chance, then how about letting me change your hairdo?").
		String()
}

func (r SalonSeamus) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.ExperimentalStyleHair(item.AmoriaHairStyleCouponExperimental),
		care.ColorCareRandom(item.AmoriaHairColorCouponRegular, r.Initial),
	}
}

func (r SalonSeamus) ExperimentalStyleHair(coupon uint32) care.ChoiceConfig {
	maleHair := []uint32{30000, 30020, 30110, 30130, 30160, 30190, 30240, 30270, 30430}
	femaleHair := []uint32{31000, 31030, 31050, 31070, 31090, 31150, 31310, 31910, 34010}
	return care.ExperimentalHairCare(coupon, maleHair, femaleHair, r.Initial)
}
