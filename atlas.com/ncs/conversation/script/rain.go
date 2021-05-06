package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Rain is located in Maple Road : Amherst (1010000)
type Rain struct {
}

func (r Rain) NPCId() uint32 {
	return npc.Rain
}

func (r Rain) Initial(l logrus.FieldLogger, c Context) State {
	return r.TownCalledAmherst(l, c)
}

func (r Rain) TownCalledAmherst(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("This is the town called ").
		BlueText().AddText("Amherst").
		BlackText().AddText(", located at the northeast part of the Maple Island. You know that Maple Island is for beginners, right? I'm glad there are only weak monsters around this place.")
	return SendNext(l, c, m.String(), r.GoToSouthPerry)
}

func (r Rain) GoToSouthPerry(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you want to get stronger, then go to ").
		BlueText().AddText("Southperry").
		BlackText().AddText(" where there's a harbor. Ride on the gigantic ship and head to the place called ").
		BlueText().AddText("Victoria Island").
		BlackText().AddText(". It's incomparable in size compared to this tiny island.")
	return SendNextPrevious(l, c, m.String(), r.ChooseYourJob, r.TownCalledAmherst)
}

func (r Rain) ChooseYourJob(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("At the Victoria Island, you can choose your job. Is it called ").
		BlueText().AddText("Perion").
		BlackText().AddText("...? I heard there's a bare, desolate town where warriors live. A highland...what kind of a place would that be?")
	return SendPrevious(l, c, m.String(), r.GoToSouthPerry)
}