package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Shane is located in Victoria Road - Ellinia (101000000)
type Shane struct {
}

func (r Shane) NPCId() uint32 {
	return npc.Shane
}

func (r Shane) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.AboveLevel(l, span)(c.CharacterId, 25) {
		return r.LevelRequirement(l, span, c)
	}

	return r.WouldYouLike(l, span, c)
}

func (r Shane) LevelRequirement(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You must be a higher level to enter the Forest of Patience.")
	return script.SendOk(l, span, c, m.String())
}

func (r Shane) WouldYouLike(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hi, i'm Shane. I can let you into the Forest of Patience for a small fee. Would you like to enter for ").
		BlueText().AddText(fmt.Sprintf("5000")).
		BlackText().AddText(" mesos?")
	return script.SendYesNo(l, span, c, m.String(), r.ValidatePayment, r.SeeYouNextTime)
}

func (r Shane) SeeYouNextTime(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright, see you next time.")
	return script.SendOk(l, span, c, m.String())
}

func (r Shane) ValidatePayment(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 5000) {
		return r.NotEnoughMeso(l, span, c)
	}

	character.GainMeso(l, span)(c.CharacterId, -5000)

	var destination uint32
	if quest.IsStarted(l)(c.CharacterId, 2050) {
		destination = _map.TheForestOfPatienceStep1
	} else if quest.IsStarted(l)(c.CharacterId, 2051) {
		destination = _map.TheForestOfPatienceStep3
	} else if character.MeetsCriteria(l, span)(c.CharacterId, character.LevelBetweenCriteria(24, 50)) {
		destination = _map.TheForestOfPatienceStep1
	} else if character.IsLevel(l, span)(c.CharacterId, 50) {
		destination = _map.TheForestOfPatienceStep3
	} else {
		l.Warnf("Unsure where to send player, returning money and warping to the same map.")
		character.GainMeso(l, span)(c.CharacterId, 5000)
		destination = c.MapId
	}
	return script.WarpById(destination, 0)(l, span, c)
}

func (r Shane) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Sorry but it doesn't like you have enough mesos!")
	return script.SendOk(l, span, c, m.String())
}
