package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Hikekuro is located in Zipangu - Plastic Surgery (801000002)
type Hikekuro struct {
}

func (r Hikekuro) NPCId() uint32 {
	return npc.Hikekuro
}

func (r Hikekuro) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Hikekuro) Hello() string {
	return message.NewBuilder().
		AddText("Well well well, welcome to the Showa Plastic Surgery! Would you like to transform your face into something new? With a ").
		BlueText().ShowItemName1(item.ShowaFaceCouponVIP).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.ShowaCosmeticLensCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the face you've always wanted~!").
		String()
}

func (r Hikekuro) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.FaceVIP(item.ShowaFaceCouponVIP),
		care.CosmeticVIPCare(item.ShowaCosmeticLensCouponVIP),
		care.CosmeticOneTimeCare(),
	}
}

func (r Hikekuro) FaceVIP(coupon uint32) care.ChoiceConfig {
	male := []uint32{20000, 20004, 20005, 20012, 20020, 20031}
	female := []uint32{21000, 21003, 21006, 21012, 21021, 21024}
	return care.FaceVIPCare(coupon, male, female)
}