package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// AmosTheStrong is located in Hidden Street - Amos' Training Ground (670010000)
type AmosTheStrong struct {
}

func (r AmosTheStrong) NPCId() uint32 {
	return npc.AmosTheStrong
}

func (r AmosTheStrong) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("My name is Amos the Strong. What would you like to do?").NewLine().
		OpenItem(0).BlueText().AddText("Enter the Amorian Challenge!!").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Trade 10 Keys for a Ticket!").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r AmosTheStrong) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Enter
	case 1:
		return r.Trade
	}
	return nil
}

func (r AmosTheStrong) Enter(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.EntranceTicket) {
		return r.MustHaveTicket(l, c)
	}
	m := message.NewBuilder().AddText("So you would like to enter the ").BlueText().AddText("Entrance").BlackText().AddText("?")
	return SendYesNo(l, c, m.String(), r.Process, r.ComeBack)
}

func (r AmosTheStrong) MustHaveTicket(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You must have an Entrance Ticket to enter.")
	return SendOk(l, c, m.String())
}

func (r AmosTheStrong) Trade(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.EntranceTicket) {
		return r.AlreadyHave(l, c)
	}

	if !character.HasItems(l)(c.CharacterId, item.LipLockKey, 10) {
		return r.GetMeKeys(l, c)
	}

	return r.ConfirmExchange(l, c)
}

func (r AmosTheStrong) ConfirmExchange(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("So you would like a Ticket?")
	return SendYesNo(l, c, m.String(), r.Exchange, r.ComeBack)
}

func (r AmosTheStrong) GetMeKeys(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Please get me 10 Keys first!")
	return SendOk(l, c, m.String())
}

func (r AmosTheStrong) AlreadyHave(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You already have an Entrance Ticket!")
	return SendOk(l, c, m.String())
}

func (r AmosTheStrong) ComeBack(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Ok come back when you're ready.")
	return SendOk(l, c, m.String())
}

func (r AmosTheStrong) Exchange(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.LipLockKey, -10)
	character.GainItem(l)(c.CharacterId, item.EntranceTicket, 1)
	return Exit()(l, c)
}

func (r AmosTheStrong) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.EntranceTicket, -1)
	return WarpById(_map.EntranceOfAmorianChallenge, 0)(l, c)
}
