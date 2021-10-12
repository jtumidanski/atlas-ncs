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

// ThiefJobInstructor is located in Victoria Road - Construction Site North of Kerning City (102040000)
type ThiefJobInstructor struct {
}

func (r ThiefJobInstructor) NPCId() uint32 {
	return npc.ThiefJobInstructor
}

func (r ThiefJobInstructor) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsCompleted(l)(c.CharacterId, 100010) {
		return r.TrueHero(l, span, c)
	} else if quest.IsCompleted(l)(c.CharacterId, 100009) {
		return r.LetYouIn(l, span, c)
	} else if quest.IsStarted(l)(c.CharacterId, 100009) {
		return r.IsntThisALetter(l, span, c)
	}
	return r.OnceYouAreReady(l, span, c)
}

func (r ThiefJobInstructor) TrueHero(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You're truly a hero!")
	return script.SendOk(l, span, c, m.String())
}

func (r ThiefJobInstructor) LetYouIn(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright I'll let you in! Defeat the monsters inside, collect 30 Dark Marbles, then strike up a conversation with a colleague of mine inside. He'll give you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(", the proof that you've passed the test. Best of luck to you.")
	return script.SendNext(l, span, c, m.String(), r.Warp)
}

func (r ThiefJobInstructor) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.ThiefsConstructionSite, 0)(l, span, c)
}

func (r ThiefJobInstructor) IsntThisALetter(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Oh, isn't this a letter from ").
		BlueText().AddText("Dark Lord").
		BlackText().AddText("?")
	return script.SendNext(l, span, c, m.String(), r.ProveYourSkills)
}

func (r ThiefJobInstructor) ProveYourSkills(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So you want to prove your skills? Very well...")
	return script.SendNextPrevious(l, span, c, m.String(), r.IfYouAreReady, r.IsntThisALetter)
}

func (r ThiefJobInstructor) IfYouAreReady(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I will give you a chance if you're ready.")
	return script.SendYesNo(l, span, c, m.String(), r.Begin, script.Exit())
}

func (r ThiefJobInstructor) Begin(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	quest.Complete(l)(c.CharacterId, 100009)
	quest.Start(l)(c.CharacterId, 100010)
	character.GainItem(l, span)(c.CharacterId, item.DarkLordsLetter, -1)
	return r.GoodLuck(l, span, c)
}

func (r ThiefJobInstructor) OnceYouAreReady(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I can show you the way once your ready for it.")
	return script.SendOk(l, span, c, m.String())
}

func (r ThiefJobInstructor) GoodLuck(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You will have to collect me ").
		BlueText().AddText("30 ").ShowItemName1(item.DarkMarble).
		BlackText().AddText(". Good luck.")
	return script.SendOk(l, span, c, m.String())
}