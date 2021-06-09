package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Ellie is located in Ludibrium - Ludibrium Plastic Surgery (220000003)
type Ellie struct {
}

func (r Ellie) NPCId() uint32 {
	return npc.Ellie
}

func (r Ellie) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Ellie) Hello() string {
	return message.NewBuilder().
		AddText("Well, hello! Welcome to the Ludibrium Plastic Surgery! Would you like to transform your face into something new? With a ").
		BlueText().ShowItemName1(item.LudibriumFaceCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the face you've always wanted~!").
		String()
}

func (r Ellie) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.FaceVIP(item.LudibriumFaceCouponVIP)}
}

func (r Ellie) FaceVIP(coupon uint32) care.ChoiceConfig {
	male := []uint32{20000, 20001, 20003, 20004, 20005, 20006, 20007, 20008, 20011, 20012, 20014, 20031}
	female := []uint32{21000, 21001, 21002, 21003, 21004, 21005, 21006, 21007, 21008, 21010, 21012, 21014}
	return care.FaceVIPCare(coupon, male, female)
}
