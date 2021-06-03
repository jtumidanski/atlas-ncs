package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// LePetitPrince is located in The Burning Sands - Dry Desert (260010500)
type LePetitPrince struct {
}

func (r LePetitPrince) NPCId() uint32 {
	return npc.LePetitPrince
}

func (r LePetitPrince) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Home is so boring... my parents ignore me so much it's unbearable. And ever since we moved from ").
		RedText().AddText("Ariant").
		BlackText().AddText(", they've been trying to get a new palace built so they don't have to live outdoors. But I love the outdoors...")
	return SendOk(l, c, m.String())
}
