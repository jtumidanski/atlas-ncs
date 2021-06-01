package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Romi is located in Orbis Park - Orbis Skin-Care (200000203)
type Romi struct {
}

func (r Romi) NPCId() uint32 {
	return npc.Romi
}

func (r Romi) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Well, hello! Welcome to the Orbis Skin-Care~! Would you like to have a firm, tight, healthy looking skin like mine?  With ").
		BlueText().ShowItemName1(item.OrbisSkinCoupon).
		BlackText().AddText(", you can let us take care of the rest and have the kind of skin you've always wanted~!").NewLine().
		OpenItem(0).AddText("Skin Care: ").ShowItemImage2(item.OrbisSkinCoupon).ShowItemName1(item.OrbisSkinCoupon).CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r Romi) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.SendSkin
	}
	return nil
}

func (r Romi) SendSkin(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!")
	return SendStyle(l, c, m.String(), r.Validate, []uint32{0, 1, 2, 3, 4})
}

func (r Romi) Validate(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasItem(l)(c.CharacterId, item.OrbisSkinCoupon) {
			return r.MissingCoupon(l, c)
		}
		return r.Process(selection)(l, c)
	}
}

func (r Romi) MissingCoupon(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Um...you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you...")
	return SendOk(l, c, m.String())
}

func (r Romi) Process(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.GainItem(l)(c.CharacterId, item.OrbisSkinCoupon, -1)
		character.SetSkin(l)(c.CharacterId, byte(selection+1))
		return r.Success(l, c)
	}
}

func (r Romi) Success(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Enjoy your new and improved skin!")
	return SendOk(l, c, m.String())
}
