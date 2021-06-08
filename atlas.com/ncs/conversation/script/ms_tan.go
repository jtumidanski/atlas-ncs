package script

import (
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
)

// MsTan is located in Victoria Road - Henesys Skin-Care (100000105)
type MsTan struct {
}

func (r MsTan) NPCId() uint32 {
	return npc.MsTan
}

func (r MsTan) HelloMessage() string {
	return message.NewBuilder().
		AddText("Well, hello! Welcome to the Henesys Skin-Care! Would you like to have a firm, tight, healthy looking skin like mine?  With a ").
		BlueText().ShowItemName1(item.HenesysSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").NewLine().
		OpenItem(0).AddText("Skin Care: ").ShowItemImage2(item.HenesysSkinCoupon).ShowItemName1(item.HenesysSkinCoupon).CloseItem().
		String()
}

func (r MsTan) Coupon() uint32 {
	return item.HenesysSkinCoupon
}

func (r MsTan) MissingMessage() string {
	return message.NewBuilder().
		AddText("Um... you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you...").
		String()
}

func (r MsTan) ChooseStyleMessage() string {
	return message.NewBuilder().
		AddText("With our specialized machine, you can see yourself after the treatment in advance. What kind of skin-treatment would you like to do? Choose the style of your liking.").
		String()
}
