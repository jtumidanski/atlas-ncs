package script

import (
	"atlas-ncs/character"
	"atlas-ncs/event"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/party"
	"github.com/sirupsen/logrus"
	"strconv"
)

// BlockedEntrance is located in Mushroom Castle - The Last Castle Tower (106021402)
type BlockedEntrance struct {
}

func (r BlockedEntrance) NPCId() uint32 {
	return npc.BlockedEntrance
}

func (r BlockedEntrance) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.TheLastCastleTower {
		if !character.QuestCompleted(l)(c.CharacterId, 2331) {
			return Exit()(l, c)
		}
		return r.ShowBossFights(l, c)
	}
	return r.LimitedSelection(l, c)
}

func (r BlockedEntrance) ShowBossFights(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		OpenItem(0).AddText("Enter to fight ").
		BlueText().AddText("King Pepe").
		BlackText().AddText(" and ").
		BlueText().AddText("Yeti Brothers").
		BlackText().AddText(".").CloseItem().NewLine().
		OpenItem(1).AddText("Enter to fight ").
		BlueText().AddText("Prime Minister").
		BlackText().AddText(".").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r BlockedEntrance) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.KingPepe
	case 1:
		return r.PrimeMinister
	}
	return nil
}

func (r BlockedEntrance) KingPepe(l logrus.FieldLogger, c Context) State {
	event.SetProperty(l)("KingPepeAndYetis", "player", strconv.Itoa(int(c.CharacterId)))
	event.StartEvent(l)(c.CharacterId, "KingPepeAndYetis")
	return Exit()(l, c)
}

func (r BlockedEntrance) PrimeMinister(l logrus.FieldLogger, c Context) State {
	p, err := party.GetParty(l)(c.CharacterId)
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve character %d party.", c.CharacterId)
	}

	ok := true
	if p != nil {
		ok = event.StartPartyEvent(l)("MK_PrimeMinister2", p.Id(), c.MapId, 1)
	} else {
		ok = event.StartEvent(l)(c.CharacterId, "MK_PrimeMinister2")
	}
	if !ok {
		return r.AnotherParty(l, c)
	}
	return Exit()(l, c)
}

func (r BlockedEntrance) AnotherParty(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Another party is already challenging the boss in this channel.")
	return SendOk(l, c, m.String())
}

func (r BlockedEntrance) LimitedSelection(l logrus.FieldLogger, c Context) State {
	questProgress := character.QuestProgressInt(l)(c.CharacterId, 2330, 3300005)
	questProgress += character.QuestProgressInt(l)(c.CharacterId, 2330, 3300006)
	questProgress += character.QuestProgressInt(l)(c.CharacterId, 2330, 3300007)

	if !(questProgress < 3 && character.QuestStarted(l)(c.CharacterId, 2330)) {
		return Exit()(l, c)
	}
	return r.ShowBossFight(l, c)
}

func (r BlockedEntrance) ShowBossFight(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		OpenItem(0).AddText("Enter to fight ").
		BlueText().AddText("King Pepe").
		BlackText().AddText(" and ").
		BlueText().AddText("Yeti Brothers").
		BlackText().AddText(".").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}
