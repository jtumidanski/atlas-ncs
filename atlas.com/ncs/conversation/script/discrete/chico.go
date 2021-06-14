package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Chico is located in Ludibrium - Ludibrium Village (220000300)
type Chico struct {
}

func (r Chico) NPCId() uint32 {
	return npc.Chico
}

func (r Chico) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return refine.NewSingleCategoryRefine(l, c, r.Hello(), r.Choices())
}

func (r Chico) Hello() string {
	m := message.NewBuilder().
		AddText("Hey there! My name is ").
		ShowNPC(npc.Chico).
		AddText(", and I am a specialist in mini-games. What kind of mini-game you want me to make?")
	return m.String()
}

func (r Chico) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Chico) Choices() []refine.RefinementChoice {
	return []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndImageList(item.ASetOfMatchCards), refine.Confirm(item.ASetOfMatchCards, r.ASetOfMatchCardsRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.BloctopusAndPinkTeddyOmokSet), refine.Confirm(item.BloctopusAndPinkTeddyOmokSet, r.BloctopusAndPinkTeddyOmokSetRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.BloctopusAndTrixterOmokSet), refine.Confirm(item.BloctopusAndTrixterOmokSet, r.BloctopusAndTrixterOmokSetRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.PinkTeddyAndTrixterOmokSet), refine.Confirm(item.PinkTeddyAndTrixterOmokSet, r.PinkTeddyAndTrixterOmokSetRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.PandaTeddyAndBlocktopusOmokSet), refine.Confirm(item.PandaTeddyAndBlocktopusOmokSet, r.PandaTeddyAndBlocktopusOmokSetRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.PandaTeddyAndPinkTeddyOmokSet), refine.Confirm(item.PandaTeddyAndPinkTeddyOmokSet, r.PandaTeddyAndPinkTeddyOmokSetRequirements())),
		r.CreateChoice(refine.ItemIdAndImageList(item.PandaTeddyAndTrixterOmokSet), refine.Confirm(item.PandaTeddyAndTrixterOmokSet, r.PandaTeddyAndTrixterOmokSetRequirements())),
	}
}

func (r Chico) ASetOfMatchCardsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4030012, Amount: 99}}, refine.SetCost(10000))
}

func (r Chico) BloctopusAndPinkTeddyOmokSetRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4030009, Amount: 1}, {ItemId: 4030013, Amount: 99}, {ItemId: 4030014, Amount: 99}}, refine.SetCost(25000))
}

func (r Chico) BloctopusAndTrixterOmokSetRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4030009, Amount: 1}, {ItemId: 4030013, Amount: 99}, {ItemId: 4030016, Amount: 99}}, refine.SetCost(25000))
}

func (r Chico) PinkTeddyAndTrixterOmokSetRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4030009, Amount: 1}, {ItemId: 4030014, Amount: 99}, {ItemId: 4030016, Amount: 99}}, refine.SetCost(25000))
}

func (r Chico) PandaTeddyAndBlocktopusOmokSetRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4030009, Amount: 1}, {ItemId: 4030015, Amount: 99}, {ItemId: 4030013, Amount: 99}}, refine.SetCost(25000))
}

func (r Chico) PandaTeddyAndPinkTeddyOmokSetRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4030009, Amount: 1}, {ItemId: 4030015, Amount: 99}, {ItemId: 4030014, Amount: 99}}, refine.SetCost(25000))
}

func (r Chico) PandaTeddyAndTrixterOmokSetRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4030009, Amount: 1}, {ItemId: 4030015, Amount: 99}, {ItemId: 4030016, Amount: 99}}, refine.SetCost(25000))
}

func (r Chico) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("There is your game set. Have fun!")
	return script.SendOk(l, c, m.String())
}

func (r Chico) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("See, I need to specify my wages to support my career, that cannot be bypassed. I will gladly help you once you've got the money.")
	return script.SendOk(l, c, m.String())
}

func (r Chico) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText("You are lacking some items for the set you want to make. Please provide them so that we can assemble the game set.")
		return script.SendOk(l, c, m.String())
	}
}

func (r Chico) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("I can't make a set for you if there's no room in your ETC inventory for it. Please free a space first and then talk to me.")
	return script.SendOk(l, c, m.String())
}