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

// HiddenNote2 is located in Victoria Road - Pet-Walking Road (100000202)
type HiddenNote2 struct {
}

func (r HiddenNote2) NPCId() uint32 {
	return npc.HiddenNote2
}

func (r HiddenNote2) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(I can see something covered in grass. Should I pull it out?)")
	return script.SendYesNo(l, span, c, m.String(), r.Yuck, r.Nah)
}

func (r HiddenNote2) Yuck(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.PetPoop, 1)
	m := message.NewBuilder().BlueText().
		AddText("(Yuck... it's pet poop!)")
	return script.SendOk(l, span, c, m.String())
}

func (r HiddenNote2) Nah(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(I didn't think much of it, so I didn't touch it.)")
	return script.SendOk(l, span, c, m.String())
}
