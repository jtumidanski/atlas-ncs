package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Bomack is located in New Leaf City Town Street - NLC Mall (600000001)
type Bomack struct {
}

func (r Bomack) NPCId() uint32 {
	return npc.Bomack
}

func (r Bomack) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Hi, there~! I'm Bomack. If you have a ").
		BlueText().ShowItemName1(item.NLCCosmeticLensCouponRegular).
		BlackText().AddText(", I can prescribe the right kind of cosmetic lenses for you. Now, what would you like to do?").
		String()
	return care.NewGenericCare(hello, r.ProvidedCare())(l, c)
}

func (r Bomack) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.CosmeticRegular(item.NLCCosmeticLensCouponRegular)}
}

func (r Bomack) CosmeticRegular(coupon uint32) care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use the regular coupon, you'll be awarded a random pair of cosmetic lenses. Are you going to use a ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(" and really make the change to your eyes?").
		String()

	next := care.WarnRandomLensColor(prompt, coupon, care.SetFace, r.Initial)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}
