package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Mimo struct {
}

func (r Mimo) NPCId() uint32 {
	return npc.Mimo
}

func (r Mimo) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r Mimo) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Wait! You'll figure the stuff out by the time you reach Lv. 10 anyway, but if you absolutely want to prepare beforehand, you may view the following information.").NewLine().NewLine().
		AddText("Tell me, what would you like to know?").NewLine().
		OpenItem(0).BlueText().AddText("About you").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Mini Map").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Quest Window").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Inventory").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("Regular Attack Hunting").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("How to Pick Up Items").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("How to Equip Items").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("Skill Window").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("How to Use Quick Slots").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("How to Break Boxes").CloseItem().NewLine().
		OpenItem(10).BlueText().AddText("How to Sit in a Chair").CloseItem().NewLine().
		OpenItem(11).BlueText().AddText("World Map").CloseItem().NewLine().
		OpenItem(12).BlueText().AddText("Quest Notifications").CloseItem().NewLine().
		OpenItem(13).BlueText().AddText("Enhancing Stats").CloseItem().NewLine().
		OpenItem(14).BlueText().AddText("Who are the Cygnus Knights?").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r Mimo) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.UnderShinsoo
	case 1:
		return r.Hint(1)
	case 2:
		return r.Hint(2)
	case 3:
		return r.Hint(3)
	case 4:
		return r.Hint(4)
	case 5:
		return r.Hint(5)
	case 6:
		return r.Hint(6)
	case 7:
		return r.Hint(7)
	case 8:
		return r.Hint(8)
	case 9:
		return r.Hint(9)
	case 10:
		return r.Hint(10)
	case 11:
		return r.Hint(11)
	case 12:
		return r.Hint(12)
	case 13:
		return r.Hint(13)
	case 14:
		return r.CygnusInfo
	}
	return nil
}

func (r Mimo) UnderShinsoo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I serve under Shinsoo, the guardian of Empress Cygnus. My master, Shinsoo, has ordered me to guide everyone who comes to Maple World to join Cygnus Knights. I will be assisting and following you around until you become a Knight or reach Lv. 11. Please let me know if you have any questions.")
	return script.SendNext(l, span, c, m.String(), r.AskAnyTime)
}

func (r Mimo) AskAnyTime(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("There is no need for you to check this info now. These are basics that you'll pick up as you play. You can always ask me questions that come up after you've reached Lv. 10, so just relax.")
	return script.SendNextPrevious(l, span, c, m.String(), script.Exit(), r.UnderShinsoo)
}

func (r Mimo) Hint(hint uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GuideHint(l)(c.CharacterId, hint)
		return script.Exit()(l, span, c)
	}
}

func (r Mimo) CygnusInfo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The Black Magician is trying to revive and conquer our peaceful Maple World. As a response to this threat, Empress Cygnus has formed a knighthood, now known as Cygnus Knights. You can become a Knight when you reach Lv. 10.")
	return script.SendOk(l, span, c, m.String())
}
