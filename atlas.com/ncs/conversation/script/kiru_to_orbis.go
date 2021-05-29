package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KiruToOrbis is located in Empress' Road - To Ereve (200090020) Empress' Road - To Orbis (200090045)
type KiruToOrbis struct {
}

func (r KiruToOrbis) NPCId() uint32 {
	return npc.KiruToOrbis
}

func (r KiruToOrbis) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Ah, such lovely winds. This should be a perfect voyage as long as no stupid customer falls off for attempting some weird skill. Of course, I'm talking about you. Please refrain from using your skills.")
	return SendOk(l, c, m.String())
}
