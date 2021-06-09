package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Dr90212 is located in Amoria - Amoria Plastic Surgery (680000003)
type Dr90212 struct {
}

func (r Dr90212) NPCId() uint32 {
	return npc.Dr90212
}

func (r Dr90212) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Dr90212) Hello() string {
	return message.NewBuilder().
		AddText("Well, hello! Welcome to Amoria Plastic Surgery! Would you like to transform your face into something new? With a ").
		BlueText().ShowItemName1(item.AmoriaFaceCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the face you've always wanted~!").
		String()
}

func (r Dr90212) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.FaceVIP(item.AmoriaFaceCouponVIP)}
}

func (r Dr90212) FaceVIP(coupon uint32) care.ChoiceConfig {
	prompt := message.NewBuilder().
		AddText("Let's see... I can totally transform your face into something new. Don't you want to try it? For ").
		BlueText().ShowItemName1(coupon).
		BlackText().AddText(", you can get the face of your liking. Take your time in choosing the face of your preference.").
		String()

	male := []uint32{20000, 20001, 20003, 20004, 20005, 20006, 20007, 20008, 20018, 20019}
	female := []uint32{21001, 21002, 21003, 21004, 21005, 21006, 21007, 21012, 21018, 21019}
	choiceSupplier := care.FaceChoices(male, female)

	vip := care.ProcessCoupon(coupon, care.SetFace, care.SetSingleUse(true))
	next := care.ShowChoices(prompt, choiceSupplier, vip)

	return care.NewChoiceConfig(next, care.FaceCouponListText(coupon), care.FaceCouponMissing(), care.FaceEnjoy())
}
