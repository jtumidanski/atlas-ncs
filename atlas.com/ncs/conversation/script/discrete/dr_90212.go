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

// Dr90212 is located in Amoria - Amoria Plastic Surgery (680000003)
type Dr90212 struct {
}

func (r Dr90212) NPCId() uint32 {
	return npc.Dr90212
}

func (r Dr90212) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, span, c)
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
	male := []uint32{20000, 20001, 20003, 20004, 20005, 20006, 20007, 20008, 20018, 20019}
	female := []uint32{21001, 21002, 21003, 21004, 21005, 21006, 21007, 21012, 21018, 21019}
	return care.FaceVIPCare(coupon, male, female)
}
