package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DrLenu is located in Victoria Road - Henesys Plastic Surgery (100000103)
type DrLenu struct {
}

func (r DrLenu) NPCId() uint32 {
	return npc.DrLenu
}

func (r DrLenu) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Hi, there~! I'm Dr. Lenu, in charge of the cosmetic lenses here at the Henesys Plastic Surgery Shop! With ").
		BlueText().ShowItemName1(item.HenesysCosmeticLensCouponRegular).
		BlackText().AddText(", you can let us take care of the rest and have the kind of beautiful look you've always craved~! Remember, the first thing everyone notices about you is the eyes, and we can help you find the cosmetic lens that most fits you! Now, what would you like to use?").
		String()
	return care.NewGenericCare(hello, r.ProvidedCare())(l, c)
}

func (r DrLenu) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.CosmeticRegular(), r.CosmeticVIP(), r.CosmeticOneTime()}
}

func (r DrLenu) CosmeticRegular() care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use the regular coupon, you'll be awarded a random pair of cosmetic lenses. Are you going to use a ").
		BlueText().ShowItemName1(item.HenesysCosmeticLensCouponRegular).
		BlackText().AddText(" and really make the change to your eyes?").
		String()

	next := care.WarnRandomLensColor(prompt, item.HenesysCosmeticLensCouponRegular, care.SetFace, r.Initial)
	return care.NewChoiceConfig(next, care.LensCouponListText(item.HenesysCosmeticLensCouponRegular), care.LensCouponMissing(), care.LensEnjoy())
}

func (r DrLenu) CosmeticVIP() care.ChoiceConfig {
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(item.HenesysCosmeticLensCouponVIP, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, care.LensColorChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponListText(item.HenesysCosmeticLensCouponVIP), care.LensCouponMissing(), care.LensEnjoy())
}

func (r DrLenu) CosmeticOneTime() care.ChoiceConfig {
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(item.HenesysCosmeticLensCouponVIP, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoicesWithNone(prompt, care.LensColorOneTimeChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponOneTimeListText(item.OneTimeCosmeticLensBlack), care.LensCouponMissing(), care.LensEnjoy())
}
