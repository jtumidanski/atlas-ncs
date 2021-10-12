package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// BigHeadward is located in Victoria Road - Henesys Hair Salon (100000104)
type BigHeadward struct {
}

func (r BigHeadward) NPCId() uint32 {
	return npc.BigHeadward
}

func (r BigHeadward) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.ProvidedCare())(l, span, c)
}

func (r BigHeadward) Hello() string {
	return message.NewBuilder().
		AddText("Hi, I'm ").
		ShowNPC(npc.BigHeadward).
		AddText(", the most charming and stylish stylist around. If you're looking for the best looking hairdos around, look no further!").
		String()
}

func (r BigHeadward) ProvidedCare() []care.ChoiceConfig {
	return []care.ChoiceConfig{r.RoyalStyleHair(), r.SpecialRoyalStyleHair()}
}

func (r BigHeadward) RoyalStyleHair() care.ChoiceConfig {
	hairStyle := message.NewBuilder().
		AddText("If you use this REGULAR coupon, your hair may transform into a random new look...do you still want to do it using ").
		BlueText().ShowItemName1(item.RoyalHairCoupon).
		BlackText().AddText(", I will do it anyways for you. But don't forget, it will be random!").
		String()

	maleHair := []uint32{character.HairBlackFoilPerm, character.HairBlackMetrosexual, character.HairBlackMohecanShaggyDo, character.HairBlackTristan, character.HairBlackMessySpike}
	femaleHair := []uint32{character.HairBlackMonica, character.HairBlackCaspia, character.HairBlackRose, character.HairBlackTheHoneybun, character.HairBlackPrincessa}
	next := care.WarnRandomStyle(hairStyle, item.RoyalHairCoupon, maleHair, femaleHair, care.SetHair, r.Initial)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(item.RoyalHairCoupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}

func (r BigHeadward) SpecialRoyalStyleHair() care.ChoiceConfig {
	hairStyle := "Using the SPECIAL coupon you can choose the style your hair will become. Pick the style that best provides you delight..."
	maleHair := []uint32{character.HairZeta, character.HairAllBack, character.HairMilitaryBuzzcut, character.HairMohawk, character.HairBlueFantasy, character.HairBlackBabbyCut, character.HairBlackGrandLionman, character.HairBlackMetroMan, character.HairBlackBowlingBall, character.HairBlackCornrow, character.HairBlackEasternMystery, character.HairBlackBoyBandCut, character.HairBlackVolumeCut}
	femaleHair := []uint32{character.HairBlackFrancesca, character.HairBlackJolie, character.HairBlackMinnie, character.HairBlackZessica, character.HairBlackGrace, character.HairBlackLagunaBeach, character.HairBlackShortTwinTails, character.HairBlackCLHair, character.HairBlackSpunkyDo, character.HairBlackVintageFlip, character.HairBlackPalmTreeHair}
	choiceSupplier := care.HairStyleChoices(maleHair, femaleHair)

	special := care.ProcessCoupon(item.SpecialRoyalHairCoupon, care.SetHair, care.SetSingleUse(true))
	next := care.ShowChoices(hairStyle, choiceSupplier, special)
	return care.NewChoiceConfig(next, care.HairStyleCouponListText(item.SpecialRoyalHairCoupon), care.HairStyleCouponMissing(), care.HairStyleEnjoy())
}
