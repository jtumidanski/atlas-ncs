package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// BossKitty is located in Zipangu - Showa Town (801000000)
type BossKitty struct {
}

func (r BossKitty) NPCId() uint32 {
	return npc.BossKitty
}

func (r BossKitty) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 8012) && !character.HasItem(l)(c.CharacterId, item.OrangeMarble) {
		m := message.NewBuilder().AddText("Did you get them all? Are you going to try to answer all of my questions?")
		return script.SendYesNo(l, c, m.String(), r.ValidateChicken, script.Exit())
	} else {
		m := message.NewBuilder().AddText("Meeeoooowww!")
		return script.SendOk(l, c, m.String())
	}
}

func (r BossKitty) ValidateChicken(l logrus.FieldLogger, c script.Context) script.State {
	if !character.HasItems(l)(c.CharacterId, item.FriedChicken, 300) {
		return r.MissingChicken(l, c)
	}
	character.GainItem(l)(c.CharacterId, item.FriedChicken, -300)
	m := message.NewBuilder().
		AddText("Good job! Now hold on a sec... Hey look! I got some food here! Help yourselves. Okay, now it's time for me to ask you some questions. I'm sure you're aware of this, but remember, if you're wrong, it's over. It's all or nothing!")
	return script.SendNext(l, c, m.String(), r.FirstQuestion)
}

func (r BossKitty) MissingChicken(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("What? No! 300! THREE. HUNDRED. No less. Hand over more if you want, but I need at least 300. Not all of us can be as big and as fed as you...")
	return script.SendOk(l, c, m.String())
}

func (r BossKitty) FirstQuestion(l logrus.FieldLogger, c script.Context) script.State {
	questions := r.QuestionSet()
	return r.AskQuestion(1, questions)(l, c)
}

func (r BossKitty) QuestionSet() []Question {
	return []Question{
		{prompt: "Which of these items does the Flaming Raccoon NOT drop?", answers: []string{"Raccoon Firewood", "Solid Horn", "Red Brick"}, correctIndex: 1},
		{prompt: "Which NPC is responsible for transporting travellers from Kerning City to Zipangu, and back?", answers: []string{"Peli", "Spinel", "Poli"}, correctIndex: 1},
		{prompt: "Which of the items sold at the Mushroom Shrine increases your attack power?", answers: []string{"Takoyaki", "Yakisoba", "Tempura"}, correctIndex: 0},
		{prompt: "Which of these items do the Extras NOT drop?", answers: []string{"Extra A's Badge", "Extra B's Corset", "Extra C's Necklace"}, correctIndex: 1},
		{prompt: "Which of these items DO NOT exist??", answers: []string{"Frozen Tuna", "Fan", "Fly Swatter"}, correctIndex: 2},
		{prompt: "What's the name of the vegetable store owner in Showa Town?", answers: []string{"Sami", "Kami", "Umi"}, correctIndex: 2},
		{prompt: "Which of these items DO exist?", answers: []string{"Cloud Fox's Tooth", "Ghost's Bouquet", "Dark Cloud Fox's Tail"}, correctIndex: 2},
		{prompt: "What is the name of the strongest boss in the Mushroom Shrine?", answers: []string{"Black Crow", "Blue Mushmom", "Himegami"}, correctIndex: 0},
		{prompt: "Which one of these items has a mis-matched class or level description?", answers: []string{"Bamboo Spear - Warrior-only Weapon", "Pico-Pico Hammer - One-handed Sword", "Mystic Cane - Level 51 equip"}, correctIndex: 0},
		{prompt: "Which of these noodles are NOT being sold by Robo at the Mushroom Shrine?", answers: []string{"Kinoko Ramen (Pig Skull)", "Kinoko Ramen (Salt)", "Mushroom Miso Ramen"}, correctIndex: 2},
		{prompt: "Which of these NPCs do NOT stand in front of Showa Movie Theater?", answers: []string{"Skye", "Furano", "Shinta"}, correctIndex: 2},
	}
}

func (r BossKitty) AskQuestion(num int, questions []Question) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		selected := rand.Intn(len(questions))
		question := questions[selected]

		m := message.NewBuilder().
			AddText(fmt.Sprintf("Question no. %d:", num)).
			AddText(question.prompt)
		for i, answer := range question.answers {
			m = m.OpenItem(i).BlueText().AddText(answer).CloseItem().NewLine()
		}
		return script.SendListSelection(l, c, m.String(), r.Validate(num, question, append(questions[:selected], questions[selected+1:]...)))
	}
}

func (r BossKitty) Validate(num int, question Question, remaining []Question) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		return func(l logrus.FieldLogger, c script.Context) script.State {
			if question.correctIndex != uint32(selection) {
				return r.Incorrect(l, c)
			}
			if len(remaining) == 0 {
				return r.AllCorrect(l, c)
			} else {
				return r.AskQuestion(num+1, remaining)(l, c)
			}
		}
	}
}

func (r BossKitty) AllCorrect(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Dang, you answered all the questions right. I may not like humans in general, but I HATE breaking a promise, so, as promised, here's the Orange Marble.")
	return script.SendNext(l, c, m.String(), r.Reward)
}

func (r BossKitty) Incorrect(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm...all humans make mistakes anyway! If you want to take another crack at it, then bring me 300 Fried Chicken.")
	return script.SendOk(l, c, m.String())
}

func (r BossKitty) Reward(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.OrangeMarble, 1)
	return r.YouCanLeave(l, c)
}

func (r BossKitty) YouCanLeave(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Our business is concluded, thank you very much! You can leave now!")
	return script.SendOk(l, c, m.String())
}
