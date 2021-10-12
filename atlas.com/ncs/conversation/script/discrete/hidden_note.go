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

// HiddenNote is located in Victoria Road - Pet-Walking Road (100000202)
type HiddenNote struct {
}

func (r HiddenNote) NPCId() uint32 {
	return npc.HiddenNote
}

func (r HiddenNote) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 4646) {
		return r.CouldNotFindAnything(l, span, c)
	}

	if character.HasItem(l, span)(c.CharacterId, item.HiddenNote) {
		return r.Eww(l, span, c)
	}

	return r.SeeSomething(l, span, c)
}

func (r HiddenNote) Eww(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(What's this... eww... a pet's poop was in there!)")
	return script.SendOk(l, span, c, m.String())
}

func (r HiddenNote) SeeSomething(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(I can see something covered in grass. Should I pull it out?)")
	return script.SendYesNo(l, span, c, m.String(), r.AwardNote, r.Covered)
}

func (r HiddenNote) CouldNotFindAnything(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(I couldn't find anything.)")
	return script.SendOk(l, span, c, m.String())
}

func (r HiddenNote) AwardNote(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.HiddenNote, 1)
	m := message.NewBuilder().
		AddText("I found the item that Pet Trainer Bartos hid... this note.")
	return script.SendOk(l, span, c, m.String())
}

func (r HiddenNote) Covered(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(I didn't touch this hidden item covered in grass)")
	return script.SendOk(l, span, c, m.String())
}
