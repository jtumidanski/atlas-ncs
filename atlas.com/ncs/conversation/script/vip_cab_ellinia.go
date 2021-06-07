package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// VIPCabEllinia is located in Victoria Road - Ellinia (101000000)
type VIPCabEllinia struct {
}

func (r VIPCabEllinia) NPCId() uint32 {
	return npc.VIPCabEllinia
}

func (r VIPCabEllinia) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r VIPCabEllinia) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hi there! This cab is for VIP customers only. Instead of just taking you to different towns like the regular cabs, we offer a much better service worthy of VIP class. It's a bit pricey, but... for only 10,000 mesos, we'll take you safely to the").
		BlueText().AddText("Ant Tunnel").
		BlackText().AddText(".")
	return SendNextExit(l, c, m.String(), r.Cost, r.MoreToOffer)
}

func (r VIPCabEllinia) MoreToOffer(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("This town also has a lot to offer. Find us if and when you feel the need to go to the Ant Tunnel Park.")
	return SendNext(l, c, m.String(), Exit())
}

func (r VIPCabEllinia) Cost(l logrus.FieldLogger, c Context) State {
	if character.IsBeginnerTree(l)(c.CharacterId) {
		return r.BeginnerCost(l, c)
	}
	return r.RegularCost(l, c)
}

func (r VIPCabEllinia) BeginnerCost(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("We have a special 90% discount for beginners. The Ant Tunnel is located deep inside in the dungeon that's placed at the center of the Victoria Island, where the 24 Hr Mobile Store is. Would you like to go there for ").
		BlueText().AddText("1,000 mesos").
		BlackText().AddText("?")
	return SendYesNoExit(l, c, m.String(), r.ValidatePayment(1000), r.MoreToOffer, r.MoreToOffer)
}

func (r VIPCabEllinia) RegularCost(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The regular fee applies for all non-beginners. The Ant Tunnel is located deep inside in the dungeon that's placed at the center of the Victoria Island, where 24 Hr Mobile Store is. Would you like to go there for ").
		BlueText().AddText("10,000 mesos").
		BlackText().AddText("?")
	return SendYesNoExit(l, c, m.String(), r.ValidatePayment(10000), r.MoreToOffer, r.MoreToOffer)
}

func (r VIPCabEllinia) ValidatePayment(cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, cost) {
			return r.NotEnoughMeso(l, c)
		}
		return r.ProcessPayment(cost)(l, c)
	}
}

func (r VIPCabEllinia) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It looks like you don't have enough mesos. Sorry but you won't be able to use this without it.")
	return SendNext(l, c, m.String(), Exit())
}

func (r VIPCabEllinia) ProcessPayment(cost uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, int32(cost))
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment of VIP Cab use by character %d.", c.CharacterId)
			return Exit()(l, c)
		}
		return WarpById(_map.AntTunnelPark, 0)(l, c)
	}
}
