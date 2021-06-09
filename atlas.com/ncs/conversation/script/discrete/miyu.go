package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/care"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Miyu is located in  Ludibrium - Ludibrium Hair Salon (220000004)
type Miyu struct {
}

func (r Miyu) NPCId() uint32 {
	return npc.Miyu
}

func (r Miyu) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return care.NewGenericCare(r.Hello(), []care.ChoiceConfig{r.StyleHair(), care.ColorCareChoice(item.LudibriumHairColorCouponVIP)})(l, c)
}

func (r Miyu) Hello() string {
	return message.NewBuilder().
		AddText("Welcome, welcome, welcome to the Ludibrium Hair Salon! Do you, by any chance, have a ").
		BlueText().ShowItemName1(item.LudibriumHairStyleCouponVIP).
		BlackText().AddText(" or a ").
		BlueText().ShowItemName1(item.LudibriumHairColorCouponVIP).
		BlackText().AddText(" If so, how about letting me take care of your hair? Please choose what you want to do with it...").
		String()
}

func (r Miyu) StyleHair() care.ChoiceConfig {
	hairStyle := care.StylePrompt(item.LudibriumHairStyleCouponVIP)
	maleHair := []uint32{30160, 30190, 30250, 30640, 30660, 30840, 30870, 30990}
	femaleHair := []uint32{31270, 31290, 31550, 31680, 31810, 31830, 31840, 31870}

	vip := care.ProcessCoupon(item.LudibriumHairStyleCouponVIP, care.SetHair, care.SetSingleUse(true))
	membership := care.ProcessCoupon(item.LudibriumHairMembership, care.SetHair, care.SetSingleUse(false), care.SetFailFunction(vip))
	next := care.ShowChoices(hairStyle, care.HairStyleChoices(maleHair, femaleHair), membership)
	return care.NewChoiceConfig(next, care.SetListText(care.StyleListEntry(item.LudibriumHairStyleCouponVIP)))
}
