package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// TrainerBartos is located in Victoria Road - Pet-Walking Road (100000202)
type TrainerBartos struct {
}

func (r TrainerBartos) NPCId() uint32 {
	return npc.TrainerBartos
}

func (r TrainerBartos) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r TrainerBartos) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Do you have any business with me?").NewLine().
		OpenItem(0).BlueText().AddText("Please tell me about this place.").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("I'm here through a word from Mar the Fairy...").CloseItem().NewLine()
	return script.SendListSelection(l, span, c, m.String(), r.WhatBusiness)
}

func (r TrainerBartos) WhatBusiness(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.AboutThisPlace
	case 1:
		return r.WordFromMar
	}
	return nil
}

func (r TrainerBartos) AboutThisPlace(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.BartosLetter) {
		return r.GetThatLetter(l, span, c)
	}
	return r.WannaTrainYourPet(l, span, c)
}

func (r TrainerBartos) WordFromMar(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey, are you sure you've met ").
		BlueText().AddText("Mar the Fairy").
		BlackText().AddText("? Don't lie to me if you've never met her before because it's obvious. That wasn't even a good lie!!")
	return script.SendOk(l, span, c, m.String())
}

func (r TrainerBartos) WannaTrainYourPet(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This is the road where you can go take a walk with your pet. You can just walk around with it, or you can train your pet to go through the obstacles here. If you aren't too close with your pet yet, that may present a problem and he will not follow your command as much... So, what do you think? Wanna train your pet?")
	return script.SendYesNo(l, span, c, m.String(), r.GiveItem, r.TooBusy)
}

func (r TrainerBartos) GiveItem(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.BartosLetter, 1)
	m := message.NewBuilder().
		AddText("Ok, here's the letter. He wouldn't know I sent you if you just went there straight, so go through the obstacles with your pet, go to the very top, and then talk to Trainer Frod to give him the letter. It won't be hard if you pay attention to your pet while going through obstacles. Good luck!")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r TrainerBartos) TooBusy(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm ... too busy to do it right now? If you feel like doing it, though, come back and find me.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r TrainerBartos) GetThatLetter(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Get that letter, jump over obstacles with your pet, and take that letter to my brother Trainer Frod. Give him the letter and something good is going to happen to your pet.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}
