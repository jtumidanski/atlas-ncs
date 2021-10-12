package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// SuspiciousMan is located in Dungeon   - Stairway to the Underground Temple (105100000)
type SuspiciousMan struct {
}

func (r SuspiciousMan) NPCId() uint32 {
	return npc.SuspiciousMan
}

func (r SuspiciousMan) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello, ").
		ShowCharacterName().
		AddText(". I can exchange your Balrog Leathers.").NewLine().NewLine().
		OpenItem(0).RedText().AddText("Redeem items").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r SuspiciousMan) Selection(_ int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Well, okay. These are what you can redeem...").
			OpenItem(0).ShowItemImage1(item.BalrogsSTRScroll30).ShowItemName2(item.BalrogsSTRScroll30).CloseItem().NewLine().
			OpenItem(1).ShowItemImage1(item.BalrogsINTScroll30).ShowItemName2(item.BalrogsINTScroll30).CloseItem().NewLine().
			OpenItem(2).ShowItemImage1(item.BalrogsLUKScroll30).ShowItemName2(item.BalrogsLUKScroll30).CloseItem().NewLine().
			OpenItem(3).ShowItemImage1(item.BalrogsDEXScroll30).ShowItemName2(item.BalrogsDEXScroll30).CloseItem().NewLine().
			OpenItem(4).ShowItemImage1(item.BalrogsHPScroll30).ShowItemName2(item.BalrogsHPScroll30).CloseItem().NewLine().
			OpenItem(5).ShowItemImage1(item.BalrogsMPScroll30).ShowItemName2(item.BalrogsMPScroll30).CloseItem().NewLine().
			OpenItem(6).ShowItemImage1(item.BalrogsSpeedScroll30).ShowItemName2(item.BalrogsSpeedScroll30).CloseItem().NewLine().
			OpenItem(7).ShowItemImage1(item.BalrogsJumpScroll30).ShowItemName2(item.BalrogsJumpScroll30).CloseItem().NewLine().
			OpenItem(8).ShowItemImage1(item.BalrogsAccuracyScroll30).ShowItemName2(item.BalrogsAccuracyScroll30).CloseItem().NewLine().
			OpenItem(9).ShowItemImage1(item.BalrogsAvoidabilityScroll30).ShowItemName2(item.BalrogsAvoidabilityScroll30).CloseItem().NewLine().
			OpenItem(10).ShowItemImage1(item.BalrogsDefenseScroll30).ShowItemName2(item.BalrogsDefenseScroll30).CloseItem().NewLine().
			OpenItem(11).ShowItemImage1(item.BalrogsTwilightScroll5).ShowItemName2(item.BalrogsTwilightScroll5).CloseItem().NewLine()
		return script.SendListSelection(l, span, c, m.String(), r.Redeem)
	}
}

func (r SuspiciousMan) Redeem(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Validate(item.BalrogsSTRScroll30)
	case 1:
		return r.Validate(item.BalrogsINTScroll30)
	case 2:
		return r.Validate(item.BalrogsLUKScroll30)
	case 3:
		return r.Validate(item.BalrogsDEXScroll30)
	case 4:
		return r.Validate(item.BalrogsHPScroll30)
	case 5:
		return r.Validate(item.BalrogsMPScroll30)
	case 6:
		return r.Validate(item.BalrogsSpeedScroll30)
	case 7:
		return r.Validate(item.BalrogsJumpScroll30)
	case 8:
		return r.Validate(item.BalrogsAccuracyScroll30)
	case 9:
		return r.Validate(item.BalrogsAvoidabilityScroll30)
	case 10:
		return r.Validate(item.BalrogsDefenseScroll30)
	case 11:
		return r.Validate(item.BalrogsTwilightScroll5)
	}
	return nil
}

func (r SuspiciousMan) Validate(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.CanHold(l)(c.CharacterId, itemId) {
			return r.MakeRoom(l, span, c)
		}
		if !character.HasItem(l, span)(c.CharacterId, item.PieceOfBalrogLeather) {
			return r.MissingLeather(l, span, c)
		}
		return r.Process(l, span, c, itemId)
	}
}

func (r SuspiciousMan) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context, itemId uint32) script.State {
	character.GainItem(l, span)(c.CharacterId, item.PieceOfBalrogLeather, -1)
	character.GainItem(l, span)(c.CharacterId, itemId, 1)
	return r.Success(l, span, c)
}

func (r SuspiciousMan) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Thank you for your redemption")
	return script.SendOk(l, span, c, m.String())
}

func (r SuspiciousMan) MissingLeather(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't have enough leathers.")
	return script.SendOk(l, span, c, m.String())
}

func (r SuspiciousMan) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please make room")
	return script.SendOk(l, span, c, m.String())
}
