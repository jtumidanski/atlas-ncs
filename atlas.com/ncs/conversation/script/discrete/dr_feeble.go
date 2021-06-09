package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DrFeeble is located in Victoria Road - Henesys Plastic Surgery (100000103)
type DrFeeble struct {
}

func (r DrFeeble) NPCId() uint32 {
	return npc.DrFeeble
}

func (r DrFeeble) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Hi, I pretty much shouldn't be doing this, but with a ").
		BlueText().ShowItemName1(item.HenesysFaceCouponRegular).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()
	return care.NewGenericCare(hello, r.ProvidedCare())(l, c)
}

func (r DrFeeble) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.PlasticSurgery()}
}

func (r DrFeeble) PlasticSurgery() care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("If you use the regular coupon, your face may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(item.HenesysFaceCouponRegular).
		BlackText().AddText("?").
		String()

	maleFace := []uint32{20000, 20005, 20008, 20012, 20016, 20022, 20032}
	femaleFace := []uint32{21000, 21002, 21008, 21014, 21020, 21024, 21029}

	next := care.WarnRandomFace(prompt, item.HenesysFaceCouponRegular, maleFace, femaleFace, care.SetFace, r.Initial)
	return care.NewChoiceConfig(next, care.FaceCouponListText(item.HenesysFaceCouponRegular), care.FaceCouponMissing(), care.FaceEnjoy())
}
