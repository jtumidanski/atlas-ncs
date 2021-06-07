package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ThiefJobInstructor is located in Victoria Road - Construction Site North of Kerning City (102040000)
type ThiefJobInstructor struct {
}

func (r ThiefJobInstructor) NPCId() uint32 {
	return npc.ThiefJobInstructor
}

func (r ThiefJobInstructor) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestCompleted(l)(c.CharacterId, 100010) {
		return r.TrueHero(l, c)
	} else if character.QuestCompleted(l)(c.CharacterId, 100009) {
		return r.LetYouIn(l, c)
	} else if character.QuestStarted(l)(c.CharacterId, 100009) {
		return r.IsntThisALetter(l, c)
	}
	return r.OnceYouAreReady(l, c)
}

func (r ThiefJobInstructor) TrueHero(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You're truly a hero!")
	return SendOk(l, c, m.String())
}

func (r ThiefJobInstructor) LetYouIn(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Alright I'll let you in! Defeat the monsters inside, collect 30 Dark Marbles, then strike up a conversation with a colleague of mine inside. He'll give you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(", the proof that you've passed the test. Best of luck to you.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r ThiefJobInstructor) Warp(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.ThiefsConstructionSite, 0)(l, c)
}

func (r ThiefJobInstructor) IsntThisALetter(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh, isn't this a letter from ").
		BlueText().AddText("Dark Lord").
		BlackText().AddText("?")
	return SendNext(l, c, m.String(), r.ProveYourSkills)
}

func (r ThiefJobInstructor) ProveYourSkills(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So you want to prove your skills? Very well...")
	return SendNextPrevious(l, c, m.String(), r.IfYouAreReady, r.IsntThisALetter)
}

func (r ThiefJobInstructor) IfYouAreReady(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I will give you a chance if you're ready.")
	return SendYesNo(l, c, m.String(), r.Begin, Exit())
}

func (r ThiefJobInstructor) Begin(l logrus.FieldLogger, c Context) State {
	character.CompleteQuest(l)(c.CharacterId, 100009)
	character.StartQuest(l)(c.CharacterId, 100010)
	character.GainItem(l)(c.CharacterId, item.DarkLordsLetter, -1)
	return r.GoodLuck(l, c)
}

func (r ThiefJobInstructor) OnceYouAreReady(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I can show you the way once your ready for it.")
	return SendOk(l, c, m.String())
}

func (r ThiefJobInstructor) GoodLuck(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You will have to collect me ").
		BlueText().AddText("30 ").ShowItemName1(item.DarkMarble).
		BlackText().AddText(". Good luck.")
	return SendOk(l, c, m.String())
}