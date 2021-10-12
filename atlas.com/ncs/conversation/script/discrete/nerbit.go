package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Nerbit is located in New Leaf City Town Street - NLC Mall (600000001)
type Nerbit struct {
}

func (r Nerbit) NPCId() uint32 {
	return npc.Nerbit
}

func (r Nerbit) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, span, c)
}

func (r Nerbit) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I pretty much shouldn't be doing this, but with a ").
		BlueText().ShowItemName1(item.NLCFaceCouponRegular).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()
}

func (r Nerbit) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.PlasticSurgery(item.NLCFaceCouponRegular)}
}

func (r Nerbit) PlasticSurgery(coupon uint32) care.ChoiceConfig {
	maleFace := []uint32{20001, 20008, 20011, 20013, 20024, 20029, 20032}
	femaleFace := []uint32{21000, 21007, 21011, 21012, 21017, 21020, 21022}
	return care.FaceRegularCare(coupon, maleFace, femaleFace, r.Initial)
}