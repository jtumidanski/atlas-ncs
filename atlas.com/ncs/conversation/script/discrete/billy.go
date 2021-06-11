package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Billy is located in Premium Road - Kerning City Internet Cafe (193000000)
type Billy struct {
}

func (r Billy) NPCId() uint32 {
	return npc.Billy
}

func (r Billy) Initial(l logrus.FieldLogger, c script.Context) script.State {
	tiers := r.GetTiers()
	return r.Hello(tiers)(l, c)
}

func (r Billy) Hello(tiers []Tier) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("The ").
			BlueText().AddText("Internet Cafe Party Quest").
			BlackText().AddText(" rewards players with ticket-like ").
			BlueText().AddText("figure erasers").
			BlackText().AddText(", that can be used on the vending machine to retrieve prizes. By further increasing the stakes, one can get better prizes, separated by ").
			RedText().AddText("tiers").
			BlackText().AddText(".").NewLine().
			AddText("The possible rewards for each tier are depicted here:").NewLine()
		for i, tier := range tiers {
			m = m.OpenItem(i).BlueText().AddText(tier.Name).CloseItem()
		}
		return script.SendListSelection(l, c, m.String(), r.Selection(tiers))
	}
}

func (r Billy) GetTiers() []Tier {
	return []Tier{
		r.Tier1(),
		r.Tier2(),
		r.Tier3(),
		r.Tier4(),
		r.Tier5(),
		r.Tier6(),
	}
}

func (r Billy) Selection(tiers []Tier) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		if selection < 0 || int(selection) >= len(tiers) {
			return script.Exit()
		}
		tier := tiers[selection]
		m := message.NewBuilder().
			AddText("The following items are being awarded at ").
			BlueText().AddText(tier.Name).NewLine().NewLine()
		for i, prize := range tier.Prizes {
			qty := ""
			if prize.Amount > 1 {
				qty = fmt.Sprintf(" ( %d )", prize.Amount)
			}
			m = m.OpenItem(i).ShowItemImage2(prize.ItemId).AddText(" ").ShowItemName1(prize.ItemId).AddText(qty).CloseItem()
		}
		return func(l logrus.FieldLogger, c script.Context) script.State {
			return script.SendPrevious(l, c, m.String(), r.Hello(tiers))
		}
	}
}

func (r Billy) Tier1() Tier {
	return Tier{
		Name: "Tier 1",
		Prizes: []Prize{
			{ItemId: 1302021, Amount: 1},
			{ItemId: 1302024, Amount: 1},
			{ItemId: 1302033, Amount: 1},
			{ItemId: 1082150, Amount: 1},
			{ItemId: 1002419, Amount: 1},
			{ItemId: 2022053, Amount: 20},
			{ItemId: 2022054, Amount: 20},
			{ItemId: 2020032, Amount: 20},
			{ItemId: 2022057, Amount: 20},
			{ItemId: 2022096, Amount: 20},
			{ItemId: 2022097, Amount: 25},
			{ItemId: 2022192, Amount: 25},
			{ItemId: 2020030, Amount: 25},
			{ItemId: 2010005, Amount: 50},
			{ItemId: 2022041, Amount: 50},
			{ItemId: 2030000, Amount: 12},
			{ItemId: 2040100, Amount: 1},
			{ItemId: 2040004, Amount: 1},
			{ItemId: 2040207, Amount: 1},
			{ItemId: 2048004, Amount: 1},
			{ItemId: 4031203, Amount: 3},
			{ItemId: 4000021, Amount: 4},
			{ItemId: 4003005, Amount: 2},
			{ItemId: 4003000, Amount: 2},
			{ItemId: 4003001, Amount: 1},
			{ItemId: 4010000, Amount: 2},
			{ItemId: 4010001, Amount: 2},
			{ItemId: 4010002, Amount: 2},
			{ItemId: 4010005, Amount: 2},
			{ItemId: 4020004, Amount: 2},
		},
	}
}

func (r Billy) Tier2() Tier {
	return Tier{
		Name: "Tier 2",
		Prizes: []Prize{
			{ItemId: 1022073, Amount: 1},
			{ItemId: 1012098, Amount: 1},
			{ItemId: 1012101, Amount: 1},
			{ItemId: 1012102, Amount: 1},
			{ItemId: 1012103, Amount: 1},
			{ItemId: 2022055, Amount: 40},
			{ItemId: 2022056, Amount: 40},
			{ItemId: 2022103, Amount: 40},
			{ItemId: 2020029, Amount: 40},
			{ItemId: 2020032, Amount: 60},
			{ItemId: 2020031, Amount: 60},
			{ItemId: 2022191, Amount: 60},
			{ItemId: 2022016, Amount: 60},
			{ItemId: 2043300, Amount: 1},
			{ItemId: 2043110, Amount: 1},
			{ItemId: 2043800, Amount: 1},
			{ItemId: 2041001, Amount: 1},
			{ItemId: 2040903, Amount: 1},
			{ItemId: 4031203, Amount: 4},
			{ItemId: 4000021, Amount: 6},
			{ItemId: 4003005, Amount: 7},
			{ItemId: 4003000, Amount: 5},
			{ItemId: 4003001, Amount: 2},
			{ItemId: 4010000, Amount: 4},
			{ItemId: 4010001, Amount: 4},
			{ItemId: 4010003, Amount: 3},
			{ItemId: 4010004, Amount: 3},
			{ItemId: 4020004, Amount: 4},
			{ItemId: 3010004, Amount: 1},
			{ItemId: 3010005, Amount: 1},
		},
	}
}

func (r Billy) Tier3() Tier {
	return Tier{
		Name: "Tier 3",
		Prizes: []Prize{
			{ItemId: 1302058, Amount: 1},
			{ItemId: 1372008, Amount: 1},
			{ItemId: 1422030, Amount: 1},
			{ItemId: 1422031, Amount: 1},
			{ItemId: 1022082, Amount: 1},
			{ItemId: 2022279, Amount: 65},
			{ItemId: 2022120, Amount: 40},
			{ItemId: 2001001, Amount: 40},
			{ItemId: 2001002, Amount: 40},
			{ItemId: 2022071, Amount: 25},
			{ItemId: 2022189, Amount: 25},
			{ItemId: 2040914, Amount: 1},
			{ItemId: 2041001, Amount: 1},
			{ItemId: 2041041, Amount: 1},
			{ItemId: 2041308, Amount: 1},
			{ItemId: 4031203, Amount: 10},
			{ItemId: 4000030, Amount: 7},
			{ItemId: 4003005, Amount: 10},
			{ItemId: 4003000, Amount: 8},
			{ItemId: 4010004, Amount: 5},
			{ItemId: 4010006, Amount: 5},
			{ItemId: 4020000, Amount: 5},
			{ItemId: 4020006, Amount: 5},
			{ItemId: 3010002, Amount: 1},
			{ItemId: 3010003, Amount: 1},
		},
	}
}

func (r Billy) Tier4() Tier {
	return Tier{
		Name: "Tier 4",
		Prizes: []Prize{
			{ItemId: 1332029, Amount: 1},
			{ItemId: 1472027, Amount: 1},
			{ItemId: 1462032, Amount: 1},
			{ItemId: 1492019, Amount: 1},
			{ItemId: 2022045, Amount: 45},
			{ItemId: 2022048, Amount: 40},
			{ItemId: 2022094, Amount: 25},
			{ItemId: 2022123, Amount: 20},
			{ItemId: 2022058, Amount: 60},
			{ItemId: 2041304, Amount: 1},
			{ItemId: 2041019, Amount: 1},
			{ItemId: 2040826, Amount: 1},
			{ItemId: 2040758, Amount: 1},
			{ItemId: 4000030, Amount: 10},
			{ItemId: 4003005, Amount: 10},
			{ItemId: 4003000, Amount: 20},
			{ItemId: 4010007, Amount: 5},
			{ItemId: 4011003, Amount: 1},
			{ItemId: 4021003, Amount: 1},
			{ItemId: 3010016, Amount: 1},
			{ItemId: 3010017, Amount: 1},
		},
	}
}

func (r Billy) Tier5() Tier {
	return Tier{
		Name: "Tier 5",
		Prizes: []Prize{
			{ItemId: 1382015, Amount: 1},
			{ItemId: 1382016, Amount: 1},
			{ItemId: 1442044, Amount: 1},
			{ItemId: 1382035, Amount: 1},
			{ItemId: 2022310, Amount: 20},
			{ItemId: 2022068, Amount: 40},
			{ItemId: 2022069, Amount: 40},
			{ItemId: 2022190, Amount: 30},
			{ItemId: 2022047, Amount: 30},
			{ItemId: 2040727, Amount: 1},
			{ItemId: 2040924, Amount: 1},
			{ItemId: 2040501, Amount: 1},
			{ItemId: 4000030, Amount: 20},
			{ItemId: 4003005, Amount: 20},
			{ItemId: 4003000, Amount: 25},
			{ItemId: 4011003, Amount: 3},
			{ItemId: 4011006, Amount: 2},
			{ItemId: 4021004, Amount: 3},
			{ItemId: 3010099, Amount: 1},
		},
	}
}

func (r Billy) Tier6() Tier {
	return Tier{
		Name: "Tier 6",
		Prizes: []Prize{
			{ItemId: 1442046, Amount: 1},
			{ItemId: 1432018, Amount: 1},
			{ItemId: 1102146, Amount: 1},
			{ItemId: 1102145, Amount: 1},
			{ItemId: 2022094, Amount: 35},
			{ItemId: 2022544, Amount: 15},
			{ItemId: 2022123, Amount: 20},
			{ItemId: 2022310, Amount: 20},
			{ItemId: 2040727, Amount: 1},
			{ItemId: 2041058, Amount: 1},
			{ItemId: 2040817, Amount: 1},
			{ItemId: 4000030, Amount: 30},
			{ItemId: 4003005, Amount: 30},
			{ItemId: 4003000, Amount: 30},
			{ItemId: 4011007, Amount: 1},
			{ItemId: 4021009, Amount: 1},
			{ItemId: 4011008, Amount: 3},
			{ItemId: 3010098, Amount: 1},
		},
	}
}

type Tier struct {
	Name   string
	Prizes []Prize
}

type Prize struct {
	ItemId uint32
	Amount uint32
}
