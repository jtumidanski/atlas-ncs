package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Robin is located in Maple Road : Snail Hunting Ground I (40000)
type Robin struct {
}

func (r Robin) NPCId() uint32 {
	return 2003
}

func (r Robin) Start(l logrus.FieldLogger, c Context) {
	m := message.NewBuilder().
		AddText("Now...ask me any questions you may have on traveling!!").AddNewLine().
		OpenItem(0).BlueText().AddText("How do I move?").CloseItem().AddNewLine().
		OpenItem(1).BlueText().AddText("How do I take down the monsters?").CloseItem().AddNewLine().
		OpenItem(2).BlueText().AddText("How can I pick up an item?").CloseItem().AddNewLine().
		OpenItem(3).BlueText().AddText("What happens when I die?").CloseItem().AddNewLine().
		OpenItem(4).BlueText().AddText("When can I choose a job?").CloseItem().AddNewLine().
		OpenItem(5).BlueText().AddText("Tell me more about this island!").CloseItem().AddNewLine().
		OpenItem(6).BlueText().AddText("What should I do to become a Warrior?").CloseItem().AddNewLine().
		OpenItem(7).BlueText().AddText("What should I do to become a Bowman?").CloseItem().AddNewLine().
		OpenItem(8).BlueText().AddText("What should I do to become a Magician?").CloseItem().AddNewLine().
		OpenItem(9).BlueText().AddText("What should I do to become a Thief?").CloseItem().AddNewLine().
		OpenItem(10).BlueText().AddText("How do I raise the character stats? (S)").CloseItem().AddNewLine().
		OpenItem(11).BlueText().AddText("How do I check the items that I just picked up?").CloseItem().AddNewLine().
		OpenItem(12).BlueText().AddText("How do I put on an item?").CloseItem().AddNewLine().
		OpenItem(13).BlueText().AddText("How do I check out the items that I'm wearing?").CloseItem().AddNewLine().
		OpenItem(14).BlueText().AddText("What are skills? (K)").CloseItem().AddNewLine().
		OpenItem(15).BlueText().AddText("How do I get to Victoria Island?").CloseItem().AddNewLine().
		OpenItem(16).BlueText().AddText("What are mesos?").CloseItem().
		BlackText()
	npc.Processor(l).SendSimple(c.CharacterId, c.NPCId, m.String())
}

func (r Robin) Continue(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) {
	if mode == 0 && theType == 0 {
		r.Start(l, c)
		return
	}
	if mode == 1 && theType == 0 {
		npc.Processor(l).Dispose(c.CharacterId)
		return
	}

	if selection == 0 {
		m := message.NewBuilder().
			AddText("Alright this is how you move. Use ").
			BlueText().AddText("left, right arrow").
			BlackText().AddText(" to move around the flatland and slanted roads, and press ").
			BlueText().AddText("Alt").
			BlackText().AddText(" to jump. A select number of shoes improve your speed and jumping abilities.")
		npc.Processor(l).SendNext(c.CharacterId, c.NPCId, m.String())
	}

}
