package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Cloy is located in Victoria Road : Henesys Park (100000200)
type Cloy struct {
}

func (r Cloy) NPCId() uint32 {
	return npc.Cloy
}

func (r Cloy) Initial(l logrus.FieldLogger, c Context) State {
	return r.AskQuestions(l, c)
}

func (r Cloy) AskQuestions(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmm... are you raising one of my kids by any chance? I perfected a spell that uses Water of Life to blow life into a doll. People call it the ").
		BlueText().AddText("Pet").
		BlackText().AddText(". If you have one with you, feel free to ask me questions.")
	return SendNext(l, c, m.String(), r.WhatDoYouWantToKnow)
}

func (r Cloy) WhatDoYouWantToKnow(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("What do you want to know more of?").AddNewLine().BlueText().
		OpenItem(0).AddText("Tell me more about Pets.").CloseItem().AddNewLine().
		OpenItem(1).AddText("How do I raise Pets?").CloseItem().AddNewLine().
		OpenItem(2).AddText("Do Pets die too?").CloseItem().AddNewLine().
		OpenItem(3).AddText("What are the commands for Brown and Black Kitty?").CloseItem().AddNewLine().
		OpenItem(4).AddText("What are the commands for Brown Puppy?").CloseItem().AddNewLine().
		OpenItem(5).AddText("What are the commands for Pink and White Bunny?").CloseItem().AddNewLine().
		OpenItem(6).AddText("What are the commands for Mini Kargo?").CloseItem().AddNewLine().
		OpenItem(7).AddText("What are the commands for Rudolph and Dasher?").CloseItem().AddNewLine().
		OpenItem(8).AddText("What are the commands for Black Pig?").CloseItem().AddNewLine().
		OpenItem(9).AddText("What are the commands for Panda?").CloseItem().AddNewLine().
		OpenItem(10).AddText("What are the commands for Husky?").CloseItem().AddNewLine().
		OpenItem(11).AddText("What are the commands for Dino Boy and Dino Girl?").CloseItem().AddNewLine().
		OpenItem(12).AddText("What are the commands for Monkey?").CloseItem().AddNewLine().
		OpenItem(13).AddText("What are the commands for Turkey?").CloseItem().AddNewLine().
		OpenItem(14).AddText("What are the commands for White Tiger?").CloseItem().AddNewLine().
		OpenItem(15).AddText("What are the commands for Penguin?").CloseItem().AddNewLine().
		OpenItem(16).AddText("What are the commands for Golden Pig?").CloseItem().AddNewLine().
		OpenItem(17).AddText("What are the commands for Robot?").CloseItem().AddNewLine().
		OpenItem(18).AddText("What are the commands for Mini Yeti?").CloseItem().AddNewLine().
		OpenItem(19).AddText("What are the commands for Jr. Balrog?").CloseItem().AddNewLine().
		OpenItem(20).AddText("What are the commands for Baby Dragon?").CloseItem().AddNewLine().
		OpenItem(21).AddText("What are the commands for Green/Red/Blue Dragon?").CloseItem().AddNewLine().
		OpenItem(22).AddText("What are the commands for Black Dragon?").CloseItem().AddNewLine().
		OpenItem(23).AddText("What are the commands for Jr. Reaper?").CloseItem().AddNewLine().
		OpenItem(24).AddText("What are the commands for Porcupine?").CloseItem().AddNewLine().
		OpenItem(25).AddText("What are the commands for Snowman?").CloseItem().AddNewLine().
		OpenItem(26).AddText("What are the commands for Skunk?").CloseItem().AddNewLine().
		OpenItem(27).AddText("Please teach me about transferring pet ability points.").CloseItem()
	return SendListSelection(l, c, m.String(), r.InfoSelection)
}

func (r Cloy) InfoSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.MoreAboutPets
	case 1:
		return r.Commands
	case 2:
		return r.Dying
	case 3:
		return r.Kitty
	case 4:
		return r.Puppy
	case 5:
		return r.Bunny
	case 6:
		return r.Kargo
	case 7:
		return r.Rudolph
	case 8:
		return r.BlackPig
	case 9:
		return r.Panda
	case 10:
		return r.Husky
	case 11:
		return r.Dino
	case 12:
		return r.Monkey
	case 13:
		return r.Turkey
	case 14:
		return r.WhiteTiger
	case 15:
		return r.Penguin
	case 16:
		return r.GoldenPig
	case 17:
		return r.Robot
	case 18:
		return r.MiniYeti
	case 19:
		return r.JrBalrog
	case 20:
		return r.BabyDragon
	case 21:
		return r.ColoredDragon
	case 22:
		return r.BlackDragon
	case 23:
		return r.JrReaper
	case 24:
		return r.Porcupine
	case 25:
		return r.Snowman
	case 26:
		return r.Skunk
	case 27:
		return r.Transfer
	}
	return nil
}

func (r Cloy) MoreAboutPets(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So you want to know more about Pets. Long ago I made a doll, sprayed Water of Life on it, and cast spell on it to create a magical animal. I know it sounds unbelievable, but it's a doll that became an actual living thing. They understand and follow people very well.")
	return SendNext(l, c, m.String(), r.CantGiveTooMuchLife)
}

func (r Cloy) OhYeah(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh yeah, they'll react when you give them special commands. You can scold them, love them... it all").AddNewLine().
		AddText("depends on how you take care of them. They are afraid to leave their masters so be nice to them, show them love. They can get sad and lonely fast...")
	return SendNextPrevious(l, c, m.String(), Exit(), r.CantGiveTooMuchLife)
}

func (r Cloy) Commands(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Depending on the command you give, pets can love it, hate, and display other kinds of reactions to it. If you give the pet a command and it follows you well, your intimacy goes up. Double click on the pet and you can check the intimacy, level, fullness and etc...")
	return SendNext(l, c, m.String(), r.TryHardRaisingIt)
}

func (r Cloy) Dying(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Dying... well, they aren't technically ALIVE per se, so I don't know if dying is the right term to use. They are dolls with my magical power and the power of Water of Life to become a live object. Of course while it's alive, it's just like a live animal...")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Kitty(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Brown Kitty and Black Kitty").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid, ihateyou, dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1~30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, say, chat").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("cutie").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand, rise").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Puppy(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Brown Puppy").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid, ihateyou, baddog, dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1~30)").AddNewLine().
		BlueText().AddText("pee").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, say, chat").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("down").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand, rise").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Bunny(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Pink Bunny and White Bunny").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand, rise").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1~30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, say, chat").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("hug").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("sleep, sleepy, gotobed").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Kargo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Mini Kargo").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand, rise").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1~30)").AddNewLine().
		BlueText().AddText("pee").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, say, chat").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("thelook, charisma").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("down").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("goodboy, goodgirl").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Rudolph(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Rudolph and Dasher").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid, ihateyou, dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("merryxmas, merrychristmas").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1~30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, say, chat").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("lonely, alone").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("cutie").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("mush, go").
		BlackText().AddText(" (Level 21 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) BlackPig(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Black Pig").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1~30)").AddNewLine().
		BlueText().AddText("hand").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid, ihateyou, dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("smile").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("thelook, charisma").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Panda(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Panda").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("chill, relax").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand, rise").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("letsplay").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("meh, bleh").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("sleep").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Husky(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Husky").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid, ihateyou, baddog, dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("hand").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("down").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand, rise").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Dino(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Dino Boy and Dino Girl").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badboy, badgirl").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("smile, laugh").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid, ihateyou, dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("cutie").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("sleep, nap, sleepy").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Monkey(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Monkey").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("rest").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badboy, badgirl").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("pee").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("play").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("melong").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("sleep, gotobed, sleepy").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Turkey(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Turkey").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("no, rudeboy, mischief").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, gobble").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("yes, goodboy").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("sleepy, birdnap, doze").
		BlackText().AddText(" (Level 20 ~ 30)").AddNewLine().
		BlueText().AddText("birdeye, thanksgiving, fly, friedbird, imhungry").
		BlackText().AddText(" (Level 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) WhiteTiger(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("White Tiger").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badboy, badgirl").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("rest, chill").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid, ihateyou, dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("actsad, sadlook").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("wait").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Penguin(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Penguin").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badboy, badgirl").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand, rise").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("hug, hugme").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("wing, hand").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("sleep").
		BlackText().AddText(" (Level 20 ~ 30)").AddNewLine().
		BlueText().AddText("kiss, smooch, muah").
		BlackText().AddText(" (Level 20 ~ 30)").AddNewLine().
		BlueText().AddText("fly").
		BlackText().AddText(" (Level 20 ~ 30)").AddNewLine().
		BlueText().AddText("cute, adorable").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) GoldenPig(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Golden Pig").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badboy, badgirl").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("loveme, hugme").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("sleep, sleepy, gotobed").
		BlackText().AddText(" (Level 21 ~ 30)").AddNewLine().
		BlueText().AddText("ignore / impressed / outofhere").
		BlackText().AddText(" (Level 21 ~ 30)").AddNewLine().
		BlueText().AddText("roll, showmethemoney").
		BlackText().AddText(" (Level 21 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Robot(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Robot").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("up, stand, rise").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid, ihateyou, dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("attack, charge").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("good, thelook, charisma").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("speack, talk, chat, say").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("disguise, change, transform").
		BlackText().AddText(" (Level 11 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) MiniYeti(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("1012005_MINI_YETI=These are the commands for ").
		RedText().AddText("Mini Yeti").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad, no, badboy, badgirl").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("dance, boogie, shakeit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("cute, cutie, pretty, adorable").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou, likeyou, mylove").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("sleep, nap, sleepy, gotobed").
		BlackText().AddText(" (Level 11 ~ 30)\n")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) JrBalrog(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Jr. Balrog").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("liedown").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("no|bad|badgirl|badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou|mylove|likeyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("cute|cutie|pretty|adorable").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("smirk|crooked|laugh").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("melong").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("good|thelook|charisma").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("speak|talk|chat|say").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("sleep|nap|sleepy").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("gas").
		BlackText().AddText(" (Level 21 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) BabyDragon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Baby Dragon").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("no|bad|badgirl|badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou|loveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid|ihateyou|dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("cutie").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("talk|chat|say").
		BlackText().AddText(" (Level 11 ~ 30)").AddNewLine().
		BlueText().AddText("sleep|sleepy|gotobed").
		BlackText().AddText(" (Level 11 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) ColoredDragon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Green/Red/Blue Dragon").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("no|bad|badgirl|badboy").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou|loveyou").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("stupid|ihateyou|dummy").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("talk|chat|say").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("sleep|sleepy|gotobed").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("change").
		BlackText().AddText(" (Level 21 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) BlackDragon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Black Dragon").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("no|bad|badgirl|badboy").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou|loveyou").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("poop").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("stupid|ihateyou|dummy").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("talk|chat|say").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("sleep|sleepy|gotobed").
		BlackText().AddText(" (Level 15 ~ 30)").AddNewLine().
		BlueText().AddText("cutie, change").
		BlackText().AddText(" (Level 21 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) JrReaper(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Jr. Reaper").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("no|bad|badgirl|badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("playdead, poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk|chat|say").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou, hug").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("smellmyfeet, rockout, boo").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("trickortreat").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("monstermash").
		BlackText().AddText(" (Level 1 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Porcupine(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Porcupine").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("no|bad|badgirl|badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("iloveyou|hug|goodboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk|chat|say").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("cushion|sleep|knit|poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("comb|beach").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("treeninja").
		BlackText().AddText(" (Level 20 ~ 30)").AddNewLine().
		BlueText().AddText("dart").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Snowman(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Snowman").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("stupid, ihateyou, dummy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("loveyou, mylove, ilikeyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("merrychristmas").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("cutie, adorable, cute, pretty").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("comb, beach/bad, no, badgirl, badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk, chat, say/sleep, sleepy, gotobed").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("chang").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Skunk(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These are the commands for ").
		RedText().AddText("Skunk").
		BlackText().AddText(". The level mentioned next to the command shows the pet level required for it to respond.").AddNewLine().
		BlueText().AddText("sit").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("bad/no/badgirl/badboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("restandrelax, poop").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("talk/chat/say, iloveyou").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("snuggle/hug, sleep, goodboy").
		BlackText().AddText(" (Level 1 ~ 30)").AddNewLine().
		BlueText().AddText("fatty, blind, badbreath").
		BlackText().AddText(" (Level 10 ~ 30)").AddNewLine().
		BlueText().AddText("suitup, bringthefunk").
		BlackText().AddText(" (Level 20 ~ 30)")
	return SendNext(l, c, m.String(), Exit())
}

func (r Cloy) Transfer(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("In order to transfer the pet ability points, closeness and level, Pet AP Reset Scroll is required. If you take this").AddNewLine().
		AddText("scroll to Mar the Fairy in Ellinia, she will transfer the level and closeness of the pet to another one. I am especially giving it to you because I can feel your heart for your pet. However, I can't give this out for free. I can give you this book for 250,000 mesos. Oh, I almost forgot! Even if you have this book, it is no use if you do not have a new pet to transfer the Ability points.")
	return SendNext(l, c, m.String(), r.PurchaseConfirmation)
}

func (r Cloy) PurchaseConfirmation(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("250,000 mesos will be deducted. Do you really want to buy?")
	return SendYesNo(l, c, m.String(), r.ValidatePurchase, Exit())
}

func (r Cloy) ValidatePurchase(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 250000) || !character.CanHold(l)(c.CharacterId, item.PetAPResetScroll) {
		return r.TransactionFailure(l, c)
	}

	err := character.GainMeso(l)(c.CharacterId, -250000)
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve mesos for purchase by character %d.", c.CharacterId)
	}
	character.GainItem(l)(c.CharacterId, item.PetAPResetScroll, 1)
	return Exit()(l, c)
}

func (r Cloy) TransactionFailure(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Please check if your inventory has empty slot or you don't have enough mesos.")
	return SendOk(l, c, m.String())
}

func (r Cloy) NeedPetFood(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh yes! Pets can't eat the normal human food. Instead my disciple ").
		BlueText().AddText("Doofus").
		BlackText().AddText(" sells ").
		BlueText().AddText("Pet Food").
		BlackText().AddText(" at the Henesys Market so if you need food for your pet, find Henesys. It'll be a good idea to buy the food in advance and feed the pet before it gets really hungry.")
	return SendNextPrevious(l, c, m.String(), r.RegularBasis, nil)
}

func (r Cloy) RegularBasis(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh, and if you don't feed the pet for a long period of time, it goes back home by itself. You can take it out of its home and feed it but it's not really good for the pet's health, so try feeding him on a regular basis so it doesn't go down to that level, alright? I think this will do.")
	return SendNextPrevious(l, c, m.String(), Exit(), r.NeedPetFood)
}

func (r Cloy) CantGiveTooMuchLife(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("But Water of Life only comes out little at the very bottom of the World Tree, so I can't give him too much time in life... I know, it's very unfortunate... but even if it becomes a doll again I can always bring life back into it so be good to it while you're with it.")
	return SendNextPrevious(l, c, m.String(), r.OhYeah, r.MoreAboutPets)
}

func (r Cloy) TryHardRaisingIt(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Talk to the pet, pay attention to it and its intimacy level will go up and eventually his overall level will go up too. As the intimacy level rises, the pet's overall level will rise soon after. As the overall level rises, one day the pet may even talk like a person a little bit, so try hard raising it. Of course it won't be easy doing so...")
	return SendNextPrevious(l, c, m.String(), r.Commands, r.HaveHunger)
}

func (r Cloy) HaveHunger(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It may be a live doll but they also have life so they can feel the hunger too. ").
		BlueText().AddText("Fullness").
		BlackText().AddText(" shows the level of hunger the pet's in. 100 is the max, and the lower it gets, it means that the pet is getting hungrier. After a while, it won't even follow your command and be on the offensive, so watch out over that.")
	return SendNextPrevious(l, c, m.String(), r.NeedPetFood, r.TryHardRaisingIt)
}

func (r Cloy) AfterSomeTime(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("After some time... that's correct, they stop moving. They just turn back to being a doll, after the effect of magic dies down and Water of Life dries out. But that doesn't mean it's stopped forever, because once you pour Water of Life over, it's going to be back alive.")
	return SendNextPrevious(l, c, m.String(), r.SadToSee, nil)
}

func (r Cloy) SadToSee(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Even if it someday moves again, it's sad to see them stop altogether. Please be nice to them while they are alive and moving. Feed them well, too. Isn't it nice to know that there's something alive that follows and listens to only you?")
	return SendNextPrevious(l, c, m.String(), Exit(), r.AfterSomeTime)
}