package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Pata is located in Mu Lung - Mu Lung (250000000)
type Pata struct {
}

func (r Pata) NPCId() uint32 {
	return npc.Pata
}

func (r Pata) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Pata) Hello() string {
	return message.NewBuilder().
		AddText("Hey, I'm Pata, and I am a renowned plastic surgeon and cosmetic lens expert here in Mu Lung. I believe your face and eyes are the most important features in your body, and with ").
		BlueText().ShowItemName1(item.MuLungPlasticSurgeryCouponVIP).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.MuLungCosmeticLensCouponVIP).
		BlackText().AddText(", I can prescribe the right kind of facial care and cosmetic lenses for you. Now, what would you like to use?").
		String()
}

func (r Pata) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.FaceVIP(item.MuLungPlasticSurgeryCouponVIP), r.CosmeticVIP(item.MuLungCosmeticLensCouponVIP), r.CosmeticOneTime()}
}

func (r Pata) FaceVIP(coupon uint32) care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("Let's see... I can totally transform your face into something new. Don't you want to try it? For ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(", you can get the face of your liking. Take your time in choosing the face of your preference.").
		String()

	male := []uint32{20000, 20001, 20004, 20005, 20006, 20007, 20009, 20012, 20022, 20028, 20031}
	female := []uint32{21000, 21003, 21005, 21006, 21008, 21009, 21011, 21012, 21023, 21024, 21026}
	choiceSupplier := care.FaceChoices(male, female)

	vip := care.ProcessCoupon(coupon, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, choiceSupplier, vip)

	return care.NewChoiceConfig(next, care.FaceCouponListText(coupon), care.FaceCouponMissing(), care.FaceEnjoy())
}

func (r Pata) CosmeticVIP(coupon uint32) care.ChoiceConfig {
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(coupon, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, care.LensColorChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponListText(coupon), care.LensCouponMissing(), care.LensEnjoy())
}

func (r Pata) CosmeticOneTime() care.ChoiceConfig {
	//TODO coupon consumption might need to be reviewed
	prompt := "With our specialized machine, you can see yourself after the treatment in advance. What kind of lens would you like to wear? Choose the style of your liking."

	special := care.ProcessCoupon(item.OneTimeCosmeticLensBlack, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoicesWithNone(prompt, care.LensColorOneTimeChoices, special)
	return care.NewChoiceConfig(next, care.LensCouponOneTimeListText(item.OneTimeCosmeticLensBlack), care.LensCouponMissing(), care.LensEnjoy())
}
