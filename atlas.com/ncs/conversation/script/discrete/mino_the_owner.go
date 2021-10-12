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

// MinoTheOwner is located in Orbis Park - Orbis Hair Salon (200000202)
type MinoTheOwner struct {
}

func (r MinoTheOwner) NPCId() uint32 {
	return npc.MinoTheOwner
}

func (r MinoTheOwner) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), r.CareOptions())(l, span, c)
}

func (r MinoTheOwner) CareOptions() []care.ChoiceConfig {
	return []care.ChoiceConfig{
		r.StyleHair(item.OrbisHairStyleCouponVIP, item.OrbisHairMembershipCoupon),
		care.ColorCareChoice(item.OrbisHairColorCouponVIP),
	}
}

func (r MinoTheOwner) Hello() string {
	return message.NewBuilder().
		AddText("Hello I'm Mino. If you have either a ").
		BlueText().ShowItemName1(item.OrbisHairStyleCouponVIP).
		BlackText().AddText(", then please let me take care of your hair. Choose what you want to do with it.").
		String()
}

func (r MinoTheOwner) StyleHair(coupon uint32, membershipCoupon uint32) care.ChoiceConfig {
	maleHair := []uint32{character.HairBlackFoilPerm, character.HairBlackMetrosexual, character.HairBlackMohecanShaggyDo, character.HairBlackTristan, character.HairBlackMessySpike}
	femaleHair := []uint32{character.HairBlackMonica, character.HairBlackCaspia, character.HairBlackRose, character.HairBlackTheHoneybun, character.HairBlackPrincessa}
	return care.VIPHairCareWithMembership(coupon, membershipCoupon, maleHair, femaleHair)
}
