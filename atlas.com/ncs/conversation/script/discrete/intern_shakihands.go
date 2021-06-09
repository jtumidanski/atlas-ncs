package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// InternShakihands is located in Amoria - Amoria Plastic Surgery (680000003)
type InternShakihands struct {
}

func (r InternShakihands) NPCId() uint32 {
	return npc.InternShakihands
}

func (r InternShakihands) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r InternShakihands) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I pretty much shouldn't be doing this, but with a ").
		BlueText().ShowItemName1(item.AmoriaFaceCouponRegular).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()
}

func (r InternShakihands) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.PlasticSurgery(item.AmoriaFaceCouponRegular)}
}

func (r InternShakihands) PlasticSurgery(coupon uint32) care.ChoiceConfig {
	maleFace := []uint32{20002, 20005, 20007, 20011, 20014, 20027, 20029}
	femaleFace := []uint32{21001, 21005, 21007, 21017, 21018, 21020, 21022}
	return care.FaceRegularCare(coupon, maleFace, femaleFace, r.Initial)
}
