package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Mos is located in Leafre - Leafre (240000000)
type Mos struct {
}

func (r Mos) NPCId() uint32 {
	return npc.Mos
}

func (r Mos) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return refine.NewGenericRefine(l, c, r.Hello(), r.Categories(l, c))
}

func (r Mos) Hello() string {
	return "A dragon's power is not to be underestimated. If you like, I can add its power to one of your weapons. However, the weapon must be powerful enough to hold its potential..."
}

func (r Mos) Categories(l logrus.FieldLogger, c script.Context) []refine.ListItem {
	base := []refine.ListItem{
		r.WhatIsAStimulator(),
		r.Warrior(),
		r.Bowman(),
		r.Magician(),
		r.Thief(),
		r.Pirate(),
		r.WarriorStimulator(),
		r.BowmanStimulator(),
		r.MagicianStimulator(),
		r.ThiefStimulator(),
		r.PirateStimulator(),
	}
	if character.QuestStarted(l)(c.CharacterId, 7301) || character.QuestStarted(l)(c.CharacterId, 7303) {
		return append(base, r.CorniansDagger())
	}
	return base
}

func (r Mos) WhatIsAStimulator() refine.ListItem {
	return refine.ListItem{
		ListText:       "What's a stimulator?",
		SelectionState: r.StimulatorInfo,
	}
}

func (r Mos) StimulatorInfo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("A stimulator is a special potion that I can add into the process of creating certain items. It gives it stats as though it had dropped from a monster. However, it is possible to have no change, and it is also possible for the item to be below average. There's also a 10% chance of not getting any item when using a stimulator, so please choose wisely.")
	return script.SendOk(l, c, m.String())
}

func (r Mos) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, r.GenericRefinementConfig())
}

func (r Mos) GenericRefinementConfig() refine.TerminalConfig {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
		StimulatorError:  r.StimulatorError,
	}
	return config
}

func (r Mos) Warrior() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonCarbella, " - Lv. 110 One-Handed Sword"), refine.Confirm(item.DragonCarbella, r.DragonCarbellaRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonAxe, " - Lv. 110 One-Handed Axe"), refine.Confirm(item.DragonAxe, r.DragonAxeRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonMace, " - Lv. 110 One-Handed BW"), refine.Confirm(item.DragonMace, r.DragonMaceRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonClaymore, " - Lv. 110 Two-Handed Sword"), refine.Confirm(item.DragonClaymore, r.DragonClaymoreRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonBattleAxe, " - Lv. 110 Two-Handed Axe"), refine.Confirm(item.DragonBattleAxe, r.DragonBattleAxeRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonFlame, " - Lv. 110 Two-Handed BW"), refine.Confirm(item.DragonFlame, r.DragonFlameRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonFaltizan, " - Lv. 110 Spear"), refine.Confirm(item.DragonFaltizan, r.DragonFaltizanRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonChelbird, " - Lv. 110 Polearm"), refine.Confirm(item.DragonChelbird, r.DragonChelbirdRequirements())),
	}
	prompt := refine.PromptCategory("Very well, then which Warrior weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Warrior weapon",
		SelectionState: prompt,
	}
}

func (r Mos) Bowman() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonShinerBow, " - Lv. 110 Bow"), refine.Confirm(item.DragonShinerBow, r.DragonShinerBowRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonShinerCross, " - Lv. 110 Crossbow"), refine.Confirm(item.DragonShinerCross, r.DragonShinerCrossRequirements())),
	}
	prompt := refine.PromptCategory("Very well, then which Bowman weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Bowman weapon",
		SelectionState: prompt,
	}
}

func (r Mos) Magician() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonWand, " - Lv. 108 Wand"), refine.Confirm(item.DragonWand, r.DragonWandRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonStaff, " - Lv. 110 Staff"), refine.Confirm(item.DragonStaff, r.DragonStaffRequirements())),
	}
	prompt := refine.PromptCategory("Very well, then which Magician weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Magician weapon",
		SelectionState: prompt,
	}
}

func (r Mos) Thief() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonKanzir, " - Lv. 110 STR Dagger"), refine.Confirm(item.DragonKanzir, r.DragonKanzirRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonKreda, " - Lv. 110 LUK Dagger"), refine.Confirm(item.DragonKreda, r.DragonKredaRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonGreenSleve, " - Lv. 110 Claw"), refine.Confirm(item.DragonGreenSleve, r.DragonGreenSleveRequirements())),
	}
	prompt := refine.PromptCategory("Very well, then which Thief weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Thief weapon",
		SelectionState: prompt,
	}
}

func (r Mos) Pirate() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonSlashClaw, " - Lv. 110 Knuckle"), refine.Confirm(item.DragonSlashClaw, r.DragonSlashClawRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonfireRevolver, " - Lv. 110 Gun"), refine.Confirm(item.DragonfireRevolver, r.DragonfireRevolverRequirements())),
	}
	prompt := refine.PromptCategory("Very well, then which Pirate weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Pirate weapon",
		SelectionState: prompt,
	}
}

func (r Mos) WarriorStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonCarbella, " - Lv. 110 One-Handed Sword"), refine.Confirm(item.DragonCarbella, r.DragonCarbellaRequirements().AddRequirement(item.OneHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonAxe, " - Lv. 110 One-Handed Axe"), refine.Confirm(item.DragonAxe, r.DragonAxeRequirements().AddRequirement(item.OneHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonMace, " - Lv. 110 One-Handed BW"), refine.Confirm(item.DragonMace, r.DragonMaceRequirements().AddRequirement(item.OneHandedBluntWeaponForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonClaymore, " - Lv. 110 Two-Handed Sword"), refine.Confirm(item.DragonClaymore, r.DragonClaymoreRequirements().AddRequirement(item.TwoHandedSwordForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonBattleAxe, " - Lv. 110 Two-Handed Axe"), refine.Confirm(item.DragonBattleAxe, r.DragonBattleAxeRequirements().AddRequirement(item.TwoHandedAxeForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonFlame, " - Lv. 110 Two-Handed BW"), refine.Confirm(item.DragonFlame, r.DragonFlameRequirements().AddRequirement(item.TwoHandedMaceForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonFaltizan, " - Lv. 110 Spear"), refine.Confirm(item.DragonFaltizan, r.DragonFaltizanRequirements().AddRequirement(item.SpearForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonChelbird, " - Lv. 110 Polearm"), refine.Confirm(item.DragonChelbird, r.DragonChelbirdRequirements().AddRequirement(item.PoleArmForgingStimulator, 1))),
	}
	prompt := refine.PromptCategory("Very well, then which Warrior weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Warrior weapon with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Mos) BowmanStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonShinerBow, " - Lv. 110 Bow"), refine.Confirm(item.DragonShinerBow, r.DragonShinerBowRequirements().AddRequirement(item.BowProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonShinerCross, " - Lv. 110 Crossbow"), refine.Confirm(item.DragonShinerCross, r.DragonShinerCrossRequirements().AddRequirement(item.CrossbowProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Very well, then which Bowman weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Bowman weapon with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Mos) MagicianStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonWand, " - Lv. 108 Wand"), refine.Confirm(item.DragonWand, r.DragonWandRequirements().AddRequirement(item.WandProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonStaff, " - Lv. 110 Staff"), refine.Confirm(item.DragonStaff, r.DragonStaffRequirements().AddRequirement(item.StaffProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Very well, then which Magician weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Magician weapon with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Mos) ThiefStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonKanzir, " - Lv. 110 STR Dagger"), refine.Confirm(item.DragonKanzir, r.DragonKanzirRequirements().AddRequirement(item.DaggerForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonKreda, " - Lv. 110 LUK Dagger"), refine.Confirm(item.DragonKreda, r.DragonKredaRequirements().AddRequirement(item.DaggerForgingStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonGreenSleve, " - Lv. 110 Claw"), refine.Confirm(item.DragonGreenSleve, r.DragonGreenSleveRequirements().AddRequirement(item.ClawProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Very well, then which Thief weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Thief weapon with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Mos) PirateStimulator() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonSlashClaw, " - Lv. 110 Knuckle"), refine.Confirm(item.DragonSlashClaw, r.DragonSlashClawRequirements().AddRequirement(item.KnucklerProductionStimulator, 1))),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DragonfireRevolver, " - Lv. 110 Gun"), refine.Confirm(item.DragonfireRevolver, r.DragonfireRevolverRequirements().AddRequirement(item.GunProductionStimulator, 1))),
	}
	prompt := refine.PromptCategory("Very well, then which Pirate weapon shall receive a dragon's power?", choices)
	return refine.ListItem{
		ListText:       "Create a Pirate Weapon with a Stimulator",
		SelectionState: prompt,
	}
}

func (r Mos) CorniansDagger() refine.ListItem {
	return refine.ListItem{
		ListText:       message.NewBuilder().AddText("Make ").ShowItemName1(item.CorniansDagger).String(),
		SelectionState: r.Sneak,
	}
}

func (r Mos) DragonCarbellaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1302056, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 8}}, refine.SetCost(120000))
}

func (r Mos) DragonAxeRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1312030, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 8}}, refine.SetCost(120000))
}

func (r Mos) DragonMaceRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1322045, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 8}}, refine.SetCost(120000))
}

func (r Mos) DragonClaymoreRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1402035, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 8}}, refine.SetCost(120000))
}

func (r Mos) DragonBattleAxeRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1412021, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 8}}, refine.SetCost(120000))
}

func (r Mos) DragonFlameRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1422027, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 8}}, refine.SetCost(120000))
}

func (r Mos) DragonFaltizanRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1432030, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 8}}, refine.SetCost(120000))
}

func (r Mos) DragonChelbirdRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1442044, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 8}}, refine.SetCost(120000))
}

func (r Mos) DragonShinerBowRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1452019, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 3}, {ItemId: 4005002, Amount: 5}}, refine.SetCost(120000))
}

func (r Mos) DragonShinerCrossRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1462015, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 5}, {ItemId: 4005002, Amount: 3}}, refine.SetCost(120000))
}

func (r Mos) DragonWandRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1372010, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005001, Amount: 6}, {ItemId: 4005003, Amount: 2}}, refine.SetCost(120000))
}

func (r Mos) DragonStaffRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1382035, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005001, Amount: 6}, {ItemId: 4005003, Amount: 2}}, refine.SetCost(120000))
}

func (r Mos) DragonKanzirRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1332051, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 5}, {ItemId: 4005002, Amount: 3}}, refine.SetCost(120000))
}

func (r Mos) DragonKredaRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1332052, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005002, Amount: 3}, {ItemId: 4005003, Amount: 5}}, refine.SetCost(120000))
}

func (r Mos) DragonGreenSleveRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1472053, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005002, Amount: 2}, {ItemId: 4005003, Amount: 6}}, refine.SetCost(120000))
}

func (r Mos) DragonSlashClawRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1482012, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 5}, {ItemId: 4005002, Amount: 3}}, refine.SetCost(120000))
}

func (r Mos) DragonfireRevolverRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 1492012, Amount: 1}, {ItemId: 4000244, Amount: 20}, {ItemId: 4000245, Amount: 25}, {ItemId: 4005000, Amount: 3}, {ItemId: 4005002, Amount: 5}}, refine.SetCost(120000))
}

func (r Mos) StimulatorError(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Unfortunately, the dragon's essence has... conflicted with your weapon. My apologies for your loss.")
	return script.SendOk(l, c, m.String())
}

func (r Mos) Success(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("The process is complete. Treat your weapon well, lest you bring the wrath of the dragons upon you.")
	return script.SendOk(l, c, m.String())
}

func (r Mos) CannotAfford(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("My fee is for the good of all of Leafre. If you cannot pay it, then begone.")
	return script.SendOk(l, c, m.String())
}

func (r Mos) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText("I'm afraid that without the correct items, the dragon's essence would... not make for a very reliable weapon. Please bring the correct items next time.")
		return script.SendOk(l, c, m.String())
	}
}

func (r Mos) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, c, m.String())
}

func (r Mos) Sneak(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("\"Oh, are you trying to sneak into these lizards to save Moira? I will support your cause wherever I can. Bring me a couple of resources and I will make you an almost identical piece of ").
		ShowItemName1(item.CorniansDagger).AddText(".")
	return script.SendNext(l, c, m.String(), refine.Confirm(item.CorniansDagger, r.CorniansDaggerRequirements())(r.GenericRefinementConfig()))
}

func (r Mos) CorniansDaggerRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011001, Amount: 1}, {ItemId: 4011002, Amount: 1}, {ItemId: 4001079, Amount: 1}}, refine.SetCost(25000))
}