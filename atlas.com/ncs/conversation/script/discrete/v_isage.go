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
	male := []uint32{20000, 20001, 20003, 20004, 20005, 20006, 20008, 20012, 20031}
	female := []uint32{21001, 21002, 21003, 21004, 21005, 21006, 21008, 21012, 21016}
	return care.FaceVIPCare(coupon, male, female)
}
