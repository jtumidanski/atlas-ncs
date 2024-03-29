package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// FallenKnight is located in Hidden Street - Cave of Black Witches (924010200)
type FallenKnight struct {
}

func (r FallenKnight) NPCId() uint32 {
	return npc.FallenKnight
}

func (r FallenKnight) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		BlueText().ShowNPC(npc.Eleanor).
		BlackText().AddText("... The black witch... Trapped me here... There's no time now, she's already on her way to ").
		RedText().AddText("attack Ereve").
		BlackText().AddText("!")
	return script.SendNext(l, span, c, m.String(), r.Ready)
}

func (r FallenKnight) Ready(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Fellow Knight, you must reach to ").
		RedText().AddText("Ereve").
		BlackText().AddText(" right now, ").
		RedText().AddText("the Empress is in danger").
		BlackText().AddText("!! Even in this condition, I can still Magic Warp you there. When you're ready talk to me. ").
		BlueText().AddText("Are you ready to face Eleanor?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, script.Exit())
}

func (r FallenKnight) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if _map.CharacterCount(l)(c.WorldId, c.ChannelId, _map.QuietEreve) > 0 {
		return r.AlreadyChallenging(l, span, c)
	}
	return script.WarpById(_map.QuietEreve, 0)(l, span, c)
}

func (r FallenKnight) AlreadyChallenging(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("There's someone already challenging her. Please wait awhile.")
	return script.SendOk(l, span, c, m.String())
}
