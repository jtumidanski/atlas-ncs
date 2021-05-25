package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// HenesysForest is located in Victoria Road - The Rain-Forest East of Henesys (100020000)
type HenesysForest struct {
}

func (r HenesysForest) NPCId() uint32 {
	return npc.HenesysForest
}

func (r HenesysForest) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestNotStarted(l)(c.CharacterId, 20706) {
		return r.NothingSuspicious(l, c)
	} else if character.QuestStarted(l)(c.CharacterId, 20706) {
		return r.Complete(l, c)
	} else if character.QuestCompleted(l)(c.CharacterId, 20706) {
		return r.AlreadyComplete(l, c)
	}
	return Exit()(l, c)
}

func (r HenesysForest) NothingSuspicious(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It looks like there's nothing suspicious in the area.")
	return SendNext(l, c, m.String(), Exit())
}

func (r HenesysForest) Complete(l logrus.FieldLogger, c Context) State {
	character.ForceCompleteQuest(l)(c.CharacterId, 20706)
	m := message.NewBuilder().
		AddText("You have spotted the shadow! Better report to ").ShowNPC(npc.Roca).AddText(".")
	return SendNext(l, c, m.String(), Exit())
}

func (r HenesysForest) AlreadyComplete(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The shadow has already been spotted. Better report to ").ShowNPC(npc.Roca).AddText(".")
	return SendNext(l, c, m.String(), Exit())
}
