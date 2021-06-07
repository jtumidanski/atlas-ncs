package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Lukan is located in Phantom Forest - Phantom Road (610010003)
type Lukan struct {
}

func (r Lukan) NPCId() uint32 {
	return npc.Lukan
}

func (r Lukan) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestCompleted(l)(c.CharacterId, 8223) {
		return r.BraveAdventurer(l, c)
	}

	if !character.CanHold(l)(c.CharacterId, item.CrimsonwoodKeystone) {
		return r.NeedInventoryRoom(l, c)
	}

	return r.GiveKeystone(l, c)
}

func (r Lukan) GiveKeystone(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.CrimsonwoodKeystone, 1)
	return r.DoNotLoseAgain(l, c)
}

func (r Lukan) DoNotLoseAgain(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So you lost your key, right? Very well, I will craft you another one, but please don't lose it again. It is fundamental to enter the Inner Sanctum, inside the Keep.")
	return SendOk(l, c, m.String())
}

func (r Lukan) NeedInventoryRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Please make a slot on your SETUP ready for the key I have to give to you. It is fundamental to enter the Inner Sanctum, inside the Keep.")
	return SendOk(l, c, m.String())
}

func (r Lukan) BraveAdventurer(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("O, brave adventurer. The Stormcasters house, from which I belong, guards the surrounding area of Yore, this landscape, from the forces of the Twisted Masters' guard that daily threatens the citizens. Please help us on the defense of Yore.")
	return SendOk(l, c, m.String())
}
