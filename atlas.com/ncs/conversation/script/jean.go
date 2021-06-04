package script

import (
	"atlas-ncs/event"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Jean is located in Victoria Road - Lith Harbor (104000000)
type Jean struct {
}

func (r Jean) NPCId() uint32 {
	return npc.Jean
}

func (r Jean) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hey, I'm ").
		BlueText().ShowNPC(npc.Jean).
		BlackText().AddText(". I am waiting for my brother ").
		BlueText().ShowNPC(npc.Paul).
		BlackText().AddText(". He is supposed to be here by now...")
	return SendNext(l, c, m.String(), r.WhatShouldIDo)
}

func (r Jean) WhatShouldIDo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmm... What should I do? The event will start, soon... Many people went to participate in the event, so we better be hurry...")
	return SendNextPrevious(l, c, m.String(), r.GoWithMe, r.Initial)
}

func (r Jean) GoWithMe(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hey... Why don't you go with me? I think my brother will come with other people.").NewLine().
		OpenItem(0).BlueText().AddText("e1. What kind of an event is it?").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("e2. Explain the event game to me.").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("e3. Alright, let's go!").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r Jean) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Event
	case 1:
		return r.Explanation
	case 2:
		return r.Go
	}
	return nil
}

func (r Jean) Event(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("All this month, MapleStory Global is celebrating its 3rd anniversary! The GM's will be holding surprise GM Events throughout the event, so stay on your toes and make sure to participate in at least one of the events for great prizes!")
	return SendOk(l, c, m.String())
}

func (r Jean) Explanation(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("There are many games for this event. It will help you a lot to know how to play the game before you play it. Choose the one you want to know more of! ").NewLine().
		OpenItem(0).BlueText().AddText("Ola Ola").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("MapleStory Maple Physical Fitness Test").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Snow Ball").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Coconut Harvest").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("OX Quiz").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("Treasure Hunt").CloseItem()
	return SendListSelection(l, c, m.String(), r.Game)
}

func (r Jean) Game(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.OlaOla
	case 1:
		return r.Fitness
	case 2:
		return r.Snowball
	case 3:
		return r.Coconut
	case 4:
		return r.OxQuiz
	case 5:
		return r.TreasureHunt
	}
	return nil
}

func (r Jean) OlaOla(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		BlueText().AddText("[Ola Ola]").
		BlackText().AddText(" is a game where participants climb ladders to reach the top. Climb your way up and move to the next level by choosing the correct portal out of the numerous portals available.").NewLine().NewLine().
		AddText("The game consists of three levels, and the time limit is ").
		BlueText().AddText("6 MINUTES").
		BlackText().AddText(". During [Ola Ola], you ").
		BlueText().AddText("won't be able to jump, teleport, haste, or boost your speed using potions or items").
		BlackText().AddText(". There are also trick portals that'll lead you to a strange place, so please be aware of those.")
	return SendOk(l, c, m.String())
}

func (r Jean) Fitness(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		BlueText().AddText("[MapleStory Physical Fitness Test] is a race through an obstacle course").
		BlackText().AddText(" much like the Forest of Patience. You can win it by overcoming various obstacles and reach the final destination within the time limit. ").NewLine().NewLine().
		AddText("The game consists of four levels, and the time limit is ").
		BlueText().AddText("15 MINUTES").
		BlackText().AddText(". During [MapleStory Physical Fitness Test], you won't be able to use teleport or haste.")
	return SendOk(l, c, m.String())
}

func (r Jean) Snowball(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		BlueText().AddText("[Snowball]").
		BlackText().AddText(" consists of two teams, Maple Team and Story Team, and the two teams duke it out to see ").
		BlueText().AddText("which team rolled the snowball farther and bigger in a limited time").
		BlackText().AddText(". If the game cannot be decided within the time period, then the team that rolled the snowball farther wins. ").NewLine().NewLine().
		AddText("To roll up the snow, attack it by pressing ").
		BlueText().AddText("Ctrl").
		BlackText().AddText(". All long-ranged attacks and skill-based attacks will not work here, ").
		BlueText().AddText("only the close-range attacks will work").
		BlackText().AddText(". ").NewLine().NewLine().
		AddText("If a character touches the snowball, he/she'll be sent back to the starting point. Attack the snowman in front of the starting point to prevent the opposing team from rolling the snow forward. This is where a well-planned strategy works, as the team will decide whether to attack the snowball or the snowman.")
	return SendOk(l, c, m.String())
}

func (r Jean) Coconut(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		BlueText().AddText("[Coconut Harvest]").
		BlackText().AddText(" consists of two teams, Maple Team and Story Team, and the two teams duke it out to see ").
		BlueText().AddText("which team gathers up the most coconuts").
		BlackText().AddText(". The time limit is ").
		BlueText().AddText("5 MINUTES").
		BlackText().AddText(". If the game ends in a tie, an additional 2 minutes will be awarded to determine the winner. If, for some reason, the score stays tied, then the game will end in a draw. ").NewLine().NewLine().
		AddText("All long-range attacks and skill-based attacks will not work here, ").
		BlueText().AddText("only the close-range attacks will work").
		BlackText().AddText(". If you don't have a weapon for the close-range attacks, you can purchase them through an NPC within the event map. No matter the level of character, the weapon, or skills, all damages applied will be the same.").NewLine().NewLine().
		AddText("Beware of the obstacles and traps within the map. If the character dies during the game, the character will be eliminated from the game. The player who strikes last before the coconut drops wins. Only the coconuts that hit the ground counts, which means the ones that do not fall off the tree, or the occasional explosion of the coconuts WILL NOT COUNT. There's also a hidden portal at one of the shells at the bottom of the map, so use that wisely!")
	return SendOk(l, c, m.String())
}

func (r Jean) OxQuiz(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		BlueText().AddText("[OX Quiz]").
		BlackText().AddText(" is a game of MapleStory smarts through X's and O's. Once you join the game, turn on the mini map by pressing ").
		BlueText().AddText("M").
		BlackText().AddText(" to see where the X and O are. A total of ").
		RedText().AddText("10 questions").
		BlackText().AddText(" will be given, and the character that answers them all correctly wins the game. ").NewLine().NewLine().
		AddText("Once the question is given, use the ladder to enter the area where the correct answer may be, be it X or O. If the character does not choose an answer or is hanging on the ladder past the time limit, the character will be eliminated. Please hold your position until [CORRECT] is off the screen before moving on. To prevent cheating of any kind, all types of chatting will be turned off during the OX Quiz.")
	return SendOk(l, c, m.String())
}

func (r Jean) TreasureHunt(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		BlueText().AddText("[Treasure Hunt]").
		BlackText().AddText(" is a game in which your goal is to find the ").
		BlueText().AddText("treasure scrolls").
		BlackText().AddText(" that are hidden all over the map ").
		RedText().AddText("in 10 minutes").
		BlackText().AddText(". There will be a number of mysterious treasure chests hidden away, and once you break them apart, many items will surface from the chest. Your job is to pick out the treasure scroll from those items. ").NewLine().
		AddText("Treasure chests can be destroyed using ").
		BlueText().AddText("regular attacks").
		BlackText().AddText(", and once you have the treasure scroll in possession, you can trade it for the Scroll of Secrets through an NPC that's in charge of trading items. The trading NPC can be found on the Treasure Hunt map, but you can also trade your scroll through ").
		BlueText().AddText("Vikin").
		BlackText().AddText(" of Lith Harbor.").NewLine().NewLine().
		AddText("This game has its share of hidden portals and hidden teleporting spots. To use them, press the ").
		BlueText().AddText("up arrow").
		BlackText().AddText(" at a certain spot, and you'll be teleported to a different place. Try jumping around, for you may also run into hidden stairs or ropes. There will also be a treasure chest that'll take you to a hidden spot, and a hidden chest that can only be found through the hidden portal, so try looking around.").NewLine().NewLine().
		AddText("During the game of Treasure Hunt, all attack skills will be ").
		RedText().AddText("disabled").
		BlackText().AddText(", so please break the treasure chest with the regular attack.")
	return SendOk(l, c, m.String())
}

func (r Jean) Go(l logrus.FieldLogger, c Context) State {
	ok := event.JoinEvent(l)(c.WorldId, c.ChannelId, c.CharacterId)
	if ok {
		return Exit()(l, c)
	} else {
		return r.NotYetStarted(l, c)
	}
}

func (r Jean) NotYetStarted(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Either the event has not been started, you already have the ").
		BlueText().AddText("Scroll of Secrets").
		BlackText().AddText(", or you have already participated in this event within the last 24 hours. Please try again later!")
	return SendOk(l, c, m.String())
}
