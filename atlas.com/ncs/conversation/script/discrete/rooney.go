package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
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

func (r Rooney) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.Hello(l, c)
}

func (r Rooney) Hello(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Santa told me to go to here, only he didn't told me when...  I hope I'm here on the right time! Oh! By the way, I'm Rooney, I can take you to ").
		BlueText().AddText("HappyVille").
		BlackText().AddText(". Are you ready to go?")
	return script.SendYesNo(l, c, m.String(), r.Warp, script.Exit())
}

func (r Rooney) Warp(l logrus.FieldLogger, c script.Context) script.State {
	character.SaveLocation(l)(c.CharacterId, "HAPPYVILLE")
	return script.WarpById(_map.Happyville, 0)(l, c)
}
