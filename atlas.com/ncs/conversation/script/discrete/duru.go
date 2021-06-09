package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
)

// Duru is located in Ellin Forest - Altaire Camp (300000000)
type Duru struct {
}

func (r Duru) NPCId() uint32 {
	return npc.Duru
}

func (r Duru) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The monsters are advancing.. I can't fight. I was badly injured by the Primitive Boars...").NewLine().
		OpenItem(0).BlueText().AddText("Hey, take these boar hides. You can recover from them.").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Duru) Selection(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if !character.HasItems(l)(c.CharacterId, item.ToughLeather, 100) {
			return r.NeedAtLeast100(l, c)
		}
		count := character.ItemQuantity(l)(c.CharacterId, item.ToughLeather)
		inventoryMax := int32(math.Floor(float64(count) / 100))
		productionMax := int32(math.Min(float64(inventoryMax), 300))
		m := message.NewBuilder().
			AddText("Hey, that's a good idea! I can give you ").
			ShowItemImage2(item.PerfectPitch).ShowItemName1(item.PerfectPitch).
			AddText(" for each 100 ").
			ShowItemImage2(item.ToughLeather).ShowItemName1(item.ToughLeather).
			AddText(" you give me. How many do you want? (Current Items: ").
			AddText(fmt.Sprintf("%d", count)).
			AddText(")")
		return script.SendGetNumber(l, c, m.String(), r.ProcessQuantity, 1, productionMax, productionMax)
	}
}

func (r Duru) NeedAtLeast100(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't have enough... I need at least 100.")
	return script.SendOk(l, c, m.String())
}

func (r Duru) ProcessQuantity(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if selection <= 0 {
			m := message.NewBuilder().AddText("I cannot make less than zero items.")
			return script.SendOk(l, c, m.String())
		}
		count := character.ItemQuantity(l)(c.CharacterId, item.ToughLeather)
		inventoryMax := int32(math.Floor(float64(count) / 100))
		if selection > inventoryMax {
			m := message.NewBuilder().AddText("You lack the resources to produce that many.")
			return script.SendOk(l, c, m.String())
		}
		if !character.CanHoldAll(l)(c.CharacterId, item.PerfectPitch, uint32(selection)) {
			m := message.NewBuilder().AddText("Please make some space in ETC tab.")
			return script.SendOk(l, c, m.String())
		}
		character.GainItem(l)(c.CharacterId, item.ToughLeather, -selection*100)
		character.GainItem(l)(c.CharacterId, item.PerfectPitch, selection)
		m := message.NewBuilder().AddText("Thanks!")
		return script.SendOk(l, c, m.String())
	}
}