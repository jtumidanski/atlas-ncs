package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DrRoberts is located in Amoria - Amoria Plastic Surgery (680000003) 
type DrRoberts struct {
}

func (r DrRoberts) NPCId() uint32 {
	return npc.DrRoberts
}

func (r DrRoberts) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Hi, there~! I'm Dr.Roberts, in charge of the cosmetic lenses here at the Amoria Plastic Surgery Shop! With ").
		BlueText().ShowItemName1(item.AmoriaCosmeticLensCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.AmoriaCosmeticLensCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the kind of beautiful look you've always craved~! Remember, the first thing everyone notices about you is the eyes, and we can help you find the cosmetic lens that most fits you! Now, what would you like to use?").
		String()
	return care.NewGenericCare(hello, r.ProvidedCare())(l, c)
}

func (r DrRoberts) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.CosmeticRegular(item.AmoriaCosmeticLensCouponRegular), r.CosmeticVIP(item.AmoriaCosmeticLensCouponVIP), r.CosmeticOneTime()}
}

func (r DrRoberts) CosmeticRegular(coupon uint32) care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use the regular coupon, you'll be awarded a random pair of cosmetic lenses. Are you going to use a ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and really make the change to your eyes?").
		String()

	next := care.WarnRandomLensColor(prompt, coupon, care.SetFace, r.Initial)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r DrRoberts) CosmeticVIP(coupon uint32) care.ChoiceConfig {
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(coupon, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, care.LensColorChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r DrRoberts) CosmeticOneTime() care.ChoiceConfig {
	//TODO coupon consumption might need to be reviewed
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(item.OneTimeCosmeticLensBlack, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoicesWithNone(prompt, care.LensColorOneTimeChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponOneTimeListText(item.OneTimeCosmeticLensBlack), care.LensCouponMissing(), care.LensEnjoy())
}
