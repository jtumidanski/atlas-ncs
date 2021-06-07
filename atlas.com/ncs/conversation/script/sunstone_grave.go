package script

import (
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

func (r SunstoneGrave) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("(This tombstone keeps emitting ever odder light waves the more I stare to it...)")
	return SendOk(l, c, m.String())
}
