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

// DrBosch is located in Ludibrium - Ludibrium Plastic Surgery (220000003)
type DrBosch struct {
}

func (r DrBosch) NPCId() uint32 {
	return npc.DrBosch
}

func (r DrBosch) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, span, c)
}

func (r DrBosch) Hello() string {
	return message.NewBuilder().
		AddText("Um... hi, I'm Dr. Bosch, and I am a cosmetic lens expert here at the Ludibrium Plastic Surgery Shop. I believe your eyes are the most important feature in your body, and with ").
		BlueText().ShowItemName1(item.LudibriumCosmeticLensCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.LudibriumCosmeticLensCouponVIP).
		BlackText().AddText(", I can prescribe the right kind of cosmetic lenses for you. Now, what would you like to use?").
		String()
}

func (r DrBosch) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		care.CosmeticRegularCare(item.LudibriumCosmeticLensCouponRegular, r.Initial),
		care.CosmeticVIPCare(item.LudibriumCosmeticLensCouponVIP),
		care.CosmeticOneTimeCare(),
	}
}