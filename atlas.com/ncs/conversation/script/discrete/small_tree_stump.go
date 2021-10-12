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

// SmallTreeStump is located in Victoria Road - Top of the Tree That Grew (101010103)
type SmallTreeStump struct {
}

func (r SmallTreeStump) NPCId() uint32 {
	return npc.SmallTreeStump
}

func (r SmallTreeStump) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 20716) {
		return r.NeverEndingFlow(l, span, c)
	}
	if character.HasItem(l, span)(c.CharacterId, item.ClearTreeSap) {
		return r.NeverEndingFlow(l, span, c)
	}
	if !character.CanHold(l)(c.CharacterId, item.ClearTreeSap) {
		return r.MakeRoom(l, span, c)
	}

	return r.GainTreeSap(l, span, c)
}

func (r SmallTreeStump) GainTreeSap(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.ClearTreeSap, 1)
	return r.ShowSuccess(l, span, c)
}

func (r SmallTreeStump) ShowSuccess(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You bottled up some of the clear tree sap.  ").
		ShowItemImage2(item.ClearTreeSap)
	return script.SendOk(l, span, c, m.String())
}

func (r SmallTreeStump) NeverEndingFlow(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("A never ending flow of sap is coming from this small tree stump.")
	return script.SendOk(l, span, c, m.String())
}

func (r SmallTreeStump) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Make sure you have a free spot in your ETC inventory.")
	return script.SendOk(l, span, c, m.String())
}
