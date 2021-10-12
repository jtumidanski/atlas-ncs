package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// MongFromKong is located in Victoria Road - Kerning City (103000000)
type MongFromKong struct {
}

func (r MongFromKong) NPCId() uint32 {
	return npc.MongFromKong
}

func (r MongFromKong) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r MongFromKong) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So, are you going to use the Internet Cafe? There is a fee to use the spaces there, that is ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 5000)).
		BlackText().AddText(". Are you going to enter the Cafe?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, script.Exit())
}

func (r MongFromKong) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 5000) {
		return r.NotEnough(l, span, c)
	}
	return r.Warp(l, span, c)
}

func (r MongFromKong) NotEnough(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Oh, you don't have the money, right? Sorry, I can't let you in.")
	return script.SendOk(l, span, c, m.String())
}

func (r MongFromKong) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -5000)
	return script.WarpById(_map.KerningCityInternetCafe, 0)(l, span, c)
}
