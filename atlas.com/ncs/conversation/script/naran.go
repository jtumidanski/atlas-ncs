package script

import (
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
)

// Naran is located in Mu Lung - Mu Lung (250000000)
type Naran struct {
}

func (r Naran) NPCId() uint32 {
	return npc.Naran
}

func (r Naran) HelloMessage() string {
	return message.NewBuilder().
		AddText("Well, hello! Welcome to the Mu Lung Skin-Care! Would you like to have a firm, tight, healthy looking skin like mine?  With a ").
		BlueText().ShowItemName1(item.MuLungSkinCareCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").NewLine().
		OpenItem(0).AddText("Skin Care: ").ShowItemImage2(item.AriantSkinCareCoupon).ShowItemName1(item.AriantSkinCareCoupon).CloseItem().
		String()
}

func (r Naran) Coupon() uint32 {
	return item.MuLungSkinCareCoupon
}

func (r Naran) MissingMessage() string {
	return message.NewBuilder().
		AddText("Um... you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you...").
		String()
}

func (r Naran) ChooseStyleMessage() string {
	return message.NewBuilder().
		AddText("With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!").
		String()
}
