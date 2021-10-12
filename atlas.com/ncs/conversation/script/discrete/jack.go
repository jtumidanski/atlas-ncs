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

// Jack is located in The Nautilus - Top Floor - Hallway (120000100)
type Jack struct {
}

func (r Jack) NPCId() uint32 {
	return npc.Jack
}

func (r Jack) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.DirtyTreasureMap) {
		return r.ScratchScratch(l, span, c)
	}
	return r.CanIKeepIt(l, span, c)
}

func (r Jack) ScratchScratch(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("(Scratch scratch...)")
	return script.SendOk(l, span, c, m.String())
}

func (r Jack) CanIKeepIt(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey, nice ").
		BlueText().AddText("Treasure Map").
		BlackText().AddText(" you have there? ").
		RedText().AddText("Can I keep it").
		BlackText().AddText(" for the Nautilus crew, if you don't need it any longer?")
	return script.SendYesNo(l, span, c, m.String(), r.Remove, script.Exit())
}

func (r Jack) Remove(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.DirtyTreasureMap, -1)
	return script.Exit()(l, span, c)
}
