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

// MotherMilkCow1 is located in Hidden Chamber - The Nautilus - Stable (912000100)
type MotherMilkCow1 struct {
}

func (r MotherMilkCow1) NPCId() uint32 {
	return npc.MotherMilkCow1
}

func (r MotherMilkCow1) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.ProgressInt(l)(c.CharacterId, 2180, 1) == 1 {
		return r.CheckAnotherCow(l, span, c)
	}
	if character.HasItem(l, span)(c.CharacterId, item.MilkJug) {
		return r.ProcessEmpty(l, span, c)
	} else if character.HasItem(l, span)(c.CharacterId, item.MilkJugOneThird) {
		return r.ProcessOneThird(l, span, c)
	} else if character.HasItem(l, span)(c.CharacterId, item.MilkJugTwoThird) {
		return r.ProcessTwoThird(l, span, c)
	}
	return script.Exit()(l, span, c)
}

func (r MotherMilkCow1) CheckAnotherCow(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You have taken milk from this cow recently, check another cow.")
	return script.SendOk(l, span, c, m.String())
}

func (r MotherMilkCow1) ProcessEmpty(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.MilkJug, -1)
	character.GainItem(l, span)(c.CharacterId, item.MilkJugOneThird, 1)
	quest.SetProgress(l)(c.CharacterId, 2180, 1, 1)
	m := message.NewBuilder().
		AddText("Now filling up the bottle with milk. The bottle is now 1/3 full of milk.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r MotherMilkCow1) ProcessOneThird(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.MilkJugOneThird, -1)
	character.GainItem(l, span)(c.CharacterId, item.MilkJugTwoThird, 1)
	quest.SetProgress(l)(c.CharacterId, 2180, 1, 1)
	m := message.NewBuilder().
		AddText("Now filling up the bottle with milk. The bottle is now 2/3 full of milk.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r MotherMilkCow1) ProcessTwoThird(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.MilkJugTwoThird, -1)
	character.GainItem(l, span)(c.CharacterId, item.MilkJugFull, 1)
	quest.SetProgress(l)(c.CharacterId, 2180, 1, 1)
	m := message.NewBuilder().
		AddText("Now filling up the bottle with milk. The bottle is now completely full of milk.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}
