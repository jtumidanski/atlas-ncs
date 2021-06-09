package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// EnigmaTombstone is located in MesoGears - Enigma Chamber (600020600)
type EnigmaTombstone struct {
}

func (r EnigmaTombstone) NPCId() uint32 {
	return npc.EnigmaTombstone
}

func (r EnigmaTombstone) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("(This enigmatic tombstone keeps emitting strange forces... Better look another way.)")
	return script.SendOk(l, c, m.String())
}
