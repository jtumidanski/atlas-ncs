package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Rupi is located in Hidden Street - Happyville (209000000)
type Rupi struct {
}

func (r Rupi) NPCId() uint32 {
	return npc.Rupi
}

func (r Rupi) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you want to get out of Happyville?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r Rupi) Warp(l logrus.FieldLogger, c Context) State {
	mapId := character.SavedLocation(l)(c.CharacterId, "HAPPYVILLE")
	return WarpById(mapId, 0)(l, c)
}
