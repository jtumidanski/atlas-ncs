package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DenmaTheOwner is located in Victoria Road - Henesys Plastic Surgery (100000103)
type DenmaTheOwner struct {
}

func (r DenmaTheOwner) NPCId() uint32 {
	return npc.DenmaTheOwner
}

func (r DenmaTheOwner) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r DenmaTheOwner) Hello() string {
	return message.NewBuilder().
		AddText("Well, hello! Welcome to the Henesys Plastic Surgery! Would you like to transform your face into something new? With a ").
		BlueText().ShowItemName1(item.HenesysFaceCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the face you've always wanted~!").
		String()
}

func (r DenmaTheOwner) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.FaceVIP(item.HenesysFaceCouponVIP)}
}

func (r DenmaTheOwner) FaceVIP(coupon uint32) care.ChoiceConfig {
	male := []uint32{20000, 20001, 20003, 20004, 20005, 20006, 20007, 20008, 20012, 20014, 20015, 20022, 20028, 20031}
	female := []uint32{21000, 21001, 21002, 21003, 21004, 21005, 21006, 21007, 21008, 21012, 21013, 21014, 21023, 21026}
	return care.FaceVIPCare(coupon, male, female)
}
