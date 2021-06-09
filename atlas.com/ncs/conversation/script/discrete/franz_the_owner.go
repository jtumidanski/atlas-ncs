package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// FranzTheOwner is located in Orbis Park - Orbis Plastic Surgery (200000201)
type FranzTheOwner struct {
}

func (r FranzTheOwner) NPCId() uint32 {
	return npc.FranzTheOwner
}

func (r FranzTheOwner) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.Hello(l, c)
}

func (r FranzTheOwner) Hello(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Well well well, welcome to the Orbis Plastic Surgery! Would you like to transform your face into something new? With a ").
		BlueText().ShowItemName1(item.OrbisFaceCouponVIP).
		BlackText().AddText(", you can let us take care of the rest and have the face you've always wanted~!").NewLine().
		OpenItem(0).AddText("Plastic Surgery: ").ShowItemImage2(item.OrbisFaceCouponVIP).ShowItemName1(item.OrbisFaceCouponVIP).CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r FranzTheOwner) Selection(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		faces := make([]uint32, 0)
		gender := character.GetGender(l)(c.CharacterId)
		if gender == character.GenderMale {
			faces = r.BaseMaleFaces()
		} else if gender == character.GenderFemale {
			faces = r.BaseFemaleFaces()
		}
		faces = ApplyEyeColor(l)(c.CharacterId, faces)
		faces = FilterCurrentFace(l)(c.CharacterId, faces)

		m := message.NewBuilder().
			AddText("I can totally transform your face into something new... how about giving us a try? For ").
			BlueText().ShowItemName1(item.OrbisFaceCouponVIP).
			BlackText().AddText(", you can get the face of your liking...take your time in choosing the face of your preference.")
		return script.SendStyle(l, c, m.String(), r.ProcessSelection(faces), faces)
	}
}

func ApplyEyeColor(l logrus.FieldLogger) func(characterId uint32, faces []uint32) []uint32 {
	return func(characterId uint32, faces []uint32) []uint32 {
		face := character.GetFace(l)(characterId)
		color := face%1000 - face%100
		results := make([]uint32, 0)
		for _, f := range faces {
			results = append(results, f+color)
		}
		return results
	}
}

func FilterCurrentFace(l logrus.FieldLogger) func(characterId uint32, faces []uint32) []uint32 {
	return func(characterId uint32, faces []uint32) []uint32 {
		current := character.GetFace(l)(characterId)
		results := make([]uint32, 0)
		for _, f := range faces {
			if f != current {
				results = append(results, f)
			}
		}
		return results
	}
}

func (r FranzTheOwner) BaseMaleFaces() []uint32 {
	return []uint32{20000, 20001, 20003, 20004, 20005, 20006, 20007, 20008, 20012, 20014, 20022, 20028, 20031}
}

func (r FranzTheOwner) BaseFemaleFaces() []uint32 {
	return []uint32{21000, 21001, 21002, 21003, 21004, 21005, 21006, 21007, 21008, 21012, 21014, 21023, 21026}
}

func (r FranzTheOwner) ProcessSelection(options []uint32) script.ProcessSelection {
	return func(selection int32) script.StateProducer {
		return func(l logrus.FieldLogger, c script.Context) script.State {
			if !character.HasItem(l)(c.CharacterId, item.OrbisFaceCouponVIP) {
				return r.MissingCoupon(l, c)
			}
			return r.PerformChange(options[selection])(l, c)
		}
	}
}

func (r FranzTheOwner) MissingCoupon(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmm ... it looks like you don't have the coupon specifically for this place. Sorry to say this, but without the coupon, there's no plastic surgery for you...")
	return script.SendOk(l, c, m.String())
}

func (r FranzTheOwner) PerformChange(choice uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		character.GainItem(l)(c.CharacterId, item.OrbisFaceCouponVIP, -1)
		character.SetFace(l)(c.CharacterId, choice)
		return r.Enjoy(l, c)
	}
}

func (r FranzTheOwner) Enjoy(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Enjoy your new and improved face!")
	return script.SendOk(l, c, m.String())
}
