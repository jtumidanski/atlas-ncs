package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Naran is located in Mu Lung - Mu Lung (250000000)
type Naran struct {
}

func (r Naran) NPCId() uint32 {
	return npc.Naran
}

func (r Naran) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Well, hello! Welcome to the Mu Lung Skin-Care! Would you like to have a firm, tight, healthy looking skin like mine?  With a ").
		BlueText().ShowItemName1(item.MuLungSkinCareCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").
		String()
	return care.NewGenericSkinCare(item.MuLungSkinCareCoupon, hello)(l, c)
}
