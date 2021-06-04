package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MrPickall is located in Victoria Road - Kerning City (103000000)
type MrPickall struct {
}

func (r MrPickall) NPCId() uint32 {
	return npc.MrPickall
}

func (r MrPickall) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Should you have a locked box you want to open, bring it to me.")
	return SendOk(l, c, m.String())
}
