package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// SubwayTrashCan3 is located in Kerning City Subway - Line 1 <Area 1> (103000101)
type SubwayTrashCan3 struct {
}

func (r SubwayTrashCan3) NPCId() uint32 {
	return npc.SubwayTrashCan3
}

func (r SubwayTrashCan3) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r SubwayTrashCan3) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 20710) {
		return r.JustATrashCan(l, span, c)
	}

	if character.HasItem(l, span)(c.CharacterId, item.BubblingDoll) {
		return r.JustATrashCan(l, span, c)
	}

	if !character.CanHold(l)(c.CharacterId, item.BubblingDoll) {
		return r.NotEnoughSpace(l, span, c)
	}

	return r.GiveItem(l, span, c)
}

func (r SubwayTrashCan3) JustATrashCan(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Just a trash can sitting there.")
	return script.SendOk(l, span, c, m.String())
}

func (r SubwayTrashCan3) NotEnoughSpace(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Not enough space in your ETC inventory.")
	return script.SendOk(l, span, c, m.String())
}

func (r SubwayTrashCan3) GiveItem(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.BubblingDoll, 1)
	m := message.NewBuilder().
		AddText("You have found a ").
		BlueText().ShowItemName1(item.BubblingDoll).
		BlackText().AddText(" in the trash can!")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}
