package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// AmosTheStrong is located in Hidden Street - Amos' Training Ground (670010000)
type AmosTheStrong struct {
}

func (r AmosTheStrong) NPCId() uint32 {
	return npc.AmosTheStrong
}

func (r AmosTheStrong) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("My name is Amos the Strong. What would you like to do?").NewLine().
		OpenItem(0).BlueText().AddText("Enter the Amorian Challenge!!").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Trade 10 Keys for a Ticket!").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r AmosTheStrong) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Enter
	case 1:
		return r.Trade
	}
	return nil
}

func (r AmosTheStrong) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.EntranceTicket) {
		return r.MustHaveTicket(l, span, c)
	}
	m := message.NewBuilder().AddText("So you would like to enter the ").BlueText().AddText("Entrance").BlackText().AddText("?")
	return script.SendYesNo(l, span, c, m.String(), r.Process, r.ComeBack)
}

func (r AmosTheStrong) MustHaveTicket(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You must have an Entrance Ticket to enter.")
	return script.SendOk(l, span, c, m.String())
}

func (r AmosTheStrong) Trade(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.EntranceTicket) {
		return r.AlreadyHave(l, span, c)
	}

	if !character.HasItems(l, span)(c.CharacterId, item.LipLockKey, 10) {
		return r.GetMeKeys(l, span, c)
	}

	return r.ConfirmExchange(l, span, c)
}

func (r AmosTheStrong) ConfirmExchange(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("So you would like a Ticket?")
	return script.SendYesNo(l, span, c, m.String(), r.Exchange, r.ComeBack)
}

func (r AmosTheStrong) GetMeKeys(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please get me 10 Keys first!")
	return script.SendOk(l, span, c, m.String())
}

func (r AmosTheStrong) AlreadyHave(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You already have an Entrance Ticket!")
	return script.SendOk(l, span, c, m.String())
}

func (r AmosTheStrong) ComeBack(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Ok come back when you're ready.")
	return script.SendOk(l, span, c, m.String())
}

func (r AmosTheStrong) Exchange(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.LipLockKey, -10)
	character.GainItem(l, span)(c.CharacterId, item.EntranceTicket, 1)
	return script.Exit()(l, span, c)
}

func (r AmosTheStrong) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.EntranceTicket, -1)
	return script.WarpById(_map.EntranceOfAmorianChallenge, 0)(l, span, c)
}
