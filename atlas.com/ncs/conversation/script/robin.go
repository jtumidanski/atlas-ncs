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

func (r Robin) Start(l logrus.FieldLogger, c Context) {
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
}

func (r Robin) Continue(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) {
	if mode == 0 && theType == 0 {
		r.Start(l, c)
		return
	}
	if mode == 1 && theType == 0 {
		npc.Processor(l).Dispose(c.CharacterId)
		return
	}

	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)

	if selection == 0 {
		m := message.NewBuilder().
			AddText("Alright this is how you move. Use ").
			BlueText().AddText("left, right arrow").
			BlackText().AddText(" to move around the flatland and slanted roads, and press ").
			BlueText().AddText("Alt").
			BlackText().AddText(" to jump. A select number of shoes improve your speed and jumping abilities.")
		conversation.SendNext(m.String())
	} else if selection == 1 {
		conversation.SendNext("Here's how to take down a monster. Every monster possesses an HP of its own and you'll take them down by attacking with either a weapon or through spells. Of course the stronger they are, the harder it is to take them down.")
	} else if selection == 2 {
		m := message.NewBuilder().
			AddText("This is how you gather up an item. Once you take down a monster, an item will be dropped to the ground. When that happens, stand in front of the item and press ").
			BlueText().AddText("Z").
			BlackText().AddText(" or ").
			BlueText().AddText("0 on the NumPad").
			BlackText().AddText(" to acquire the item.")
		conversation.SendNext(m.String())
	} else if selection == 3 {
		conversation.SendNext("Curious to find out what happens when you die? You'll become a ghost when your HP reaches 0. There will be a tombstone in that place and you won't be able to move, although you still will be able to chat.")
	} else if selection == 4 {
		conversation.SendNext("When do you get to choose your job? Hahaha, take it easy, my friend. Each job has a requirement set for you to meet. Normally a level between 8 and 10 will do, so work hard.")
	} else if selection == 5 {
		conversation.SendNext("Want to know about this island? It's called Maple Island and it floats in the air. It's been floating in the sky for a while so the nasty monsters aren't really around. It's a very peaceful island, perfect for beginners!")
	} else if selection == 6 {
		m := message.NewBuilder().
			AddText("You want to become a #bWarrior#k? Hmmm, then I suggest you head over to Victoria Island. Head over to a warrior-town called ").
			RedText().AddText("Perion").
			BlackText().AddText(" and see ").
			BlueText().ShowNPC(10202).
			BlackText().AddText(". He'll teach you all about becoming a true warrior. Ohh, and one VERY important thing: You'll need to be at least level 10 in order to become a warrior!!")
		conversation.SendNext(m.String())
	} else if selection == 7 {
		m := message.NewBuilder().
			AddText("You want to become a ").
			BlueText().AddText("Bowman").
			BlackText().AddText("? You'll need to go to Victoria Island to make the job advancement. Head over to a bowman-town called ").
			RedText().AddText("Henesys").
			BlackText().AddText(" and talk to the beautiful ").
			BlueText().ShowNPC(10200).
			BlackText().AddText(" and learn the in's and out's of being a bowman. Ohh, and one VERY important thing: You'll need to be at least level 10 in order to become a bowman!!")
		conversation.SendNext(m.String())
	} else if selection == 8 {
		m := message.NewBuilder().
			AddText("You want to become a ").
			BlueText().AddText("Magician").
			BlackText().AddText("? For you to do that, you'll have to head over to Victoria Island. Head over to a magician-town called ").
			RedText().AddText("Ellinia").
			BlackText().AddText(", and at the very top lies the Magic Library. Inside, you'll meet the head of all wizards, ").
			BlueText().ShowNPC(10201).
			BlackText().AddText(", who'll teach you everything about becoming a wizard.")
		conversation.SendNext(m.String())
	} else if selection == 9 {
		m := message.NewBuilder().
			AddText("You want to become a ").
			BlueText().AddText("Thief").
			BlackText().AddText("? In order to become one, you'll have to head over to Victoria Island. Head over to a thief-town called ").
			RedText().AddText("Kerning City").
			BlackText().AddText(", and on the shadier side of town, you'll see a thief's hideaway. There, you'll meet ").
			BlueText().ShowNPC(10203).
			BlackText().AddText(" who'll teach you everything about being a thief. Ohh, and one VERY important thing: You'll need to be at least level 10 in order to become a thief!!")
		conversation.SendNext(m.String())
	} else if selection == 10 {
		m := message.NewBuilder().
			AddText("You want to know how to raise your character's ability stats? First press ").
			BlueText().AddText("S").
			BlackText().AddText(" to check out the ability window. Every time you level up, you'll be awarded 5 ability points aka AP's. Assign those AP's to the ability of your choice. It's that simple.")
		conversation.SendNext(m.String())
	} else if selection == 11 {
		m := message.NewBuilder().
			AddText("You want to know how to check out the items you've picked up, huh? When you defeat a monster, it'll drop an item on the ground, and you may press ").
			BlueText().AddText("Z").
			BlackText().AddText(" to pick up the item. That item will then be stored in your item inventory, and you can take a look at it by simply pressing ").
			BlueText().AddText("I").
			BlackText().AddText(".")
		conversation.SendNext(m.String())
	} else if selection == 12 {
		m := message.NewBuilder().
			AddText("You want to know how to wear the items, right? Press ").
			BlueText().AddText("I").
			BlackText().AddText(" to check out your item inventory. Place your mouse cursor on top of an item and double-click on it to put it on your character. If you find yourself unable to wear the item, chances are your character does not meet the level & stat requirements. You can also put on the item by opening the equipment inventory (").
			BlueText().AddText("E").
			BlackText().AddText(") and dragging the item into it. To take off an item, double-click on the item at the equipment inventory.")
		conversation.SendNext(m.String())
	} else if selection == 13 {
		m := message.NewBuilder().
			AddText("You want to check on the equipped items, right? Press ").
			BlueText().AddText("E").
			BlackText().AddText(" to open the equipment inventory, where you'll see exactly what you are wearing right at the moment. To take off an item, double-click on the item. The item will then be sent to the item inventory.")
		conversation.SendNext(m.String())
	} else if selection == 14 {
		m := message.NewBuilder().
			AddText("The special 'abilities' you get after acquiring a job are called skills. You'll acquire skills that are specifically for that job. You're not at that stage yet, so you don't have any skills yet, but just remember that to check on your skills, press ").
			BlueText().AddText("K").
			BlackText().AddText(" to open the skill book. It'll help you down the road.")
		conversation.SendNext(m.String())
	} else if selection == 15 {
		conversation.SendNext("How do you get to Victoria Island? On the east of this island there's a harbor called Southperry. There, you'll find a ship that flies in the air. In front of the ship stands the captain. Ask him about it.")
	} else if selection == 16 {
		conversation.SendNext("It's the currency used in MapleStory. You may purchase items through mesos. To earn them, you may either defeat the monsters, sell items at the store, or complete quests...")
	}

}
