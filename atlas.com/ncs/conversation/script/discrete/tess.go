package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math"
)

// Tess is located in Ellin Forest - Altaire Camp (300000000)
type Tess struct {
}

func (r Tess) NPCId() uint32 {
	return npc.Tess
}

func (r Tess) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("These monsters are a piece of cake! One hit with my sword and I will kill them... better get a sword first.").NewLine().
		OpenItem(0).BlueText().AddText("Hey, take these tree trunks. You can build a better sword with them.").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r Tess) Selection(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasItems(l, span)(c.CharacterId, item.TreeTrunk, 100) {
			return r.NeedAtLeast100(l, span, c)
		}
		count := character.ItemQuantity(l)(c.CharacterId, item.TreeTrunk)
		inventoryMax := int32(math.Floor(float64(count) / 100))
		productionMax := int32(math.Min(float64(inventoryMax), 300))
		m := message.NewBuilder().
			AddText("Hey, that's a good idea! I can give you ").
			ShowItemImage2(item.PerfectPitch).ShowItemName1(item.PerfectPitch).
			AddText(" for each 100 ").
			ShowItemImage2(item.TreeTrunk).ShowItemName1(item.TreeTrunk).
			AddText(" you give me. How many do you want? (Current Items: ").
			AddText(fmt.Sprintf("%d", count)).
			AddText(")")
		return script.SendGetNumber(l, span, c, m.String(), r.ProcessQuantity, 1, productionMax, productionMax)
	}
}

func (r Tess) NeedAtLeast100(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't have enough... I need at least 100.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tess) ProcessQuantity(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if selection <= 0 {
			m := message.NewBuilder().AddText("I cannot make less than zero items.")
			return script.SendOk(l, span, c, m.String())
		}
		count := character.ItemQuantity(l)(c.CharacterId, item.TreeTrunk)
		inventoryMax := int32(math.Floor(float64(count) / 100))
		if selection > inventoryMax {
			m := message.NewBuilder().AddText("You lack the resources to produce that many.")
			return script.SendOk(l, span, c, m.String())
		}
		if !character.CanHoldAll(l)(c.CharacterId, item.PerfectPitch, uint32(selection)) {
			m := message.NewBuilder().AddText("Please make some space in ETC tab.")
			return script.SendOk(l, span, c, m.String())
		}
		character.GainItem(l, span)(c.CharacterId, item.TreeTrunk, -selection*100)
		character.GainItem(l, span)(c.CharacterId, item.PerfectPitch, selection)
		m := message.NewBuilder().AddText("Thanks!")
		return script.SendOk(l, span, c, m.String())
	}
}