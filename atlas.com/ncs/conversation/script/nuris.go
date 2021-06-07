package script

import (
	"atlas-ncs/event"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Nuris is located in Sharenian - Returning Path (990001100)
type Nuris struct {
}

func (r Nuris) NPCId() uint32 {
	return npc.Nuris
}

func (r Nuris) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("It seems you have finished exploring Sharenian Keep, yes? Are you going to return to the recruitment map now?")
	return SendYesNo(l, c, m.String(), r.Process, Exit())
}

func (r Nuris) Process(l logrus.FieldLogger, c Context) State {
	if event.Cleared(l)(c.CharacterId) {
		ok := event.GiveEventReward(l)(c.CharacterId)
		if !ok {
			return r.MakeRoom(l, c)
		}
	}
	return Warp(_map.ExcavationSiteCamp)(l, c)
}

func (r Nuris) MakeRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It seems you don't have a free slot in either your ").
		RedText().AddText("Equip").
		BlackText().AddText(", ").
		RedText().AddText("Use").
		BlackText().AddText(" or ").
		RedText().AddText("Etc").
		BlackText().AddText(" inventories. Please make some room first.")
	return SendOk(l, c, m.String())
}
