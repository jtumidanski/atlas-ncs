package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// VIsage is located in New Leaf City Town Street - NLC Mall (600000001)
type VIsage struct {
}

func (r VIsage) NPCId() uint32 {
	return npc.VIsage
}

func (r VIsage) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r VIsage) Hello() string {
	return message.NewBuilder().
		AddText("Well, hello! Welcome to the New Leaf City Plastic Surgery! Would you like to transform your face into something new? With a ").
		BlueText().ShowItemName1(item.NLCFaceCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the face you've always wanted~!").
		String()
}

func (r VIsage) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.FaceVIP(item.NLCFaceCouponVIP)}
}

func (r VIsage) FaceVIP(coupon uint32) care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("Let's see... I can totally transform your face into something new. Don't you want to try it? For ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(", you can get the face of your liking. Take your time in choosing the face of your preference.").
		String()

	male := []uint32{20000, 20001, 20003, 20004, 20005, 20006, 20008, 20012, 20031}
	female := []uint32{21001, 21002, 21003, 21004, 21005, 21006, 21008, 21012, 21016}
	choiceSupplier := care.FaceChoices(male, female)

	vip := care.ProcessCoupon(coupon, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, choiceSupplier, vip)

	return care.NewChoiceConfig(next, care.FaceCouponListText(coupon), care.FaceCouponMissing(), care.FaceEnjoy())
}
