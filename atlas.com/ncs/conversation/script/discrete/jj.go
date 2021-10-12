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

// JJ is located in New Leaf City Town Street - NLC Mall (600000001)
type JJ struct {
}

func (r JJ) NPCId() uint32 {
	return npc.JJ
}

func (r JJ) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Hey, there~! I'm J.J.! I'm in charge of the cosmetic lenses here at NLC Shop! If you have a ").
		BlueText().ShowItemName1(item.NLCCosmeticLensCouponVIP).
		BlackText().AddText(", I can get you the best cosmetic lenses you have ever had! Now, what would you like to do?").
		String()
	return care.NewGenericCare(hello, r.ProvidedCare())(l, span, c)
}

func (r JJ) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		care.CosmeticVIPCare(item.NLCCosmeticLensCouponVIP),
		care.CosmeticOneTimeCare(),
	}
}