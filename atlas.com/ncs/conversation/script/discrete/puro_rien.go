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

// PuroRien is located in Snow Island - Dangerous Forest (140020300)
type PuroRien struct {
}

func (r PuroRien) NPCId() uint32 {
	return npc.PuroRien
}

func (r PuroRien) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r PuroRien) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Are you trying to leave Rien? Board this ship and I'll take you from ").
		BlueText().AddText("Rien").
		BlackText().AddText(" to ").
		BlueText().AddText("Lith Harbor").
		BlackText().AddText(" and back. for a ").
		BlueText().AddText("fee of 800").
		BlackText().AddText(" Mesos. Would you like to head over to Lith Harbor now? It'll take about a minute to get there.").NewLine().
		OpenItem(0).BlueText().AddText(" Lith Harbor (800 mesos)").CloseItem()
	return script.SendListSelectionExit(l, span, c, m.String(), r.Selection, r.LetMeKnow)
}

func (r PuroRien) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Validate
	}
	return nil
}

func (r PuroRien) LetMeKnow(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("OK. If you ever change your mind, please let me know.")
	return script.SendOk(l, span, c, m.String())
}

func (r PuroRien) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 800) {
		return r.NotEnoughMeso(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r PuroRien) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmm... Are you sure you have ").
		BlueText().AddText("800").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return script.SendOk(l, span, c, m.String())
}

func (r PuroRien) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -800)
	npc.WarpById(l, span)(c.WorldId, c.ChannelId, c.CharacterId, _map.ToLithHarbor, 0)
	return nil
}