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

// DrRoberts is located in Amoria - Amoria Plastic Surgery (680000003) 
type DrRoberts struct {
}

func (r DrRoberts) NPCId() uint32 {
	return npc.DrRoberts
}

func (r DrRoberts) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Hi, there~! I'm Dr.Roberts, in charge of the cosmetic lenses here at the Amoria Plastic Surgery Shop! With ").
		BlueText().ShowItemName1(item.AmoriaCosmeticLensCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.AmoriaCosmeticLensCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the kind of beautiful look you've always craved~! Remember, the first thing everyone notices about you is the eyes, and we can help you find the cosmetic lens that most fits you! Now, what would you like to use?").
		String()
	return care.NewGenericCare(hello, r.ProvidedCare())(l, span, c)
}

func (r DrRoberts) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		care.CosmeticRegularCare(item.AmoriaCosmeticLensCouponRegular, r.Initial),
		care.CosmeticVIPCare(item.AmoriaCosmeticLensCouponVIP),
		care.CosmeticOneTimeCare(),
	}
}