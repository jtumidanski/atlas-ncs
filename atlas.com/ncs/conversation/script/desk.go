package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Desk is located in Magatia - Home of the Missing Alchemist (261000001)
type Desk struct {
}

func (r Desk) NPCId() uint32 {
	return npc.Desk
}

func (r Desk) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3311) {
		progress := character.QuestProgressInt(l)(c.CharacterId, 3311, 0)
		if progress == 4 {
			progress = 7
		} else {
			progress = 5
		}
		character.SetQuestProgress(l)(c.CharacterId, 3311, 0, uint32(progress))
		m := message.NewBuilder().AddText("The diary of Dr. De Lang. A lot of formulas and pompous scientific texts can be found all way through the pages, but it is worth noting that in the last entry (3 weeks ago), it is written that he concluded the researches on an improvement on the blueprints for the Neo Huroids, thus making the last preparations to show it to the 'society'... No words after this...")
		return SendOk(l, c, m.String(), AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
	}
	if character.QuestStarted(l)(c.CharacterId, 3322) && !character.HasItem(l)(c.CharacterId, item.SilverPendant) {
		if character.CanHold(l)(c.CharacterId, item.SilverPendant) {
			character.GainItem(l)(c.CharacterId, item.SilverPendant, 1)
		} else {
			m := message.NewBuilder().AddText("Your inventory is full, make sure a ETC slot is available for the item.")
			return SendOk(l, c, m.String())
		}
	}
	return Exit()(l, c)
}
