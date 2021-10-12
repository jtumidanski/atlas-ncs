package care

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Hello   string
	Choices []ChoiceConfig
}

func NewChoiceConfig(nextState ChoiceStateProducer, configurators ...ChoiceConfigurator) ChoiceConfig {
	choiceConfig := &ChoiceConfig{
		ListText:      "",
		NextState:     nextState,
		MissingCoupon: "",
		Enjoy:         "",
	}
	for _, configurator := range configurators {
		configurator(choiceConfig)
	}
	return *choiceConfig
}

func SetListText(value string) ChoiceConfigurator {
	return func(c *ChoiceConfig) {
		c.ListText = value
	}
}

func SetEnjoy(value string) ChoiceConfigurator {
	return func(c *ChoiceConfig) {
		c.Enjoy = value
	}
}

func SetMissingCoupon(value string) ChoiceConfigurator {
	return func(c *ChoiceConfig) {
		c.MissingCoupon = value
	}
}

type ChoiceConfig struct {
	ListText      string
	NextState     ChoiceStateProducer
	MissingCoupon string
	Enjoy         string
}

type Configurator func(c *Config)

type ChoiceConfigurator func(c *ChoiceConfig)

// NewGenericCare Creates a configurable choice generic care script.
func NewGenericCare(hello string, choices []ChoiceConfig, configurators ...Configurator) script.StateProducer {
	config := &Config{
		Hello:   hello,
		Choices: choices,
	}
	for _, configurator := range configurators {
		configurator(config)
	}
	return Initial(config)
}

// NewGenericSkinCare Creates a single choice generic skin care script.
func NewGenericSkinCare(coupon uint32, hello string, configurators ...Configurator) script.StateProducer {
	return NewGenericCare(hello, []ChoiceConfig{SkinCareChoice(coupon)}, configurators...)
}

// StateProducer Produces script.StateProducer which leverages the provided care Config
type StateProducer func(config *Config) script.StateProducer

// ChoiceStateProducer Produces script.StateProduce which leverages the provided ChoiceConfig
type ChoiceStateProducer func(config ChoiceConfig) script.StateProducer

// The Initial state in the generic care providers state machine. This will say hello, and present character with configured
// care options.
func Initial(config *Config) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText(config.Hello).NewLine()
		for i, choice := range config.Choices {
			m = m.OpenItem(i).BlueText().AddText(choice.ListText).CloseItem().NewLine()
		}
		return script.SendListSelection(l, span, c, m.String(), Selection(config))
	}
}

// Selection processes the initial care selection input by the character. If it is a valid selection, it will execute the
// configured state.
func Selection(config *Config) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		if selection < 0 || int(selection) >= len(config.Choices) {
			return script.Exit()
		}
		choiceConfig := config.Choices[selection]
		return choiceConfig.NextState(choiceConfig)
	}
}

// ChoicesSupplier supplies the choices to be considered when providing care.
type ChoicesSupplier func(l logrus.FieldLogger, span opentracing.Span, c script.Context) []uint32

// FixedChoices supplies the choices as a fixed slice.
func FixedChoices(values []uint32) ChoicesSupplier {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) []uint32 {
		return values
	}
}

// ChoiceSupplier supplies the choice of care the character will have applied to them.
type ChoiceSupplier func() uint32

// FixedChoice supplies the choice as a fixed value
func FixedChoice(value uint32) ChoiceSupplier {
	return func() uint32 {
		return value
	}
}

// ChoiceConsumer uses the script state and care choice to perform an action.
type ChoiceConsumer func(l logrus.FieldLogger, c script.Context, choice uint32)

// ChoiceHandlerProducer produces a ChoiceStateProducer which acts upon a ChoiceSupplier.
type ChoiceHandlerProducer func(s ChoiceSupplier) ChoiceStateProducer

func ShowChoices(prompt string, s ChoicesSupplier, next ChoiceHandlerProducer) ChoiceStateProducer {
	return func(config ChoiceConfig) script.StateProducer {
		return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
			choices := s(l, span, c)
			if len(choices) == 0 {
				l.Errorf("Zero choices available for care.")
				return script.SendOk(l, span, c, message.NewBuilder().AddText("No styles available.").String())
			}
			return script.SendStyle(l, span, c, prompt, ChoiceSelection(config)(choices, next), choices)
		}
	}
}

func ShowChoicesWithNone(prompt string, s ChoicesSupplier, next ChoiceHandlerProducer) ChoiceStateProducer {
	return func(config ChoiceConfig) script.StateProducer {
		return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
			choices := s(l, span, c)
			if len(choices) == 0 {
				m := message.NewBuilder().AddText(config.MissingCoupon)
				return script.SendOk(l, span, c, m.String())
			}
			return script.SendStyle(l, span, c, prompt, ChoiceSelection(config)(choices, next), choices)
		}
	}
}

// ChoiceSelection processes the care selection made by the character. If it is valid, calls the ChoiceHandlerProducer provided to act upon it.
func ChoiceSelection(config ChoiceConfig) func(choices []uint32, next ChoiceHandlerProducer) script.ProcessSelection {
	return func(choices []uint32, next ChoiceHandlerProducer) script.ProcessSelection {
		return func(selection int32) script.StateProducer {
			if selection < 0 || int(selection) >= len(choices) {
				return script.Exit()
			}
			choice := choices[selection]
			return next(FixedChoice(choice))(config)
		}
	}
}

// CouponConfigurator supports configuring CouponConfig options
type CouponConfigurator func(c *CouponConfig)

// CouponConfig allows different interactions based on the type of coupon
type CouponConfig struct {
	singleUse bool
	fail      ChoiceHandlerProducer
}

// SetSingleUse dictates whether the coupon is considered a single use coupon
func SetSingleUse(value bool) CouponConfigurator {
	return func(c *CouponConfig) {
		c.singleUse = value
	}
}

// SetFailFunction dictates the next ChoiceHandlerProducer to act on if the existing one fails.
func SetFailFunction(next ChoiceHandlerProducer) CouponConfigurator {
	return func(c *CouponConfig) {
		c.fail = next
	}
}

// ProcessCoupon acts upon the provided coupon. If the character possesses the coupon provided, will act upon the choice made by the character.
func ProcessCoupon(coupon uint32, consume ChoiceConsumer, configurators ...CouponConfigurator) ChoiceHandlerProducer {
	config := &CouponConfig{singleUse: true, fail: MissingCoupon}
	for _, configurator := range configurators {
		configurator(config)
	}

	return func(s ChoiceSupplier) ChoiceStateProducer {
		return func(careConfig ChoiceConfig) script.StateProducer {
			return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
				if character.HasItem(l, span)(c.CharacterId, coupon) {
					if config.singleUse {
						character.GainItem(l, span)(c.CharacterId, coupon, -1)
					}
					consume(l, c, s())
					return Enjoy(careConfig)(l, span, c)
				}
				return config.fail(s)(careConfig)(l, span, c)
			}
		}
	}
}

// Enjoy is a terminal state which tells the character they should enjoy care provided.
func Enjoy(config ChoiceConfig) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText(config.Enjoy)
		return script.SendOk(l, span, c, m.String())
	}
}

// MissingCoupon is a terminal state which tells the character they are missing the coupon needed for care.
func MissingCoupon(_ ChoiceSupplier) ChoiceStateProducer {
	return func(config ChoiceConfig) script.StateProducer {
		return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
			m := message.NewBuilder().AddText(config.MissingCoupon)
			return script.SendOk(l, span, c, m.String())
		}
	}
}
