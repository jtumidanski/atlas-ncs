package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Nerbit is located in New Leaf City Town Street - NLC Mall (600000001)
type Nerbit struct {
}

func (r Nerbit) NPCId() uint32 {
	return npc.Nerbit
}

func (r Nerbit) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Nerbit) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I pretty much shouldn't be doing this, but with a ").
		BlueText().ShowItemName1(item.NLFFaceCouponRegular).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()
}

func (r Nerbit) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.PlasticSurgery(item.NLFFaceCouponRegular)}
}

func (r Nerbit) PlasticSurgery(coupon uint32) care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use the regular coupon, your face may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText("?").
		String()

	maleFace := []uint32{20001, 20008, 20011, 20013, 20024, 20029, 20032}
	femaleFace := []uint32{21000, 21007, 21011, 21012, 21017, 21020, 21022}

	next := care.WarnRandomFace(prompt, coupon, maleFace, femaleFace, care.SetFace, r.Initial)
	return care.NewChoiceConfig(next, care.FaceCouponListText(coupon), care.FaceCouponMissing(), care.FaceEnjoy())
}