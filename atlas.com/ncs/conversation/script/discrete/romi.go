package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Romi is located in Orbis Park - Orbis Skin-Care (200000203)
type Romi struct {
}

func (r Romi) NPCId() uint32 {
	return npc.Romi
}

func (r Romi) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Well, hello! Welcome to the Orbis Skin-Care~! Would you like to have a firm, tight, healthy looking skin like mine?  With ").
		BlueText().ShowItemName1(item.OrbisSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").
		String()
	return care.NewGenericSkinCare(item.OrbisSkinCoupon, hello)(l, c)
}
