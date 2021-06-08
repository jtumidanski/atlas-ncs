package script

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Natalie is located in Victoria Road - Henesys Hair Salon (100000104)
type Natalie struct {
}

func (r Natalie) NPCId() uint32 {
	return npc.Natalie
}

func (r Natalie) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Natalie) Hello(l logrus.FieldLogger, c Context) State {
	hello := message.NewBuilder().
		AddText("I'm the head of this hair salon. If you have a ").
		BlueText().ShowItemName1(item.HenesysHairStyleCouponVIP).
		BlackText().AddText(" or a ").
		BlueText().ShowItemName1(item.HenesysHairColorCouponVIP).
		BlackText().AddText(" allow me to take care of your hairdo. Please choose the one you want.").
		String()

	return care.NewGenericCare(hello, []care.ChoiceConfig{r.StyleHair(), care.ColorCareChoice(item.HenesysHairColorCouponVIP)})(l, c)
}

func (r Natalie) StyleHair() care.ChoiceConfig {
	listText := care.StyleListEntry(item.HenesysHairStyleCouponVIP)
	hairStyle := care.StylePrompt(item.HenesysHairStyleCouponVIP)
	maleHair := []uint32{character.HairBlackCatalyst, character.HairBlackTopknot, character.HairBlackWind, character.HairBlackShaggyWax, character.HairBlackAcorn, character.HairBlackAranCut, character.HairBlackTheCoco}
	femaleHair := []uint32{character.HairBlackAngelica, character.HairBlackChantelle, character.HairBlackFourtailBraids, character.HairBlackCrazyMedusa, character.HairBlackFrizzleDizzle, character.HairBlackAranHair, character.HairBlackFullBangs}

	vip := care.ProcessCoupon(item.HenesysHairStyleCouponVIP, care.SetHair, care.SetSingleUse(true))
	membership := care.ProcessCoupon(item.HenesysHairMembershipCoupon, care.SetHair, care.SetSingleUse(false), care.SetFailFunction(vip))
	return care.ChoiceConfig{
		ListText:      listText,
		NextState:     care.ShowChoices(hairStyle, care.HairStyleChoices(maleHair, femaleHair), membership),
		MissingCoupon: "Hmmm...it looks like you don't have our designated coupon...I'm afraid I can't give you a haircut without it. I'm sorry...",
		Enjoy:         "Enjoy your new and improved hairstyle!",
	}
}
