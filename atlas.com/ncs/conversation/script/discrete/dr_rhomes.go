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

// DrRhomes is located in Orbis Park - Orbis Plastic Surgery (200000201)
type DrRhomes struct {
}

func (r DrRhomes) NPCId() uint32 {
	return npc.DrRhomes
}

func (r DrRhomes) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, span, c)
}

func (r DrRhomes) Hello() string {
	return message.NewBuilder().
		AddText("Hello, I'm Dr. Rhomes, head of the cosmetic lens department here at the Orbis Plastic Surgery Shop.").NewLine().
		AddText("My goal here is to add personality to everyone's eyes through the wonders of cosmetic lenses, and with ").
		BlueText().ShowItemName1(item.OrbisCosmeticLensCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.OrbisCosmeticLensCouponVIP).
		BlackText().AddText(", I can do the same for you, too! Now, what would you like to use?").
		String()
}

func (r DrRhomes) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		care.CosmeticRegularCare(item.OrbisCosmeticLensCouponRegular, r.Initial),
		care.CosmeticVIPCare(item.OrbisCosmeticLensCouponVIP),
		care.CosmeticOneTimeCare(),
	}
}