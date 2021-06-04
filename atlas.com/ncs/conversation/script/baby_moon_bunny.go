package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// BabyMoonBunny is located in Hidden Street - Lunar World (922230000)
type BabyMoonBunny struct {
}

func (r BabyMoonBunny) NPCId() uint32 {
	return npc.BabyMoonBunny
}

func (r BabyMoonBunny) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.Henesys {
		m := message.NewBuilder().
			AddText("There! Did you see that? You didn't? A UFO just passed... there!! Look, someone is getting dragged into the UFO... arrrrrrgh, it's Gaga! ").
			RedText().AddText("Gaga just got kidnapped by a UFO!")
		return SendNext(l, c, m.String(), r.Validate)
	}
	return Exit()(l, c)
}

func (r BabyMoonBunny) Validate(l logrus.FieldLogger, c Context) State {
	if !character.IsLevel(l)(c.CharacterId, 12) {
		return r.LevelRequirement(l, c)
	}
	return r.WhatDoWeDo(l, c)
}

func (r BabyMoonBunny) WhatDoWeDo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("What do we do now? It's just a rumor yet, but... I've heard that scary things happen to you if you get kidnapped by aliens... may be that's what happening to Gaga right now! Please, please rescue Gaga! \\r\\n #bGaga may be a bit undetermined and clueless, but#k he has a really good heart. I can't let something terrible happen to him. Right! Grandpa from the moon might know how to rescue him! I will send you to the moon, so please go meet Grandpa and rescue Gaga!!!")
	return SendYesNo(l, c, m.String(), r.ThankYou, Exit())
}

func (r BabyMoonBunny) LevelRequirement(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh! It seems you don't reach the level requirements to save Gaga. Please come back when you are level 12 or higher.")
	return SendOk(l, c, m.String())
}

func (r BabyMoonBunny) ThankYou(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Thank you so much. Please rescue Gaga! Grandpa from the moon will help you.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r BabyMoonBunny) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.MoonCorner, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.MoonCorner, c.NPCId)
	}
	return Exit()(l, c)
}
