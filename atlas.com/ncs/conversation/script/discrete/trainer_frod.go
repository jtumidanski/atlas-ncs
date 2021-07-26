package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/pet"
	"github.com/sirupsen/logrus"
)

// TrainerFrod is located in Victoria Road - Pet-Walking Road (100000202)
type TrainerFrod struct {
}

func (r TrainerFrod) NPCId() uint32 {
	return npc.TrainerFrod
}

func (r TrainerFrod) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.HasItem(l)(c.CharacterId, item.BartosLetter) {
		return r.BrothersLetter(l, c)
	}
	return r.BrotherToldMe(l, c)
}

func (r TrainerFrod) BrothersLetter(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Eh, that's my brother's letter! Probably scolding me for thinking I'm not working and stuff...Eh? Ahhh...you followed my brother's advice and trained your pet and got up here, huh? Nice!! Since you worked hard to get here, I'll boost your intimacy level with your pet.")
	return script.SendNext(l, c, m.String(), r.Action)
}

func (r TrainerFrod) BrotherToldMe(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("My brother told me to take care of the pet obstacle course, but ... since I'm so far away from him, I can't help but wanting to goof around ...hehe, since I don't see him in sight, might as well just chill for a few minutes.")
	return script.SendOk(l, c, m.String())
}

func (r TrainerFrod) Action(l logrus.FieldLogger, c script.Context) script.State {
	if pet.HasPets(l)(c.CharacterId) {
		return r.CompleteCourse(l, c)
	}
	return r.CompleteCourseNoPets(l, c)
}

func (r TrainerFrod) CompleteCourseNoPets(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm ... did you really get here with your pet? These obstacles are for pets. What are you here for without it?? Get outta here!")
	return script.SendNextPrevious(l, c, m.String(), script.Exit(), r.BrothersLetter)
}

func (r TrainerFrod) CompleteCourse(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.BartosLetter, -1)
	pet.GainCloseness(l)(c.CharacterId, 2)
	m := message.NewBuilder().
		AddText("What do you think? Don't you think you have gotten much closer with your pet? If you have time, train your pet again on this obstacle course...of course, with my brother's permission.")
	return script.SendNextPrevious(l, c, m.String(), script.Exit(), r.BrothersLetter)
}
