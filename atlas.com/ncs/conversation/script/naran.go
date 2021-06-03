package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Naran is located in Mu Lung - Mu Lung (250000000)
type Naran struct {
}

func (r Naran) NPCId() uint32 {
	return npc.Naran
}

func (r Naran) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Naran) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Well, hello! Welcome to the Mu Lung Skin-Care! Would you like to have a firm, tight, healthy looking skin like mine?  With a ").
		BlueText().ShowItemName1(item.MuLungSkinCareCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").NewLine().
		OpenItem(0).AddText("Skin Care: ").ShowItemImage2(item.MuLungSkinCareCoupon).ShowItemName1(item.MuLungSkinCareCoupon).CloseItem()
	return SendListSelection(l, c, m.String(), r.Choices)
}

func (r Naran) Choices(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.SkinCoupon
	}
	return nil
}

func (r Naran) SkinCoupon(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.MuLungSkinCareCoupon) {
		return r.MissingCoupon(l, c)
	}

	return r.ChooseStyle(l, c)
}

func (r Naran) MissingCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Um... you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you...")
	return SendOk(l, c, m.String())
}

func (r Naran) ChooseStyle(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!")
	return SendStyle(l, c, m.String(), r.ProcessSelection, []uint32{0, 1, 2, 3, 4})
}

func (r Naran) ProcessSelection(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.GainItem(l)(c.CharacterId, item.MuLungSkinCareCoupon, -1)
		character.SetSkin(l)(c.CharacterId, byte(selection))
		return r.Enjoy(l, c)
	}
}

func (r Naran) Enjoy(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Enjoy your new and improved skin!")
	return SendOk(l, c, m.String())
}
