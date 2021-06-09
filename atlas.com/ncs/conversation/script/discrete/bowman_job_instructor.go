package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// BowmanJobInstructor is located in Warning Street - The Road to the Dungeon (106010000)
type BowmanJobInstructor struct {
}

func (r BowmanJobInstructor) NPCId() uint32 {
	return npc.BowmanJobInstructor
}

func (r BowmanJobInstructor) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.QuestCompleted(l)(c.CharacterId, 100001) {
		return r.TrueHero(l, c)
	} else if character.QuestCompleted(l)(c.CharacterId, 100000) {
		return r.LetYouIn(l, c)
	} else if character.QuestStarted(l)(c.CharacterId, 100000) {
		return r.IsntThisALetter(l, c)
	}
	return r.OnceYouAreReady(l, c)
}

func (r BowmanJobInstructor) TrueHero(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You're truly a hero!")
	return script.SendOk(l, c, m.String())
}

func (r BowmanJobInstructor) LetYouIn(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright I'll let you in! Defeat the monsters inside, collect 30 Dark Marbles, then strike up a conversation with a colleague of mine inside. He'll give you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(", the proof that you've passed the test. Best of luck to you.")
	return script.SendNext(l, c, m.String(), script.WarpById(_map.AntTunnelForBowman, 0))
}

func (r BowmanJobInstructor) IsntThisALetter(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Oh, isn't this a letter from ").
		BlueText().AddText("Athena").
		BlackText().AddText("?")
	return script.SendNext(l, c, m.String(), r.ProveYourSkills)
}

func (r BowmanJobInstructor) ProveYourSkills(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So you want to prove your skills? Very well...")
	return script.SendNextPrevious(l, c, m.String(), r.IfYouAreReady, r.IsntThisALetter)
}

func (r BowmanJobInstructor) IfYouAreReady(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I will give you a chance if you're ready.")
	return script.SendYesNo(l, c, m.String(), r.Begin, script.Exit())
}

func (r BowmanJobInstructor) Begin(l logrus.FieldLogger, c script.Context) script.State {
	character.CompleteQuest(l)(c.CharacterId, 100000)
	character.StartQuest(l)(c.CharacterId, 100001)
	character.GainItem(l)(c.CharacterId, item.AthenaPiercesLetter, -1)
	return r.GoodLuck(l, c)
}

func (r BowmanJobInstructor) OnceYouAreReady(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I can show you the way once your ready for it.")
	return script.SendOk(l, c, m.String())
}

func (r BowmanJobInstructor) GoodLuck(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You will have to collect me ").
		BlueText().AddText("30 ").ShowItemName1(item.DarkMarble).
		BlackText().AddText(". Good luck.")
	return script.SendOk(l, c, m.String())
}