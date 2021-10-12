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

// Eric is located in Singapore - CBD (540000000)
type Eric struct {
}

func (r Eric) NPCId() uint32 {
	return npc.Eric
}

func (r Eric) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), []care.ChoiceConfig{r.StyleHair(item.CBDHairStyleCouponVIP), care.ColorCareChoice(item.CBDHairColorCouponVIP)})(l, span, c)
}

func (r Eric) Hello() string {
	return message.NewBuilder().
		AddText("Welcome to the Quick-Hand Hair-Salon!. Do you, by any chance, have ").
		BlueText().ShowItemName1(item.CBDHairStyleCouponVIP).
		BlackText().AddText(" or a ").
		BlueText().ShowItemName1(item.CBDHairColorCouponVIP).
		BlackText().AddText("? If so, how about letting me take care of your hair? Please what you want to do with it.").
		String()
}

func (r Eric) StyleHair(coupon uint32) care.ChoiceConfig {
	hairStyle := care.StylePrompt(coupon)
	maleHair := []uint32{30000, 30020, 30110, 30120, 30270, 30290, 30310, 30670, 30840}
	femaleHair := []uint32{31010, 31050, 31110, 31120, 31240, 31250, 31280, 31670, 31810}

	vip := care.ProcessCoupon(coupon, care.SetHair, care.SetSingleUse(true))
	next := care.ShowChoices(hairStyle, care.HairStyleChoices(maleHair, femaleHair), vip)
	return care.NewChoiceConfig(next, care.SetListText(care.StyleListEntry(coupon)))
}
