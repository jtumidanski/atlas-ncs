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

// TrashCan is located in The Nautilus - Top Floor - Hallway (120000100)
type TrashCan struct {
}

func (r TrashCan) NPCId() uint32 {
	return npc.TrashCan
}

func (r TrashCan) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsCompleted(l)(c.CharacterId, 2162) && !character.HasItem(l, span)(c.CharacterId, item.CrumpledLetter) {
		return r.Validate(l, span, c)
	}
	return script.Exit()(l, span, c)
}

func (r TrashCan) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.CrumpledLetter) {
		return r.InventoryFull(l, span, c)
	}
	return r.Give(l, span, c)
}

func (r TrashCan) InventoryFull(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("(You see a Crumpled Paper standing out of the trash can. It's content seems important, but you can't retrieve it since your inventory is full.)")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r TrashCan) Give(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.CrumpledLetter, 1)
	m := message.NewBuilder().
		AddText("(You retrieved a Crumpled Paper standing out of the trash can. It's content seems important.)")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}
