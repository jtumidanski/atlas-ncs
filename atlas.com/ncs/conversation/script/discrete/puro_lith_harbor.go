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

// PuroLithHarbor is located in Victoria Road - Lith Harbor (104000000)
type PuroLithHarbor struct {
}

func (r PuroLithHarbor) NPCId() uint32 {
	return npc.PuroLithHarbor
}

func (r PuroLithHarbor) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r PuroLithHarbor) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Are you thinking about leaving Victoria Island and heading to our town? If you board this ship, I can take you from ").
		BlueText().AddText("Lith Harbor").
		BlackText().AddText(" to ").
		BlueText().AddText("Rien").
		BlackText().AddText(" and back. But you must pay a ").
		BlueText().AddText("fee of 800").
		BlackText().AddText(" Mesos. Would you like to go to Rien? It'll take about a minute to get there.").NewLine().
		OpenItem(0).BlueText().AddText(" Rien (800 mesos)").CloseItem()
	return script.SendListSelectionExit(l, span, c, m.String(), r.Selection, r.LetMeKnow)
}

func (r PuroLithHarbor) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Validate
	}
	return nil
}

func (r PuroLithHarbor) LetMeKnow(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("OK. If you ever change your mind, please let me know.")
	return script.SendOk(l, span, c, m.String())
}

func (r PuroLithHarbor) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 800) {
		return r.NotEnoughMeso(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r PuroLithHarbor) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmm... Are you sure you have ").
		BlueText().AddText("800").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return script.SendOk(l, span, c, m.String())
}

func (r PuroLithHarbor) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -800)
	npc.WarpById(l, span)(c.WorldId, c.ChannelId, c.CharacterId, _map.ToRien, 0)
	return nil
}
