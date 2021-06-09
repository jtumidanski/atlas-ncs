package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Miranda is located in New Leaf City Town Street - NLC Mall (600000001)
type Miranda struct {
}

func (r Miranda) NPCId() uint32 {
	return npc.Miranda
}

func (r Miranda) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Well, hello! Welcome to the NLC Skin-Care! Would you like to have a firm, tight, healthy looking skin like mine?  With ").
		BlueText().ShowItemName1(item.NLCSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").
		String()
	return care.NewGenericSkinCare(item.NLCSkinCoupon, hello)(l, c)
}
