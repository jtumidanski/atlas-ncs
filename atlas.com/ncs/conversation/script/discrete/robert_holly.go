package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// RobertHolly is located in Ludibrium - Ludibrium (220000000)
type RobertHolly struct {
}

func (r RobertHolly) NPCId() uint32 {
	return npc.RobertHolly
}

func (r RobertHolly) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I hope I can make as much as yesterday... well, hello! Don't you want to extend your buddy list? You look like someone who'd have a whole lot of friends... well, what do you think? With some money I can make it happen for you. Remember, though, it only applies to one character at a time, so it won't affect any of your other characters on your account. Do you want to extend your buddy list?")
	return script.SendYesNo(l, span, c, m.String(), r.Alright, r.ISee)
}

func (r RobertHolly) ISee(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I see... you don't have as many friends as I thought you would. Hahaha, just kidding! Anyway if you feel like changing your mind, please feel free to come back and we'll talk business. If you make a lot of friends, then you know ... hehe ...")
	return script.SendOk(l, span, c, m.String())
}

func (r RobertHolly) Alright(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright, good call! It's not that expensive actually. ").
		BlueText().AddText("240,000 mesos and I'll add 5 more slots to your buddy list").
		BlackText().AddText(". And no, I won't be selling them individually. Once you buy it, it's going to be permanently on your buddy list. So if you're one of those that needs more space there, then you might as well do it. What do you think? Will you spend 240,000 mesos for it?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, r.NotAsManyFriends)
}

func (r RobertHolly) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.BuddyCapacity(l)(c.CharacterId) >= 50 || !character.HasMeso(l, span)(c.CharacterId, 240000) {
		return r.AreYouSure(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r RobertHolly) AreYouSure(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey... are you sure you have ").
		BlueText().AddText("240,000 mesos").
		BlackText().AddText("? If so, then check and see if you have extended your buddy list to the max. Even if you pay up, the most you can have on your buddy list is ").
		BlueText().AddText("50").
		BlackText().AddText(".")
	return script.SendOk(l, span, c, m.String())
}

func (r RobertHolly) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -240000)
	err := character.IncreaseBuddyCapacity(l)(c.CharacterId, 5)
	if err != nil {
		l.WithError(err).Errorf("Error increasing buddy capacity for character %d.", c.CharacterId)
	}

	m := message.NewBuilder().
		AddText("Alright! Your buddy list will have 5 extra slots by now. Check and see for it yourself. And if you still need more room on your buddy list, you know who to find. Of course, it isn't going to be for free ... well, so long ...")
	return script.SendOk(l, span, c, m.String())
}

func (r RobertHolly) NotAsManyFriends(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I see... I don't think you don't have as many friends as I thought you would. If not, you just don't have 240,000 mesos with you right this minute? Anyway, if you ever change your mind, come back and we'll talk business. That is, of course, once you have get some financial relief. .. hehe ...")
	return script.SendOk(l, span, c, m.String())
}
