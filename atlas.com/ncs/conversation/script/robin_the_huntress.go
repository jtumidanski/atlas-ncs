package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// RobinTheHuntress is located in Amoria - Cherished Visage Photos (680000300) and Amoria - The Love Pinata (680000401)
type RobinTheHuntress struct {
}

func (r RobinTheHuntress) NPCId() uint32 {
	return npc.RobinTheHuntress
}

func (r RobinTheHuntress) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId != _map.TheLovePinata {
		if c.MapId != _map.UntamedHeartsHuntingGround {
			m := message.NewBuilder().
				AddText("Hello, where would you like to go?").NewLine().
				OpenItem(0).BlueText().AddText("Untamed Hearts Hunting Ground").CloseItem().NewLine().
				OpenItem(1).BlueText().AddText("Please warp me out.").CloseItem()
			return SendListSelectionExit(l, c, m.String(), r.UntamedHeartsSelection, r.Goodbye)
		} else {
			m := message.NewBuilder().
				AddText("Hello, where would you like to go?").NewLine().
				OpenItem(0).BlueText().AddText("I have 7 keys. Bring me to smash boxes").CloseItem().NewLine().
				OpenItem(1).BlueText().AddText("Please warp me out.").CloseItem()
			return SendListSelectionExit(l, c, m.String(), r.SmashBoxesSelection, r.Goodbye)
		}
	} else {
		m := message.NewBuilder().
			AddText("Hello, do you want to go back now? Returning here again will cost you #rother 7 keys").
			BlackText().AddText(".").NewLine().
			OpenItem(0).BlueText().AddText("Please warp me back to the training grounds.")
		return SendListSelectionExit(l, c, m.String(), r.LovePinataSelection, r.Goodbye)
	}
}

func (r RobinTheHuntress) UntamedHeartsSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.ValidateUntamedHearts
	case 1:
		return r.WarpToWeddingExit
	}
	return nil
}

func (r RobinTheHuntress) Goodbye(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Goodbye then.")
	return SendOk(l, c, m.String())
}

func (r RobinTheHuntress) SmashBoxesSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.ValidateSmashBoxes
	case 1:
		return r.WarpToWeddingExit
	}
	return nil
}

func (r RobinTheHuntress) WarpToUntamedHearts(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.UntamedHeartsHuntingGround, 0)(l, c)
}

func (r RobinTheHuntress) LovePinataSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.WarpToUntamedHearts
	}
	return nil
}

func (r RobinTheHuntress) WarpToWeddingExit(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.WeddingExitMap, 0)(l, c)
}

func (r RobinTheHuntress) ValidateUntamedHearts(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.GoldenMapleLeaf) {
		return r.MissingGoldenMapleLeaf(l, c)
	}
	return r.WarpToUntamedHearts(l, c)
}

func (r RobinTheHuntress) MissingGoldenMapleLeaf(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It seems like you lost your ").
		BlueText().ShowItemName1(item.GoldenMapleLeaf).
		BlackText().AddText(". I'm sorry, but I can't let you proceed to the hunting grounds without that.")
	return SendOk(l, c, m.String())
}

func (r RobinTheHuntress) ValidateSmashBoxes(l logrus.FieldLogger, c Context) State {
	if !character.HasItems(l)(c.CharacterId, item.GoldenKey, 7) {
		return r.MissingKeys(l, c)
	}
	return r.ProcessSmashBoxes(l, c)
}

func (r RobinTheHuntress) ProcessSmashBoxes(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.GoldenKey, -7)
	return WarpById(_map.TheLovePinata, 0)(l, c)
}

func (r RobinTheHuntress) MissingKeys(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It seems like you don't have 7 Keys. Kill the cakes and candles in the Untamed Heart Hunting Ground to get keys.")
	return SendOk(l, c, m.String())
}
