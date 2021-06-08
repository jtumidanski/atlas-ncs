package script

import (
	"atlas-ncs/conversation/script/generic/skin"
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

func (r Lila) Initial(l logrus.FieldLogger, c Context) State {
	hello := message.NewBuilder().
		AddText("Hohoh~ welcome welcome. Welcome to Ariant Skin Care. You have stepped into a renowned Skin Care shop that even the Queen herself frequents this place. If you have ").
		BlueText().ShowItemName1(item.AriantSkinCareCoupon).
		BlackText().AddText(" with you, we'll take care of the rest. How about letting work on your skin today?").
		String()
	missingCoupon := skin.CareConfiguratorMissingCoupon("Hmmm... I don't think you have our Skin Care coupon with you. Without it, I can't give you the treatment")
	return skin.NewGenericCare(item.AriantSkinCareCoupon, hello, missingCoupon)(l, c)
}
