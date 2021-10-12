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

// DrLenu is located in Victoria Road - Henesys Plastic Surgery (100000103)
type DrLenu struct {
}

func (r DrLenu) NPCId() uint32 {
	return npc.DrLenu
}

func (r DrLenu) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Hi, there~! I'm Dr. Lenu, in charge of the cosmetic lenses here at the Henesys Plastic Surgery Shop! With ").
		BlueText().ShowItemName1(item.HenesysCosmeticLensCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.HenesysCosmeticLensCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the kind of beautiful look you've always craved~! Remember, the first thing everyone notices about you is the eyes, and we can help you find the cosmetic lens that most fits you! Now, what would you like to use?").
		String()
	return care.NewGenericCare(hello, r.ProvidedCare())(l, span, c)
}

func (r DrLenu) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		care.CosmeticRegularCare(item.HenesysCosmeticLensCouponRegular, r.Initial),
		care.CosmeticVIPCare(item.HenesysCosmeticLensCouponVIP),
		care.CosmeticOneTimeCare(),
	}
}