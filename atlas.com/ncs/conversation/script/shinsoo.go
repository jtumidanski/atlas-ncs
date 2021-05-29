package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Shinsoo is located in Empress' Road - Ereve (130000000)
type Shinsoo struct {
}

func (r Shinsoo) NPCId() uint32 {
	return npc.Shinsoo
}

func (r Shinsoo) Initial(l logrus.FieldLogger, c Context) State {
	if character.MeetsCriteria(l)(c.CharacterId, character.IsCygnusCriteria(), character.IsJobBranchCriteria(2)) {
		return r.GiveBlessing(l, c)
	} else {
		return r.DoNotStop(l, c)
	}
}

func (r Shinsoo) GiveBlessing(l logrus.FieldLogger, c Context) State {
	character.UseItem(l)(c.CharacterId, item.ShinsoosBlessing)
	m := message.NewBuilder().
		AddText("Let me cast you my blessings, my Knight. Please protect the world of Maple....")
	return SendOk(l, c, m.String())
}

func (r Shinsoo) DoNotStop(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Don't stop training. Every ounce of your energy is required to protect the world of Maple....")
	return SendOk(l, c, m.String())
}