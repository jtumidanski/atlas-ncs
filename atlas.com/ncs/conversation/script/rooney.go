package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Rooney is located in All Towns
type Rooney struct {
}

func (r Rooney) NPCId() uint32 {
	return npc.Rooney
}

func (r Rooney) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Rooney) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Santa told me to go to here, only he didn't told me when...  I hope I'm here on the right time! Oh! By the way, I'm Rooney, I can take you to ").
		BlueText().AddText("HappyVille").
		BlackText().AddText(". Are you ready to go?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r Rooney) Warp(l logrus.FieldLogger, c Context) State {
	character.SaveLocation(l)(c.CharacterId, "HAPPYVILLE")
	return WarpById(_map.Happyville, 0)(l, c)
}
