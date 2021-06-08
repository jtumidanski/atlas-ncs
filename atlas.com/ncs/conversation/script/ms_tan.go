package script

import (
	"atlas-ncs/conversation/script/generic/skin"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MsTan is located in Victoria Road - Henesys Skin-Care (100000105)
type MsTan struct {
}

func (r MsTan) NPCId() uint32 {
	return npc.MsTan
}

func (r MsTan) Initial(l logrus.FieldLogger, c Context) State {
	hello := message.NewBuilder().
		AddText("Well, hello! Welcome to the Henesys Skin-Care! Would you like to have a firm, tight, healthy looking skin like mine?  With a ").
		BlueText().ShowItemName1(item.HenesysSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").
		String()
	return skin.NewGenericCare(item.HenesysSkinCoupon, hello)(l, c)
}
