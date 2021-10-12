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

// Kelvin is located in Singapore - CBD (540000000)
type Kelvin struct {
}

func (r Kelvin) NPCId() uint32 {
	return npc.Kelvin
}

func (r Kelvin) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, span, c)
}

func (r Kelvin) Hello() string {
	return message.NewBuilder().
		AddText("Let's see...I can totally transform your face into something new. Don't you want to try it? For ").
		BlueText().ShowItemName1(item.CBDFaceCouponVIP).
		BlackText().AddText(", you can get the face of your liking. Take your time in choosing the face of your preference...").
		String()
}

func (r Kelvin) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.FaceVIP(item.CBDFaceCouponVIP)}
}

func (r Kelvin) FaceVIP(coupon uint32) care.ChoiceConfig {
	male := []uint32{20005, 20012, 20013, 20020, 20021, 20026}
	female := []uint32{21006, 21009, 21011, 21012, 21021, 21025}
	return care.FaceVIPCare(coupon, male, female)
}
