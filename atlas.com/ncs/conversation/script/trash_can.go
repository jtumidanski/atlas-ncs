package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// TrashCan is located in The Nautilus - Top Floor - Hallway (120000100)
type TrashCan struct {
}

func (r TrashCan) NPCId() uint32 {
	return npc.TrashCan
}

func (r TrashCan) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestCompleted(l)(c.CharacterId, 2162) && !character.HasItem(l)(c.CharacterId, item.CrumpledLetter) {
		return r.Validate(l, c)
	}
	return Exit()(l, c)
}

func (r TrashCan) Validate(l logrus.FieldLogger, c Context) State {
	if !character.CanHold(l)(c.CharacterId, item.CrumpledLetter) {
		return r.InventoryFull(l, c)
	}
	return r.Give(l, c)
}

func (r TrashCan) InventoryFull(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("(You see a Crumpled Paper standing out of the trash can. It's content seems important, but you can't retrieve it since your inventory is full.)")
	return SendNext(l, c, m.String(), Exit())
}

func (r TrashCan) Give(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.CrumpledLetter, 1)
	m := message.NewBuilder().
		AddText("(You retrieved a Crumpled Paper standing out of the trash can. It's content seems important.)")
	return SendNext(l, c, m.String(), Exit())
}
