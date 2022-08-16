package discrete

import (
	"atlas-ncs/character/location"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Gaga struct {
}

func (r Gaga) NPCId() uint32 {
	return npc.Gaga
}

func (r Gaga) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey, traveler! I am ").
		ShowNPC(npc.Gaga).
		AddText(", and my job is to recruit travelers like you, who are eager for new challenges daily. Right now, my team is holding contests that thoroughly tests the mental and physical capabilities of adventurers like you.")
	return script.SendNext(l, span, c, m.String(), r.BossFights)
}

func (r Gaga) BossFights(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("These contests involve ").
		BlueText().AddText("sequential boss fights").
		BlackText().AddText(", with some resting spots between some sections. These will require some strategy time and enough supplies at hand, as they are not common fights.")
	return script.SendNextPrevious(l, span, c, m.String(), r.Confirm, r.Initial)
}

func (r Gaga) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If you feel you are powerful enough, you can join others like you at where we are hosting the contests of power. ... So, what is your decision? Will you come to where the contests are being held right now?")
	return script.SendAcceptDecline(l, span, c, m.String(), r.VeryWell, script.Exit())
}

func (r Gaga) VeryWell(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Very well. Remember, there you can assemble a team or take on the fighting on your own, it's up to you. Good luck!")
	return script.SendOkTrigger(l, span, c, m.String(), r.Warp)
}

func (r Gaga) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	location.SaveLocation(l, span)(c.CharacterId, "BOSSPQ")
	return script.WarpByName(_map.ExclusiveTrainingCenter, "out00")(l, span, c)
}
