package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/generic/refine"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Gordon is located in El Nath - El Nath Market (211000100)
type Gordon struct {
}

func (r Gordon) NPCId() uint32 {
	return npc.Gordon
}

func (r Gordon) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return refine.NewGenericRefine(l, span, c, r.Hello(), r.Categories())
}

func (r Gordon) Hello() string {
	return "Hello there. El Nath winters are incredibly cold, you're going to need a warm pair of shoes to survive."
}

func (r Gordon) Categories() []refine.ListItem {
	return []refine.ListItem{
		r.Warrior(),
		r.Bowman(),
		r.Magician(),
		r.Thief(),
	}
}

func (r Gordon) CreateChoice(listTextProvider refine.RefinementListTextProvider, selectionPrompt refine.TerminalState) refine.RefinementChoice {
	config := refine.TerminalConfig{
		Success:          r.Success,
		MesoError:        r.CannotAfford,
		RequirementError: r.MissingSomething,
		InventoryError:   r.MakeRoom,
	}
	return refine.CreateRefinementChoice(listTextProvider, selectionPrompt, config)
}

func (r Gordon) Warrior() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.SapphireCamelBoots, " - Warrior Lv. 60"), refine.Confirm(item.SapphireCamelBoots, r.SapphireCamelBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrihalconCamelBoots, " - Warrior Lv. 60"), refine.Confirm(item.OrihalconCamelBoots, r.OrihalconCamelBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodCamelBoots, " - Warrior Lv. 60"), refine.Confirm(item.BloodCamelBoots, r.BloodCamelBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueCarzenBoots, " - Warrior Lv. 70"), refine.Confirm(item.BlueCarzenBoots, r.BlueCarzenBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleCarzenBoots, " - Warrior Lv. 70"), refine.Confirm(item.PurpleCarzenBoots, r.PurpleCarzenBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkCarzenBoots, " - Warrior Lv. 70"), refine.Confirm(item.DarkCarzenBoots, r.DarkCarzenBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedRiversBoots, " - Warrior Lv. 80"), refine.Confirm(item.RedRiversBoots, r.RedRiversBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueRiversBoots, " - Warrior Lv. 80"), refine.Confirm(item.BlueRiversBoots, r.BlueRiversBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkRiversBoots, " - Warrior Lv. 80"), refine.Confirm(item.DarkRiversBoots, r.DarkRiversBootsRequirements())),
	}
	prompt := refine.PromptCategory("Warrior shoes? Okay, then which set?", choices)
	return refine.ListItem{ListText: "Create Warrior shoes", SelectionState: prompt}
}

func (r Gordon) Bowman() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedGoreBoots, " - Bowman Lv. 60"), refine.Confirm(item.RedGoreBoots, r.RedGoreBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueGoreBoots, " - Bowman Lv. 60"), refine.Confirm(item.BlueGoreBoots, r.BlueGoreBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenGoreBoots, " - Bowman Lv. 60"), refine.Confirm(item.GreenGoreBoots, r.GreenGoreBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueElfBoots, " - Bowman Lv. 70"), refine.Confirm(item.BlueElfBoots, r.BlueElfBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BeigeElfBoots, " - Bowman Lv. 70"), refine.Confirm(item.BeigeElfBoots, r.BeigeElfBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenElfBoots, " - Bowman Lv. 70"), refine.Confirm(item.GreenElfBoots, r.GreenElfBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkElfBoots, " - Bowman Lv. 70"), refine.Confirm(item.DarkElfBoots, r.DarkElfBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueWingBoots, " - Bowman Lv. 80"), refine.Confirm(item.BlueWingBoots, r.BlueWingBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedWingBoots, " - Bowman Lv. 80"), refine.Confirm(item.RedWingBoots, r.RedWingBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenWingBoots, " - Bowman Lv. 80"), refine.Confirm(item.GreenWingBoots, r.GreenWingBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkWingBoots, " - Bowman Lv. 80"), refine.Confirm(item.DarkWingBoots, r.DarkWingBootsRequirements())),
	}
	prompt := refine.PromptCategory("Bowman shoes? Okay, then which set?", choices)
	return refine.ListItem{ListText: "Create Bowman shoes", SelectionState: prompt}
}

func (r Gordon) Magician() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PinkGoldrunners, " - Magician Lv. 60"), refine.Confirm(item.PinkGoldrunners, r.PinkGoldrunnersRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenGoldrunners, " - Magician Lv. 60"), refine.Confirm(item.GreenGoldrunners, r.GreenGoldrunnersRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.OrangeGoldrunners, " - Magician Lv. 60"), refine.Confirm(item.OrangeGoldrunners, r.OrangeGoldrunnersRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueGoldrunners, " - Magician Lv. 60"), refine.Confirm(item.BlueGoldrunners, r.BlueGoldrunnersRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueLapizSandals, " - Magician Lv. 70"), refine.Confirm(item.BlueLapizSandals, r.BlueLapizSandalsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedLapizSandals, " - Magician Lv. 70"), refine.Confirm(item.RedLapizSandals, r.RedLapizSandalsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BrownLapizSandals, " - Magician Lv. 70"), refine.Confirm(item.BrownLapizSandals, r.BrownLapizSandalsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldLapizSandals, " - Magician Lv. 70"), refine.Confirm(item.GoldLapizSandals, r.GoldLapizSandalsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenEnigmaShoes, " - Magician Lv. 80"), refine.Confirm(item.GreenEnigmaShoes, r.GreenEnigmaShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleEnigmaShoes, " - Magician Lv. 80"), refine.Confirm(item.PurpleEnigmaShoes, r.PurpleEnigmaShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkEnigmaShoes, " - Magician Lv. 80"), refine.Confirm(item.DarkEnigmaShoes, r.DarkEnigmaShoesRequirements())),
	}
	prompt := refine.PromptCategory("Magician shoes? Okay, then which set?", choices)
	return refine.ListItem{ListText: "Create Magician shoes", SelectionState: prompt}
}

func (r Gordon) Thief() refine.ListItem {
	choices := []refine.RefinementChoice{
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BloodMossBoots, " - Thief Lv. 60"), refine.Confirm(item.BloodMossBoots, r.BloodMossBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GoldMossBoots, " - Thief Lv. 60"), refine.Confirm(item.GoldMossBoots, r.GoldMossBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkMossBoots, " - Thief Lv. 60"), refine.Confirm(item.DarkMossBoots, r.DarkMossBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.PurpleMystiqueShoes, " - Thief Lv. 70"), refine.Confirm(item.PurpleMystiqueShoes, r.PurpleMystiqueShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.BlueMystiqueShoes, " - Thief Lv. 70"), refine.Confirm(item.BlueMystiqueShoes, r.BlueMystiqueShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedMystiqueShoes, " - Thief Lv. 70"), refine.Confirm(item.RedMystiqueShoes, r.RedMystiqueShoesRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.GreenPirateBoots, " - Thief Lv. 80"), refine.Confirm(item.GreenPirateBoots, r.GreenPirateBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.RedPirateBoots, " - Thief Lv. 80"), refine.Confirm(item.RedPirateBoots, r.RedPirateBootsRequirements())),
		r.CreateChoice(refine.ItemIdAndDescriptionList(item.DarkPirateBoots, " - Thief Lv. 80"), refine.Confirm(item.DarkPirateBoots, r.DarkPirateBootsRequirements())),
	}
	prompt := refine.PromptCategory("Thief shoes? Okay, then which set?", choices)
	return refine.ListItem{ListText: "Create Thief shoes", SelectionState: prompt}
}

func (r Gordon) SapphireCamelBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4021005, Amount: 8}, {ItemId: 4000030, Amount: 80}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(60000))
}

func (r Gordon) OrihalconCamelBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4011005, Amount: 8}, {ItemId: 4000030, Amount: 80}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(60000))
}

func (r Gordon) BloodCamelBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021008, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4021000, Amount: 8}, {ItemId: 4000030, Amount: 80}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(60000))
}

func (r Gordon) BlueCarzenBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005000, Amount: 1}, {ItemId: 4005002, Amount: 3}, {ItemId: 4011002, Amount: 5}, {ItemId: 4000048, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) PurpleCarzenBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005000, Amount: 2}, {ItemId: 4005002, Amount: 2}, {ItemId: 4011005, Amount: 5}, {ItemId: 4000048, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) DarkCarzenBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005000, Amount: 3}, {ItemId: 4005002, Amount: 1}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000048, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) RedRiversBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005000, Amount: 2}, {ItemId: 4005002, Amount: 3}, {ItemId: 4021000, Amount: 7}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 65}}, refine.SetCost(80000))
}

func (r Gordon) BlueRiversBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005000, Amount: 3}, {ItemId: 4005002, Amount: 2}, {ItemId: 4021002, Amount: 7}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 65}}, refine.SetCost(80000))
}

func (r Gordon) DarkRiversBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005000, Amount: 4}, {ItemId: 4005002, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 65}}, refine.SetCost(80000))
}

func (r Gordon) RedGoreBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 5}, {ItemId: 4021000, Amount: 8}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) BlueGoreBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 5}, {ItemId: 4021005, Amount: 8}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) GreenGoreBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4011006, Amount: 5}, {ItemId: 4021003, Amount: 8}, {ItemId: 4021007, Amount: 1}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) BlueElfBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 1}, {ItemId: 4005000, Amount: 3}, {ItemId: 4021005, Amount: 5}, {ItemId: 4000055, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) BeigeElfBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 2}, {ItemId: 4005000, Amount: 2}, {ItemId: 4021004, Amount: 5}, {ItemId: 4000055, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) GreenElfBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 2}, {ItemId: 4005000, Amount: 2}, {ItemId: 4021003, Amount: 5}, {ItemId: 4000055, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) DarkElfBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 3}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000055, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) BlueWingBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 2}, {ItemId: 4005000, Amount: 3}, {ItemId: 4021002, Amount: 7}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) RedWingBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 3}, {ItemId: 4005000, Amount: 2}, {ItemId: 4021000, Amount: 7}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) GreenWingBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 4}, {ItemId: 4005000, Amount: 1}, {ItemId: 4021003, Amount: 7}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) DarkWingBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 5}, {ItemId: 4021008, Amount: 2}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) PinkGoldrunnersRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 4}, {ItemId: 4011005, Amount: 5}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) GreenGoldrunnersRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 4}, {ItemId: 4021003, Amount: 5}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) OrangeGoldrunnersRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 4}, {ItemId: 4011003, Amount: 5}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) BlueGoldrunnersRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021009, Amount: 1}, {ItemId: 4011006, Amount: 4}, {ItemId: 4021002, Amount: 5}, {ItemId: 4000030, Amount: 70}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) BlueLapizSandalsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005001, Amount: 1}, {ItemId: 4005003, Amount: 3}, {ItemId: 4021002, Amount: 5}, {ItemId: 4000051, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) RedLapizSandalsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005001, Amount: 2}, {ItemId: 4005003, Amount: 2}, {ItemId: 4021000, Amount: 5}, {ItemId: 4000051, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) BrownLapizSandalsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005001, Amount: 2}, {ItemId: 4005003, Amount: 2}, {ItemId: 4011003, Amount: 5}, {ItemId: 4000051, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) GoldLapizSandalsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005001, Amount: 3}, {ItemId: 4005003, Amount: 1}, {ItemId: 4011006, Amount: 3}, {ItemId: 4000051, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) GreenEnigmaShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005001, Amount: 2}, {ItemId: 4005003, Amount: 3}, {ItemId: 4021003, Amount: 7}, {ItemId: 4000030, Amount: 85}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) PurpleEnigmaShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005001, Amount: 3}, {ItemId: 4005003, Amount: 2}, {ItemId: 4021001, Amount: 7}, {ItemId: 4000030, Amount: 85}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) DarkEnigmaShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005001, Amount: 4}, {ItemId: 4005003, Amount: 1}, {ItemId: 4021008, Amount: 2}, {ItemId: 4000030, Amount: 85}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) BloodMossBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021007, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4021000, Amount: 8}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) GoldMossBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021007, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4011006, Amount: 5}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) DarkMossBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4021007, Amount: 1}, {ItemId: 4011007, Amount: 1}, {ItemId: 4021008, Amount: 1}, {ItemId: 4000030, Amount: 75}, {ItemId: 4003000, Amount: 50}}, refine.SetCost(60000))
}

func (r Gordon) PurpleMystiqueShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005003, Amount: 1}, {ItemId: 4005000, Amount: 3}, {ItemId: 4021001, Amount: 5}, {ItemId: 4000051, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) BlueMystiqueShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005003, Amount: 1}, {ItemId: 4005002, Amount: 3}, {ItemId: 4021005, Amount: 5}, {ItemId: 4000051, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) RedMystiqueShoesRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 1}, {ItemId: 4005003, Amount: 3}, {ItemId: 4021000, Amount: 5}, {ItemId: 4000051, Amount: 100}, {ItemId: 4003000, Amount: 55}}, refine.SetCost(70000))
}

func (r Gordon) GreenPirateBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005000, Amount: 3}, {ItemId: 4005003, Amount: 2}, {ItemId: 4021003, Amount: 7}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) RedPirateBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005002, Amount: 3}, {ItemId: 4005003, Amount: 2}, {ItemId: 4021000, Amount: 7}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) DarkPirateBootsRequirements() refine.Requirements {
	return refine.NewRequirements([]refine.Requirement{{ItemId: 4005003, Amount: 3}, {ItemId: 4005002, Amount: 2}, {ItemId: 4021008, Amount: 7}, {ItemId: 4000030, Amount: 90}, {ItemId: 4003000, Amount: 60}}, refine.SetCost(80000))
}

func (r Gordon) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("All done. Stay warm!")
	return script.SendOk(l, span, c, m.String())
}

func (r Gordon) CannotAfford(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm afraid you cannot afford my services.")
	return script.SendOk(l, span, c, m.String())
}

func (r Gordon) MissingSomething(_ uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText("I only make quality goods, which I cannot do without the proper materials.")
		return script.SendOk(l, span, c, m.String())
	}
}

func (r Gordon) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Check your inventory for a free slot first.")
	return script.SendOk(l, span, c, m.String())
}
