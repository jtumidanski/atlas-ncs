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

// KiruOrbis is located in Orbis - Station (200000161)
type KiruOrbis struct {
}

func (r KiruOrbis) NPCId() uint32 {
	return npc.KiruOrbis
}

func (r KiruOrbis) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This ship will head towards ").
		BlueText().AddText("Ereve").
		BlackText().AddText(", an island where you'll find crimson leaves soaking up the sun, the gentle breeze that glides past the stream, and the Empress of Maple Cygnus. If you're interested in joining the Cygnus Knights, Then you should definitely pay a visit here. Are you interested in visiting Ereve?, The Trip will cost you ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos").NewLine().
		OpenItem(0).BlueText().AddText(" Ereve (1000 mesos)").CloseItem()
	return script.SendListSelectionExit(l, span, c, m.String(), r.DestinationSelection, r.LetMeKnow)
}

func (r KiruOrbis) LetMeKnow(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("OK. If you ever change your mind, please let me know.")
	return script.SendOk(l, span, c, m.String())
}

func (r KiruOrbis) DestinationSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Validate
	}
	return nil
}

func (r KiruOrbis) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 1000) {
		return r.NotEnoughMeso(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r KiruOrbis) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hmm... Are you sure you have ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return script.SendOk(l, span, c, m.String())
}

func (r KiruOrbis) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -1000)
	return script.Warp(_map.EmpressRoadToEreveFromOrbis)(l, span, c)
}
