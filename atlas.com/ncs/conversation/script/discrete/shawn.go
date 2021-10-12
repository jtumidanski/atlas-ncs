package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Shawn is located in Victoria Road - Excavation Site <Camp> (101030104)
type Shawn struct {
}

func (r Shawn) NPCId() uint32 {
	return npc.Shawn
}

func (r Shawn) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("We, the Union of Guilds, have been trying to decipher 'Emerald Tablet,' a treasured old relic, for a long time. As a result, we have found out that Sharenian, the mysterious country from the past, lay asleep here. We also found out that clues of ").
		ShowItemName1(item.Rubian).
		AddText(", a legendary, mythical jewelry, may be here at the remains of Sharenian. This is why the Union of Guilds have opened Guild Quest to ultimately find ").
		ShowItemName1(item.Rubian).
		AddText(".").NewLine().
		OpenItem(0).BlueText().AddText("What's Sharenian?").CloseItem().NewLine().
		OpenItem(1).BlueText().ShowItemName1(item.Rubian).AddText("? What's that?").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Guild Quest?").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("No, I'm fine now.").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r Shawn) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.LiterateCivilization
	case 1:
		return r.LegendaryJewel
	case 2:
		return r.Before
	case 3:
		return r.Exit
	}
	return nil
}

func (r Shawn) LiterateCivilization(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Sharenian was a literate civilization from the past that had control over every area of the Victoria Island. The Temple of Golem, the Shrine in the deep part of the Dungeon, and other old architectural constructions where no one knows who built it are indeed made during the Sharenian times.")
	return script.SendNext(l, span, c, m.String(), r.LastKing)
}

func (r Shawn) LastKing(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The last king of Sharenian was a gentleman named Sharen III, and apparently he was a very wise and compassionate king. But one day, the whole kingdom collapsed, and there was no explanation made for it.")
	return script.SendNextPrevious(l, span, c, m.String(), r.AnyMore, r.LiterateCivilization)
}

func (r Shawn) AnyMore(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you have any other questions?").NewLine().
		OpenItem(0).BlueText().AddText("What's Sharenian?").CloseItem().NewLine().
		OpenItem(1).BlueText().ShowItemName1(item.Rubian).AddText("? What's that?").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Guild Quest?").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("No, I'm fine now.").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r Shawn) LegendaryJewel(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		ShowItemName1(item.Rubian).
		AddText(" is a legendary jewel that brings eternal youth to the one that possesses it. Ironically, it seems like everyone that had ").
		ShowItemName1(item.Rubian).
		AddText(" ended up downtrodden, which should explain the downfall of Sharenian.")
	return script.SendNext(l, span, c, m.String(), r.UltimateGoal)
}

func (r Shawn) UltimateGoal(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The ultimate goal of this Guild Quest is to explore Sharenian and find ").
		ShowItemName1(item.Rubian).
		AddText(". This is not a task where power solves everything. Teamwork is more important here.")
	return script.SendNextPrevious(l, span, c, m.String(), r.AnyMore, r.LegendaryJewel)
}

func (r Shawn) Before(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I've sent groups of explorers to Sharenian before, but none of them ever came back, which prompted us to start the Guild Quest. We've been waiting for guilds that are strong enough to take on tough challenges, guilds like yours.")
	return script.SendNext(l, span, c, m.String(), r.AnyMore)
}

func (r Shawn) Exit(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Really? If you have anything else to ask, please feel free to talk to me.")
	return script.SendOk(l, span, c, m.String())
}
