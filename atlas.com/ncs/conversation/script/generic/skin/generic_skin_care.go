package skin

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type CareConfig struct {
	coupon        uint32
	hello         string
	missingCoupon string
	chooseStyle   string
	enjoy         string
}

type CareConfigurator func(c *CareConfig)

func CareConfiguratorMissingCoupon(value string) CareConfigurator {
	return func(c *CareConfig) {
		c.missingCoupon = value
	}
}

func NewGenericCare(coupon uint32, hello string, configurators ...CareConfigurator) script.StateProducer {
	config := &CareConfig{
		coupon:        coupon,
		hello:         hello,
		missingCoupon: "Um... you don't have the skin-care coupon you need to receive the treatment. Sorry, but I am afraid we can't do it for you...",
		chooseStyle:   "With our specialized machine, you can see the way you'll look after the treatment PRIOR to the procedure. What kind of a look are you looking for? Go ahead and choose the style of your liking~!",
		enjoy:         "Enjoy your new and improved skin!",
	}
	for _, configurator := range configurators {
		configurator(config)
	}
	return GenericCare{}.Hello(config)
}

type GenericCare struct {
}

func (r GenericCare) Hello(config *CareConfig) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText(config.hello).NewLine().
			OpenItem(0).AddText("Skin Care: ").ShowItemImage2(config.coupon).ShowItemName1(config.coupon).CloseItem()
		return script.SendListSelection(l, c, m.String(), r.Choices(config))
	}
}

func (r GenericCare) Choices(config *CareConfig) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		switch selection {
		case 0:
			return r.SkinCoupon(config)
		}
		return nil
	}
}

func (r GenericCare) SkinCoupon(config *CareConfig) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if !character.HasItem(l)(c.CharacterId, config.coupon) {
			return r.MissingCoupon(config)(l, c)
		}

		return r.ChooseStyle(config)(l, c)
	}
}

func (r GenericCare) MissingCoupon(config *CareConfig) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText(config.missingCoupon)
		return script.SendOk(l, c, m.String())
	}
}

func (r GenericCare) ChooseStyle(config *CareConfig) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText(config.chooseStyle)
		return script.SendStyle(l, c, m.String(), r.ProcessSelection(config), []uint32{0, 1, 2, 3, 4})
	}
}

func (r GenericCare) ProcessSelection(config *CareConfig) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		return func(l logrus.FieldLogger, c script.Context) script.State {
			character.GainItem(l)(c.CharacterId, config.coupon, -1)
			character.SetSkin(l)(c.CharacterId, byte(selection))
			return r.Enjoy(config)(l, c)
		}
	}
}

func (r GenericCare) Enjoy(config *CareConfig) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().AddText(config.enjoy)
		return script.SendOk(l, c, m.String())
	}
}
