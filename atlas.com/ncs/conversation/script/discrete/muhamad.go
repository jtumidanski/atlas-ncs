package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Muhamad is located in Ariant - The Town of Ariant (260000200)
type Muhamad struct {
}

func (r Muhamad) NPCId() uint32 {
	return npc.Muhamad
}

func (r Muhamad) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Are you here to refine the ores of a mineral or a jewel? It doesn't matter how many ores you have, if you don't have them refined by a master like me, then they won't see the light of day. What do you think, do you want to refine them right now?")
	return script.SendYesNo(l, c, m.String(), r.What, r.ComeBack)
}

func (r Muhamad) ComeBack(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If you aren't in a hurry, then please come back in a bit. As you can see, there's so much work going on right now that I can't possibly give them to you on time.")
	return script.SendOk(l, c, m.String())
}

func (r Muhamad) What(l logrus.FieldLogger, c script.Context) script.State {
	return refine.NewGenericRefine(l, c, r.Hello(), r.Categories())
}

func (r Muhamad) Hello() string {
	return "I like your attitude! Let's just take care of this right now. What kind of ores would you like to refine?"
}

func (r Muhamad) Categories() []refine.ListItem {
	return []refine.ListItem{
		r.Mineral(),
		r.Jewel(),
		r.Crystal(),
	}
}

func (r Muhamad) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Muhamad) Mineral() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.BronzePlate), refine.HowMany(item.BronzePlate, r.BronzePlateRequirements())),
		r.CreateChoice(refine.ItemIdList(item.SteelPlate), refine.HowMany(item.SteelPlate, r.SteelPlateRequirements())),
		r.CreateChoice(refine.ItemIdList(item.MithrilPlate), refine.HowMany(item.MithrilPlate, r.MithrilPlateRequirements())),
		r.CreateChoice(refine.ItemIdList(item.AdamantiumPlate), refine.HowMany(item.AdamantiumPlate, r.AdamantiumPlateRequirements())),
		r.CreateChoice(refine.ItemIdList(item.SilverPlate), refine.HowMany(item.SilverPlate, r.SilverPlateRequirements())),
		r.CreateChoice(refine.ItemIdList(item.OrihalconPlate), refine.HowMany(item.OrihalconPlate, r.OrihalconPlateRequirements())),
		r.CreateChoice(refine.ItemIdList(item.GoldPlate), refine.HowMany(item.GoldPlate, r.GoldPlateRequirements())),
		r.CreateChoice(refine.ItemIdList(item.Lidium), refine.HowMany(item.Lidium, r.LidiumRequirements())),
	}
	prompt := refine.PromptCategory("Which mineral would you like to refine?", choices)
	return refine.ListItem{ListText: "Refine a mineral ore", SelectionState: prompt}
}

func (r Muhamad) Jewel() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.Garnet), refine.HowMany(item.Garnet, r.GarnetRequirements())),
		r.CreateChoice(refine.ItemIdList(item.Amethyst), refine.HowMany(item.Amethyst, r.AmethystRequirements())),
		r.CreateChoice(refine.ItemIdList(item.AquaMarine), refine.HowMany(item.AquaMarine, r.AquamarineRequirements())),
		r.CreateChoice(refine.ItemIdList(item.Emerald), refine.HowMany(item.Emerald, r.EmeraldRequirements())),
		r.CreateChoice(refine.ItemIdList(item.Opal), refine.HowMany(item.Opal, r.OpalRequirements())),
		r.CreateChoice(refine.ItemIdList(item.Sapphire), refine.HowMany(item.Sapphire, r.SapphireRequirements())),
		r.CreateChoice(refine.ItemIdList(item.Topaz), refine.HowMany(item.Topaz, r.TopazRequirements())),
		r.CreateChoice(refine.ItemIdList(item.Diamond), refine.HowMany(item.Diamond, r.DiamondRequirements())),
		r.CreateChoice(refine.ItemIdList(item.BlackCrystal), refine.HowMany(item.BlackCrystal, r.BlackCrystalRequirements())),
	}
	prompt := refine.PromptCategory("Which jewel would you like to refine?", choices)
	return refine.ListItem{ListText: "Refine jewel ores", SelectionState: prompt}
}

func (r Muhamad) Crystal() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.PowerCrystal), refine.HowMany(item.PowerCrystal, r.PowerCrystalRequirements())),
		r.CreateChoice(refine.ItemIdList(item.WisdomCrystal), refine.HowMany(item.WisdomCrystal, r.WisdomCrystalRequirements())),
		r.CreateChoice(refine.ItemIdList(item.DEXCrystal), refine.HowMany(item.DEXCrystal, r.DEXCrystalRequirements())),
		r.CreateChoice(refine.ItemIdList(item.LUKCrystal), refine.HowMany(item.LUKCrystal, r.LUKCrystalRequirements())),
	}
	prompt := refine.PromptCategory("A crystal? That's a rare item indeed. Don't worry, I can refine it just as well as others. Which crystal would you like to refine?", choices)
	return refine.ListItem{ListText: "Refine crystal ores", SelectionState: prompt}
}

func (r Muhamad) BronzePlateRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4010000, Amount: 10}}, refine.SetCost(270))
}

func (r Muhamad) SteelPlateRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4010001, Amount: 10}}, refine.SetCost(270))
}

func (r Muhamad) MithrilPlateRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4010002, Amount: 10}}, refine.SetCost(270))
}

func (r Muhamad) AdamantiumPlateRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4010003, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) SilverPlateRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4010004, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) OrihalconPlateRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4010005, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) GoldPlateRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4010006, Amount: 10}}, refine.SetCost(720))
}

func (r Muhamad) LidiumRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4010007, Amount: 10}}, refine.SetCost(270))
}

func (r Muhamad) GarnetRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4020000, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) AmethystRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4020001, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) AquamarineRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4020002, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) EmeraldRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4020003, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) OpalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4020004, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) SapphireRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4020005, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) TopazRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4020006, Amount: 10}}, refine.SetCost(450))
}

func (r Muhamad) DiamondRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4020007, Amount: 10}}, refine.SetCost(900))
}

func (r Muhamad) BlackCrystalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4020008, Amount: 10}}, refine.SetCost(2700))
}

func (r Muhamad) PowerCrystalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4004000, Amount: 10}}, refine.SetCost(4500))
}

func (r Muhamad) WisdomCrystalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4004001, Amount: 10}}, refine.SetCost(4500))
}

func (r Muhamad) DEXCrystalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4004002, Amount: 10}}, refine.SetCost(4500))
}

func (r Muhamad) LUKCrystalRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4004003, Amount: 10}}, refine.SetCost(4500))
}

func (r Muhamad) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("There, finished. What do you think, a piece of art, isn't it? Well, if you need anything else, you know where to find me.")
	return script.SendOk(l, c, m.String())
}

func (r Muhamad) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm afraid you cannot afford my services.")
	return script.SendOk(l, c, m.String())
}

func (r Muhamad) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText("Please check and see if you have all the necessary items with you.")
		return script.SendOk(l, c, m.String())
	}
}

func (r Muhamad) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm afraid you are short in inventory slots for this.")
	return script.SendOk(l, c, m.String())
}
