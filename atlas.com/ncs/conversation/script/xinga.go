package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Xinga is located in Victoria Island - Kerning Airport (540010100)
type Xinga struct {
}

func (r Xinga) NPCId() uint32 {
	return npc.Xinga
}

func (r Xinga) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The plane will be taking off soon, will you leave now? You will have to buy the plane ticket again to come in here.")
	return SendYesNo(l, c, m.String(), r.NotRefundable, Exit())
}

func (r Xinga) NotRefundable(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The ticket is not refundable, hope to see you again!")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r Xinga) Warp(l logrus.FieldLogger, c Context) State {
	return Warp(_map.KerningCity)(l, c)
}
