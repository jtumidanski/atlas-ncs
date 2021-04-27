package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Robin is located in Maple Road : Snail Hunting Ground I (40000)
type Robin struct {
}

func (r Robin) NPCId() uint32 {
	return 2003
}

func (r Robin) Initial() State {
	return r.AskMeAnything
}

func (r Robin) AskMeAnything(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)

	m := message.NewBuilder().
		AddText("Now...ask me any questions you may have on traveling!!").AddNewLine().
		OpenItem(0).BlueText().AddText("How do I move?").CloseItem().AddNewLine().
		OpenItem(1).BlueText().AddText("How do I take down the monsters?").CloseItem().AddNewLine().
		OpenItem(2).BlueText().AddText("How can I pick up an item?").CloseItem().AddNewLine().
		OpenItem(3).BlueText().AddText("What happens when I die?").CloseItem().AddNewLine().
		OpenItem(4).BlueText().AddText("When can I choose a job?").CloseItem().AddNewLine().
		OpenItem(5).BlueText().AddText("Tell me more about this island!").CloseItem().AddNewLine().
		OpenItem(6).BlueText().AddText("What should I do to become a Warrior?").CloseItem().AddNewLine().
		OpenItem(7).BlueText().AddText("What should I do to become a Bowman?").CloseItem().AddNewLine().
		OpenItem(8).BlueText().AddText("What should I do to become a Magician?").CloseItem().AddNewLine().
		OpenItem(9).BlueText().AddText("What should I do to become a Thief?").CloseItem().AddNewLine().
		OpenItem(10).BlueText().AddText("How do I raise the character stats? (S)").CloseItem().AddNewLine().
		OpenItem(11).BlueText().AddText("How do I check the items that I just picked up?").CloseItem().AddNewLine().
		OpenItem(12).BlueText().AddText("How do I put on an item?").CloseItem().AddNewLine().
		OpenItem(13).BlueText().AddText("How do I check out the items that I'm wearing?").CloseItem().AddNewLine().
		OpenItem(14).BlueText().AddText("What are skills? (K)").CloseItem().AddNewLine().
		OpenItem(15).BlueText().AddText("How do I get to Victoria Island?").CloseItem().AddNewLine().
		OpenItem(16).BlueText().AddText("What are mesos?").CloseItem().
		BlackText()
	conversation.SendSimple(m.String())
	return ListSelection(GenericExit, r.ProcessListSelection)
}

func (r Robin) ProcessListSelection(selection int32) State {
	switch selection {
	case 0:
		return r.HowToMove
	case 1:
		return r.HowToTakeDownMonster
	case 2:
		return r.HowToGather
	case 3:
		return r.WhenYouDie
	case 4:
		return r.ChoosingAJob
	case 5:
		return r.TheIsland
	case 6:
		return r.WarriorToDo
	case 7:
		return r.BowmanToDo
	case 8:
		return r.MagicianToDo
	case 9:
		return r.ThiefToDo
	case 10:
		return r.RaiseStats
	case 11:
		return r.CheckItems
	case 12:
		return r.WearItems
	case 13:
		return r.CheckEquipment
	case 14:
		return r.SpecialAbilities
	case 15:
		return r.GetToVictoria
	case 16:
		return r.Mesos
	}
	return nil
}

func (r Robin) HowToMove(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Alright this is how you move. Use ").
		BlueText().AddText("left, right arrow").
		BlackText().AddText(" to move around the flatland and slanted roads, and press ").
		BlueText().AddText("Alt").
		BlackText().AddText(" to jump. A select number of shoes improve your speed and jumping abilities.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.AttackMonsters)
}

func (r Robin) HowToTakeDownMonster(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	conversation.SendNext("Here's how to take down a monster. Every monster possesses an HP of its own and you'll take them down by attacking with either a weapon or through spells. Of course the stronger they are, the harder it is to take them down.")
	return Next(GenericExit, r.JobAdvancement)

}

func (r Robin) HowToGather(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("This is how you gather up an item. Once you take down a monster, an item will be dropped to the ground. When that happens, stand in front of the item and press ").
		BlueText().AddText("Z").
		BlackText().AddText(" or ").
		BlueText().AddText("0 on the NumPad").
		BlackText().AddText(" to acquire the item.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.FullInventory)
}

func (r Robin) WhenYouDie(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	conversation.SendNext("Curious to find out what happens when you die? You'll become a ghost when your HP reaches 0. There will be a tombstone in that place and you won't be able to move, although you still will be able to chat.")
	return Next(GenericExit, r.BeginnerDeath)

}

func (r Robin) ChoosingAJob(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	conversation.SendNext("When do you get to choose your job? Hahaha, take it easy, my friend. Each job has a requirement set for you to meet. Normally a level between 8 and 10 will do, so work hard.")
	return Next(GenericExit, r.HowToAdvance)

}

func (r Robin) TheIsland(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	conversation.SendNext("Want to know about this island? It's called Maple Island and it floats in the air. It's been floating in the sky for a while so the nasty monsters aren't really around. It's a very peaceful island, perfect for beginners!")
	return Next(GenericExit, r.PowerfulPlayer)

}

func (r Robin) WarriorToDo(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("You want to become a #bWarrior#k? Hmmm, then I suggest you head over to Victoria Island. Head over to a warrior-town called ").
		RedText().AddText("Perion").
		BlackText().AddText(" and see ").
		BlueText().ShowNPC(10202).
		BlackText().AddText(". He'll teach you all about becoming a true warrior. Ohh, and one VERY important thing: You'll need to be at least level 10 in order to become a warrior!!")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.AskMeAnything)

}

func (r Robin) BowmanToDo(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("You want to become a ").
		BlueText().AddText("Bowman").
		BlackText().AddText("? You'll need to go to Victoria Island to make the job advancement. Head over to a bowman-town called ").
		RedText().AddText("Henesys").
		BlackText().AddText(" and talk to the beautiful ").
		BlueText().ShowNPC(10200).
		BlackText().AddText(" and learn the in's and out's of being a bowman. Ohh, and one VERY important thing: You'll need to be at least level 10 in order to become a bowman!!")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.AskMeAnything)

}

func (r Robin) MagicianToDo(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("You want to become a ").
		BlueText().AddText("Magician").
		BlackText().AddText("? For you to do that, you'll have to head over to Victoria Island. Head over to a magician-town called ").
		RedText().AddText("Ellinia").
		BlackText().AddText(", and at the very top lies the Magic Library. Inside, you'll meet the head of all wizards, ").
		BlueText().ShowNPC(10201).
		BlackText().AddText(", who'll teach you everything about becoming a wizard.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.MagicianSpecial)

}

func (r Robin) ThiefToDo(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("You want to become a ").
		BlueText().AddText("Thief").
		BlackText().AddText("? In order to become one, you'll have to head over to Victoria Island. Head over to a thief-town called ").
		RedText().AddText("Kerning City").
		BlackText().AddText(", and on the shadier side of town, you'll see a thief's hideaway. There, you'll meet ").
		BlueText().ShowNPC(10203).
		BlackText().AddText(" who'll teach you everything about being a thief. Ohh, and one VERY important thing: You'll need to be at least level 10 in order to become a thief!!")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.AskMeAnything)

}

func (r Robin) RaiseStats(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("You want to know how to raise your character's ability stats? First press ").
		BlueText().AddText("S").
		BlackText().AddText(" to check out the ability window. Every time you level up, you'll be awarded 5 ability points aka AP's. Assign those AP's to the ability of your choice. It's that simple.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.AbilityExplanation)

}

func (r Robin) CheckItems(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("You want to know how to check out the items you've picked up, huh? When you defeat a monster, it'll drop an item on the ground, and you may press ").
		BlueText().AddText("Z").
		BlackText().AddText(" to pick up the item. That item will then be stored in your item inventory, and you can take a look at it by simply pressing ").
		BlueText().AddText("I").
		BlackText().AddText(".")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.AskMeAnything)

}

func (r Robin) WearItems(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("You want to know how to wear the items, right? Press ").
		BlueText().AddText("I").
		BlackText().AddText(" to check out your item inventory. Place your mouse cursor on top of an item and double-click on it to put it on your character. If you find yourself unable to wear the item, chances are your character does not meet the level & stat requirements. You can also put on the item by opening the equipment inventory (").
		BlueText().AddText("E").
		BlackText().AddText(") and dragging the item into it. To take off an item, double-click on the item at the equipment inventory.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.AskMeAnything)

}

func (r Robin) CheckEquipment(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("You want to check on the equipped items, right? Press ").
		BlueText().AddText("E").
		BlackText().AddText(" to open the equipment inventory, where you'll see exactly what you are wearing right at the moment. To take off an item, double-click on the item. The item will then be sent to the item inventory.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.AskMeAnything)

}

func (r Robin) SpecialAbilities(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("The special 'abilities' you get after acquiring a job are called skills. You'll acquire skills that are specifically for that job. You're not at that stage yet, so you don't have any skills yet, but just remember that to check on your skills, press ").
		BlueText().AddText("K").
		BlackText().AddText(" to open the skill book. It'll help you down the road.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.AskMeAnything)

}

func (r Robin) GetToVictoria(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	conversation.SendNext("How do you get to Victoria Island? On the east of this island there's a harbor called Southperry. There, you'll find a ship that flies in the air. In front of the ship stands the captain. Ask him about it.")
	return Next(GenericExit, r.OneLastPiece)

}

func (r Robin) Mesos(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	conversation.SendNext("It's the currency used in MapleStory. You may purchase items through mesos. To earn them, you may either defeat the monsters, sell items at the store, or complete quests...")
	return Next(GenericExit, r.AskMeAnything)

}

func (r Robin) AttackMonsters(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("In order to attack the monsters, you'll need to be equipped with a weapon. When equipped, press ").
		BlueText().AddText("Ctrl").
		BlackText().AddText(" to use the weapon. With the right timing, you'll be able to easily take down the monsters.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.AskMeAnything, r.HowToMove)

}

func (r Robin) JobAdvancement(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Once you make the job advancement, you'll acquire different kinds of skills, and you can assign them to HotKeys for easier access. If it's an attacking skill, you don't need to press Ctrl to attack, just press the button assigned as a HotKey.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.AskMeAnything, r.HowToTakeDownMonster)

}

func (r Robin) FullInventory(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Remember, though, that if your item inventory is full, you won't be able to acquire more. So if you have an item you don't need, sell it so you can make something out of it. The inventory may expand once you make the job advancement.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.AskMeAnything, r.HowToGather)

}

func (r Robin) BeginnerDeath(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("There isn't much to lose when you die if you are just a beginner. Once you have a job, however, it's a different story. You'll lose a portion of your EXP when you die, so make sure you avoid danger and death at all cost.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.AskMeAnything, r.WhenYouDie)

}

func (r Robin) HowToAdvance(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Level isn't the only thing that determines the advancement, though. You also need to boost up the levels of a particular ability based on the occupation. For example, to be a warrior, your STR has to be over 35, and so forth, you know what I'm saying? Make sure you boost up the abilities that has direct implications to your job.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.AskMeAnything, r.ChoosingAJob)

}

func (r Robin) PowerfulPlayer(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("But, if you want to be a powerful player, better not think about staying here for too long. You won't be able to get a job anyway. Underneath this island lies an enormous island called Victoria Island. That place is so much bigger than here, it's not even funny.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.AskMeAnything, r.TheIsland)
}

func (r Robin) MagicianSpecial(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Oh by the way, unlike other jobs, to become a magician you only need to be at level 8. What comes with making the job advancement early also comes with the fact that it takes a lot to become a true powerful magician. Think long and carefully before choosing your path.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.AskMeAnything, r.MagicianToDo)
}

func (r Robin) AbilityExplanation(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Place your mouse cursor on top of all abilities for a brief explanation. For example, STR for warriors, DEX for bowman, INT for magician, and LUK for thief. That itself isn't everything you need to know, so you'll need to think long and hard on how to emphasize your character's strengths through assigning the points.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.AskMeAnything, r.RaiseStats)
}

func (r Robin) OneLastPiece(l logrus.FieldLogger, c Context, _ byte, _ byte, _ int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Oh yeah! One last piece of information before I go. If you are not sure where you are, always press ").
		BlueText().AddText("W").
		BlackText().AddText(". The world map will pop up with the locator showing where you stand. You won't have to worry about getting lost with that.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.AskMeAnything, r.GetToVictoria)
}
