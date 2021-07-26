package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// Shulynch is located in The Nautilus - Training Room (120000104)
type Shulynch struct {
}

func (r Shulynch) NPCId() uint32 {
	return npc.Shulynch
}

func (r Shulynch) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 6410) {
		return r.AnyBusiness(l, c)
	}
	return r.LetsGoSave(l, c)
}

func (r Shulynch) AnyBusiness(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey, do you have any business with me?")
	return script.SendOk(l, c, m.String())
}

func (r Shulynch) LetsGoSave(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Let's go save ").
		RedText().ShowNPC(npc.Delli).
		BlackText().AddText("?")
	return script.SendYesNo(l, c, m.String(), r.Warp, script.Exit())
}

func (r Shulynch) Warp(l logrus.FieldLogger, c script.Context) script.State {
	return script.WarpById(_map.LookingForDelli1, 0)(l, c)
}
