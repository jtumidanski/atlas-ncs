package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SalonSeamus is located in Amoria - Amoria Hair Salon (680000002)
type SalonSeamus struct {
}

func (r SalonSeamus) NPCId() uint32 {
	return npc.SalonSeamus
}

func (r SalonSeamus) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, c)
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
	return []care.ChoiceConfig{r.ExperimentalStyleHair(item.AmoriaHairStyleCouponExperimental), care.ColorCareRandom(item.AmoriaHairColorCouponRegular, r.Initial)}
}

func (r SalonSeamus) ExperimentalStyleHair(coupon uint32) care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use the EXP coupon your hair will change RANDOMLY with a chance to obtain a new experimental style that even you didn't think was possible. Are you going to use ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and really change your hairstyle?").
		String()

	maleHair := []uint32{30000, 30020, 30110, 30130, 30160, 30190, 30240, 30270, 30430}
	femaleHair := []uint32{31000, 31030, 31050, 31070, 31090, 31150, 31310, 31910, 34010}
	next := care.WarnRandomStyle(hairStyle, coupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}