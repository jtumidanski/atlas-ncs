package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Brittany is located in Victoria Road - Henesys Hair Salon (100000104)
type Brittany struct {
}

func (r Brittany) NPCId() uint32 {
	return npc.Brittany
}

func (r Brittany) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, c)
}

func (r Brittany) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.RegularStyleHair(item.HenesysHairStyleCouponREG),
		r.ExperimentalStyleHair(item.HenesysHairStyleCouponEXP),
		care.ColorCareRandom(item.HenesysHairColorCouponREG, r.Initial),
	}
}

func (r Brittany) Hello() string {
	return message.NewBuilder().
		AddText("I'm Brittany the assistant. If you have ").
		BlueText().ShowItemName1(item.HenesysHairStyleCouponREG).
		BlackText().AddText(", ").
		BlueText().ShowItemName1(item.HenesysHairStyleCouponEXP).
		BlackText().AddText(" or ").
		BlueText().ShowItemName1(item.HenesysHairColorCouponREG).
		BlackText().AddText(" by any chance, then how about letting me change your hairdo?").
		String()
}

func (r Brittany) RegularStyleHair(coupon uint32) care.ChoiceConfig {
	maleHair := []uint32{character.HairBlackCatalyst, character.HairBlackTopknot, character.HairBlackWind, character.HairBlackShaggyWax, character.HairBlackAcorn, character.HairBlackTheMoRawk, character.HairBlackAranCut, character.HairBlackTheCoco}
	femaleHair := []uint32{character.HairBlackStella, character.HairBlackRainbow, character.HairBlackAngelica, character.HairBlackChantelle, character.HairBlackFourtailBraids, character.HairBlackCrazyMedusa, character.HairBlackAranHair, character.HairBlackFullBangs}
	return care.RegularHairCare(coupon, maleHair, femaleHair, r.Initial)
}

func (r Brittany) ExperimentalStyleHair(coupon uint32) care.ChoiceConfig {
	maleHair := []uint32{character.HairBlackBuzz, character.HairBlackTopknot, character.HairBlackWind, character.HairBlackShaggyWax, character.HairBlackAcorn, character.HairBlackTheMoRawk, character.HairBlackAranCut, character.HairBlackTheCoco}
	femaleHair := []uint32{character.HairBlackStella, character.HairBlackAngelica, character.HairBlackChantelle, character.HairBlackFourtailBraids, character.HairSkinHead, character.HairBlackCrazyMedusa, character.HairBlackAranHair, character.HairBlackFullBangs}
	return care.ExperimentalHairCare(coupon, maleHair, femaleHair, r.Initial)
}
