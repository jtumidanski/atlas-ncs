package script

import (
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
)

// Gina is located in Ludibrium - Ludibrium Skin Care (220000005)
type Gina struct {
}

func (r Gina) NPCId() uint32 {
	return npc.Gina
}

func (r Gina) HelloMessage() string {
	return message.NewBuilder().
		AddText("Oh, hello! Welcome to the Ludibrium Skin-Care! Are you interested in getting tanned and looking sexy? How about a beautiful, snow-white skin? If you have ").
		BlueText().ShowItemName1(item.LudibriumSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always dreamed of!").NewLine().
		OpenItem(0).AddText("Skin Care: ").ShowItemImage2(item.LudibriumSkinCoupon).ShowItemName1(item.LudibriumSkinCoupon).CloseItem().
		String()
}

func (r Gina) Coupon() uint32 {
	return item.LudibriumSkinCoupon
}

func (r Gina) MissingMessage() string {
	return message.NewBuilder().
		AddText("Um... you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you...").
		String()
}

func (r Gina) ChooseStyleMessage() string {
	return message.NewBuilder().
		AddText("With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!").
		String()
}
