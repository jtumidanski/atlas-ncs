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

// Gina is located in Ludibrium - Ludibrium Skin Care (220000005)
type Gina struct {
}

func (r Gina) NPCId() uint32 {
	return npc.Gina
}

func (r Gina) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Oh, hello! Welcome to the Ludibrium Skin-Care! Are you interested in getting tanned and looking sexy? How about a beautiful, snow-white skin? If you have ").
		BlueText().ShowItemName1(item.LudibriumSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always dreamed of!").
		String()
	return care.NewGenericSkinCare(item.LudibriumSkinCoupon, hello)(l, span, c)
}
