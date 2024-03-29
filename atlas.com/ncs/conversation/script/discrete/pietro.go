package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Pietro is located in Hidden Street - Receiving the Reward For the Event (109050000)
type Pietro struct {
}

func (r Pietro) NPCId() uint32 {
	return npc.Pietro
}

func (r Pietro) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Bam bam bam bam!! You have won the game from the ").NewLine().
		BlueText().AddText("EVENT").
		BlackText().AddText(". Congratulations on making it this far!")
	return script.SendNext(l, span, c, m.String(), r.Prize)
}

func (r Pietro) Prize(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You'll be awarded the ").
		BlueText().AddText("Scroll of Secrets").
		BlackText().AddText(" as the winning prize. On the scroll, it has secret information written in ancient characters.")
	return script.SendNext(l, span, c, m.String(), r.SomethingGood)
}

func (r Pietro) SomethingGood(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The Scroll of Secrets can be deciphered by ").
		RedText().AddText("Chun Ji").
		BlackText().AddText(" or ").NewLine().
		RedText().AddText("Geanie").
		BlackText().AddText(" at Ludibrium. Bring it with you and something good's bound to happen.")
	return script.SendNext(l, span, c, m.String(), r.Validate)
}

func (r Pietro) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.ScrollOfSecrets) {
		return r.MakeRoom(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r Pietro) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.ScrollOfSecrets, 1)
	mapId := character.SavedLocation(l)(c.CharacterId, "EVENT")
	return script.WarpById(mapId, 0)(l, span, c)
}

func (r Pietro) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I think your Etc window is full. Please make room, then talk to me.")
	return script.SendOk(l, span, c, m.String())
}
