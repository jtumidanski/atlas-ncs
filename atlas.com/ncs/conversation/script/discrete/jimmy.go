package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Jimmy is located in Singapore - CBD (540000000)
type Jimmy struct {
}

func (r Jimmy) NPCId() uint32 {
	return npc.Jimmy
}

func (r Jimmy) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, c)
}

func (r Jimmy) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.RegularStyleHair(item.CBDHairStyleCouponRegular),
		care.ColorCareRandom(item.CBDHairColorCouponRegular, r.Initial),
	}
}

func (r Jimmy) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I'm the assistant here. Dont worry, I'm plenty good enough for this. If you have ").
		BlueText().ShowItemName1(item.CBDHairStyleCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.CBDHairColorCouponRegular).
		BlackText().AddText(" by any chance, then allow me to take care of the rest?").
		String()
}

func (r Jimmy) RegularStyleHair(coupon uint32) care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use this REGULAR coupon, your hair may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()

	maleHair := []uint32{30110, 30180, 30260, 30290, 30300, 30350, 30470, 30720, 30840}
	femaleHair := []uint32{31110, 31200, 31250, 31280, 31600, 31640, 31670, 31810, 34020}
	next := care.WarnRandomStyle(hairStyle, coupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}
