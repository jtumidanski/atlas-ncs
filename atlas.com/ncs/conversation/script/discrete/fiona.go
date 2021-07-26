package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// Fiona is located in Phantom Forest - Dead Man's Gorge (610010004)
type Fiona struct {
}

func (r Fiona) NPCId() uint32 {
	return npc.Fiona
}

func (r Fiona) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !quest.IsCompleted(l)(c.CharacterId, 8225) {
		m := message.NewBuilder().AddText("Step aside, novice, we're doing business here.")
		return script.SendOk(l, c, m.String())
	}
	return refine.NewGenericRefine(l, c, r.Hello(), r.Categories())
}

func (r Fiona) Hello() string {
	return "Hey, partner! If you have the right goods, I can turn it into something very nice..."
}

func (r Fiona) Categories() []refine.ListItem {
	return []refine.ListItem{
		r.Forging(),
		r.Upgrading(),
	}
}

func (r Fiona) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Fiona) Forging() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.BalancedFury), refine.Confirm(item.BalancedFury, r.BalancedFuryRequirements())),
		r.CreateChoice(refine.ItemIdList(item.CrimsonArcanon), refine.Confirm(item.CrimsonArcanon, r.CrimsonArcanonRequirements())),
		r.CreateChoice(refine.ItemIdList(item.CrimsonArcglaive), refine.Confirm(item.CrimsonArcglaive, r.CrimsonArcglaiveRequirements())),
		r.CreateChoice(refine.ItemIdList(item.CrimsonArclancer), refine.Confirm(item.CrimsonArclancer, r.CrimsonArclancerRequirements())),
	}
	prompt := refine.PromptCategory("So, what kind of weapon would you like me to forge?", choices)
	return refine.ListItem{ListText: "Weapon Forging", SelectionState: prompt}
}

func (r Fiona) Upgrading() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdList(item.DawnRavensClaw), refine.Confirm(item.DawnRavensClaw, r.DawnRavensClawRequirements())),
		r.CreateChoice(refine.ItemIdList(item.NightRavensClaw), refine.Confirm(item.NightRavensClaw, r.NightRavensClawRequirements())),
		r.CreateChoice(refine.ItemIdList(item.DuskRavensClaw), refine.Confirm(item.DuskRavensClaw, r.DuskRavensClawRequirements())),
		r.CreateChoice(refine.ItemIdList(item.DawnRavensBeak), refine.Confirm(item.DawnRavensBeak, r.DawnRavensBeakRequirements())),
		r.CreateChoice(refine.ItemIdList(item.NightRavensBeak), refine.Confirm(item.NightRavensBeak, r.NightRavensBeakRequirements())),
		r.CreateChoice(refine.ItemIdList(item.DuskRavensBeak), refine.Confirm(item.DuskRavensBeak, r.DuskRavensBeakRequirements())),
		r.CreateChoice(refine.ItemIdList(item.DawnRavensEye), refine.Confirm(item.DawnRavensEye, r.DawnRavensEyeRequirements())),
		r.CreateChoice(refine.ItemIdList(item.NightRavensEye), refine.Confirm(item.NightRavensEye, r.NightRavensEyeRequirements())),
		r.CreateChoice(refine.ItemIdList(item.DuskRavensEye), refine.Confirm(item.DuskRavensEye, r.DuskRavensEyeRequirements())),
		r.CreateChoice(refine.ItemIdList(item.DawnRavensWing), refine.Confirm(item.DawnRavensWing, r.DawnRavensWingRequirements())),
		r.CreateChoice(refine.ItemIdList(item.NightRavensWing), refine.Confirm(item.NightRavensWing, r.NightRavensWingRequirements())),
		r.CreateChoice(refine.ItemIdList(item.DuskRavensWing), refine.Confirm(item.DuskRavensWing, r.DuskRavensWingRequirements())),
	}
	prompt := refine.PromptCategory("An upgraded weapon? Of course, but note that upgrades won't carry over to the new item...", choices)
	return refine.ListItem{ListText: "Weapon Upgrading", SelectionState: prompt}
}

func (r Fiona) BalancedFuryRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032015, Amount: 1}, {ItemId: 4032016, Amount: 1}, {ItemId: 4032017, Amount: 1}, {ItemId: 4021008, Amount: 100}, {ItemId: 4032005, Amount: 30}}, refine.SetCost(70000))
}

func (r Fiona) CrimsonArcanonRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032016, Amount: 1}, {ItemId: 4032017, Amount: 1}, {ItemId: 4032004, Amount: 400}, {ItemId: 4032005, Amount: 10}, {ItemId: 4032012, Amount: 30}, {ItemId: 4005001, Amount: 4}}, refine.SetCost(70000))
}

func (r Fiona) CrimsonArcglaiveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032015, Amount: 1}, {ItemId: 4032017, Amount: 1}, {ItemId: 4032004, Amount: 500}, {ItemId: 4032005, Amount: 40}, {ItemId: 4032012, Amount: 20}, {ItemId: 4005000, Amount: 4}}, refine.SetCost(70000))
}

func (r Fiona) CrimsonArclancerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032015, Amount: 1}, {ItemId: 4032016, Amount: 1}, {ItemId: 4032004, Amount: 300}, {ItemId: 4032005, Amount: 75}, {ItemId: 4032012, Amount: 10}, {ItemId: 4005002, Amount: 4}}, refine.SetCost(70000))
}

func (r Fiona) DawnRavensClawRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032017, Amount: 1}, {ItemId: 4005001, Amount: 10}, {ItemId: 4021008, Amount: 20}}, refine.SetCost(75000))
}

func (r Fiona) NightRavensClawRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032015, Amount: 1}, {ItemId: 4005002, Amount: 10}, {ItemId: 4021008, Amount: 30}}, refine.SetCost(50000))
}

func (r Fiona) DuskRavensClawRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032016, Amount: 1}, {ItemId: 4005000, Amount: 5}, {ItemId: 4021008, Amount: 20}}, refine.SetCost(50000))
}

func (r Fiona) DawnRavensBeakRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032017, Amount: 1}, {ItemId: 4005001, Amount: 10}, {ItemId: 4021008, Amount: 20}}, refine.SetCost(75000))
}

func (r Fiona) NightRavensBeakRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032015, Amount: 1}, {ItemId: 4005002, Amount: 10}, {ItemId: 4021008, Amount: 30}}, refine.SetCost(50000))
}

func (r Fiona) DuskRavensBeakRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032016, Amount: 1}, {ItemId: 4005000, Amount: 5}, {ItemId: 4021008, Amount: 20}}, refine.SetCost(50000))
}

func (r Fiona) DawnRavensEyeRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032017, Amount: 1}, {ItemId: 4005001, Amount: 10}, {ItemId: 4021008, Amount: 20}}, refine.SetCost(75000))
}

func (r Fiona) NightRavensEyeRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032015, Amount: 1}, {ItemId: 4005002, Amount: 10}, {ItemId: 4021008, Amount: 30}}, refine.SetCost(50000))
}

func (r Fiona) DuskRavensEyeRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032016, Amount: 1}, {ItemId: 4005000, Amount: 5}, {ItemId: 4021008, Amount: 20}}, refine.SetCost(50000))
}

func (r Fiona) DawnRavensWingRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032017, Amount: 1}, {ItemId: 4005001, Amount: 10}, {ItemId: 4021008, Amount: 20}}, refine.SetCost(75000))
}

func (r Fiona) NightRavensWingRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032015, Amount: 1}, {ItemId: 4005002, Amount: 10}, {ItemId: 4021008, Amount: 30}}, refine.SetCost(50000))
}

func (r Fiona) DuskRavensWingRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4032016, Amount: 1}, {ItemId: 4005000, Amount: 5}, {ItemId: 4021008, Amount: 20}}, refine.SetCost(50000))
}

func (r Fiona) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("All done. If you need anything else... Well, I'm not going anywhere.")
	return script.SendOk(l, c, m.String())
}

func (r Fiona) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I am afraid you don't have enough to pay me, partner. Please check this out first, ok?")
	return script.SendOk(l, c, m.String())
}

func (r Fiona) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText("Hey, I need those items to craft properly, you know?")
		return script.SendOk(l, c, m.String())
	}
}

func (r Fiona) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, c, m.String())
}
