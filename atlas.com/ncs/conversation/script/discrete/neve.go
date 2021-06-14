package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Neve is located in Orbis - Orbis Park (200000200)
type Neve struct {
}

func (r Neve) NPCId() uint32 {
	return npc.Neve
}

func (r Neve) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return refine.NewGenericRefine(l, c, r.Hello(), r.Categories())
}

func (r Neve) Hello() string {
	return "Hello there. I'm Orbis' number one glove maker. Would you like me to make you something?"
}

func (r Neve) Categories() []refine.ListItem {
	return []refine.ListItem{
		r.Warrior(),
		r.Bowman(),
		r.Magician(),
		r.Thief(),
	}
}

func (r Neve) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Neve) Warrior() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeHusk, " - Warrior Lv. 70"), refine.Confirm(item.BronzeHusk, r.BronzeHuskRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MithrilHusk, " - Warrior Lv. 70"), refine.Confirm(item.MithrilHusk, r.MithrilHuskRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkHusk, " - Warrior Lv. 70"), refine.Confirm(item.DarkHusk, r.DarkHuskRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SapphireEmperor, " - Warrior Lv. 80"), refine.Confirm(item.SapphireEmperor, r.SapphireEmperorRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.EmeraldEmperor, " - Warrior Lv. 80"), refine.Confirm(item.EmeraldEmperor, r.EmeraldEmperorRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodEmperor, " - Warrior Lv. 80"), refine.Confirm(item.BloodEmperor, r.BloodEmperorRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkEmperor, " - Warrior Lv. 80"), refine.Confirm(item.DarkEmperor, r.DarkEmperorRequirements())),
	}
	prompt := refine.PromptCategory("Warrior glove? Okay, then which one?", choices)
	return refine.ListItem{ListText: "Create or upgrade a Warrior glove", SelectionState: prompt}
}

func (r Neve) Bowman() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueEyes, " - Bowman Lv. 70"), refine.Confirm(item.BlueEyes, r.BlueEyesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldEyes, " - Bowman Lv. 70"), refine.Confirm(item.GoldEyes, r.GoldEyesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkEyes, " - Bowman Lv. 70"), refine.Confirm(item.DarkEyes, r.DarkEyesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedCordon, " - Bowman Lv. 80"), refine.Confirm(item.RedCordon, r.RedCordonRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueCordon, " - Bowman Lv. 80"), refine.Confirm(item.BlueCordon, r.BlueCordonRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenCordon, " - Bowman Lv. 80"), refine.Confirm(item.GreenCordon, r.GreenCordonRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkCordon, " - Bowman Lv. 80"), refine.Confirm(item.DarkCordon, r.DarkCordonRequirements())),
	}
	prompt := refine.PromptCategory("Bowman glove? Okay, then which one?", choices)
	return refine.ListItem{ListText: "Create or upgrade a Bowman glove", SelectionState: prompt}
}

func (r Neve) Magician() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownLorin, " - Magician Lv. 70"), refine.Confirm(item.BrownLorin, r.BrownLorinRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueLorin, " - Magician Lv. 70"), refine.Confirm(item.BlueLorin, r.BlueLorinRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkLorin, " - Magician Lv. 70"), refine.Confirm(item.DarkLorin, r.DarkLorinRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenClarity, " - Magician Lv. 80"), refine.Confirm(item.GreenClarity, r.GreenClarityRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueClarity, " - Magician Lv. 80"), refine.Confirm(item.BlueClarity, r.BlueClarityRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkClarity, " - Magician Lv. 80"), refine.Confirm(item.DarkClarity, r.DarkClarityRequirements())),
	}
	prompt := refine.PromptCategory("Magician glove? Okay, then which one?", choices)
	return refine.ListItem{ListText: "Create or upgrade a Magician glove", SelectionState: prompt}
}

func (r Neve) Thief() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BronzeRover, " - Thief Lv. 70"), refine.Confirm(item.BronzeRover, r.BronzeRoverRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverRover, " - Thief Lv. 70"), refine.Confirm(item.SilverRover, r.SilverRoverRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldRover, " - Thief Lv. 70"), refine.Confirm(item.GoldRover, r.GoldRoverRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenLarceny, " - Thief Lv. 80"), refine.Confirm(item.GreenLarceny, r.GreenLarcenyRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleLarceny, " - Thief Lv. 80"), refine.Confirm(item.PurpleLarceny, r.PurpleLarcenyRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkLarceny, " - Thief Lv. 80"), refine.Confirm(item.DarkLarceny, r.DarkLarcenyRequirements())),
	}
	prompt := refine.PromptCategory("Thief glove? Okay, then which one?", choices)
	return refine.ListItem{ListText: "Create or upgrade a Thief glove", SelectionState: prompt}
}

func (r Neve) BronzeHuskRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005000, Amount: 2}, {ItemId: 4011000, Amount: 8}, {ItemId: 4011006, Amount: 3}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(90000))
}

func (r Neve) MithrilHuskRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082103, Amount: 1}, {ItemId: 4011002, Amount: 6}, {ItemId: 4021006, Amount: 4}}, refine.SetCost(90000))
}

func (r Neve) DarkHuskRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082103, Amount: 1}, {ItemId: 4021006, Amount: 8}, {ItemId: 4021008, Amount: 3}}, refine.SetCost(100000))
}

func (r Neve) SapphireEmperorRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005000, Amount: 2}, {ItemId: 4005002, Amount: 1}, {ItemId: 4021005, Amount: 8}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(100000))
}

func (r Neve) EmeraldEmperorRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082114, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4005002, Amount: 1}, {ItemId: 4021003, Amount: 7}}, refine.SetCost(110000))
}

func (r Neve) BloodEmperorRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082114, Amount: 1}, {ItemId: 4005002, Amount: 3}, {ItemId: 4021000, Amount: 8}}, refine.SetCost(110000))
}

func (r Neve) DarkEmperorRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082114, Amount: 1}, {ItemId: 4005000, Amount: 2}, {ItemId: 4005002, Amount: 1}, {ItemId: 4021008, Amount: 4}}, refine.SetCost(120000))
}

func (r Neve) BlueEyesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 2}, {ItemId: 4021005, Amount: 8}, {ItemId: 4011004, Amount: 3}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(90000))
}

func (r Neve) GoldEyesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082106, Amount: 1}, {ItemId: 4021006, Amount: 5}, {ItemId: 4011006, Amount: 3}}, refine.SetCost(90000))
}

func (r Neve) DarkEyesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082106, Amount: 1}, {ItemId: 4021007, Amount: 2}, {ItemId: 4021008, Amount: 3}}, refine.SetCost(100000))
}

func (r Neve) RedCordonRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 2}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021000, Amount: 8}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(100000))
}

func (r Neve) BlueCordonRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082109, Amount: 1}, {ItemId: 4005002, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021005, Amount: 7}}, refine.SetCost(110000))
}

func (r Neve) GreenCordonRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082109, Amount: 1}, {ItemId: 4005002, Amount: 1}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021003, Amount: 7}}, refine.SetCost(110000))
}

func (r Neve) DarkCordonRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082109, Amount: 1}, {ItemId: 4005002, Amount: 2}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021008, Amount: 4}}, refine.SetCost(120000))
}

func (r Neve) BrownLorinRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005001, Amount: 2}, {ItemId: 4011000, Amount: 6}, {ItemId: 4011004, Amount: 6}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(90000))
}

func (r Neve) BlueLorinRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082098, Amount: 1}, {ItemId: 4021002, Amount: 6}, {ItemId: 4021007, Amount: 2}}, refine.SetCost(90000))
}

func (r Neve) DarkLorinRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082098, Amount: 1}, {ItemId: 4021008, Amount: 3}, {ItemId: 4011006, Amount: 3}}, refine.SetCost(100000))
}

func (r Neve) GreenClarityRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005001, Amount: 2}, {ItemId: 4005003, Amount: 1}, {ItemId: 4021003, Amount: 8}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(100000))
}

func (r Neve) BlueClarityRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082121, Amount: 1}, {ItemId: 4005001, Amount: 1}, {ItemId: 4005003, Amount: 1}, {ItemId: 4021005, Amount: 7}}, refine.SetCost(110000))
}

func (r Neve) DarkClarityRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082121, Amount: 1}, {ItemId: 4005001, Amount: 2}, {ItemId: 4005003, Amount: 1}, {ItemId: 4021008, Amount: 4}}, refine.SetCost(120000))
}

func (r Neve) BronzeRoverRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005003, Amount: 2}, {ItemId: 4011000, Amount: 6}, {ItemId: 4011003, Amount: 6}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(90000))
}

func (r Neve) SilverRoverRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082095, Amount: 1}, {ItemId: 4011004, Amount: 6}, {ItemId: 4021007, Amount: 2}}, refine.SetCost(90000))
}

func (r Neve) GoldRoverRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082095, Amount: 1}, {ItemId: 4021007, Amount: 3}, {ItemId: 4011006, Amount: 3}}, refine.SetCost(100000))
}

func (r Neve) GreenLarcenyRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005003, Amount: 2}, {ItemId: 4005002, Amount: 1}, {ItemId: 4011002, Amount: 8}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(100000))
}

func (r Neve) PurpleLarcenyRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082118, Amount: 1}, {ItemId: 4005003, Amount: 1}, {ItemId: 4005002, Amount: 1}, {ItemId: 4021001, Amount: 7}}, refine.SetCost(110000))
}

func (r Neve) DarkLarcenyRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1082118, Amount: 1}, {ItemId: 4005003, Amount: 2}, {ItemId: 4005002, Amount: 1}, {ItemId: 4021000, Amount: 8}}, refine.SetCost(120000))
}

func (r Neve) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Done. If you need anything else, just ask again.")
	return script.SendOk(l, c, m.String())
}

func (r Neve) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm afraid you cannot afford my services.")
	return script.SendOk(l, c, m.String())
}

func (r Neve) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText("I'm afraid that substitute items are unacceptable, if you want your gloves made properly.")
		return script.SendOk(l, c, m.String())
	}
}

func (r Neve) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, c, m.String())
}
