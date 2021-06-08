package script

import (
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
)

// Romi is located in Orbis Park - Orbis Skin-Care (200000203)
type Romi struct {
}

func (r Romi) NPCId() uint32 {
	return npc.Romi
}

func (r Romi) HelloMessage() string {
	return message.NewBuilder().
		AddText("Well, hello! Welcome to the Orbis Skin-Care~! Would you like to have a firm, tight, healthy looking skin like mine?  With ").
		BlueText().ShowItemName1(item.OrbisSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").NewLine().
		OpenItem(0).AddText("Skin Care: ").ShowItemImage2(item.OrbisSkinCoupon).ShowItemName1(item.OrbisSkinCoupon).CloseItem().
		String()
}

func (r Romi) Coupon() uint32 {
	return item.OrbisSkinCoupon
}

func (r Romi) MissingMessage() string {
	return message.NewBuilder().
		AddText("Um... you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you...").
		String()
}

func (r Romi) ChooseStyleMessage() string {
	return message.NewBuilder().
		AddText("With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!").
		String()
}
