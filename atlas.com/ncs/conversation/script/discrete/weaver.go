package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Weaver is located in Ludibrium - Ludibrium Pet Walkway (220000006)
type Weaver struct {
}

func (r Weaver) NPCId() uint32 {
	return npc.Weaver
}

func (r Weaver) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This is the road where you can go take a walk with your pet. You can walk around with it, or you can train your pet to go through obstacles here. If you aren't too close with your pet yet, that may present a problem and he will not follow your command as much... So, what do you think? Wanna train your pet?")
	return script.SendYesNo(l, c, m.String(), r.Continue, r.TooBusy)
}

func (r Weaver) TooBusy(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmm ... too busy to do it right now? If you feel like doing it, though, come back and find me.")
	return script.SendOk(l, c, m.String())
}

func (r Weaver) Continue(l logrus.FieldLogger, c script.Context) script.State {
	if character.HasItem(l)(c.CharacterId, item.WeaversLetter) {
		return r.Info(l, c)
	}
	return r.HereIsTheLetter(l, c)
}

func (r Weaver) HereIsTheLetter(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.WeaversLetter, 1)
	m := message.NewBuilder().
		AddText("Ok, here's the letter. He wouldn't know I sent you if you just went there straight, so go through the obstacles with your pet, go to the very top, and then talk to Trainer Frod to give him the letter. It won't be hard if you pay attention to your pet while going through obstacles. Good luck!")
	return script.SendOk(l, c, m.String())
}

func (r Weaver) Info(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Get that letter, jump over obstacles with your pet, and take that letter to my brother Trainer Frod. Give him the letter and something good is going to happen to your pet.")
	return script.SendOk(l, c, m.String())
}
