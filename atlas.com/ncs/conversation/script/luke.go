package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Luke is located in Warning Street - Henesys Dungeon Entrance (106010100)
type Luke struct {
}

func (r Luke) NPCId() uint32 {
	return npc.Luke
}

func (r Luke) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 28177) && !character.HasItem(l)(c.CharacterId, item.TrainingInstructorsBadge) {
		if character.CanHold(l)(c.CharacterId, item.TrainingInstructorsBadge) {
			return r.GiveItem(l, c)
		}
		return r.MakeRoom(l, c)
	}
	return r.Snooze(l, c)
}

func (r Luke) Snooze(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Zzzzzz...")
	return SendOk(l, c, m.String())
}

func (r Luke) MakeRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hey, make a slot available before talking to me.")
	return SendOk(l, c, m.String())
}

func (r Luke) GiveItem(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.TrainingInstructorsBadge, 1)
	return r.Success(l, c)
}

func (r Luke) Success(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Huh, are you looking for me? Chief Stan sent you here, right? But hey, I am not the suspect you seek. If I have some proof? Here, take this and return it to ").
		BlueText().ShowNPC(npc.ChiefStan).
		BlackText().AddText(".")
	return SendOk(l, c, m.String())
}