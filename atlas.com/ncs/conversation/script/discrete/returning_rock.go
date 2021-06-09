package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type ReturningRock struct {
}

func (r ReturningRock) NPCId() uint32 {
	return npc.ReturningRock
}

func (r ReturningRock) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Would you like to exit the Guild Quest?")
	return script.SendYesNo(l, c, m.String(), r.Warp, script.Exit())
}

func (r ReturningRock) Warp(l logrus.FieldLogger, c script.Context) script.State {
	return script.Warp(_map.ExcavationSiteCamp)(l, c)
}
