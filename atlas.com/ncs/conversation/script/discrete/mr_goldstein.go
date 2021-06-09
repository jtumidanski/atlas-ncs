package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MrGoldstein is located in Victoria Road - Lith Harbor (104000000)
type MrGoldstein struct {
}

func (r MrGoldstein) NPCId() uint32 {
	return npc.MrGoldstein
}

func (r MrGoldstein) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.ExtendBuddyList(l, c)
}

func (r MrGoldstein) ExtendBuddyList(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I hope I can make as much as yesterday... well, hello! Don't you want to extend your buddy list? You look like someone who'd have a whole lot of friends... well, what do you think? With some money I can make it happen for you. Remember, though, it only applies to one character at a time, so it won't affect any of your other characters on your account. Do you want to extend your buddy list?")
	return script.SendYesNo(l, c, m.String(), r.OkGood, r.QuickExit)
}

func (r MrGoldstein) QuickExit(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I see... you don't have as many friends as I thought you would. Hahaha, just kidding! Anyway if you feel like changing your mind, please feel free to come back and we'll talk business. If you make a lot of friends, then you know ... hehe ...")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r MrGoldstein) OkGood(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright, good call! It's not that expensive actually. ").
		BlueText().AddText("240,000 mesos and I'll add 5 more slots to your buddy list").
		BlackText().AddText(". And no, I won't be selling them individually. Once you buy it, it's going to be permanently on your buddy list. So if you're one of those that needs more space there, then you might as well do it. What do you think? Will you spend 240,000 mesos for it?")
	return script.SendYesNo(l, c, m.String(), r.ConfirmPayment, r.Exit)
}

func (r MrGoldstein) Exit(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I see... I don't think you don't have as many friends as I thought you would. If not, you just don't have 240,000 mesos with you right this minute? Anyway, if you ever change your mind, come back and we'll talk business. That is, of course, once you have get some financial relief. .. hehe ...")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r MrGoldstein) ConfirmPayment(l logrus.FieldLogger, c script.Context) script.State {
	if !character.HasMeso(l)(c.CharacterId, 240000) || !(character.BuddyCapacity(l)(c.CharacterId) >= 50) {
		return r.AreYouSure(l, c)
	}
	return r.ProcessPayment(l, c)
}

func (r MrGoldstein) AreYouSure(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey... are you sure you have ").
		BlueText().AddText("240,000 mesos").
		BlackText().AddText("? If so, then check and see if you have extended your buddy list to the max. Even if you pay up, the most you can have on your buddy list is ").
		BlueText().AddText("50").
		BlackText().AddText(".")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r MrGoldstein) ProcessPayment(l logrus.FieldLogger, c script.Context) script.State {
	err := character.GainMeso(l)(c.CharacterId, -240000)
	if err != nil {
		l.WithError(err).Errorf("Unable to take payment by character %d for buddy list increase.", c.CharacterId)
		return r.Error(l, c)
	}
	err = character.IncreaseBuddyCapacity(l)(c.CharacterId, 5)
	if err != nil {
		l.WithError(err).Errorf("Unable to increase buddy capacity of character %d, refunding payment.", c.CharacterId)
		err = character.GainMeso(l)(c.CharacterId, 240000)
		if err != nil {
			l.WithError(err).Errorf("Unable to take payment by character %d for buddy list increase.", c.CharacterId)
		}
		return r.Error(l, c)
	}
	return r.SendSuccess(l, c)
}

func (r MrGoldstein) SendSuccess(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright! Your buddy list will have 5 extra slots by now. Check and see for it yourself. And if you still need more room on your buddy list, you know who to find. Of course, it isn't going to be for free ... well, so long ...")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r MrGoldstein) Error(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("There was an issue increasing the size of your buddy list. Please try again later.")
	return script.SendNext(l, c, m.String(), script.Exit())
}
