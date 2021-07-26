package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// Mo is located in Phantom Forest - Dead Man's Gorge (610010004)
type Mo struct {
}

func (r Mo) NPCId() uint32 {
	return npc.Mo
}

func (r Mo) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !quest.IsCompleted(l)(c.CharacterId, 8224) {
		m := message.NewBuilder().AddText("Hm, at who do you think you are looking at?")
		return script.SendOk(l, c, m.String())
	}
	npc.OpenShop(l)(c.CharacterId, 9201099)
	return script.Exit()(l, c)
}
