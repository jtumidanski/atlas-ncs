package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Carta is located in Aqua Road - Carta's Cave (230040001)
type Carta struct {
}

func (r Carta) NPCId() uint32 {
	return npc.Carta
}

func (r Carta) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 6301) {
		return r.DoNotFoolAround(l, c)
	}
	if !character.HasItem(l)(c.CharacterId, item.MiniaturePianus) {
		return r.MustPossessItem(l, c)
	}
	return r.Process(l, c)
}

func (r Carta) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.MiniaturePianus, -1)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.WarpedDimension, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.WarpedDimension, c.NPCId)
	}
	return Exit()(l, c)
}

func (r Carta) MustPossessItem(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("In order to open the crack of dimension you will have to possess one piece of Miniature Pianus. Those could be gained by defeating a Pianus.")
	return SendOk(l, c, m.String())
}

func (r Carta) DoNotFoolAround(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I'm ").
		BlueText().AddText("Carta the sea-witch.").
		BlackText().AddText(" Don't fool around with me, as I'm known for my habit of turning people into worms.")
	return SendOk(l, c, m.String())
}
