package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KonpeiExit is located in Zipangu - Near the Hideout (801040000)
type KonpeiExit struct {
}

func (r KonpeiExit) NPCId() uint32 {
	return npc.KonpeiExit
}

func (r KonpeiExit) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Here you are, right in front of the hideout! What? You want to").NewLine().
		AddText("return to ").
		ShowMap(_map.ShowaTown).
		AddText("?")
	return script.SendYesNo(l, c, m.String(), script.Warp(_map.ShowaTown), r.TalkToMe)
}

func (r KonpeiExit) TalkToMe(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If you want to return to ").
		ShowMap(_map.ShowaTown).
		AddText(", then talk to me.")
	return script.SendOk(l, c, m.String())
}
