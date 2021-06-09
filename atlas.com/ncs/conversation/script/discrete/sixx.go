package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Sixx is located in Singapore - CBD (540000000)
type Sixx struct {
}

func (r Sixx) NPCId() uint32 {
	return npc.Sixx
}

func (r Sixx) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, c)
}

func (r Sixx) Hello() string {
	return message.NewBuilder().
		AddText("Hi, there! I'm Sixx, in charge of Da Yan Jing Lens Shop here at CBD! With ").
		BlueText().ShowItemName1(item.CBDCosmeticLensCouponRegular).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.CBDCosmeticLensCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the kind of beautiful look you've always craved! Remember, the first thing everyone notices about you are the eyes, and we can help you find the cosmetic lens that most fits you! Now, what would you like to use?").
		String()
}

func (r Sixx) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		care.CosmeticRegularCare(item.CBDCosmeticLensCouponRegular, r.Initial),
		care.CosmeticVIPCare(item.CBDCosmeticLensCouponVIP),
		care.CosmeticOneTimeCare(),
	}
}
