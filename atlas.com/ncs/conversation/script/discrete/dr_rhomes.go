package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DrRhomes is located in Orbis Park - Orbis Plastic Surgery (200000201)
type DrRhomes struct {
}

func (r DrRhomes) NPCId() uint32 {
	return npc.DrRhomes
}

func (r DrRhomes) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r DrRhomes) Hello() string {
	return message.NewBuilder().
		AddText("Hello, I'm Dr. Rhomes, head of the cosmetic lens department here at the Orbis Plastic Surgery Shop.").NewLine().
		AddText("My goal here is to add personality to everyone's eyes through the wonders of cosmetic lenses, and with ").
		BlueText().ShowItemName1(item.OrbisCosmeticLensCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.OrbisCosmeticLensCouponVIP).
		BlackText().AddText(", I can do the same for you, too! Now, what would you like to use?").
		String()
}

func (r DrRhomes) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.CosmeticRegular(item.OrbisCosmeticLensCouponRegular), r.CosmeticVIP(item.OrbisCosmeticLensCouponVIP), r.CosmeticOneTime()}
}

func (r DrRhomes) CosmeticRegular(coupon uint32) care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use the regular coupon, you'll be awarded a random pair of cosmetic lenses. Are you going to use a ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and really make the change to your eyes?").
		String()

	next := care.WarnRandomLensColor(prompt, coupon, care.SetFace, r.Initial)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r DrRhomes) CosmeticVIP(coupon uint32) care.ChoiceConfig {
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(coupon, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, care.LensColorChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r DrRhomes) CosmeticOneTime() care.ChoiceConfig {
	//TODO coupon consumption might need to be reviewed
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(item.OneTimeCosmeticLensBlack, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoicesWithNone(prompt, care.LensColorOneTimeChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponOneTimeListText(item.OneTimeCosmeticLensBlack), care.LensCouponMissing(), care.LensEnjoy())
}
