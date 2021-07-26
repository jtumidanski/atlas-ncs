package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// Carta is located in Aqua Road - Carta's Cave (230040001)
type Carta struct {
}

func (r Carta) NPCId() uint32 {
	return npc.Carta
}

func (r Carta) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 6301) {
		return r.DoNotFoolAround(l, c)
	}
	if !character.HasItem(l)(c.CharacterId, item.MiniaturePianus) {
		return r.MustPossessItem(l, c)
	}
	return r.Process(l, c)
}

func (r Carta) Process(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.MiniaturePianus, -1)
	return script.WarpById(_map.WarpedDimension, 0)(l, c)
}

func (r Carta) MustPossessItem(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("In order to open the crack of dimension you will have to possess one piece of Miniature Pianus. Those could be gained by defeating a Pianus.")
	return script.SendOk(l, c, m.String())
}

func (r Carta) DoNotFoolAround(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I'm ").
		BlueText().AddText("Carta the sea-witch.").
		BlackText().AddText(" Don't fool around with me, as I'm known for my habit of turning people into worms.")
	return script.SendOk(l, c, m.String())
}
