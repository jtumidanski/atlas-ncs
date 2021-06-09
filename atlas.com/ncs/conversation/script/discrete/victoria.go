package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Victoria is located in Amoria - Amoria (680000000)
type Victoria struct {
}

func (r Victoria) NPCId() uint32 {
	return npc.Victoria
}

func (r Victoria) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Welcome to the Cathedral.").NewLine().NewLine().
		BlueText().AddText("Couples").
		BlackText().AddText(" wanting to marry on the Cathedral should first ").
		BlueText().AddText("arrange a reservation").
		BlackText().AddText(" with ").
		RedText().ShowNPC(npc.AssistantNicole).
		BlackText().AddText(". When the arranged time comes ").
		BlueText().AddText("both").
		BlackText().AddText(" must show up here, on the same channel from the reservation, and start the ceremony (there is a 10-minutes fault policy to this) by talking to ").
		RedText().ShowNPC(npc.HighPriestJohn).
		BlackText().AddText(". Once arranged, both can ").
		BlueText().AddText("distribute tickets").
		BlackText().AddText(" to friends or acquaintances to become the guests for the marriage.").NewLine().NewLine().
		AddText("The ceremony will start accepting ").
		BlueText().AddText("guests").
		BlackText().AddText(" after the groom and the bride has entered the building. Show the ").
		BlueText().ShowItemName1(item.GoldenMapleLeaf).
		BlackText().AddText(" to ").
		RedText().ShowNPC(npc.AssistantNicole).
		BlackText().AddText(" to access the inner rooms. No one without a ticket is allowed to enter the stage!")
	return script.SendOk(l, c, m.String())
}
