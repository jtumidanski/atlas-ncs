package script

import (
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
)

// Lila is located in The Burning Road - Ariant (260000000)
type Lila struct {
}

func (r Lila) NPCId() uint32 {
	return npc.Lila
}

func (r Lila) HelloMessage() string {
	return message.NewBuilder().
		AddText("Hohoh~ welcome welcome. Welcome to Ariant Skin Care. You have stepped into a renowned Skin Care shop that even the Queen herself frequents this place. If you have ").
		BlueText().ShowItemName1(item.AriantSkinCareCoupon).
		BlackText().AddText(" with you, we'll take care of the rest. How about letting us work on your skin today?").NewLine().
		OpenItem(0).AddText("Skin Care: ").ShowItemImage2(item.AriantSkinCareCoupon).ShowItemName1(item.AriantSkinCareCoupon).CloseItem().
		String()
}

func (r Lila) Coupon() uint32 {
	return item.AriantSkinCareCoupon
}

func (r Lila) MissingMessage() string {
	return message.NewBuilder().
		AddText("Hmmm... I don't think you have our Skin Care coupon with you. Without it, I can't give you the treatment").
		String()
}

func (r Lila) ChooseStyleMessage() string {
	return message.NewBuilder().
		AddText("With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!").
		String()
}
