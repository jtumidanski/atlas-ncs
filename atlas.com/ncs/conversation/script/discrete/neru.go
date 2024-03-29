package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/pet"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Neru is located in Ludibrium - Ludibrium Pet Walkway (220000006)
type Neru struct {
}

func (r Neru) NPCId() uint32 {
	return npc.Neru
}

func (r Neru) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.WeaversLetter) {
		return r.BrothersLeter(l, span, c)
	}
	return r.Chill(l, span, c)
}

func (r Neru) BrothersLeter(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Eh, that's my brother's letter! Probably scolding me for thinking I'm not working and stuff...Eh? Ahhh...you followed my brother's advice and trained your pet and got up here, huh? Nice!! Since you worked hard to get here, I'll boost your intimacy level with your pet.")
	return script.SendNext(l, span, c, m.String(), r.Validate)
}

func (r Neru) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !pet.HasPets(l)(c.CharacterId) {
		return r.GetOutOfHere(l, span, c)
	}
	return r.RaiseCloseness(l, span, c)
}

func (r Neru) GetOutOfHere(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm ... did you really get here with your pet? These obstacles are for pets. What are you here for without it?? Get outta here!")
	return script.SendOk(l, span, c, m.String())
}

func (r Neru) RaiseCloseness(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.WeaversLetter, -1)
	pet.GainCloseness(l)(c.CharacterId, 4)
	return r.TrainAgain(l, span, c)
}

func (r Neru) TrainAgain(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("What do you think? Don't you think you have gotten much closer with your pet? If you have time, train your pet again on this obstacle course...of course, with my brother's permission.")
	return script.SendOk(l, span, c, m.String())
}

func (r Neru) Chill(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("My brother told me to take care of the pet obstacle course, but ... since I'm so far away from him, I can't help but wanting to goof around ...hehe, since I don't see him in sight, might as well just chill for a few minutes.")
	return script.SendOk(l, span, c, m.String())
}
