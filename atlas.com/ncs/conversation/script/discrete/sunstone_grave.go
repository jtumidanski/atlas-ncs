package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SunstoneGrave is located in MesoGears - Fire Chamber (600020400)
type SunstoneGrave struct {
}

func (r SunstoneGrave) NPCId() uint32 {
	return npc.SunstoneGrave
}

func (r SunstoneGrave) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("(This tombstone keeps emitting ever odder light waves the more I stare to it...)")
	return script.SendOk(l, c, m.String())
}
