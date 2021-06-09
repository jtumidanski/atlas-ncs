package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/event"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"strconv"
)

// Aura is located in Adobis's Mission I - Unknown Dead Mine (280010000)
type Aura struct {
}

func (r Aura) NPCId() uint32 {
	return npc.Aura
}

func (r Aura) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if event.Cleared(l)(c.CharacterId) {
		return r.ReceivePrize(l, c)
	} else {
		return r.Hello(l, c)
	}
}

func (r Aura) Hello(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("...").NewLine().
		OpenItem(0).BlueText().AddText("What am I supposed to do here?").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("I brought items!").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("I want to get out!").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Aura) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.WhatAmISupposedToDo
	case 1:
		return r.IBroughtItems
	case 2:
		return r.GetOut
	}
	return nil
}

func (r Aura) WhatAmISupposedToDo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("To reveal the power of Zakum, you'll have to recreate its core. Hidden somewhere in this dungeon is a ").
		BlueText().AddText("Fire Ore").
		BlackText().AddText(" which is one of the necessary materials for that core. Find it, and bring it to me.").NewLine().NewLine().
		AddText("Oh, and could you do me a favour? There's also a number of ").
		BlueText().AddText("Paper Documents").
		BlackText().AddText(" lying under rocks around here. If you can get 30 of them, I can reward you for your efforts.")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r Aura) IBroughtItems(l logrus.FieldLogger, c script.Context) script.State {
	if !event.Leader(l)(c.CharacterId) {
		return r.HaveLeaderBring(l, c)
	}
	if !character.HasItem(l)(c.CharacterId, item.FireOre) {
		return r.BringFireOre(l, c)
	}
	if character.HasItems(l)(c.CharacterId, item.PaperDocument, 30) {
		return r.EachMemberGetsAPiece2(l, c)
	}
	return r.EachMemberGetsAPiece(l, c)
}

func (r Aura) HaveLeaderBring(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please let your leader bring the materials to me to complete this ordeal.")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r Aura) BringFireOre(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please bring the ").
		BlueText().AddText("Fire Ore").
		BlackText().AddText(" with you.")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r Aura) EachMemberGetsAPiece2(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So, you brought the fire ore and the documents with you? In that case, I can give to you and to each member of your party a piece of it, that should be more than enough to make the core of Zakum. As well, since you ").
		RedText().AddText("brought the documents").
		BlackText().AddText(" with you, I can also provide you a special item which will ").
		BlueText().AddText("bring you to the mine's entrance at any time").
		BlackText().AddText(". Make sure your whole party has room in their inventory before proceeding.")
	return script.SendYesNo(l, c, m.String(), r.Process(true), script.Exit())
}

func (r Aura) EachMemberGetsAPiece(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So, you brought the fire ore with you? In that case, I can give to you and to each member of your party a piece of it, that should be more than enough to make the core of Zakum. Make sure your whole party has room in their inventory before proceeding.")
	return script.SendYesNo(l, c, m.String(), r.Process(false), script.Exit())
}

func (r Aura) Process(hasAll bool) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		character.GainItem(l)(c.CharacterId, item.FireOre, -1)
		if hasAll {
			character.GainItem(l)(c.CharacterId, item.PaperDocument, -30)
			event.SetProperty(l)(strconv.Itoa(int(c.NPCId)), "gotDocuments", "1")
			event.GiveParticipantsExperience(l)(20000)
		} else {
			event.GiveParticipantsExperience(l)(12000)
		}
		event.Clear(l)
		return script.Exit()(l, c)
	}
}

func (r Aura) GetOut(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Are you sure you want to exit? If you're the party leader, your party will also be removed from the mines.")
	return script.SendYesNo(l, c, m.String(), r.WarpExit, script.Exit())
}

func (r Aura) WarpExit(l logrus.FieldLogger, c script.Context) script.State {
	//TODO warp rest of party?
	return script.WarpById(_map.TheDoorToZakum, 0)(l, c)
}

func (r Aura) ReceivePrize(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You completed this ordeal, now receive your prize.")
	return script.SendNext(l, c, m.String(), r.ProcessPrize)
}

func (r Aura) ProcessPrize(l logrus.FieldLogger, c script.Context) script.State {
	if event.GetProperty(l)(strconv.Itoa(int(c.NPCId)), "gotDocuments") == "1" {
		return r.GotDocumentsPrize(l, c)
	}
	return r.NoDocumentsPrize(l, c)
}

func (r Aura) GotDocumentsPrize(l logrus.FieldLogger, c script.Context) script.State {
	if !event.ReceivedReward(l)(c.CharacterId) {
		if !character.CanHoldThese(l)(c.CharacterId, character.Item{ItemId: item.ReturnScrollToDeadMine, Quantity: 5}, character.Item{ItemId: item.PieceOfFireOre, Quantity: 1}) {
			return r.MakeRoom(l, c)
		}
		character.GainItem(l)(c.CharacterId, item.ReturnScrollToDeadMine, 5)
		character.GainItem(l)(c.CharacterId, item.PieceOfFireOre, 1)
		event.SetRewardReceived(l)(c.CharacterId)
		return script.Exit()(l, c)
	}
	return r.AlreadyReceived(l, c)
}

func (r Aura) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Make sure you have room in your inventory before proceeding.")
	return script.SendOk(l, c, m.String())
}

func (r Aura) AlreadyReceived(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You have already received your share. You can now exit the mines through the portal over there.")
	return script.SendOk(l, c, m.String())
}

func (r Aura) NoDocumentsPrize(l logrus.FieldLogger, c script.Context) script.State {
	if !event.ReceivedReward(l)(c.CharacterId) {
		if !character.CanHoldAll(l)(c.CharacterId, item.PieceOfFireOre, 1) {
			return r.MakeRoom(l, c)
		}
		character.GainItem(l)(c.CharacterId, item.PieceOfFireOre, 1)
		event.SetRewardReceived(l)(c.CharacterId)
		return script.Exit()(l, c)
	}
	return r.AlreadyReceived(l, c)
}
