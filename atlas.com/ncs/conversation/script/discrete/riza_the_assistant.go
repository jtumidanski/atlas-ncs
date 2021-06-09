package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// RizaTheAssistant is located in Orbis Park - Orbis Plastic Surgery (200000201)
type RizaTheAssistant struct {
}

func (r RizaTheAssistant) NPCId() uint32 {
	return npc.RizaTheAssistant
}

func (r RizaTheAssistant) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r RizaTheAssistant) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I pretty much shouldn't be doing this, but with a ").
		BlueText().ShowItemName1(item.OrbisFaceCouponRegular).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()
}

func (r RizaTheAssistant) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.PlasticSurgery(item.OrbisFaceCouponRegular)}
}

func (r RizaTheAssistant) PlasticSurgery(coupon uint32) care.ChoiceConfig {
	maleFace := []uint32{20003, 20011, 20021, 20022, 20023, 20027, 20031}
	femaleFace := []uint32{21004, 21007, 21010, 21012, 21020, 21021, 21030}
	return care.FaceRegularCare(coupon, maleFace, femaleFace, r.Initial)
}
