package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// FallenKnight is located in Hidden Street - Cave of Black Witches (924010200)
type FallenKnight struct {
}

func (r FallenKnight) NPCId() uint32 {
	return npc.FallenKnight
}

func (r FallenKnight) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		BlueText().ShowNPC(npc.Eleanor).
		BlackText().AddText("... The black witch... Trapped me here... There's no time now, she's already on her way to ").
		RedText().AddText("attack Ereve").
		BlackText().AddText("!")
	return SendNext(l, c, m.String(), r.Ready)
}

func (r FallenKnight) Ready(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Fellow Knight, you must reach to ").
		RedText().AddText("Ereve").
		BlackText().AddText(" right now, ").
		RedText().AddText("the Empress is in danger").
		BlackText().AddText("!! Even in this condition, I can still Magic Warp you there. When you're ready talk to me. ").
		BlueText().AddText("Are you ready to face Eleanor?")
	return SendYesNo(l, c, m.String(), r.Validate, Exit())
}

func (r FallenKnight) Validate(l logrus.FieldLogger, c Context) State {
	if _map.CharacterCount(l)(c.WorldId, c.ChannelId, _map.QuietEreve) > 0 {
		return r.AlreadyChallenging(l, c)
	}
	return WarpById(_map.QuietEreve, 0)(l, c)
}

func (r FallenKnight) AlreadyChallenging(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("There's someone already challenging her. Please wait awhile.")
	return SendOk(l, c, m.String())
}
