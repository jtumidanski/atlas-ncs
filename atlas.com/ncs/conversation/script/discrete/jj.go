package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// JJ is located in New Leaf City Town Street - NLC Mall (600000001)
type JJ struct {
}

func (r JJ) NPCId() uint32 {
	return npc.JJ
}

func (r JJ) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Hey, there~! I'm J.J.! I'm in charge of the cosmetic lenses here at NLC Shop! If you have a ").
		BlueText().ShowItemName1(item.NLCCosmeticLensCouponVIP).
		BlackText().AddText(", I can get you the best cosmetic lenses you have ever had! Now, what would you like to do?").
		String()
	return care.NewGenericCare(hello, r.ProvidedCare())(l, c)
}

func (r JJ) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.CosmeticVIP(item.NLCCosmeticLensCouponVIP), r.CosmeticOneTime()}
}

func (r JJ) CosmeticVIP(coupon uint32) care.ChoiceConfig {
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(coupon, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, care.LensColorChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r JJ) CosmeticOneTime() care.ChoiceConfig {
	//TODO coupon consumption might need to be reviewed
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(item.OneTimeCosmeticLensBlack, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoicesWithNone(prompt, care.LensColorOneTimeChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponOneTimeListText(item.OneTimeCosmeticLensBlack), care.LensCouponMissing(), care.LensEnjoy())
}
