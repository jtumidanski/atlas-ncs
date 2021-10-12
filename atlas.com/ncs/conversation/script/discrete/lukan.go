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

// Lukan is located in Phantom Forest - Phantom Road (610010003)
type Lukan struct {
}

func (r Lukan) NPCId() uint32 {
	return npc.Lukan
}

func (r Lukan) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsCompleted(l)(c.CharacterId, 8223) {
		return r.BraveAdventurer(l, span, c)
	}

	if !character.CanHold(l)(c.CharacterId, item.CrimsonwoodKeystone) {
		return r.NeedInventoryRoom(l, span, c)
	}

	return r.GiveKeystone(l, span, c)
}

func (r Lukan) GiveKeystone(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.CrimsonwoodKeystone, 1)
	return r.DoNotLoseAgain(l, span, c)
}

func (r Lukan) DoNotLoseAgain(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So you lost your key, right? Very well, I will craft you another one, but please don't lose it again. It is fundamental to enter the Inner Sanctum, inside the Keep.")
	return script.SendOk(l, span, c, m.String())
}

func (r Lukan) NeedInventoryRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please make a slot on your SETUP ready for the key I have to give to you. It is fundamental to enter the Inner Sanctum, inside the Keep.")
	return script.SendOk(l, span, c, m.String())
}

func (r Lukan) BraveAdventurer(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("O, brave adventurer. The Stormcasters house, from which I belong, guards the surrounding area of Yore, this landscape, from the forces of the Twisted Masters' guard that daily threatens the citizens. Please help us on the defense of Yore.")
	return script.SendOk(l, span, c, m.String())
}
