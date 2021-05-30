package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// TutorialLilin is located in Snow Island - Ice Cave (140090000)
type TutorialLilin struct {
}

func (r TutorialLilin) NPCId() uint32 {
	return npc.TutorialLilin
}

func (r TutorialLilin) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.IceCave {
		if !character.AreaInfo(l)(c.CharacterId, 21019, "helper=clear") {
			return r.FinallyAwoken(l, c)
		} else {
			return r.AreYouAlright(l, c)
		}
	}
	return r.AnythingStillCurious(l, c)
}

func (r TutorialLilin) FinallyAwoken(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You've finally awoken...!")
	return SendNextSpeaker(l, c, m.String(), npc.SpeakerUnknown, r.AndYouAre)
}

func (r TutorialLilin) AndYouAre(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("And you are...?")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerCharacterLeft, r.TheHero, r.FinallyAwoken)
}

func (r TutorialLilin) TheHero(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The hero who fought against the Black Magician... I've been waiting for you to wake up!")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerUnknown, r.WhatAreYouTalkingAbout, r.AndYouAre)
}

func (r TutorialLilin) WhatAreYouTalkingAbout(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Who... Who are you? And what are you talking about?")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerCharacterLeft, r.AndWhoAmI, r.TheHero)
}

func (r TutorialLilin) AndWhoAmI(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("And who am I...? I can't remember anything... Ouch, my head hurts!")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerCharacterLeft, r.ShowIntro, r.WhatAreYouTalkingAbout)
}

func (r TutorialLilin) ShowIntro(l logrus.FieldLogger, c Context) State {
	character.ShowIntro(l)(c.CharacterId, "Effect/Direction1.img/aranTutorial/face")
	character.ShowIntro(l)(c.CharacterId, "Effect/Direction1.img/aranTutorial/ClickLilin")
	character.SetAreaInfo(l)(c.CharacterId, 21019, "helper=clear")
	return Exit()(l, c)
}

func (r TutorialLilin) AreYouAlright(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Are you alright?")
	return SendNextSpeaker(l, c, m.String(), npc.SpeakerUnknown, r.ICannotRemember)
}

func (r TutorialLilin) ICannotRemember(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I can't remember anything. Where am I? And who are you...?")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerCharacterLeft, r.StayCalm, r.AreYouAlright)
}

func (r TutorialLilin) StayCalm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Stay calm. There is no need to panic. You can't remember anything because the curse of the Black Magician erased your memory. I'll tell you everything you need to know...step by step.")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerUnknown, r.YouAreAHero, r.ICannotRemember)
}

func (r TutorialLilin) YouAreAHero(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You're a hero who fought the Black Magician and saved Maple World hundreds of years ago. But at the very last moment, the curse of the Black Mage put you to sleep for a long, long time. That's when you lost all of your memories.")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerUnknown, r.ThisIsRien, r.StayCalm)
}

func (r TutorialLilin) ThisIsRien(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("This island is called Rien, and it's where the Black Magician trapped you. Despite its name, this island is always covered in ice and snow because of the Black Magician's curse. You were found deep inside the Ice Cave.")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerUnknown, r.MyNameIs, r.YouAreAHero)
}

func (r TutorialLilin) MyNameIs(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("My name is Lilin and I belong to the clan of Rien. The Rien Clan has been waiting for a hero to return for a long time now, and we finally found you. You've finally returned!")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerUnknown, r.SaidTooMuch, r.ThisIsRien)
}

func (r TutorialLilin) SaidTooMuch(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I've said too much. It's okay if you don't really understand everything I just told you. You'll get it eventually. For now, ").
		BlueText().AddText("you should head to town").
		BlackText().AddText(". I'll stay by your side and help you until you get there.")
	return SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerUnknown, r.Warp, r.MyNameIs)
}

func (r TutorialLilin) Warp(l logrus.FieldLogger, c Context) State {
	character.SpawnGuide(l)(c.CharacterId)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ColdForest1, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ColdForest1, c.NPCId)
	}
	return Exit()(l, c)
}

func (r TutorialLilin) AnythingStillCurious(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Is there anything you're still curious about? If so, I'll try to explain it better. ").NewLine().
		OpenItem(0).BlueText().AddText("Who am I?").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Where am I?").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Who are you?").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Tell me what I have to do.").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Tell me about my Inventory.").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("How do I advance my skills?").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("I want to know how to equip items.").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("How do I use quick slots?").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("How can I open breakable containers?").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("I want to sit in a chair but I forgot how.").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r TutorialLilin) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.WhoAmI
	case 1:
		return r.WhereAmI
	case 2:
		return r.WhoAreYou
	case 3:
		return r.GetToTown
	case 4:
		return r.GuideHint(14)
	case 5:
		return r.GuideHint(15)
	case 6:
		return r.GuideHint(16)
	case 7:
		return r.GuideHint(17)
	case 8:
		return r.GuideHint(18)
	case 9:
		return r.GuideHint(19)
	}
	return nil
}

func (r TutorialLilin) WhoAmI(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You are one of the heroes that saved Maple World from the Black Magician hundreds of years ago. You've lost your memory due to the curse of the Black Mage.")
	return SendOk(l, c, m.String())
}

func (r TutorialLilin) WhereAmI(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("This island is called Rien, and this is where the Black Magician's curse put you to sleep. It's a small island covered in ice and snow, and the majority of the residents are Penguins.")
	return SendOk(l, c, m.String())
}

func (r TutorialLilin) WhoAreYou(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I'm Lilin, a clan member of Rien, and I've been waiting for your return as the prophecy foretold. I'll be your guide for now.")
	return SendOk(l, c, m.String())
}

func (r TutorialLilin) GetToTown(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Let's not waste any more time and just get to town. I'll give you the details when we get there.")
	return SendOk(l, c, m.String())
}

func (r TutorialLilin) GuideHint(hint uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.GuideHint(l)(c.CharacterId, hint)
		return Exit()(l, c)
	}
}
