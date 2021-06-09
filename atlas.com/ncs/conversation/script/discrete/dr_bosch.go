package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DrBosch is located in Ludibrium - Ludibrium Plastic Surgery (220000003)
type DrBosch struct {
}

func (r DrBosch) NPCId() uint32 {
	return npc.DrBosch
}

func (r DrBosch) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r DrBosch) Hello() string {
	return message.NewBuilder().
		AddText("Um... hi, I'm Dr. Bosch, and I am a cosmetic lens expert here at the Ludibrium Plastic Surgery Shop. I believe your eyes are the most important feature in your body, and with ").
		BlueText().ShowItemName1(item.LudibriumCosmeticLensCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.LudibriumCosmeticLensCouponVIP).
		BlackText().AddText(", I can prescribe the right kind of cosmetic lenses for you. Now, what would you like to use?").
		String()
}

func (r DrBosch) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.CosmeticRegular(item.LudibriumCosmeticLensCouponRegular), r.CosmeticVIP(item.LudibriumCosmeticLensCouponVIP), r.CosmeticOneTime()}
}

func (r DrBosch) CosmeticRegular(coupon uint32) care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use the regular coupon, you'll be awarded a random pair of cosmetic lenses. Are you going to use a ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and really make the change to your eyes?").
		String()

	next := care.WarnRandomLensColor(prompt, coupon, care.SetFace, r.Initial)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r DrBosch) CosmeticVIP(coupon uint32) care.ChoiceConfig {
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(coupon, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, care.LensColorChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r DrBosch) CosmeticOneTime() care.ChoiceConfig {
	//TODO coupon consumption might need to be reviewed
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(item.OneTimeCosmeticLensBlack, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoicesWithNone(prompt, care.LensColorOneTimeChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponOneTimeListText(item.OneTimeCosmeticLensBlack), care.LensCouponMissing(), care.LensEnjoy())
}