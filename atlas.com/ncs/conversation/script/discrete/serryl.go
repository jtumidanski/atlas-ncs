package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Serryl is located in The Nautilus - Mid Floor - Hallway (120000200)
//TODO level confirm prompt
type Serryl struct {
}

func (r Serryl) NPCId() uint32 {
	return npc.Serryl
}

func (r Serryl) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return refine.NewGenericRefine(l, c, r.Hello(), r.Categories())
}

func (r Serryl) Hello() string {
	return "What? You want to make your own weapons and gloves? Seriously... it's tough to do it by yourself if you don't have experience... I'll help you out. I've been a pirate for 20 years, and for 20 years I have made various items for the crew here. It's easy for me."
}

func (r Serryl) Categories() []refine.ListItem {
	return []refine.ListItem{
		r.Knuckler(),
		r.Gun(),
		r.Gloves(),
	}
}

func (r Serryl) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Serryl) Knuckler() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.LeatherArms, " (Level limit: 15, Pirate)"), refine.Confirm(item.LeatherArms, r.LeatherArmsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DoubleTailKnuckler, " (Level limit: 20, Pirate)"), refine.Confirm(item.DoubleTailKnuckler, r.DoubleTailKnucklerRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.NormanGrip, " (Level limit: 25, Pirate)"), refine.Confirm(item.NormanGrip, r.NormanGripRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PrimeHands, " (Level limit: 30, Pirate)"), refine.Confirm(item.PrimeHands, r.PrimeHandsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SilverMaiden, " (Level limit: 35, Pirate)"), refine.Confirm(item.SilverMaiden, r.SilverMaidenRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.Neozard, " (Level limit: 40, Pirate)"), refine.Confirm(item.Neozard, r.NeozardRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.FuryClaw, " (Level limit: 50, Pirate)"), refine.Confirm(item.FuryClaw, r.FuryClawRequirements())),
	}
	prompt := refine.PromptCategory("As long as you bring in the materials required, I'll make you a fine Knuckler. Which Knuckler would you like to make?", choices)
	return refine.ListItem{ListText: "Make a Knuckler", SelectionState: prompt}
}

func (r Serryl) Gun() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DellingerSpecial, " (Level limit: 15, Pirate)"), refine.Confirm(item.DellingerSpecial, r.DellingerSpecialRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.TheNegotiator, " (Level limit: 20, Pirate)"), refine.Confirm(item.TheNegotiator, r.TheNegotiatorRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldenHook, " (Level limit: 25, Pirate)"), refine.Confirm(item.GoldenHook, r.GoldenHookRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.ColdMind, " (Level limit: 30, Pirate)"), refine.Confirm(item.ColdMind, r.ColdMindRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.ShootingStar, " (Level limit: 35, Pirate)"), refine.Confirm(item.ShootingStar, r.ShootingStarRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.LunarShooter, " (Level limit: 40, Pirate)"), refine.Confirm(item.LunarShooter, r.LunarShooterRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.MrRasfelt, " (Level limit: 50, Pirate)"), refine.Confirm(item.MrRasfelt, r.MrRasfeltRequirements())),
	}
	prompt := refine.PromptCategory("As long as you bring in the materials required, I'll make you a fine Gun. Which Gun would you like to make?", choices)
	return refine.ListItem{ListText: "Make a Gun", SelectionState: prompt}
}

func (r Serryl) Gloves() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.GreenLaggerHalfglove), refine.Confirm(item.GreenLaggerHalfglove, r.GreenLaggerHalfgloveRequirements())),
		r.CreateChoice(refine.ItemIdList(item.BrownLeatherArmourGlove), refine.Confirm(item.BrownLeatherArmourGlove, r.BrownLeatherArmourGloveRequirements())),
		r.CreateChoice(refine.ItemIdList(item.HardLeatherGlove), refine.Confirm(item.HardLeatherGlove, r.HardLeatherGloveRequirements())),
		r.CreateChoice(refine.ItemIdList(item.YellowTartis), refine.Confirm(item.YellowTartis, r.YellowTartisRequirements())),
		r.CreateChoice(refine.ItemIdList(item.BrownJewelled), refine.Confirm(item.BrownJewelled, r.BrownJewelledRequirements())),
		r.CreateChoice(refine.ItemIdList(item.BrownBarbee), refine.Confirm(item.BrownBarbee, r.BrownBarbeeRequirements())),
		r.CreateChoice(refine.ItemIdList(item.BrownRoyce), refine.Confirm(item.BrownRoyce, r.BrownRoyceRequirements())),
		r.CreateChoice(refine.ItemIdList(item.BlackSchult), refine.Confirm(item.BlackSchult, r.BlackSchultRequirements())),
	}
	prompt := refine.PromptCategory("As long as you bring in the materials required, I'll make you a fine glove. Which glove would you like to make?", choices)
	return refine.ListItem{ListText: "Make a pair of gloves", SelectionState: prompt}
}

func (r Serryl) LeatherArmsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 20}}, refine.SetCost(1000))
}

func (r Serryl) DoubleTailKnucklerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 1}, {ItemId: 4011000, Amount: 1}, {ItemId: 4000021, Amount: 10}, {ItemId: 4003000, Amount: 5}}, refine.SetCost(2000))
}

func (r Serryl) NormanGripRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 2}, {ItemId: 4011001, Amount: 1}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(5000))
}

func (r Serryl) PrimeHandsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 1}, {ItemId: 4011001, Amount: 1}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(15000))
}

func (r Serryl) SilverMaidenRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 2}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 30}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(30000))
}

func (r Serryl) NeozardRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 1}, {ItemId: 4011001, Amount: 1}, {ItemId: 4021000, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(50000))
}

func (r Serryl) FuryClawRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000039, Amount: 150}, {ItemId: 4011000, Amount: 1}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000030, Amount: 20}, {ItemId: 4000021, Amount: 20}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(100000))
}

func (r Serryl) DellingerSpecialRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 1}, {ItemId: 4003000, Amount: 5}, {ItemId: 4003001, Amount: 1}}, refine.SetCost(1000))
}

func (r Serryl) TheNegotiatorRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 1}, {ItemId: 4003000, Amount: 10}, {ItemId: 4003001, Amount: 5}, {ItemId: 4000021, Amount: 10}}, refine.SetCost(2000))
}

func (r Serryl) GoldenHookRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 2}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(5000))
}

func (r Serryl) ColdMindRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 10}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(15000))
}

func (r Serryl) ShootingStarRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 10}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 5}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(30000))
}

func (r Serryl) LunarShooterRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011004, Amount: 1}, {ItemId: 4011001, Amount: 2}, {ItemId: 4000021, Amount: 10}, {ItemId: 4003000, Amount: 20}}, refine.SetCost(50000))
}

func (r Serryl) MrRasfeltRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 1}, {ItemId: 4011004, Amount: 2}, {ItemId: 4011001, Amount: 4}, {ItemId: 4000030, Amount: 30}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(100000))
}

func (r Serryl) GreenLaggerHalfgloveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 15}, {ItemId: 4021003, Amount: 1}}, refine.SetCost(1000))
}

func (r Serryl) BrownLeatherArmourGloveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 35}}, refine.SetCost(8000))
}

func (r Serryl) HardLeatherGloveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 2}, {ItemId: 4000021, Amount: 20}}, refine.SetCost(15000))
}

func (r Serryl) YellowTartisRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021006, Amount: 2}, {ItemId: 4000021, Amount: 50}, {ItemId: 4003000, Amount: 10}}, refine.SetCost(25000))
}

func (r Serryl) BrownJewelledRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 3}, {ItemId: 4000021, Amount: 60}, {ItemId: 4003000, Amount: 15}}, refine.SetCost(30000))
}

func (r Serryl) BrownBarbeeRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4000021, Amount: 80}, {ItemId: 4011000, Amount: 3}, {ItemId: 4011001, Amount: 3}, {ItemId: 4003000, Amount: 25}}, refine.SetCost(40000))
}

func (r Serryl) BrownRoyceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011000, Amount: 3}, {ItemId: 4000021, Amount: 20}, {ItemId: 4000030, Amount: 40}, {ItemId: 4003000, Amount: 30}}, refine.SetCost(50000))
}

func (r Serryl) BlackSchultRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011007, Amount: 1}, {ItemId: 4021008, Amount: 1}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 50}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(70000))
}

func (r Serryl) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("All done. If you need anything else... Well, I'm not going anywhere.")
	return script.SendOk(l, c, m.String())
}

func (r Serryl) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You cannot afford to make this item.")
	return script.SendOk(l, c, m.String())
}

func (r Serryl) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText("Check and make sure you have all the necessary items to make this.")
		return script.SendOk(l, c, m.String())
	}
}

func (r Serryl) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Make sure your Equips inventory has room. I can't give you the item if your inventory is full, you know.")
	return script.SendOk(l, c, m.String())
}
