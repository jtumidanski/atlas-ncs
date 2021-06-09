package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Noma is located in Mu Lung - Mu Lung (250000000)
type Noma struct {
}

func (r Noma) NPCId() uint32 {
	return npc.Noma
}

func (r Noma) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Noma) Hello() string {
	return message.NewBuilder().
		AddText("Hey, I'm Noma, and I am assisting Pata in changing faces and applying lenses as my internship studies. With ").
		BlueText().ShowItemName1(item.MuLungPlasticSurgeryCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.MuLungCosmeticLensCouponRegular).
		BlackText().AddText(", I can change the way you look. Now, what would you like to use?").
		String()
}

func (r Noma) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.PlasticSurgery(item.MuLungPlasticSurgeryCouponRegular),
		care.CosmeticRegularCare(item.MuLungCosmeticLensCouponRegular, r.Initial),
	}
}

func (r Noma) PlasticSurgery(coupon uint32) care.ChoiceConfig {
	maleFace := []uint32{20002, 20005, 20007, 20011, 20014, 20017, 20029}
	femaleFace := []uint32{21001, 21010, 21013, 21018, 21020, 21021, 21030}
	return care.FaceRegularCare(coupon, maleFace, femaleFace, r.Initial)
}