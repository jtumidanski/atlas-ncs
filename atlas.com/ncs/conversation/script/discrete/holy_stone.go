package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// HolyStone is located in Hidden Street - Holy Ground at the Snowfield (211040401)
type HolyStone struct {
}

func (r HolyStone) NPCId() uint32 {
	return npc.HolyStone
}

type Question struct {
	prompt       string
	answers      []string
	correctIndex uint32
}

func (r HolyStone) Questions() []Question {
	return []Question{
		//Questions Related to CHARACTERS
		{prompt: "In MapleStory, what is the EXP needed to level up from Lv1 to Lv2?", answers: []string{"20", "15", "4", "12", "16"}, correctIndex: 1},
		{prompt: "In 1st job adv. which of the following is WRONG requirement?", answers: []string{"Magician - Level 8", "Pirate - 20 DEX or more", "Archer - 25 DEX or more", "Thief - 20 LUK or more", "Swordman - 35 STR or more"}, correctIndex: 3},
		{prompt: "When you hit by monster, which of the following is not fully explained?", answers: []string{"Sealed - skills become disabled", "Undead - turns undead & halved recovery amounts", "Weaken - slow down moving speed", "Cursed - EXP received are decreased", "Stunned - cannot move"}, correctIndex: 2},
		{prompt: "For the 1st job adv. Which job fully states the job adv. requirement?", answers: []string{"Pirate - 25 LUK", "Magician - Level 10", "Thief - 25 LUK", "Warrior - 30 STR", "Bowman - 25 DEX"}, correctIndex: 4},
		//Questions Related to ITEMS
		{prompt: "Which of following monsters got CORRECT item corresponding to the monster?", answers: []string{"Royal cactus - Needle", "Wild Boar - Boar fang", "Lazy Buffy - Buffy hat", "Chipmunk - Nut", "Stirge - Stirge's wing"}, correctIndex: 4},
		{prompt: "Which of following monsters got WRONG item corresponding to the monster?", answers: []string{"Greatest Oldies - Greatest oldies", "Nependeath - Nependeath's leaf", "Ghost stump - Seedling", "Sparker - Seal tooth", "Miner Zombie - Zombie's lost tooth"}, correctIndex: 1},
		{prompt: "In GM Event, how many FRUIT CAKE you can get as reward?", answers: []string{"20", "200", "5", "25", "100"}, correctIndex: 2},
		{prompt: "Which of following potions got CORRECT info.?", answers: []string{"Warrior Elixir - Attack +5 for 3 minutes", "Pure Water - Recover 700 MP", "Cake - Recover 150 HP & MP", "Salad - Recover 300 MP", "Pizza - Recover 400 HP"}, correctIndex: 4},
		{prompt: "Which of following potions got WRONG info.?", answers: []string{"Mana Elixir - Recover 300 MP", "Tonic - Cures state of weakness", "Apple - Recover 30 HP", "Sunrise Dew - Recover 3000 MP", "Ramen - Recover 1000 HP"}, correctIndex: 3},
		//Questions Related to MONSTERS
		{prompt: "Green Mushroom, Tree Stump, Bubbling, Axe Stump, Octopus, which is highest level of all?", answers: []string{"Tree Stump", "Bubbling", "Axe Stump", "Octopus", "Green Mushroom"}, correctIndex: 2},
		{prompt: "Which monster will be seen during the ship trip to Orbis/Ellinia?", answers: []string{"Werewolf", "Slime", "Crimson Balrog", "Zakum", "Star Pixie"}, correctIndex: 2},
		{prompt: "Maple Island doesn't have which following monsters?", answers: []string{"Shroom", "Blue Snail", "Slime", "Red Snail", "Pig"}, correctIndex: 4},
		{prompt: "Which monster is not at Victoria Island and Sleepywood?", answers: []string{"Evil Eye", "Sentinel", "Jr. Balrog", "Ghost Stump", "Snail"}, correctIndex: 1},
		{prompt: "El Nath doesn't have which following monsters?", answers: []string{"Dark Yeti", "Dark Ligator", "Yeti & Pepe", "Bain", "Coolie Zombie"}, correctIndex: 1},
		{prompt: "Which of following monsters can fly?", answers: []string{"Malady", "Ligator", "Cold Eye", "Meerkat", "Alishar"}, correctIndex: 0},
		{prompt: "Which of these monsters will you NOT be facing in Ossyria?", answers: []string{"Lunar Pixie", "Lioner", "Cellion", "Croco", "Hector"}, correctIndex: 3},
		{prompt: "Which monster has not appeared in Maple Island?", answers: []string{"Snail", "Shroom", "Evil Eye", "Orange Mushroom", "Blue Snail"}, correctIndex: 2},
		//Questions Related to QUESTS
		{prompt: "Which material doesn't need for awaken Hero's Gladius?", answers: []string{"Flaming Feather", "Old Gladius", "Piece of Ice", "Ancient Scroll", "Fairy Wing"}, correctIndex: 4},
		{prompt: "Which of following quests can be repeated?", answers: []string{"Mystery of Niora Hospital", "Rightful Donation Culture", "The Ghost Whereabout", "Arwen and the Glass Shoe", "Maya and the Weird Medicine"}, correctIndex: 3},
		{prompt: "Which of following are not 2nd job adv.?", answers: []string{"Mage", "Cleric", "Assassin", "Gunslinger", "Fighter"}, correctIndex: 0},
		{prompt: "Which of following is the highest level quest?", answers: []string{"Cupid's Courier", "Lost in the Ocean", "Alcaster and the Dark Crystal", "Eliminating the Drumming Bunny", "War of Pang Pang"}, correctIndex: 2},
		//Questions Related to TOWN/NPC
		{prompt: "Which town is not at Victoria Island?", answers: []string{"Florina Beach or Nautilus", "Amherst or Southperry", "Kerning City & Square", "Perion or Ellinia", "Sleepywood"}, correctIndex: 1},
		{prompt: "Which is the first NPC you meet in Maple Island?", answers: []string{"Sera", "Heena", "Lucas", "Roger", "Shanks"}, correctIndex: 1},
		{prompt: "Which NPC cannot be seen in El Nath?", answers: []string{"Vogen", "Sophia", "Pedro", "Master Sergeant Fox", "Rumi"}, correctIndex: 1},
		{prompt: "Which NPC cannot be seen in El Nath snowfield?", answers: []string{"Hidden Rock", "Glibber", "Jeff", "Holy Stone", "Elma the Housekeeper"}, correctIndex: 4},
		{prompt: "Which NPC cannot be seen in Perion?", answers: []string{"Ayan", "Sophia", "Mr. Smith", "Francois", "Manji"}, correctIndex: 3},
		{prompt: "Which NPC cannot be seen in Henesys?", answers: []string{"Teo", "Vicious", "Mia", "Doofus", "Casey"}, correctIndex: 0},
		{prompt: "Which NPC cannot be seen in Ellinia?", answers: []string{"Mr. Park", "Mar the Fairy", "Roel", "Ria", "Shane"}, correctIndex: 2},
		{prompt: "Which NPC cannot be seen in Kerning City?", answers: []string{"Dr. Faymus", "Mong from Kong", "Ervine", "Luke", "Nella"}, correctIndex: 3},
		{prompt: "Which NPC is not related to pets?", answers: []string{"Doofus", "Vicious", "Patricia", "Weaver", "Cloy"}, correctIndex: 1},
		{prompt: "In Kerning City, who is the father of Alex, the runaway kid?", answers: []string{"Chief Stan", "JM From tha Streetz", "Dr. Faymus", "Vicious", "Luke"}, correctIndex: 0},
		{prompt: "Which NPC is not belong to Alpha Platoon's Network of Communication?", answers: []string{"Staff Sergeant Charlie", "Sergeant Bravo", "Corporal Easy", "Master Sergeant Fox", "Peter"}, correctIndex: 4},
		{prompt: "What do you receive in return from giving 30 Dark Marbles to the 2nd job advancement NPC?", answers: []string{"Old Ring", "Memory Powder", "Fairy Dust", "Proof of Hero", "Scroll of Secrets"}, correctIndex: 3},
		{prompt: "Which item you give Maya at Henesys in order to cure her sickness?", answers: []string{"Apple", "Power Elixir", "Weird Medicine", "Chrysanthemum", "Orange Juice"}, correctIndex: 2},
		{prompt: "Which of following NPC is not related to item synthesis/refine?", answers: []string{"Neve", "Serryl", "Shane", "Francois", "JM From tha Streetz"}, correctIndex: 2},
		{prompt: "Which NPC cannot be seen in Maple Island?", answers: []string{"Bari", "Teo", "Pio", "Sid", "Maria"}, correctIndex: 1},
		{prompt: "Who do you see in the monitor in the navigation room with Kyrin?", answers: []string{"Lucas", "Dr. Kim", "Chief Stan", "Scadur", "Professor Foxwit"}, correctIndex: 1},
		{prompt: "You know Athena Pierce in Henesys? What color are her eyes?", answers: []string{"Blue", "Green", "Brown", "Red", "Black"}, correctIndex: 1},
		{prompt: "How many feathers are there on Dances with Barlog's Hat?", answers: []string{"7", "8", "3", "13", "16"}, correctIndex: 3},
		{prompt: "What's the color of the marble Grendel the Really Old from Ellinia carries with him?", answers: []string{"White", "Orange", "Blue", "Purple", "Green"}, correctIndex: 2},
	}
}
func (r HolyStone) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.NecklaceOfWisdom) {
		return script.Exit()(l, span, c)
	}

	if character.HasItem(l, span)(c.CharacterId, item.DarkCrystal) {
		return r.BringMe(l, span, c)
	}

	if !character.CanHold(l)(c.CharacterId, item.NecklaceOfWisdom) {
		return r.InventoryRoom(l, span, c)
	}

	return r.Alright(l, span, c)
}

func (r HolyStone) BringMe(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Bring me a ").
		BlueText().ShowItemName1(item.DarkCrystal).
		BlackText().AddText(" to proceed with the trial.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r HolyStone) InventoryRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Have a free ETC slot available before accepting this trial.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r HolyStone) Alright(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Alright... I'll be testing out your wisdom here. Answer all the questions correctly, and you will pass the test BUT, if you even lie to me once, then you'll have to start over again ok, here we go.")
	return script.SendNext(l, span, c, m.String(), r.TakeCrystal)
}

func (r HolyStone) TakeCrystal(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.DarkCrystal, -1)

	questionSet := r.GenerateQuestionSet()

	return r.AskQuestion(questionSet)(l, span, c)
}

func (r HolyStone) GenerateQuestionSet() []uint32 {
	questions := make([]uint32, 0)
	for i := 0; i < 5; i++ {
		random := uint32(rand.Intn(len(r.Questions())))
		for !r.Contains(questions, random) {
			random = uint32(rand.Intn(len(r.Questions())))
		}
		questions = append(questions, random)
	}

	return questions
}

func (r HolyStone) AskQuestion(set []uint32) script.StateProducer {
	if len(set) == 0 {
		return script.Exit()
	}
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		question := r.Questions()[set[0]]

		m := message.NewBuilder().AddText(r.GetHeading(5-len(set)+1) + question.prompt)
		for i, answer := range question.answers {
			m = m.OpenItem(i).BlueText().AddText(answer).CloseItem().NewLine()
		}
		return script.SendListSelection(l, span, c, m.String(), r.ProcessAnswer(set))
	}
}

func (r HolyStone) GetHeading(question int) string {
	switch question {
	case 1:
		return "Here's the 1st question. "
	case 2:
		return "Here's the 2nd question. "
	case 3:
		return "Here's the 3rd question. "
	default:
		return fmt.Sprintf("Here's the %dth question.", question)
	}
}

func (r HolyStone) ProcessAnswer(set []uint32) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		question := r.Questions()[set[0]]

		return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
			if question.correctIndex != uint32(selection) {
				return r.Failed(l, span, c)
			}

			if len(set) == 1 {
				return r.Complete(l, span, c)
			}
			return r.AskQuestion(set[1:])(l, span, c)
		}
	}
}

func (r HolyStone) Failed(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You have failed the question.")
	return script.SendOk(l, span, c, m.String())
}

func (r HolyStone) Complete(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.NecklaceOfWisdom, 1)
	m := message.NewBuilder().AddText("Alright. All your answers have been proven as the truth. Your wisdom has been proven.").NewLine().
		AddText("Take this necklace and go back.")
	return script.SendOk(l, span, c, m.String())
}

func (r HolyStone) Contains(questions []uint32, random uint32) bool {
	for _, v := range questions {
		if v == random {
			return true
		}
	}
	return false
}
