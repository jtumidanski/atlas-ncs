package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Saeko is located in Zipangu - Plastic Surgery (801000002)
type Saeko struct {
}

func (r Saeko) NPCId() uint32 {
	return npc.Saeko
}

func (r Saeko) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Saeko) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I pretty much shouldn't be doing this, but with a ").
		BlueText().ShowItemName1(item.ShowaFaceCouponRegular).
		BlackText().AddText(" or a ").
		BlueText().ShowItemName1(item.ShowaCosmeticLensCouponRegular).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()
}

func (r Saeko) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.PlasticSurgery(item.ShowaFaceCouponRegular),
		care.CosmeticRegularCare(item.ShowaCosmeticLensCouponRegular, r.Initial),
	}
}

func (r Saeko) PlasticSurgery(coupon uint32) care.ChoiceConfig {
	maleFace := []uint32{20000, 20016, 20019, 20020, 20021, 20024, 20026}
	femaleFace := []uint32{21000, 21002, 21009, 21016, 21022, 21025, 21027}
	return care.FaceRegularCare(coupon, maleFace, femaleFace, r.Initial)
}
