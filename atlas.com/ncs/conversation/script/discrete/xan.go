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

// Xan is located in Singapore - CBD (540000000)
type Xan struct {
}

func (r Xan) NPCId() uint32 {
	return npc.Xan
}

func (r Xan) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Well, hello! Welcome to the Lian Hua Hua Skin-Care! Would you like to have a firm, tight, healthy looking skin like mine?  With ").
		BlueText().ShowItemName1(item.CBDSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted!").
		String()
	return care.NewGenericSkinCare(item.CBDSkinCoupon, hello)(l, span, c)
}
