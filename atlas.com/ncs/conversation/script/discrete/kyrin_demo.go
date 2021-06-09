package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KyrinDemo is located in Maple Road : Split Road of Destiny (1020000)
type KyrinDemo struct {
}

func (r KyrinDemo) NPCId() uint32 {
	return npc.KyrinDemo
}

func (r KyrinDemo) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.PirateIntroduction(l, c)
}

func (r KyrinDemo) PirateIntroduction(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Pirates are blessed with outstanding dexterity and power, utilizing their guns for long-range attacks while using their power on melee combat situations. Gunslingers use elemental-based bullets for added damage, while Infighters transform to a different being for maximum effect.")
	return script.SendNext(l, c, m.String(), r.Demo)
}

func (r KyrinDemo) Demo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Pirate?")
	return script.SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r KyrinDemo) DoDemo(l logrus.FieldLogger, c script.Context) script.State {
	npc.LockUI(l)(c.CharacterId)
	return script.WarpById(_map.PirateDemo, 0)(l, c)
}

func (r KyrinDemo) SeeMeAgain(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Pirate, come see me again.")
	return script.SendNext(l, c, m.String(), script.Exit())
}
