package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Noel is located in Singapore - CBD (540000000)
type Noel struct {
}

func (r Noel) NPCId() uint32 {
	return npc.Noel
}

func (r Noel) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Noel) Hello() string {
	return message.NewBuilder().
		AddText("If you use this regular coupon, your face may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(item.CBDFaceCouponRegular).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()
}

func (r Noel) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.PlasticSurgery(item.CBDFaceCouponRegular)}
}

func (r Noel) PlasticSurgery(coupon uint32) care.ChoiceConfig {
	maleFace := []uint32{20002, 20005, 20006, 20013, 20017, 20021, 20024}
	femaleFace := []uint32{21002, 21003, 21014, 21016, 21017, 21021, 21027}
	return care.FaceRegularCare(coupon, maleFace, femaleFace, r.Initial)
}
