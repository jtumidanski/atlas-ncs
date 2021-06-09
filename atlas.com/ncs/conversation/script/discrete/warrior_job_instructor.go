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

// WarriorJobInstructor is located in Victoria Road - West Rocky Mountain IV (102020300)
type WarriorJobInstructor struct {
}

func (r WarriorJobInstructor) NPCId() uint32 {
	return npc.WarriorJobInstructor
}

func (r WarriorJobInstructor) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.QuestCompleted(l)(c.CharacterId, 100004) {
		return r.TrueHero(l, c)
	} else if character.QuestCompleted(l)(c.CharacterId, 100003) {
		return r.LetYouIn(l, c)
	} else if character.QuestStarted(l)(c.CharacterId, 100003) {
		return r.ExplainTest(l, c)
	}
	return r.OnceYouAreReady(l, c)
}

func (r WarriorJobInstructor) TrueHero(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You're truly a hero!")
	return script.SendOk(l, c, m.String())
}

func (r WarriorJobInstructor) LetYouIn(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright I'll let you in! Defeat the monsters inside, collect 30 Dark Marbles, then strike up a conversation with a colleague of mine inside. He'll give you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(", the proof that you've passed the test. Best of luck to you.")
	return script.SendNext(l, c, m.String(), r.Warp)
}

func (r WarriorJobInstructor) Warp(l logrus.FieldLogger, c script.Context) script.State {
	return script.WarpById(_map.WarriorsRockyMountain, 0)(l, c)
}

func (r WarriorJobInstructor) ExplainTest(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm...it is definitely the letter from ").
		BlueText().AddText("Dances with Balrog").
		BlackText().AddText("...so you came all the way here to take the test and make the 2nd job advancement as the warrior. Alright, I'll explain the test to you. Don't sweat it too much, it's not that complicated.")
	return script.SendNext(l, c, m.String(), r.SendYou)
}

func (r WarriorJobInstructor) SendYou(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I'll send you to a hidden map. You'll see monsters you don't normally see. They look the same like the regular ones, but with a totally different attitude. They neither boost your experience level nor provide you with item.")
	return script.SendNextPrevious(l, c, m.String(), r.AcquireMarble, r.ExplainTest)
}

func (r WarriorJobInstructor) AcquireMarble(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You'll be able to acquire a marble called ").
		BlueText().ShowItemName1(item.DarkMarble).
		BlackText().AddText(" while knocking down those monsters. It is a special marble made out of their sinister, evil minds. Collect 30 of those, and then go talk to a colleague of mine in there. That's how you pass the test.")
	return script.SendNextPrevious(l, c, m.String(), r.CannotLeave, r.SendYou)
}

func (r WarriorJobInstructor) CannotLeave(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Once you go inside, you can't leave until you take care of your mission. If you die, your experience level will decrease..so you better really buckle up and get ready...well, do you want to go for it now?")
	return script.SendYesNo(l, c, m.String(), r.Begin, script.Exit())
}

func (r WarriorJobInstructor) Begin(l logrus.FieldLogger, c script.Context) script.State {
	character.CompleteQuest(l)(c.CharacterId, 100003)
	character.StartQuest(l)(c.CharacterId, 100004)
	character.GainItem(l)(c.CharacterId, item.DancesWithBalrogsLetter, -1)
	return r.LetYouIn(l, c)
}

func (r WarriorJobInstructor) OnceYouAreReady(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I can show you the way once your ready for it.")
	return script.SendOk(l, c, m.String())
}
