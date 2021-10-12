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

// OlsonTheToySoldier is located in Ludibrium - Eos Tower 100th Floor (221024400)
type OlsonTheToySoldier struct {
}

func (r OlsonTheToySoldier) NPCId() uint32 {
	return npc.OlsonTheToySoldier
}

func (r OlsonTheToySoldier) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 3230) {
		return r.NotAllowed(l, span, c)
	}

	if event.GetProperty(l)("DollHouse", "noEntry") != "false" {
		return r.AlreadyChallenging(l, span, c)
	}
	return r.PendulumInside(l, span, c)
}

func (r OlsonTheToySoldier) NotAllowed(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("We are not allowed to let the general public wander past this point.")
	return script.SendOk(l, span, c, m.String())
}

func (r OlsonTheToySoldier) AlreadyChallenging(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Someone else is already searching the area. Please wait until the area is cleared.")
	return script.SendOk(l, span, c, m.String())
}

func (r OlsonTheToySoldier) PendulumInside(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The pendulum is hidden inside a dollhouse that looks different than the others.")
	return script.SendNext(l, span, c, m.String(), r.AreYouReady)
}

func (r OlsonTheToySoldier) AreYouReady(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Are you ready to enter the dollhouse map?")
	return script.SendYesNo(l, span, c, m.String(), r.Start, script.Exit())
}

func (r OlsonTheToySoldier) Start(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	ok := event.StartEvent(l)(c.CharacterId, "DollHouse")
	if !ok {
		return r.Hmm(l, span, c)
	}
	return script.Exit()(l, span, c)
}

func (r OlsonTheToySoldier) Hmm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hmm... The DollHouse is being challenged already, it seems. Try again later.")
	return script.SendOk(l, span, c, m.String())
}
