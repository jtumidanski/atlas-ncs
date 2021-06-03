package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Lila is located in The Burning Road - Ariant (260000000)
type Lila struct {
}

func (r Lila) NPCId() uint32 {
	return npc.Lila
}

func (r Lila) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Lila) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hohoh~ welcome welcome. Welcome to Ariant Skin Care. You have stepped into a renowned Skin Care shop that even the Queen herself frequents this place. If you have ").
		BlueText().ShowItemName1(item.AriantSkinCareCoupon).
		BlackText().AddText(" with you, we'll take care of the rest. How about letting work on your skin today?")
	return SendNext(l, c, m.String(), r.ChooseStyle)
}

func (r Lila) MissingCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmmm... I don't think you have our Skin Care coupon with you. Without it, I can't give you the treatment")
	return SendOk(l, c, m.String())
}

func (r Lila) ChooseStyle(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("With our specialized machine, you can see yourself after the treatment in advance. What kind of skin-treatment would you like to do? Choose the style of your liking...")
	return SendStyle(l, c, m.String(), r.ValidateSelection, []uint32{0, 1, 2, 3, 4})
}

func (r Lila) ValidateSelection(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasItem(l)(c.CharacterId, item.AriantSkinCareCoupon) {
			return r.MissingCoupon(l, c)
		}
		return r.ProcessSelection(l, c, selection)
	}
}

func (r Lila) ProcessSelection(l logrus.FieldLogger, c Context, selection int32) State {
	character.GainItem(l)(c.CharacterId, item.AriantSkinCareCoupon, -1)
	character.SetSkin(l)(c.CharacterId, byte(selection))
	return r.Enjoy(l, c)
}

func (r Lila) Enjoy(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Enjoy your new and improved skin!")
	return SendOk(l, c, m.String())
}
