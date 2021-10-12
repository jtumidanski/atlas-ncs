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

// DrFeeble is located in Victoria Road - Henesys Plastic Surgery (100000103)
type DrFeeble struct {
}

func (r DrFeeble) NPCId() uint32 {
	return npc.DrFeeble
}

func (r DrFeeble) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, span, c)
}

func (r DrFeeble) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I pretty much shouldn't be doing this, but with a ").
		BlueText().ShowItemName1(item.HenesysFaceCouponRegular).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()
}

func (r DrFeeble) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.PlasticSurgery(item.HenesysFaceCouponRegular)}
}

func (r DrFeeble) PlasticSurgery(coupon uint32) care.ChoiceConfig {
	maleFace := []uint32{20000, 20005, 20008, 20012, 20016, 20022, 20032}
	femaleFace := []uint32{21000, 21002, 21008, 21014, 21020, 21024, 21029}
	return care.FaceRegularCare(coupon, maleFace, femaleFace, r.Initial)
}
