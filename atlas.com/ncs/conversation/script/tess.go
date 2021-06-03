package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
)

// Tess is located in Ellin Forest - Altaire Camp (300000000)
type Tess struct {
}

func (r Tess) NPCId() uint32 {
	return npc.Tess
}

func (r Tess) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These monsters are a piece of cake! One hit with my sword and I will kill them... better get a sword first.").NewLine().
		OpenItem(0).BlueText().AddText("Hey, take these tree trunks. You can build a better sword with them.").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r Tess) Selection(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasItems(l)(c.CharacterId, item.TreeTrunk, 100) {
			return r.NeedAtLeast100(l, c)
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
		return SendGetNumber(l, c, m.String(), r.ProcessQuantity, 1, productionMax, productionMax)
	}
}

func (r Tess) NeedAtLeast100(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You don't have enough... I need at least 100.")
	return SendOk(l, c, m.String())
}

func (r Tess) ProcessQuantity(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if selection <= 0 {
			m := message.NewBuilder().AddText("I cannot make less than zero items.")
			return SendOk(l, c, m.String())
		}
		count := character.ItemQuantity(l)(c.CharacterId, item.TreeTrunk)
		inventoryMax := int32(math.Floor(float64(count) / 100))
		if selection > inventoryMax {
			m := message.NewBuilder().AddText("You lack the resources to produce that many.")
			return SendOk(l, c, m.String())
		}
		if !character.CanHoldAll(l)(c.CharacterId, item.PerfectPitch, uint32(selection)) {
			m := message.NewBuilder().AddText("Please make some space in ETC tab.")
			return SendOk(l, c, m.String())
		}
		character.GainItem(l)(c.CharacterId, item.TreeTrunk, -selection*100)
		character.GainItem(l)(c.CharacterId, item.PerfectPitch, selection)
		m := message.NewBuilder().AddText("Thanks!")
		return SendOk(l, c, m.String())
	}
}