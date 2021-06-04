package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ChunJi is located in Victoria Road - Kerning City Construction Site (103010000)
type ChunJi struct {
}

func (r ChunJi) NPCId() uint32 {
	return npc.ChunJi
}

func (r ChunJi) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Just a kid... Don't talk to me...")
	return SendOk(l, c, m.String())
}
