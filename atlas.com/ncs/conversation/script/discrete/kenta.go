package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/event"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// Kenta is located in Aquarium - Zoo (230000003)
type Kenta struct {
}

func (r Kenta) NPCId() uint32 {
	return npc.Kenta
}

func (r Kenta) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsCompleted(l)(c.CharacterId, 6002) {
		return r.ThanksForSaving(l, c)
	}
	if quest.IsStarted(l)(c.CharacterId, 6002) {
		if character.HasItems(l)(c.CharacterId, item.Pheromone, 5) && character.HasItems(l)(c.CharacterId, item.KentasReport, 5) {
			return r.ThanksForSaving(l, c)
		}
		ok := event.StartEvent(l)(c.CharacterId, "3rdJob_mount")
		if !ok {
			return r.SomeoneElse(l, c)
		}
		character.RemoveAll(l)(c.CharacterId, item.Pheromone)
		character.RemoveAll(l)(c.CharacterId, item.KentasReport)
		return script.Exit()(l, c)
	}
	return r.Restricted(l, c)
}

func (r Kenta) ThanksForSaving(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Thanks for saving the pork.")
	return script.SendOk(l, c, m.String())
}

func (r Kenta) Restricted(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Only few adventurers, from a selected public, are eligible to protect the Watch Hog.")
	return script.SendOk(l, c, m.String())
}

func (r Kenta) SomeoneElse(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("There is currently someone in this map, come back later.")
	return script.SendOk(l, c, m.String())
}
