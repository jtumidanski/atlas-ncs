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
	return 12101
}

func (r Rain) Initial() State {
	return r.TownCalledAmherst
}

func (r Rain) TownCalledAmherst(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("This is the town called ").
		BlueText().AddText("Amherst").
		BlackText().AddText(", located at the northeast part of the Maple Island. You know that Maple Island is for beginners, right? I'm glad there are only weak monsters around this place.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.GoToSouthPerry)
}

func (r Rain) GoToSouthPerry(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("If you want to get stronger, then go to ").
		BlueText().AddText("Southperry").
		BlackText().AddText(" where there's a harbor. Ride on the gigantic ship and head to the place called ").
		BlueText().AddText("Victoria Island").
		BlackText().AddText(". It's incomparable in size compared to this tiny island.")
	conversation.SendNextPrevious(m.String())
	return NextPrevious(GenericExit, r.ChooseYourJob, r.TownCalledAmherst)
}

func (r Rain) ChooseYourJob(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("At the Victoria Island, you can choose your job. Is it called ").
		BlueText().AddText("Perion").
		BlackText().AddText("...? I heard there's a bare, desolate town where warriors live. A highland...what kind of a place would that be?")
	conversation.SendPrevious(m.String())
	return Previous(GenericExit, r.GoToSouthPerry)
}