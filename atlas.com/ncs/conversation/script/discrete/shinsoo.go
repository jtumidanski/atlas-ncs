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

// Shinsoo is located in Empress' Road - Ereve (130000000)
type Shinsoo struct {
}

func (r Shinsoo) NPCId() uint32 {
	return npc.Shinsoo
}

func (r Shinsoo) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.MeetsCriteria(l, span)(c.CharacterId, character.IsCygnusCriteria(), character.IsJobBranchCriteria(2)) {
		return r.GiveBlessing(l, span, c)
	} else {
		return r.DoNotStop(l, span, c)
	}
}

func (r Shinsoo) GiveBlessing(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.UseItem(l)(c.CharacterId, item.ShinsoosBlessing)
	m := message.NewBuilder().
		AddText("Let me cast you my blessings, my Knight. Please protect the world of Maple....")
	return script.SendOk(l, span, c, m.String())
}

func (r Shinsoo) DoNotStop(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Don't stop training. Every ounce of your energy is required to protect the world of Maple....")
	return script.SendOk(l, span, c, m.String())
}
