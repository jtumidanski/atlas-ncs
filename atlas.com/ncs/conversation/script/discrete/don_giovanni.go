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

// DonGiovanni is located in Victoria Road - Kerning City Hair Salon (103000005)
type DonGiovanni struct {
}

func (r DonGiovanni) NPCId() uint32 {
	return npc.DonGiovanni
}

func (r DonGiovanni) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, span, c)
}

func (r DonGiovanni) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.Style(item.KerningCityHairStyleCouponVIP, item.KerningCityHairMembership),
		care.ColorCareChoice(item.KerningCityHairColorCouponVIP),
	}
}

func (r DonGiovanni) Hello() string {
	return message.NewBuilder().
		AddText("Hello! I'm Don Giovanni, head of the beauty salon! If you have either ").
		BlueText().ShowItemName1(item.KerningCityHairStyleCouponVIP).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.KerningCityHairColorCouponVIP).
		BlackText().AddText(", why don't you let me take care of the rest? Decide what you want to do with your hair...").
		String()
}

func (r DonGiovanni) Style(coupon uint32, membershipCoupon uint32) care.ChoiceConfig {
	maleHair := []uint32{30040, 30130, 30780, 30850, 30860, 30920, 33040}
	femaleHair := []uint32{31090, 31140, 31330, 31440, 31760, 31880, 34050}
	return care.VIPHairCareWithMembership(coupon, membershipCoupon, maleHair, femaleHair)
}
