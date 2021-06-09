package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// HenesysForest2 is located in Victoria Road - The Rain-Forest East of Henesys (100020000)
type HenesysForest2 struct {
}

func (r HenesysForest2) NPCId() uint32 {
	return npc.HenesysForest2
}

func (r HenesysForest2) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.Hello(l, c)
}

func (r HenesysForest2) Hello(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It looks like there's nothing suspicious in the area.")
	return script.SendNext(l, c, m.String(), script.Exit())
}
