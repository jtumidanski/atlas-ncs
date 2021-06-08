package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MsTan is located in Victoria Road - Henesys Skin-Care (100000105)
type MsTan struct {
}

func (r MsTan) NPCId() uint32 {
	return npc.MsTan
}

func (r MsTan) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r MsTan) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Well, hello! Welcome to the Henesys Skin-Care! Would you like to have a firm, tight, healthy looking skin like mine?  With a ").
		BlueText().ShowItemName1(item.HenesysSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").NewLine().
		OpenItem(0).AddText("Skin Care: ").ShowItemImage2(item.HenesysSkinCoupon).ShowItemName1(item.HenesysSkinCoupon).CloseItem()
	return SendListSelection(l, c, m.String(), r.Choices)
}

func (r MsTan) Choices(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.SkinCoupon
	}
	return nil
}

func (r MsTan) SkinCoupon(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.HenesysSkinCoupon) {
		return r.MissingCoupon(l, c)
	}

	return r.ChooseStyle(l, c)
}

func (r MsTan) MissingCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Um... you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you...")
	return SendOk(l, c, m.String())
}

func (r MsTan) ChooseStyle(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("With our specialized machine, you can see yourself after the treatment in advance. What kind of skin-treatment would you like to do? Choose the style of your liking.")
	return SendStyle(l, c, m.String(), r.ProcessSelection, []uint32{0, 1, 2, 3, 4})
}

func (r MsTan) ProcessSelection(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.GainItem(l)(c.CharacterId, item.HenesysSkinCoupon, -1)
		character.SetSkin(l)(c.CharacterId, byte(selection))
		return r.Enjoy(l, c)
	}
}

func (r MsTan) Enjoy(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Enjoy your new and improved skin!")
	return SendOk(l, c, m.String())
}
