package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Lira is located in Adobis's Mission I - Breath of Lava <Level 2> (280020001)
type Lira struct {
}

func (r Lira) NPCId() uint32 {
	return npc.Lira
}

func (r Lira) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Congratulations on getting this far! Well, I suppose I'd better give you the ").
		BlueText().AddText("Breath of Fire").
		BlackText().AddText(". You've certainly earned it!")
	return SendNext(l, c, m.String(), r.Validate)
}

func (r Lira) Validate(l logrus.FieldLogger, c Context) State {
	if !character.CanHold(l)(c.CharacterId, item.TheBreathOfLava) {
		return r.FreeSlot(l, c)
	}
	return r.HeadOff(l, c)
}

func (r Lira) FreeSlot(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Try freeing a slot to receive the ").
		BlueText().ShowItemName1(item.TheBreathOfLava).
		BlackText().AddText(".")
	return SendOk(l, c, m.String())
}

func (r Lira) HeadOff(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Well, time for you to head off.")
	return SendNext(l, c, m.String(), r.Process)
}

func (r Lira) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.TheBreathOfLava, 1)
	character.GainExperience(l)(c.CharacterId, 1000)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.TheDoorToZakum, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.TheDoorToZakum, c.NPCId)
	}
	return Exit()(l, c)
}
