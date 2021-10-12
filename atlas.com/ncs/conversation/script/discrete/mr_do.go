package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// MrDo is located in Mu Lung - Mu Lung (250000000)
type MrDo struct {
}

func (r MrDo) NPCId() uint32 {
	return npc.MrDo
}

func (r MrDo) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsActive(l)(c.CharacterId, 3821) && !character.HasItem(l, span)(c.CharacterId, item.PeachTreeErbPouch) && !character.HasItem(l, span)(c.CharacterId, item.BookOnHerbalMedicine) && quest.IsCompleted(l)(c.CharacterId, 3830) {
		if character.CanHold(l)(c.CharacterId, item.PeachTreeErbPouch) {
			character.GainItem(l, span)(c.CharacterId, item.PeachTreeErbPouch, 1)
			m := message.NewBuilder().
				AddText("Oh, the boy wanted you to bring him a ").
				ShowItemName1(item.PeachTreeErbPouch).
				AddText("? No problem, I was on his debt anyway. Now, tell him I am repaying the debt, OK?")
			return script.SendOk(l, span, c, m.String())
		} else {
			m := message.NewBuilder().
				AddText("Oh, the boy wanted you to bring him a ").
				ShowItemName1(item.PeachTreeErbPouch).
				AddText("? Make room at your ETC inventory first.")
			return script.SendOk(l, span, c, m.String())
		}
	}
	return refine.NewGenericRefine(l, span, c, r.Hello(), r.Categories())
}

func (r MrDo) Hello() string {
	return "I am a man of many talents. Let me know what you'd like to do."
}

func (r MrDo) Categories() []refine.ListItem {
	return []refine.ListItem{
		r.Medicine(),
		r.Scroll(),
		r.Donate(),
	}
}

func (r MrDo) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, r.GenericRefinementConfig())
}

func (r MrDo) GenericRefinementConfig() refine.TerminalConfig {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
		StimulatorError:  refine.GenericStimulatorError,
	}
	return config
}

func (r MrDo) Medicine() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndImageList(item.CannedPeach), refine.HowMany(item.CannedPeach, r.CannedPeachRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.PeachJuice), refine.HowMany(item.PeachJuice, r.PeachJuiceRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.BellflowerMedicineSoup), refine.HowMany(item.BellflowerMedicineSoup, r.BellflowerMedicineSoupRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.PillOfTunnelVision), refine.HowMany(item.PillOfTunnelVision, r.PillOfTunnelVisionRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.PillOfIntelligence), refine.HowMany(item.PillOfIntelligence, r.PillOfIntelligenceRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.SlitheringBalm), refine.HowMany(item.SlitheringBalm, r.SlitheringBalmRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.AllCurePotion), refine.HowMany(item.AllCurePotion, r.AllCurePotionRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.PeachTreeErbPouch), refine.HowMany(item.PeachTreeErbPouch, r.PeachTreeErbPouchRequirements())),
	}
	prompt := refine.PromptCategory("What kind of medicine are you interested in making?", choices)
	confirmBook := r.ValidateBook(prompt)
	return refine.ListItem{ListText: "Make a medicine", SelectionState: confirmBook}
}

func (r MrDo) Scroll() refine.ListItem {
	//TODO MrDo will at a 10% rate award a 60% scroll rather than a 100%.
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.ScrollForOneHandedSwordForAttack100), refine.Confirm(item.ScrollForOneHandedSwordForAttack100, r.ScrollForOneHandedSwordForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForOneHandedAxeForAttack100), refine.Confirm(item.ScrollForOneHandedAxeForAttack100, r.ScrollForOneHandedAxeForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForOneHandedBWForAttack100), refine.Confirm(item.ScrollForOneHandedBWForAttack100, r.ScrollForOneHandedBWForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForDaggerForAttack100), refine.Confirm(item.ScrollForDaggerForAttack100, r.ScrollForDaggerForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForWandForMagicAttack100), refine.Confirm(item.ScrollForWandForMagicAttack100, r.ScrollForWandForMagicAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForStaffForMagicAttack100), refine.Confirm(item.ScrollForStaffForMagicAttack100, r.ScrollForStaffForMagicAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForTwoHandedSwordForAttack100), refine.Confirm(item.ScrollForTwoHandedSwordForAttack100, r.ScrollForTwoHandedSwordForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForTwoHandedAxeForAttack100), refine.Confirm(item.ScrollForTwoHandedAxeForAttack100, r.ScrollForTwoHandedAxeForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForTwoHandedBWForAttack100), refine.Confirm(item.ScrollForTwoHandedBWForAttack100, r.ScrollForTwoHandedBWForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForSpearForAttack100), refine.Confirm(item.ScrollForSpearForAttack100, r.ScrollForSpearForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForPoleArmForAttack100), refine.Confirm(item.ScrollForPoleArmForAttack100, r.ScrollForPoleArmForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForBowForAttack100), refine.Confirm(item.ScrollForBowForAttack100, r.ScrollForBowForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForCrossbowForATttack100), refine.Confirm(item.ScrollForCrossbowForATttack100, r.ScrollForCrossbowForATttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForClawForAttack100), refine.Confirm(item.ScrollForClawForAttack100, r.ScrollForClawForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForKnuckleForAttack100), refine.Confirm(item.ScrollForKnuckleForAttack100, r.ScrollForKnuckleForAttack100Requirements())),
		r.CreateChoice(refine.ItemIdList(item.ScrollForGunForAttack100), refine.Confirm(item.ScrollForGunForAttack100, r.ScrollForGunForAttack100Requirements())),
	}
	prompt := refine.PromptCategory("What kind of scrolls are you interested in making?", choices)
	return refine.ListItem{ListText: "Make a scroll", SelectionState: prompt}
}

type donationChoice struct {
	itemId uint32
	lower  uint32
	upper  uint32
}

func (r MrDo) Donate() refine.ListItem {
	choices := []donationChoice{
		{itemId: item.Acorn, lower: 7, upper: 7},
		{itemId: item.Thimble, lower: 7, upper: 7},
		{itemId: item.NeedlePouch, lower: 7, upper: 8},
		{itemId: item.NeckiFlower, lower: 10, upper: 10},
		{itemId: item.NeckiSwimmingCap, lower: 11, upper: 11},
		{itemId: item.BrokenPieceOfPot, lower: 8, upper: 8},
		{itemId: item.GinsengBoiledWater, lower: 7, upper: 8},
		{itemId: item.StrawDoll, lower: 7, upper: 9},
		{itemId: item.WoodenDoll, lower: 7, upper: 8},
		{itemId: item.Bellflower, lower: 9, upper: 9},
		{itemId: item.HundredYearOldBellflower, lower: 10, upper: 10},
		{itemId: item.OldPaper, lower: 10, upper: 11},
		{itemId: item.YellowBelt, lower: 11, upper: 11},
		{itemId: item.BrokenDeerHorn, lower: 11, upper: 12},
		{itemId: item.RedBelt, lower: 13, upper: 13},
		{itemId: item.PeachSeed, lower: 13, upper: 13},
		{itemId: item.MrAllisLeather, lower: 14, upper: 14},
		{itemId: item.CatDoll, lower: 15, upper: 15},
		{itemId: item.MarkOfThePirate, lower: 15, upper: 16},
		{itemId: item.CaptainsHat, lower: 17, upper: 17},
	}
	return refine.ListItem{ListText: "Donate medicine ingredients", SelectionState: r.PromptDonation(choices)}
}

func (r MrDo) ValidateBook(prompt script.StateProducer) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasItem(l, span)(c.CharacterId, item.BookOnHerbalMedicine) {
			m := message.NewBuilder().AddText("If you want to make a medicine, you must study the Book on Herbal Medicine first. Nothing is more dangerous than practicing a medicine without proper knowledge.")
			return script.SendOk(l, span, c, m.String())
		}
		return prompt(l, span, c)
	}
}

func (r MrDo) CannedPeachRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 2022116, Amount: 3}})
}

func (r MrDo) PeachJuiceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 2022116, Amount: 3}})
}

func (r MrDo) BellflowerMedicineSoupRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000281, Amount: 10}, {ItemId: 4000293, Amount: 10}}, refine.SetCost(910))
}

func (r MrDo) PillOfTunnelVisionRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000276, Amount: 20}, {ItemId: 2002005, Amount: 1}}, refine.SetCost(950))
}

func (r MrDo) PillOfIntelligenceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000288, Amount: 20}, {ItemId: 4000292, Amount: 20}}, refine.SetCost(1940))
}

func (r MrDo) SlitheringBalmRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000295, Amount: 10}}, refine.SetCost(600))
}

func (r MrDo) AllCurePotionRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 2022131, Amount: 1}, {ItemId: 2022132, Amount: 1}}, refine.SetCost(700))
}

func (r MrDo) PeachTreeErbPouchRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000286, Amount: 20}, {ItemId: 4000287, Amount: 20}, {ItemId: 4000293, Amount: 20}}, refine.SetCost(1000))
}

func (r MrDo) ScrollForOneHandedSwordForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForOneHandedAxeForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForOneHandedBWForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForDaggerForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForWandForMagicAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForStaffForMagicAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForTwoHandedSwordForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForTwoHandedAxeForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForTwoHandedBWForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForSpearForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForPoleArmForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForBowForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForCrossbowForATttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForClawForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForKnuckleForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) ScrollForGunForAttack100Requirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4001124, Amount: 100}, {ItemId: 4010001, Amount: 10}})
}

func (r MrDo) PromptDonation(choices []donationChoice) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("So you wish to donate some medicine ingredients? This is great news! Donations will be accepted in the unit of ").
			BlueText().AddText("100").
			BlackText().AddText(". The donator will receive a marble that enables one to make a scroll. Which of these would you like to donate?")
		for i, choice := range choices {
			m = m.OpenItem(i).ShowItemImage2(choice.itemId).AddText(" ").ShowItemName1(choice.itemId).CloseItem().NewLine()
		}
		return script.SendListSelection(l, span, c, m.String(), r.ConfirmDonation(choices))
	}
}

func (r MrDo) ConfirmDonation(choices []donationChoice) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		if selection < 0 || int(selection) >= len(choices) {
			return script.Exit()
		}
		return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
			choice := choices[selection]
			m := message.NewBuilder().
				AddText("Are you sure you want to donate ").
				BlueText().AddText("100").
				ShowItemName1(choice.itemId).
				BlackText().AddText("?")
			return script.SendYesNo(l, span, c, m.String(), r.ProcessDonation(choice), r.VeryBusy)
		}
	}
}

func (r MrDo) VeryBusy(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Oh, talk to me when you have decided what you want from me. I am very busy right now.")
	return script.SendOk(l, span, c, m.String())
}

func (r MrDo) ProcessDonation(choice donationChoice) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasItems(l, span)(c.CharacterId, choice.itemId, 100) {
			return r.MissingSomething(choice.itemId)(l, span, c)
		}

		amount := choice.lower + uint32(rand.Intn(int(choice.upper-choice.lower)))
		if !character.CanHoldAll(l)(c.CharacterId, item.DrDosMarble, amount) {
			return r.MakeRoom(l, span, c)
		}

		character.GainItem(l, span)(c.CharacterId, choice.itemId, -100)
		character.GainItem(l, span)(c.CharacterId, item.DrDosMarble, int32(amount))
		return r.DonationSuccess(l, span, c)
	}
}

func (r MrDo) DonationSuccess(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Thank you for your donation.")
	return script.SendOk(l, span, c, m.String())
}

func (r MrDo) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText("You are lacking ingredients.")
		return script.SendOk(l, span, c, m.String())
	}
}

func (r MrDo) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You are lacking space in your USE inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r MrDo) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("There you go!")
	return script.SendOk(l, span, c, m.String())
}

func (r MrDo) CannotAfford(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You cannot afford to create this medicine.")
	return script.SendOk(l, span, c, m.String())
}
