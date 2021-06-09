package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type AceOfHearts struct {
}

func (r AceOfHearts) NPCId() uint32 {
	return npc.AceOfHearts
}

func (r AceOfHearts) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey adventurer! Keep it a secret, ok? We are currently manufacturing the so-called ").
		BlueText().ShowItemName1(item.ChaosScroll60).
		BlackText().AddText(", under Just-in-time marketing strategy. You needed? We're here. So, we act in two fronts: talk to me if you want a good bunch of these. It will be a ").
		BlueText().AddText("Quest").
		BlackText().AddText("-esque procedure, however I will need plenty of ").
		BlueText().AddText("hard-to-get gadgets").
		BlackText().AddText(" from you. I will require a ").
		RedText().AddText("3 days").
		BlackText().AddText(" break after the completion to start working for you again.").NewLine().
		AddText("Talk to my partner here, and he will JIT ").
		BlueText().AddText("synthesize").
		BlackText().AddText(" these scrolls for you, requiring a bunch of ").
		BlueText().AddText("low-cost items").
		BlackText().AddText(", ").
		RedText().AddText("anytime anywhere").
		BlackText().AddText(".")
	return script.SendOk(l, c, m.String())
}
