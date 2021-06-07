package script

import (
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

func (r EnigmaTombstone) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("(This enigmatic tombstone keeps emitting strange forces... Better look another way.)")
	return SendOk(l, c, m.String())
}
