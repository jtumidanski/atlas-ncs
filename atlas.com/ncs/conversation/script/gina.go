package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Gina is located in Ludibrium - Ludibrium Skin Care (220000005)
type Gina struct {
}

func (r Gina) NPCId() uint32 {
	return npc.Gina
}

func (r Gina) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Gina) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh, hello! Welcome to the Ludibrium Skin-Care! Are you interested in getting tanned and looking sexy? How about a beautiful, snow-white skin? If you have ").
		BlueText().ShowItemName1(item.LudibriumSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always dreamed of!").NewLine().
		OpenItem(0).AddText("Skin Care: ").ShowItemImage2(item.LudibriumSkinCoupon).ShowItemName1(item.LudibriumSkinCoupon).CloseItem()
	return SendListSelection(l, c, m.String(), r.Choices)
}

func (r Gina) Choices(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.SkinCoupon
	}
	return nil
}

func (r Gina) SkinCoupon(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.LudibriumSkinCoupon) {
		return r.MissingCoupon(l, c)
	}

	return r.ChooseStyle(l, c)
}

func (r Gina) MissingCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Um... you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you...")
	return SendOk(l, c, m.String())
}

func (r Gina) ChooseStyle(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!")
	return SendStyle(l, c, m.String(), r.ProcessSelection, []uint32{0, 1, 2, 3, 4})
}

func (r Gina) ProcessSelection(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.GainItem(l)(c.CharacterId, item.LudibriumSkinCoupon, -1)
		character.SetSkin(l)(c.CharacterId, byte(selection))
		return r.Enjoy(l, c)
	}
}

func (r Gina) Enjoy(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Enjoy your new and improved skin!")
	return SendOk(l, c, m.String())
}
