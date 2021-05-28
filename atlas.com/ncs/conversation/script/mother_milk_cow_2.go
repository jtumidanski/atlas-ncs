package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MotherMilkCow2 is located in Hidden Chamber - The Nautilus - Stable (912000100)
type MotherMilkCow2 struct {
}

func (r MotherMilkCow2) NPCId() uint32 {
	return npc.MotherMilkCow2
}

func (r MotherMilkCow2) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestProgressInt(l)(c.CharacterId, 2180, 1) == 2 {
		return r.CheckAnotherCow(l, c)
	}
	if character.HasItem(l)(c.CharacterId, item.MilkJug) {
		return r.ProcessEmpty(l, c)
	} else if character.HasItem(l)(c.CharacterId, item.MilkJugOneThird) {
		return r.ProcessOneThird(l, c)
	} else if character.HasItem(l)(c.CharacterId, item.MilkJugTwoThird) {
		return r.ProcessTwoThird(l, c)
	}
	return Exit()(l, c)
}

func (r MotherMilkCow2) CheckAnotherCow(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You have taken milk from this cow recently, check another cow.")
	return SendOk(l, c, m.String())
}

func (r MotherMilkCow2) ProcessEmpty(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.MilkJug, -1)
	character.GainItem(l)(c.CharacterId, item.MilkJugOneThird, 1)
	character.SetQuestProgress(l)(c.CharacterId, 2180, 1, 2)
	m := message.NewBuilder().
		AddText("Now filling up the bottle with milk. The bottle is now 1/3 full of milk.")
	return SendNext(l, c, m.String(), Exit())
}

func (r MotherMilkCow2) ProcessOneThird(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.MilkJugOneThird, -1)
	character.GainItem(l)(c.CharacterId, item.MilkJugTwoThird, 1)
	character.SetQuestProgress(l)(c.CharacterId, 2180, 1, 2)
	m := message.NewBuilder().
		AddText("Now filling up the bottle with milk. The bottle is now 2/3 full of milk.")
	return SendNext(l, c, m.String(), Exit())
}

func (r MotherMilkCow2) ProcessTwoThird(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.MilkJugTwoThird, -1)
	character.GainItem(l)(c.CharacterId, item.MilkJugFull, 1)
	character.SetQuestProgress(l)(c.CharacterId, 2180, 1, 2)
	m := message.NewBuilder().
		AddText("Now filling up the bottle with milk. The bottle is now completely full of milk.")
	return SendNext(l, c, m.String(), Exit())
}
