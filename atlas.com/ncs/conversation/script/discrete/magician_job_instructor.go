package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// MagicianJobInstructor is located in Victoria Road - The Forest North of Ellinia (101020000)
type MagicianJobInstructor struct {
}

func (r MagicianJobInstructor) NPCId() uint32 {
	return npc.MagicianJobInstructor
}

func (r MagicianJobInstructor) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsCompleted(l)(c.CharacterId, 100007) {
		return r.TrueHero(l, span, c)
	} else if quest.IsCompleted(l)(c.CharacterId, 100006) {
		return r.LetYouIn(l, span, c)
	} else if quest.IsStarted(l)(c.CharacterId, 100006) {
		return r.ExplainTest(l, span, c)
	}
	return r.OnceYouAreReady(l, span, c)
}

func (r MagicianJobInstructor) TrueHero(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You're truly a hero!")
	return script.SendOk(l, span, c, m.String())
}

func (r MagicianJobInstructor) LetYouIn(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright I'll let you in! Defeat the monsters inside, collect 30 Dark Marbles, then strike up a conversation with a colleague of mine inside. He'll give you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(", the proof that you've passed the test. Best of luck to you.")
	return script.SendNext(l, span, c, m.String(), script.WarpById(_map.MagiciansTreeDungeon, 0))
}

func (r MagicianJobInstructor) ExplainTest(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm...it is definitely the letter from ").
		BlueText().AddText("Grendell the Really Old").
		BlackText().AddText("...so you came all the way here to take the test and make the 2nd job advancement as a magician. Alright, I'll explain the test to you. Don't sweat it too much, it's not that complicated.")
	return script.SendNext(l, span, c, m.String(), r.SendYou)
}

func (r MagicianJobInstructor) SendYou(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I'll send you to a hidden map. You'll see monsters you don't normally see. They look the same like the regular ones, but with a totally different attitude. They neither boost your experience level nor provide you with item.")
	return script.SendNextPrevious(l, span, c, m.String(), r.AcquireMarble, r.ExplainTest)
}

func (r MagicianJobInstructor) AcquireMarble(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You'll be able to acquire a marble called ").
		BlueText().ShowItemName1(item.DarkMarble).
		BlackText().AddText(" while knocking down those monsters. It is a special marble made out of their sinister, evil minds. Collect 30 of those, and then go talk to a colleague of mine in there. That's how you pass the test.")
	return script.SendNextPrevious(l, span, c, m.String(), r.CannotLeave, r.SendYou)
}

func (r MagicianJobInstructor) CannotLeave(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Once you go inside, you can't leave until you take care of your mission. If you die, your experience level will decrease..so you better really buckle up and get ready...well, do you want to go for it now?")
	return script.SendYesNo(l, span, c, m.String(), r.Begin, script.Exit())
}

func (r MagicianJobInstructor) Begin(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	quest.Complete(l)(c.CharacterId, 100006)
	quest.Start(l)(c.CharacterId, 100007)
	character.GainItem(l, span)(c.CharacterId, item.GrendelTheReallyOldsLetter, -1)
	return r.LetYouIn(l, span, c)
}

func (r MagicianJobInstructor) OnceYouAreReady(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I can show you the way once your ready for it.")
	return script.SendOk(l, span, c, m.String())
}