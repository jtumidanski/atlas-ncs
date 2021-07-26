package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// HenesysForest is located in Victoria Road - The Rain-Forest East of Henesys (100020000)
type HenesysForest struct {
}

func (r HenesysForest) NPCId() uint32 {
	return npc.HenesysForest
}

func (r HenesysForest) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsNotStarted(l)(c.CharacterId, 20706) {
		return r.NothingSuspicious(l, c)
	} else if quest.IsStarted(l)(c.CharacterId, 20706) {
		return r.Complete(l, c)
	} else if quest.IsCompleted(l)(c.CharacterId, 20706) {
		return r.AlreadyComplete(l, c)
	}
	return script.Exit()(l, c)
}

func (r HenesysForest) NothingSuspicious(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It looks like there's nothing suspicious in the area.")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r HenesysForest) Complete(l logrus.FieldLogger, c script.Context) script.State {
	quest.ForceComplete(l)(c.CharacterId, 20706)
	m := message.NewBuilder().
		AddText("You have spotted the shadow! Better report to ").ShowNPC(npc.Roca).AddText(".")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r HenesysForest) AlreadyComplete(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The shadow has already been spotted. Better report to ").ShowNPC(npc.Roca).AddText(".")
	return script.SendNext(l, c, m.String(), script.Exit())
}
