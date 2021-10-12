package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// VIPCabEllinia is located in Victoria Road - Ellinia (101000000)
type VIPCabEllinia struct {
}

func (r VIPCabEllinia) NPCId() uint32 {
	return npc.VIPCabEllinia
}

func (r VIPCabEllinia) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r VIPCabEllinia) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hi there! This cab is for VIP customers only. Instead of just taking you to different towns like the regular cabs, we offer a much better service worthy of VIP class. It's a bit pricey, but... for only 10,000 mesos, we'll take you safely to the").
		BlueText().AddText("Ant Tunnel").
		BlackText().AddText(".")
	return script.SendNextExit(l, span, c, m.String(), r.Cost, r.MoreToOffer)
}

func (r VIPCabEllinia) MoreToOffer(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This town also has a lot to offer. Find us if and when you feel the need to go to the Ant Tunnel Park.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r VIPCabEllinia) Cost(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.IsBeginnerTree(l, span)(c.CharacterId) {
		return r.BeginnerCost(l, span, c)
	}
	return r.RegularCost(l, span, c)
}

func (r VIPCabEllinia) BeginnerCost(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("We have a special 90% discount for beginners. The Ant Tunnel is located deep inside in the dungeon that's placed at the center of the Victoria Island, where the 24 Hr Mobile Store is. Would you like to go there for ").
		BlueText().AddText("1,000 mesos").
		BlackText().AddText("?")
	return script.SendYesNoExit(l, span, c, m.String(), r.ValidatePayment(1000), r.MoreToOffer, r.MoreToOffer)
}

func (r VIPCabEllinia) RegularCost(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The regular fee applies for all non-beginners. The Ant Tunnel is located deep inside in the dungeon that's placed at the center of the Victoria Island, where 24 Hr Mobile Store is. Would you like to go there for ").
		BlueText().AddText("10,000 mesos").
		BlackText().AddText("?")
	return script.SendYesNoExit(l, span, c, m.String(), r.ValidatePayment(10000), r.MoreToOffer, r.MoreToOffer)
}

func (r VIPCabEllinia) ValidatePayment(cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, cost) {
			return r.NotEnoughMeso(l, span, c)
		}
		return r.ProcessPayment(cost)(l, span, c)
	}
}

func (r VIPCabEllinia) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It looks like you don't have enough mesos. Sorry but you won't be able to use this without it.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r VIPCabEllinia) ProcessPayment(cost uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainMeso(l, span)(c.CharacterId, int32(cost))
		return script.WarpById(_map.AntTunnelPark, 0)(l, span, c)
	}
}
