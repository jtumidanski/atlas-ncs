package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// IcebyrdSlimm is located in New Leaf City Town Street - New Leaf City - Town Center (600000000)
type IcebyrdSlimm struct {
}

func (r IcebyrdSlimm) NPCId() uint32 {
	return npc.IcebyrdSlimm
}

func (r IcebyrdSlimm) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestCompleted(l)(c.CharacterId, 4911) {
		return r.GoodJob(l, c)
	}
	if character.QuestCompleted(l)(c.CharacterId, 4900) || character.QuestStarted(l)(c.CharacterId, 4900) {
		return r.PayAttention(l, c)
	}
	return r.WhatUp(l, c)
}

func (r IcebyrdSlimm) WhatUp(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("What up! Name's Icebyrd Slimm, mayor of New Leaf City! Happy to see you accepted my invite. So, what can I do for you?").NewLine().
		OpenItem(0).BlueText().AddText("What is this place?").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Who is Professor Foxwit?").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("What's a Foxwit Door?").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Where are the MesoGears?").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("What is the Krakian Jungle?").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("What's a Gear Portal?").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("What do the street signs mean?").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("What's the deal with Jack Masque?").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("Lita Lawless looks like a tough cookie, what's her story?").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("When will new boroughs open up in the city?").CloseItem().NewLine().
		OpenItem(10).BlueText().AddText("I want to take the quiz!").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r IcebyrdSlimm) PayAttention(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hey, pay attention, I'm trying to quiz you on another question, fam!")
	return SendOk(l, c, m.String())
}

func (r IcebyrdSlimm) GoodJob(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Good job! You've solved all of my questions about NLC. Enjoy your trip!")
	return SendOk(l, c, m.String())
}

func (r IcebyrdSlimm) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.WhatIsThisPlace
	case 1:
		return r.Foxwit
	case 2:
		return r.FoxwitDoor
	case 3:
		return r.WhereAreMesoGears
	case 4:
		return r.KrakianJungle
	case 5:
		return r.GearPortal
	case 6:
		return r.AlwaysBuilding
	case 7:
		return r.JackMasque
	case 8:
		return r.Rekindled
	case 9:
		return r.HardAtWork
	case 10:
		return r.Quiz
	}
	return nil
}

func (r IcebyrdSlimm) WhatIsThisPlace(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I've always dreamed of building a city. Not just any city, but one where everyone was welcome. I used to live in Kerning City, so I decided to see if I could create a city. As I went along in finding the means to do so, I encountered many people, some of whom I've come to regard as friends. Like Professor Foxwit-he's our resident genius; saved him from a group of man-eating plants. Jack Masque is an old hunting buddy from Amoria-almost too smooth of a talker for his own good. Lita and I are old friends from Kerning City-she's saved me a few times with that weapon of hers; so I figured she was a perfect choice for Town Sheriff. It took a bit of persuasion, but she came to believe her destiny lies here. About our resident explorer, Barricade came searching for something; he agreed to bring whatever he found to the museum. I'd heard stories about him and his brother when I was still in Kerning City. And Elpam...well, let's just say he's not from around here. At all. We've spoken before, and he seems to mean well, so I've allowed him to stay. I just realized that I've rambled quite a bit! What else would you like to know?")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) Foxwit(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("A pretty spry guy for being 97. He's a time-traveller I ran into outside the city one day. Old guy had a bit of trouble with some jungle creatures-like they tried to eat him. In return for me saving him, he agreed to build a time museum. I get the feeling that he's come here for another reason, as he's mentioned more than a few times that New Leaf City has an interesting role to play in the future. Maybe you can find out a bit more...")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) FoxwitDoor(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Heh, I asked the same thing when I saw the Professor building them. They're warp points. Pressing Up will warp you to another location. I recommend getting the hang of them, they're our transport system.")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) WhereAreMesoGears(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The MesoGears are beneath Bigger Ben. It's a monster-infested section of Bigger Ben that Barricade discovered. It seems to reside in a separate section of the tower-quite strange if you ask me. I hear he needs a bit of help exploring it, you should see him. Be careful though, the Wolf Spiders in there are no joke.")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) KrakianJungle(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Ah...well. The Krakian Jungle is located on the outskirts of New Leaf City. Many new and powerful creatures roam those areas, so you'd better be prepared to fight if you head out there. It's at the right end of town. Rumors abound that the Jungle leads to a lost city, but we haven't found anything yet.")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) GearPortal(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Well, when John found himself in the MesoGears portion of Bigger Ben, he stood on one and went to another location. However, he could only head back and forth-they don't cycle through like the Foxwit Door. Ancient tech for you.")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) AlwaysBuilding(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Well, you'll see them just about everywhere. They're areas under construction. The Red lights mean it's not finished, but the Green lights mean it's open. Check back often, we're always building!")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) JackMasque(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Ah, Jack. You know those guys that are too cool for school? The ones who always seem to get away with everything? AND get the girl? Well, that's Jack, but without the girl. He thinks he blew his chance, and began wearing that mask to hide his true identity. My lips are sealed about who he is, but he's from Amoria. He might tell you a bit more if you ask him.")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) Rekindled(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I've known Lita for a while, though we've just recently rekindled our friendship. I didn't see her for quite a bit, but I understand why. She trained for a very, very long time as a Thief. Matter of fact, that's how we first met! I was besieged by a group of wayward Mushrooms, and she jumped in to help. When it was time to pick a sheriff, it was a no-brainer. She's made a promise to help others in their training and protect the city, so if you're interested in a bit of civic duty, speak with her.")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) HardAtWork(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Soon, my friend. Even though you can't see them, the city developers are hard at work. When they're ready, we'll open them. I know you're looking forward to it and so am I!")
	return SendNext(l, c, m.String(), r.WhatUp)
}

func (r IcebyrdSlimm) Quiz(l logrus.FieldLogger, c Context) State {
	if character.IsLevel(l)(c.CharacterId, 10) {
		character.StartQuest(l)(c.CharacterId, 4900)
		m := message.NewBuilder().AddText("No problem. I'll give you something nice if you answer them correctly!")
		return SendNext(l, c, m.String(), Exit())
	} else {
		m := message.NewBuilder().AddText("Eager, are we? How about you explore a bit more before I let you take the quiz?")
		return SendNext(l, c, m.String(), Exit())
	}
}
