package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// FranzTheOwner is located in Orbis Park - Orbis Plastic Surgery (200000201)
type FranzTheOwner struct {
}

func (r FranzTheOwner) NPCId() uint32 {
	return npc.FranzTheOwner
}

func (r FranzTheOwner) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, span, c)
}

func (r FranzTheOwner) Hello() string {
	return message.NewBuilder().
		AddText("Well well well, welcome to the Orbis Plastic Surgery! Would you like to transform your face into something new? With a ").
		BlueText().ShowItemName1(item.OrbisFaceCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the face you've always wanted~!").
		String()
}

func (r FranzTheOwner) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.FaceVIP(item.OrbisFaceCouponVIP)}
}

func (r FranzTheOwner) FaceVIP(coupon uint32) care.ChoiceConfig {
	male := []uint32{20000, 20001, 20003, 20004, 20005, 20006, 20007, 20008, 20012, 20014, 20022, 20028, 20031}
	female := []uint32{21000, 21001, 21002, 21003, 21004, 21005, 21006, 21007, 21008, 21012, 21014, 21023, 21026}
	return care.FaceVIPCare(coupon, male, female)
}
