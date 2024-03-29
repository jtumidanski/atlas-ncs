package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/event"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// LordJonathan is located in The Nautilus - Lord Jonathan's Room (120000102)
type LordJonathan struct {
}

func (r LordJonathan) NPCId() uint32 {
	return npc.LordJonathan
}

func (r LordJonathan) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 6400) {
		return r.BotherSomeoneElse(l, span, c)
	}
	return r.AQuestion(l, span, c)
}

func (r LordJonathan) BotherSomeoneElse(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Who are you talking to me? If you're just bored, go bother somebody else.")
	return script.SendOk(l, span, c, m.String())
}

func (r LordJonathan) AQuestion(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	progress := quest.ProgressInt(l)(c.CharacterId, 6400, 1)
	if progress == 0 {
		return r.FirstQuestion(l, span, c)
	} else if progress == 1 {
		return r.NextQuestion(l, span, c)
	} else {
		return r.Impressive(l, span, c)
	}
}

func (r LordJonathan) FirstQuestion(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ok then! I'll give you the first question now! You better be ready because this one's a hard one. Even the seagulls here think this one's pretty tough. It's a pretty difficult problem.")
	return script.SendNext(l, span, c, m.String(), r.IssueFirstQuestion)
}

func (r LordJonathan) IssueFirstQuestion(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("One day, I went to the ocean and caught 62 Octopi for dinner. But then some kid came by and gave me 10 Octopi as a gift! How many Octopi do I have then, in total?")
	return script.SendGetText(l, span, c, m.String(), r.OctopiResponse)
}

func (r LordJonathan) OctopiResponse(text string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if text != "72" {
			return r.NotRight(l, span, c)
		}
		return r.FirstResponseCorrect(l, span, c)
	}
}

func (r LordJonathan) NotRight(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmm, that's not quite how I recall it. Try again!")
	return script.SendOk(l, span, c, m.String())
}

func (r LordJonathan) FirstResponseCorrect(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	quest.SetProgress(l)(c.CharacterId, 6400, 1, 1)
	m := message.NewBuilder().
		AddText("What! I can't believe how incredibly smart you are! Incredible! In the seagull world, that kind of intelligence would give you a Ph.D. and then some. You're really amazing... I can't believe it... I simply can't believe it!")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r LordJonathan) NextQuestion(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Now~ Let's go onto the next question. This one is really difficult. I am going to have Bart help me on this one. You know Bart, right?")
	return script.SendNext(l, span, c, m.String(), r.TestOfWill)
}

func (r LordJonathan) TestOfWill(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I'm going to send you to an empty room in The Nautilus. You will see 9 Barts there. Hahaha~ Are they twins? No, no, certainly not. I've used a bit of magic for this test of will.")
	return script.SendNextPrevious(l, span, c, m.String(), r.TruePirate, r.NextQuestion)
}

func (r LordJonathan) TruePirate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Anyway, only one of 9 Barts is the real Bart. You know that Pirates are known for the strength of their friendships and camaraderie with their fellow pirates. If you're a true pirate, you should be able to find your own mate with ease. Alright then, I'll send you to the room where Bart is.")
	return script.SendNextPrevious(l, span, c, m.String(), r.StartBartEvent, r.TestOfWill)
}

func (r LordJonathan) StartBartEvent(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !event.StartEvent(l)(c.CharacterId, "4jaerial") {
		return r.AnotherPlayerChallenging(l, span, c)
	}
	return script.Exit()(l, span, c)
}

func (r LordJonathan) AnotherPlayerChallenging(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Another player is already challenging the test in this channel. Please try another channel, or wait for the current player to finish.")
	return script.SendOk(l, span, c, m.String())
}

func (r LordJonathan) Impressive(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ohhhh! Now that was impressive! I considered my test quite difficult, and for you to pass that... you are indeed an integral member of the Pirate family, and a friend of seagulls. We are now bonded by the mutual friendship that will last a lifetime! And, most of all, friends are there to help you out when you are in dire straits. If you are in a state of emergency, call us seagulls.")
	return script.SendNext(l, span, c, m.String(), r.AirStrike)
}

func (r LordJonathan) AirStrike(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Notify us using the skill Air Strike, and we will be there to help you out, because that's what friends are for.")
	return script.SendNextPrevious(l, span, c, m.String(), r.Passed, r.Impressive)
}

func (r LordJonathan) Passed(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You have met all my challenges, and passed! Good job!")
	return script.SendNextPrevious(l, span, c, m.String(), script.Exit(), r.AirStrike)
}
