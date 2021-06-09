package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// BulletinBoard is located in Sharenian - Excavation Site (990000000) and Victoria Road - Excavation Site <Camp> (101030104)
type BulletinBoard struct {
}

func (r BulletinBoard) NPCId() uint32 {
	return npc.BulletinBoard
}

func (r BulletinBoard) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("<Notice> ").NewLine().
		AddText("Are you part of a Guild that possesses an ample amount of courage and trust? Then take on the Guild Quest and challenge yourselves!").NewLine().NewLine().
		BlueText().AddText("To Participate :").NewLine().
		BlackText().AddText("1. The Guild must consist of at least 6 people!").NewLine().
		AddText("2. The leader of the Guild Quest must be a Master or a Jr. Master of the Guild!").NewLine().
		AddText("3. The Guild Quest may end early if the number of guild members participating falls below 6, or if the leader decides to end it early!")
	return script.SendOk(l, c, m.String())
}
