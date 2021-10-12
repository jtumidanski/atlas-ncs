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

// Everton is located in Ludibrium - Ludibrium Plastic Surgery (220000003)
type Everton struct {
}

func (r Everton) NPCId() uint32 {
	return npc.Everton
}

func (r Everton) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, span, c)
}

func (r Everton) Hello() string {
	return message.NewBuilder().
		AddText("Well, I'm bored, so I'll help out the doctor. For a ").
		BlueText().ShowItemName1(item.LudibriumFaceCouponRegular).
		BlackText().AddText(", I will change the way you look. But don't forget, it will be random!").
		String()
}

func (r Everton) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.PlasticSurgery(item.LudibriumFaceCouponRegular)}
}

func (r Everton) PlasticSurgery(coupon uint32) care.ChoiceConfig {
	maleFace := []uint32{20001, 20003, 20007, 20013, 20021, 20023, 20025}
	femaleFace := []uint32{21002, 21004, 21006, 21008, 21022, 21027, 21029}
	return care.FaceRegularCare(coupon, maleFace, femaleFace, r.Initial)
}
