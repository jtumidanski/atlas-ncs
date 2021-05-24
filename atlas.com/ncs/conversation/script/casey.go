package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Casey is located in Victoria Road - Henesys Game Park (100000203)
type Casey struct {
}

func (r Casey) NPCId() uint32 {
	return npc.Casey
}

func (r Casey) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Casey) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hey, you look like you need a breather. You should be enjoying the life, just like I am. Well, if you have a couple of items, I can trade you for an item you can play mini games with. Now... what can I do for you?").AddNewLine().
		OpenItem(0).BlueText().AddText("Create a mini game item").CloseItem().AddNewLine().
		OpenItem(1).BlueText().AddText("Explain to me what the mini games are about").CloseItem().AddNewLine()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r Casey) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.MakeMiniGame
	case 1:
		return r.LearnMore
	}
	return nil
}

func (r Casey) MakeMiniGame(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You want to make the mini game item? Mini games aren't something you can just go ahead and play right off the bat. For each mini game, you'll need a specific set of items. Which minigame it em do you want to make?").AddNewLine().
		OpenItem(0).BlueText().AddText("Omok Set").CloseItem().AddNewLine().
		OpenItem(1).BlueText().AddText("A Set of Match Cards").CloseItem().AddNewLine()
	return SendListSelection(l, c, m.String(), r.MiniGameProducer)
}

func (r Casey) MiniGameProducer(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.MakeOmok
	case 1:
		return r.MakeMatchCards
	}
	return nil
}

func (r Casey) LearnMore(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You want to learn more about the mini games? Awesome! Ask me anything. Which mini game do you want to know more about?").AddNewLine().
		OpenItem(0).BlueText().AddText("Omok").CloseItem().AddNewLine().
		OpenItem(1).BlueText().AddText("Match Cards").CloseItem().AddNewLine()
	return SendListSelection(l, c, m.String(), r.LearnMoreProducer)
}

func (r Casey) LearnMoreProducer(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.OmokRules
	case 1:
		return r.MatchCardRules
	}
	return nil
}

func (r Casey) OmokRules(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Here are the rules for Omok, so listen carefully. Omok is a game in which you and your opponent take turns laying a piece on the table until someone finds a way to lay 5 consecutive pieces in a line, be it horizontal, diagonal, or vertical. For starters, only the ones with an ").
		BlueText().AddText("Omok Set").
		BlackText().AddText(" can open a game room.")
	return SendNext(l, c, m.String(), r.OmokCost)
}

func (r Casey) MatchCardRules(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Here are the rules for Match Cards, so listen carefully. As the name suggests, Match Cards is simply finding a matching pair among the number of cards laid on the table. When all the matching pairs are found, then the person with more matching pairs will win the game. Just like Omok, you'll need ").
		BlueText().AddText("A set of Match Cards").
		BlackText().AddText(" to open the game room.")
	return SendNext(l, c, m.String(), r.MatchCardCost)
}

func (r Casey) OmokCost(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Every game of Omok will cost you ").
		RedText().AddText("100 mesos").
		BlackText().AddText(". Even if you don't have an ").
		BlueText().AddText("Omok Set").
		BlackText().AddText(", you can enter the room and play. However, if you don't possess 100 mesos, then you won't be allowed to enter in the room at all. The person opening the game room also needs 100 mesos to open the room (or else there's no game). If you run out of mesos during the game, then you're automatically kicked out of the room!")
	return SendNextPrevious(l, c, m.String(), r.HowToStart(r.OmokStart, r.OmokCost), r.OmokRules)
}

func (r Casey) MatchCardCost(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Every game of Match Cards will cost you ").
		RedText().AddText("100 mesos").
		BlackText().AddText(". Even if you don't have ").
		BlueText().AddText("A set of Match Cards").
		BlackText().AddText(", you can enter the room and play. However, if you don't possess 100 mesos, then you won't be allowed to enter in the room at all. The person opening the game room also needs 100 mesos to open the room (or else there's no game). If you run out of mesos during the game, then you're automatically kicked out of the room!")
	return SendNextPrevious(l, c, m.String(), r.HowToStart(r.MatchCardStart, r.MatchCardCost), r.MatchCardRules)
}

func (r Casey) HowToStart(next StateProducer, previous StateProducer) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("Enter the room, and when you're ready to play, click on ").
			BlueText().AddText("Ready").
			BlackText().AddText(".").AddNewLine().
			AddText("Once the visitor clicks on ").
			BlueText().AddText("Ready").
			BlackText().AddText(", the room owner can press ").
			BlueText().AddText("Start").
			BlackText().AddText(" to begin the game. If an unwanted visitor walks in, and you don't want to play with that person, the room owner has the right to kick the visitor out of the room. There will be a square box with x written on the right of that person. Click on that for a cold goodbye, okay?")
		return SendNextPrevious(l, c, m.String(), next, previous)
	}
}

func (r Casey) OmokStart(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("When the first fame starts, ").
		BlueText().AddText("the room owner goes first").
		BlackText().AddText(". Be warned that you'll be given a time limit, and you may lose your turn if you don't make your move on time. Normally, 3 x 3 is not allowed, but if there comes a point that it's absolutely necessary to put your piece there or face ending the game, then you can put it there. 3 x 3 is allowed as the last line of defense! Oh, and it won't count if it's ").
		RedText().AddText("6 or 7 straight").
		BlackText().AddText(". Only 5!")
	return SendNextPrevious(l, c, m.String(), r.RedoOrTie, r.HowToStart(r.OmokStart, r.OmokCost))
}

func (r Casey) MatchCardStart(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh, and unlike Omok, when you create the game room for Match Cards, you'll need to set your game on the number of cards you'll use for the game. There are 3 modes available, 3x4, 4x5, and 5x6, which will require 12, 20, and 30 cards respectively. Remember that you won't be able to change it up once the room is open, so if you really wish to change it up, you may have to close the room and open another one.")
	return SendNextPrevious(l, c, m.String(), r.MoveOnTime, r.HowToStart(r.MatchCardStart, r.MatchCardCost))
}

func (r Casey) RedoOrTie(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you know your back is against the wall, you can request a ").
		BlueText().AddText("Redo").
		BlackText().AddText(". If the opponent accepts your request, then you and your opponent's last moves will cancel out. If you ever feel the need to go to the bathroom, or take an extended break, you can request a #btie#k. The game will end in a tie if the opponent accepts the request. Tip: this may be a good way to keep your friendships in tact.")
	return SendNextPrevious(l, c, m.String(), r.NextGame(r.RedoOrTie), r.OmokStart)
}

func (r Casey) MoveOnTime(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("When the first game starts, ").
		BlueText().AddText("the room owner goes first.").
		BlackText().AddText(" Beware that you'll be given a time limit, and you may lose your turn if you don't make your move on time. When you find a matching pair on your turn, you'll get to keep your turn, as long as you keep finding a pair of matching cards. Use your memorizing skills to make a streak.")
	return SendNextPrevious(l, c, m.String(), r.LongerStreak, r.MatchCardStart)
}

func (r Casey) NextGame(previous StateProducer) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("When the next game starts, the loser will go first. Also, no one is allowed to leave in the middle of a game. If you do, you may need to request either a ").
			BlueText().AddText("forfeit or tie").
			BlackText().AddText(". (Of course, if you request a forfeit, you'll lose the game.) And if you click on 'Leave' in the middle of the game and call to leave after the game, you'll leave the room right after the game is over. This will be a much more useful way to leave.")
		return SendPrevious(l, c, m.String(), previous)
	}
}

func (r Casey) LongerStreak(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you and your opponent have the same number of matched pairs, then whoever had a longer streak of matched pairs will win. If you ever feel the need to go to the bathroom, or take an extended break, you can request a ").
		BlueText().AddText("tie").
		BlackText().AddText(". The game will end in a tie if the opponent accepts the request. Tip: this may be a good way to keep your friendships in tact.")
	return SendNextPrevious(l, c, m.String(), r.NextGame(r.LongerStreak), r.MoveOnTime)
}

func (r Casey) MakeOmok(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You want to play ").
		BlueText().AddText("Omok").
		BlackText().AddText(", huh? To play it, you'll need the Omok Set. Only the ones with that item can open the room for a game of Omok, and you can play this game almost anywhere except for a few places at the market place.")
	return SendNext(l, c, m.String(), r.WhichSet)
}

func (r Casey) WhichSet(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The set also differs based on what kind of pieces you want to use for the game. Which set would you like to make?").AddNewLine().
		OpenItem(0).BlueText().ShowItemName1(item.SlimeAndMushroomOmokSet).CloseItem().AddNewLine().BlackText().
		OpenItem(1).BlueText().ShowItemName1(item.SlimeAndOctopusOmokSet).CloseItem().AddNewLine().BlackText().
		OpenItem(2).BlueText().ShowItemName1(item.SlimeAndPigOmokSet).CloseItem().AddNewLine().BlackText().
		OpenItem(3).BlueText().ShowItemName1(item.OctopusAndMushroomOmokSet).CloseItem().AddNewLine().BlackText().
		OpenItem(4).BlueText().ShowItemName1(item.PigAndOctopusOmokSet).CloseItem().AddNewLine().BlackText().
		OpenItem(5).BlueText().ShowItemName1(item.PigAndMushroomOmokSet).CloseItem().AddNewLine().BlackText()
	return SendListSelection(l, c, m.String(), r.OmokSetSelection)
}

func (r Casey) OmokSetSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.ProcessSelection(item.SlimeAndMushroomOmokSet, r.SlimeAndMushroomOmokSetRequirements())
	case 1:
		return r.ProcessSelection(item.SlimeAndOctopusOmokSet, r.SlimeAndOctopusOmokSetRequirements())
	case 2:
		return r.ProcessSelection(item.SlimeAndPigOmokSet, r.SlimeAndPigOmokSetRequirements())
	case 3:
		return r.ProcessSelection(item.OctopusAndMushroomOmokSet, r.OctopusAndMushroomOmokSetRequirements())
	case 4:
		return r.ProcessSelection(item.PigAndOctopusOmokSet, r.PigAndOctopusOmokSetRequirements())
	case 5:
		return r.ProcessSelection(item.PigAndMushroomOmokSet, r.PigAndMushroomOmokSetRequirements())
	}
	return nil
}

type SetRequirements struct {
	requirements []SetRequirement
}

type SetRequirement struct {
	itemId uint32
}

func (r Casey) GetTheMaterials(itemId uint32, requirement1 uint32, requirement2 uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			BlueText().AddText("You want to make ").
			ShowItemName1(itemId).
			BlackText().AddText("? Hmm...get me the materials, and I can do just that. Listen carefully, the materials you need will be: ").
			RedText().AddText(fmt.Sprintf("%d ", 99)).ShowItemName1(requirement1).
			AddText(", ").AddText(fmt.Sprintf("%d ", 99)).ShowItemName1(requirement2).
			AddText(", 1 ").ShowItemName1(item.OmokTable).
			BlackText().AddText(". The monsters will probably drop those every once in a while...")
		return SendNext(l, c, m.String(), Exit())
	}
}

func (r Casey) ProcessSelection(itemId uint32, requirements SetRequirements) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		for _, req := range requirements.requirements {
			if !character.HasItem(l)(c.CharacterId, req.itemId) {
				return r.GetTheMaterials(itemId, requirements.requirements[0].itemId, requirements.requirements[1].itemId)(l, c)
			}
		}
		if !character.HasItem(l)(c.CharacterId, item.OmokTable) {
			return r.GetTheMaterials(itemId, requirements.requirements[0].itemId, requirements.requirements[1].itemId)(l, c)
		}

		character.GainItem(l)(c.CharacterId, requirements.requirements[0].itemId, -99)
		character.GainItem(l)(c.CharacterId, requirements.requirements[1].itemId, -99)
		character.GainItem(l)(c.CharacterId, item.OmokTable, -1)
		character.GainItem(l)(c.CharacterId, itemId, 1)
		return Exit()(l, c)
	}
}

func (r Casey) SlimeAndMushroomOmokSetRequirements() SetRequirements {
	return SetRequirements{requirements: []SetRequirement{{itemId: 4030000}, {itemId: 4030001}}}
}

func (r Casey) SlimeAndOctopusOmokSetRequirements() SetRequirements {
	return SetRequirements{requirements: []SetRequirement{{itemId: 4030000}, {itemId: 4030010}}}
}

func (r Casey) SlimeAndPigOmokSetRequirements() SetRequirements {
	return SetRequirements{requirements: []SetRequirement{{itemId: 4030000}, {itemId: 4030011}}}
}

func (r Casey) OctopusAndMushroomOmokSetRequirements() SetRequirements {
	return SetRequirements{requirements: []SetRequirement{{itemId: 4030010}, {itemId: 4030001}}}
}

func (r Casey) PigAndOctopusOmokSetRequirements() SetRequirements {
	return SetRequirements{requirements: []SetRequirement{{itemId: 4030011}, {itemId: 4030010}}}
}

func (r Casey) PigAndMushroomOmokSetRequirements() SetRequirements {
	return SetRequirements{requirements: []SetRequirement{{itemId: 4030011}, {itemId: 4030001}}}
}

func (r Casey) MakeMatchCards(l logrus.FieldLogger, c Context) State {
	if !character.HasItems(l)(c.CharacterId, item.MonsterCard, 15) {
		return r.MatchCardNeeds(l, c)
	}
	return r.ProcessMatchCardCreation(l, c)
}

func (r Casey) MatchCardNeeds(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You want ").
		BlueText().AddText("A set of Match Cards").
		BlackText().AddText("? Hmm...to make A set of Match Cards, you'll need some ").
		BlueText().AddText("Monster Cards").
		BlackText().AddText(". Monster Card can be obtained by taking out the monsters all around the island. Collect 15 Monster Cards and you can make a set of A set of Match Cards.")
	return SendNext(l, c, m.String(), Exit())
}

func (r Casey) ProcessMatchCardCreation(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.MonsterCard, -15)
	character.GainItem(l)(c.CharacterId, item.ASetOfMatchCards, 1)
	return Exit()(l, c)
}
