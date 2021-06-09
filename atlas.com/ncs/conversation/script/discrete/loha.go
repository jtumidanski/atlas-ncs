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

// Loha is located in Ellin Forest - Altaire Camp (300000000)
type Loha struct {
}

func (r Loha) NPCId() uint32 {
	return npc.Loha
}

func (r Loha) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So many injuries, so little medicine...").NewLine().
		OpenItem(0).BlueText().AddText("Hey, take these black spores. You can make better medicine with them.").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Loha) Selection(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if !character.HasItems(l)(c.CharacterId, item.BlackMushroomSpore, 100) {
			return r.NeedAtLeast100(l, c)
		}
		count := character.ItemQuantity(l)(c.CharacterId, item.BlackMushroomSpore)
		inventoryMax := int32(math.Floor(float64(count) / 100))
		productionMax := int32(math.Min(float64(inventoryMax), 300))
		m := message.NewBuilder().
			AddText("Hey, that's a good idea! I can give you ").
			ShowItemImage2(item.PerfectPitch).ShowItemName1(item.PerfectPitch).
			AddText(" for each 100 ").
			ShowItemImage2(item.BlackMushroomSpore).ShowItemName1(item.BlackMushroomSpore).
			AddText(" you give me. How many do you want? (Current Items: ").
			AddText(fmt.Sprintf("%d", count)).
			AddText(")")
		return script.SendGetNumber(l, c, m.String(), r.ProcessQuantity, 1, productionMax, productionMax)
	}
}

func (r Loha) NeedAtLeast100(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You don't have enough... I need at least 100.")
	return script.SendOk(l, c, m.String())
}

func (r Loha) ProcessQuantity(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if selection <= 0 {
			m := message.NewBuilder().AddText("I cannot make less than zero items.")
			return script.SendOk(l, c, m.String())
		}
		count := character.ItemQuantity(l)(c.CharacterId, item.BlackMushroomSpore)
		inventoryMax := int32(math.Floor(float64(count) / 100))
		if selection > inventoryMax {
			m := message.NewBuilder().AddText("You lack the resources to produce that many.")
			return script.SendOk(l, c, m.String())
		}
		if !character.CanHoldAll(l)(c.CharacterId, item.PerfectPitch, uint32(selection)) {
			m := message.NewBuilder().AddText("Please make some space in ETC tab.")
			return script.SendOk(l, c, m.String())
		}
		character.GainItem(l)(c.CharacterId, item.BlackMushroomSpore, -selection*100)
		character.GainItem(l)(c.CharacterId, item.PerfectPitch, selection)
		m := message.NewBuilder().AddText("Thanks!")
		return script.SendOk(l, c, m.String())
	}
}