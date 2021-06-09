package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Shati is located in The Burning Road - Ariant (260000000)
type Shati struct {
}

func (r Shati) NPCId() uint32 {
	return npc.Shati
}

func (r Shati) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, c)
}

func (r Shati) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.RegularStyleHair(item.AriantHairStyleCouponRegular),
		care.ColorCareRandom(item.AriantHairColorCouponRegular, r.Initial),
	}
}

func (r Shati) Hello() string {
	return message.NewBuilder().
		AddText("Hey there! I'm Shatti, and I'm Mazra's apprentice. If you have ").
		BlueText().ShowItemName1(item.AriantHairStyleCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.AriantHairColorCouponRegular).
		BlackText().AddText(" with you, how about allowing me to work on your hair?").
		String()
}

func (r Shati) RegularStyleHair(coupon uint32) care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use this REGULAR coupon, your hair may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()

	maleHair := []uint32{30150, 30170, 30180, 30320, 30330, 30410, 30460, 30680, 30800, 30820, 30900}
	femaleHair := []uint32{31090, 31190, 31330, 31340, 31400, 31420, 31520, 31620, 31650, 31660, 34000}
	next := care.WarnRandomStyle(hairStyle, coupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(coupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}