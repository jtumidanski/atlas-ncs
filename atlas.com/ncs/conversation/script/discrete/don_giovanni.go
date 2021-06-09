package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DonGiovanni is located in Victoria Road - Kerning City Hair Salon (103000005)
type DonGiovanni struct {
}

func (r DonGiovanni) NPCId() uint32 {
	return npc.DonGiovanni
}

func (r DonGiovanni) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), []care.ChoiceConfig{
		r.Style(item.KerningCityHairStyleCouponVIP, item.KerningCityHairMembership),
		care.ColorCareChoice(item.KerningCityHairColorCouponVIP),
	})(l, c)
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
	choiceSupplier := care.HairStyleChoices(maleHair, femaleHair)

	vip := care.ProcessCoupon(coupon, care.SetHair, care.SetSingleUse(true))
	membership := care.ProcessCoupon(membershipCoupon, care.SetHair, care.SetSingleUse(false), care.SetFailFunction(vip))

	prompt := care.StylePrompt(coupon)
	next := care.ShowChoices(prompt, choiceSupplier, membership)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}
