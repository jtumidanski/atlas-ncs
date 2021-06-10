package refine

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

type RefinementCategory struct {
	ListText        string
	Prompt          string
	Choices         []RefinementChoice
	SelectionPrompt CategoryStateProducer
}

type CategoryStateProducer func(RefinementCategory) script.StateProducer

type RefinementChoice struct {
	ListText        string
	SelectionPrompt TerminalState
	Config          TerminalConfig
}

type RefinementRequirements struct {
	Requirements []Requirement
	Cost         uint32
}

type Requirement struct {
	ItemId uint32
	Amount uint8
}

func NewGenericRefine(l logrus.FieldLogger, c script.Context, hello string, categories []RefinementCategory) script.State {
	m := message.NewBuilder().AddText(hello).NewLine()
	for i, category := range categories {
		m = m.OpenItem(i).BlueText().AddText(category.ListText).CloseItem().NewLine()
	}
	return script.SendListSelection(l, c, m.String(), categorySelection(categories))
}

func categorySelection(categories []RefinementCategory) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		if selection < 0 || int(selection) > len(categories) {
			return script.Exit()
		}
		category := categories[selection]
		return category.SelectionPrompt(category)
	}
}

func PromptCategory(category RefinementCategory) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText(category.Prompt)
		for i, choice := range category.Choices {
			m = m.OpenItem(i).BlueText().AddText(choice.ListText)
		}
		return script.SendListSelection(l, c, m.String(), choiceSelection(category.Choices))
	}
}

func choiceSelection(choices []RefinementChoice) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		if selection < 0 || int(selection) > len(choices) {
			return script.Exit()
		}
		choice := choices[selection]
		return choice.SelectionPrompt(choice.Config)
	}
}

type TerminalState func(config TerminalConfig) script.StateProducer

func HowMany(itemId uint32, requirements RefinementRequirements) TerminalState {
	return func(config TerminalConfig) script.StateProducer {
		return func(l logrus.FieldLogger, c script.Context) script.State {
			m := message.NewBuilder().
				AddText("So, you want me to make some ").
				ShowItemName1(itemId).
				AddText("s? In that case, how many do you want me to make?")
			return script.SendGetNumber(l, c, m.String(), quantitySelection(itemId, requirements, config), 1, 1, 100)
		}
	}
}

func Confirm(itemId uint32, requirements RefinementRequirements) TerminalState {
	return func(config TerminalConfig) script.StateProducer {
		return confirmQuantity(itemId, 1, requirements, config)
	}
}

func quantitySelection(itemId uint32, requirements RefinementRequirements, config TerminalConfig) script.ProcessNumber {
	return func(selection int32) script.StateProducer {
		return confirmQuantity(itemId, uint32(selection), requirements, config)
	}
}

func confirmQuantity(itemId uint32, amount uint32, requirements RefinementRequirements, config TerminalConfig) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("You want me to make ")
		if amount == 1 {
			m = m.AddText("a ").ShowItemName1(itemId)
		} else {
			m = m.AddText(fmt.Sprintf("%d ", amount)).ShowItemName1(itemId)
		}
		m = m.AddText("? In that case, I'm going to need specific items from you in order to make it. Make sure you have room in your inventory, though!").NewLine()
		for _, req := range requirements.Requirements {
			m = m.ShowItemImage2(req.ItemId).AddText(fmt.Sprintf(" %d ", req.Amount)).ShowItemName1(req.ItemId).NewLine()
		}
		if requirements.Cost > 0 {
			m = m.ShowItemImage2(item.MoneySack).AddText(fmt.Sprintf(" %d meso", requirements.Cost*amount))
		}
		return script.SendYesNo(l, c, m.String(), validate(itemId, amount, requirements, config), script.Exit())
	}
}

func validate(itemId uint32, amount uint32, requirements RefinementRequirements, config TerminalConfig) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if !character.CanHoldAll(l)(c.CharacterId, itemId, amount) {
			return config.InventoryError(l, c)
		}
		if !character.HasMeso(l)(c.CharacterId, requirements.Cost*amount) {
			return config.MesoError(l, c)
		}
		for _, req := range requirements.Requirements {
			if !character.HasItems(l)(c.CharacterId, req.ItemId, uint32(req.Amount)*amount) {
				return config.RequirementError(req.ItemId)(l, c)
			}
		}
		return performRefine(itemId, amount, requirements, config)(l, c)
	}
}

func performRefine(itemId uint32, amount uint32, requirements RefinementRequirements, config TerminalConfig) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		err := character.GainMeso(l)(c.CharacterId, -int32(amount*requirements.Cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment for refine.")
		}
		for _, req := range requirements.Requirements {
			character.GainItem(l)(c.CharacterId, req.ItemId, -int32(req.Amount)*int32(amount))
		}
		character.GainItem(l)(c.CharacterId, itemId, int32(amount))
		return config.Success(l, c)
	}
}

type TerminalConfig struct {
	Success          script.StateProducer
	MesoError        script.StateProducer
	RequirementError func(itemId uint32) script.StateProducer
	InventoryError   script.StateProducer
}