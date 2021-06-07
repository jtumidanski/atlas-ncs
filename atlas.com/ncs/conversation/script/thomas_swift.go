package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ThomasSwift is located in Amoria - Amoria (680000000) and Victoria Road - Henesys (100000000)
type ThomasSwift struct {
}

func (r ThomasSwift) NPCId() uint32 {
	return npc.ThomasSwift
}

func (r ThomasSwift) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.Henesys {
		m := message.NewBuilder().
			AddText("I can take you to the Amoria Village. Are you ready to go?")
		return SendYesNo(l, c, m.String(), r.HaveAGreatTime(_map.Amoria, 0), r.HangAround)
	} else {
		m := message.NewBuilder().
			AddText("I can take you back to Henesys. Are you ready to go?")
		return SendYesNo(l, c, m.String(), r.HaveAGreatTime(_map.Henesys, 5), r.HangAround)
	}
}

func (r ThomasSwift) HangAround(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Ok, feel free to hang around until you're ready to go!")
	return SendOk(l, c, m.String())
}

func (r ThomasSwift) HaveAGreatTime(mapId uint32, portalId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().AddText("I hope you had a great time! See you around!")
		return SendNext(l, c, m.String(), WarpById(mapId, portalId))
	}
}
