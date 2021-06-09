package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Lilishu is located in Mu Lung - Mu Lung Hair Salon (250000003)
type Lilishu struct {
}

func (r Lilishu) NPCId() uint32 {
	return npc.Lilishu
}

func (r Lilishu) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, c)
}

func (r Lilishu) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.ExperimentalStyleHair(item.MuLungHairStyleCouponExperimental),
		care.ColorCareRandom(item.MuLungHairColorCouponRegular, r.Initial),
	}
}

func (r Lilishu) Hello() string {
	return message.NewBuilder().
		AddText("I'm a hair assistant in this shop. If you have ").
		BlueText().ShowItemName1(item.MuLungHairStyleCouponExperimental).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.MuLungHairColorCouponRegular).
		BlackText().AddText(" by any chance, then how about letting me change your hairdo?").
		String()
}

func (r Lilishu) ExperimentalStyleHair(coupon uint32) care.ChoiceConfig {
	maleHair := []uint32{30030, 30150, 30240, 30370, 30420, 30550, 30600, 30640, 30700, 30710, 30720, 30750, 30810, 30830}
	femaleHair := []uint32{31140, 31160, 31180, 31210, 31300, 31430, 31460, 31470, 31660, 31690, 31800, 31890, 31910, 31940}
	return care.ExperimentalHairCare(coupon, maleHair, femaleHair, r.Initial)
}
