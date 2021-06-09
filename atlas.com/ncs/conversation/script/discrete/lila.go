package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Lila is located in The Burning Road - Ariant (260000000)
type Lila struct {
}

func (r Lila) NPCId() uint32 {
	return npc.Lila
}

func (r Lila) Initial(l logrus.FieldLogger, c script.Context) script.State {
	hello := message.NewBuilder().
		AddText("Hohoh~ welcome welcome. Welcome to Ariant Skin Care. You have stepped into a renowned Skin Care shop that even the Queen herself frequents this place. If you have ").
		BlueText().ShowItemName1(item.AriantSkinCareCoupon).
		BlackText().AddText(" with you, we'll take care of the rest. How about letting work on your skin today?").
		String()

	coupon := item.AriantSkinCareCoupon
	return care.NewGenericCare(hello, []care.ChoiceConfig{r.SkinCare(coupon)})(l, c)
}

func (r Lila) SkinCare(coupon uint32) care.ChoiceConfig {
	skinPrompt := care.SkinCarePrompt(coupon)
	skinColors := care.FixedChoices([]uint32{0, 1, 2, 3, 4})
	missingCoupon := care.SetMissingCoupon("Hmmm... I don't think you have our Skin Care coupon with you. Without it, I can't give you the treatment")
	couponProcessor := care.ProcessCoupon(coupon, care.SetSkin, care.SetSingleUse(true))
	next := care.ShowChoices(skinPrompt, skinColors, couponProcessor)
	choice := care.NewChoiceConfig(next, care.SetListText(care.ListEntry(coupon)), missingCoupon, care.SetEnjoy(care.EnjoyNewSkin()))
	return choice
}
