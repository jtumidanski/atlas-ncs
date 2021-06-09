package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
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

func (r BabyMoonBunny) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if c.MapId == _map.Henesys {
		m := message.NewBuilder().
			AddText("There! Did you see that? You didn't? A UFO just passed... there!! Look, someone is getting dragged into the UFO... arrrrrrgh, it's Gaga! ").
			RedText().AddText("Gaga just got kidnapped by a UFO!")
		return script.SendNext(l, c, m.String(), r.Validate)
	}
	return script.Exit()(l, c)
}

func (r BabyMoonBunny) Validate(l logrus.FieldLogger, c script.Context) script.State {
	if !character.IsLevel(l)(c.CharacterId, 12) {
		return r.LevelRequirement(l, c)
	}
	return r.WhatDoWeDo(l, c)
}

func (r BabyMoonBunny) WhatDoWeDo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("What do we do now? It's just a rumor yet, but... I've heard that scary things happen to you if you get kidnapped by aliens... may be that's what happening to Gaga right now! Please, please rescue Gaga! \\r\\n #bGaga may be a bit undetermined and clueless, but#k he has a really good heart. I can't let something terrible happen to him. Right! Grandpa from the moon might know how to rescue him! I will send you to the moon, so please go meet Grandpa and rescue Gaga!!!")
	return script.SendYesNo(l, c, m.String(), r.ThankYou, script.Exit())
}

func (r BabyMoonBunny) LevelRequirement(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Oh! It seems you don't reach the level requirements to save Gaga. Please come back when you are level 12 or higher.")
	return script.SendOk(l, c, m.String())
}

func (r BabyMoonBunny) ThankYou(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Thank you so much. Please rescue Gaga! Grandpa from the moon will help you.")
	return script.SendNext(l, c, m.String(), script.WarpById(_map.MoonCorner, 0))
}
