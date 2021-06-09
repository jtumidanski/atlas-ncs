package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
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

func (r Lira) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Congratulations on getting this far! Well, I suppose I'd better give you the ").
		BlueText().AddText("Breath of Fire").
		BlackText().AddText(". You've certainly earned it!")
	return script.SendNext(l, c, m.String(), r.Validate)
}

func (r Lira) Validate(l logrus.FieldLogger, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.TheBreathOfLava) {
		return r.FreeSlot(l, c)
	}
	return r.HeadOff(l, c)
}

func (r Lira) FreeSlot(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Try freeing a slot to receive the ").
		BlueText().ShowItemName1(item.TheBreathOfLava).
		BlackText().AddText(".")
	return script.SendOk(l, c, m.String())
}

func (r Lira) HeadOff(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Well, time for you to head off.")
	return script.SendNext(l, c, m.String(), r.Process)
}

func (r Lira) Process(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.TheBreathOfLava, 1)
	character.GainExperience(l)(c.CharacterId, 1000)
	return script.WarpById(_map.TheDoorToZakum, 0)(l, c)
}
