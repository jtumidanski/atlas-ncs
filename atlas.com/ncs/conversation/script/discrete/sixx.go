package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Sixx is located in Singapore - CBD (540000000)
type Sixx struct {
}

func (r Sixx) NPCId() uint32 {
	return npc.Sixx
}

func (r Sixx) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Sixx) Hello() string {
	return message.NewBuilder().
		AddText("Hi, there! I'm Sixx, in charge of Da Yan Jing Lens Shop here at CBD! With ").
		BlueText().ShowItemName1(item.CBDCosmeticLensCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.CBDCosmeticLensCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the kind of beautiful look you've always craved! Remember, the first thing everyone notices about you are the eyes, and we can help you find the cosmetic lens that most fits you! Now, what would you like to use?").
		String()
}

func (r Sixx) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.CosmeticRegular(item.CBDCosmeticLensCouponRegular), r.CosmeticVIP(item.CBDCosmeticLensCouponVIP), r.CosmeticOneTime()}
}

func (r Sixx) CosmeticRegular(coupon uint32) care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use the regular coupon, you'll be awarded a random pair of cosmetic lenses. Are you going to use a ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and really make the change to your eyes?").
		String()

	next := care.WarnRandomLensColor(prompt, coupon, care.SetFace, r.Initial)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r Sixx) CosmeticVIP(coupon uint32) care.ChoiceConfig {
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(coupon, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, care.LensColorChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r Sixx) CosmeticOneTime() care.ChoiceConfig {
	//TODO coupon consumption might need to be reviewed
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(item.OneTimeCosmeticLensBlack, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoicesWithNone(prompt, care.LensColorOneTimeChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponOneTimeListText(item.OneTimeCosmeticLensBlack), care.LensCouponMissing(), care.LensEnjoy())
}
