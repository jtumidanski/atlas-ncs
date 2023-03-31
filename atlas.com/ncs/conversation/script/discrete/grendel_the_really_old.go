package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/job"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// GrendelTheReallyOld is located in Victoria Road - Magic Library (101000003)
type GrendelTheReallyOld struct {
}

func (r GrendelTheReallyOld) NPCId() uint32 {
	return npc.GrendelTheReallyOld
}

func (r GrendelTheReallyOld) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.IsJob(l, span)(c.CharacterId, job.Beginner) {
		return r.FirstJobInitial(l, span, c)
	} else if character.IsLevel(l, span)(c.CharacterId, 30) && character.IsJob(l, span)(c.CharacterId, job.Magician) {
		if character.HasItem(l, span)(c.CharacterId, item.ProofOfHero) {
			return r.SecondJobNextStep(l, span, c)
		} else if character.HasItem(l, span)(c.CharacterId, item.GrendelTheReallyOldsLetter) {
			return r.GoSeeInstructor(l, span, c)
		} else {
			return r.Astonishing(l, span, c)
		}
	}
	return nil
}

func (r GrendelTheReallyOld) FirstJobInitial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Want to be a ").
		RedText().AddText("magician").
		BlackText().AddText("? There are some standards to meet. because we can't just accept EVERYONE in... ").
		BlueText().AddText("Your level should be at least 8").
		BlackText().AddText(", with getting 20 INT as your top priority. Let's see.")
	return script.SendNext(l, span, c, m.String(), r.FirstJobRequirementCheck)
}

func (r GrendelTheReallyOld) FirstJobRequirementCheck(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.MeetsCriteria(l, span)(c.CharacterId, character.IsLevelCriteria(8), character.HasIntelligenceCriteria(20)) {
		m := message.NewBuilder().
			AddText("Oh...! You look like someone that can definitely be a part of us... all you need is a little sinister mind, and... yeah... so, what do you think? Wanna be the Magician?")
		return script.SendYesNo(l, span, c, m.String(), r.AwardFirstJob, r.FirstJobInitial)
	}

	m := message.NewBuilder().
		AddText("Train a bit more until you reach the base requirements and I can show you the way of the ").
		RedText().AddText("Magician").
		BlackText().AddText(".")
	return script.SendOk(l, span, c, m.String())
}

func (r GrendelTheReallyOld) AwardFirstJob(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.BeginnerMagicWand) {
		m := message.NewBuilder().
			AddText("Make some room in your inventory and talk back to me.")
		return script.SendNext(l, span, c, m.String(), script.Exit())
	}

	character.ChangeJob(l, span)(c.CharacterId, job.Magician)
	character.GainItem(l, span)(c.CharacterId, item.BeginnerMagicWand, 1)
	character.ResetAP(l, span)(c.CharacterId)

	m := message.NewBuilder().
		AddText("Alright, from here out, you are a part of us! You'll be living the life of a wanderer at ..., but just be patient as soon, you'll be living the high life. Alright, it ain't much, but I'll give you some of my abilities... HAAAHHH!!!")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r GrendelTheReallyOld) SecondJobNextStep(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I see you have done well. I will allow you to take the next step on your long road.")
	return script.SendNext(l, span, c, m.String(), r.SecondJobPathInfo)
}

func (r GrendelTheReallyOld) SecondJobPathInfo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright, when you have made your decision, click on [I'll choose my occupation] at the bottom.").
		BlueText().NewLine().
		OpenItem(0).AddText("Please explain to me what being the Wizard (Fire / Poison) is all about.").CloseItem().NewLine().
		OpenItem(1).AddText("Please explain to me what being the Wizard (Ice / Lightning) is all about.").CloseItem().NewLine().
		OpenItem(2).AddText("Please explain to me what being the Cleric is all about.").CloseItem().NewLine().
		OpenItem(3).AddText("I'll choose my occupation!")
	return script.SendListSelection(l, span, c, m.String(), r.SecondJobPathSelection)
}

func (r GrendelTheReallyOld) SecondJobPathSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.FirePoisonWizardInfo
	case 1:
		return r.IceLightningWizardInfo
	case 2:
		return r.ClericInfo
	case 3:
		return r.SecondJobChoice
	}
	return nil
}

func (r GrendelTheReallyOld) FirePoisonWizardInfo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Magicians that master ").
		RedText().AddText("Fire/Poison-based magic").
		BlackText().NewLine().NewLine().
		BlueText().AddText("Wizards").
		BlackText().AddText(" are a active class that deal magical, elemental damage. These abilities grants them a significant advantage against enemies weak to their element. With their skills ").
		RedText().AddText("Meditation").
		BlackText().AddText(" and ").
		RedText().AddText("Slow").
		BlackText().AddText(", ").
		BlueText().AddText("Wizards").
		BlackText().AddText(" can increase their magic attack and reduce the opponent's mobility. ").
		BlueText().AddText("Fire/Poison Wizards").
		BlackText().AddText(" contains a powerful flame arrow attack and poison attack.")
	return script.SendNext(l, span, c, m.String(), r.SecondJobNextStep)
}

func (r GrendelTheReallyOld) IceLightningWizardInfo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Magicians that master ").
		RedText().AddText("Ice/Lightning-based magic").
		BlackText().NewLine().NewLine().
		BlueText().AddText("Wizards").
		BlackText().AddText(" are a active class that deal magical, elemental damage. These abilities grants them a significant advantage against enemies weak to their element. With their skills ").
		RedText().AddText("Meditation").
		BlackText().AddText(" and ").
		RedText().AddText("Slow").
		BlackText().AddText(", ").
		BlueText().AddText("Wizards").
		BlackText().AddText(" can increase their magic attack and reduce the opponent's mobility. ").
		BlueText().AddText("Ice/Lightning Wizards").
		BlackText().AddText(" have a freezing ice attack and a striking lightning attack.")
	return script.SendNext(l, span, c, m.String(), r.SecondJobNextStep)
}

func (r GrendelTheReallyOld) ClericInfo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Magicians that master ").
		RedText().AddText("Holy magic").
		BlackText().NewLine().NewLine().
		BlueText().AddText("Clerics").
		BlackText().AddText(" are a powerful supportive class, bound to be accepted into any Party. That's because the have the power to ").
		RedText().AddText("Heal").
		BlackText().AddText(" themselves and others in their party. Using ").
		RedText().AddText("Bless").
		BlackText().AddText(", ").
		BlueText().AddText("Clerics").
		BlackText().AddText(" can buff the attributes and reduce the amount of damage taken. This class is on worth going for if you find it hard to survive. ").
		BlueText().AddText("Clerics").
		BlackText().AddText(" are especially effective against undead monsters.")
	return script.SendNext(l, span, c, m.String(), r.SecondJobNextStep)
}

func (r GrendelTheReallyOld) SecondJobChoice(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Now... have you made up your mind? Please choose the job you'd like to select for your 2nd job advancement. ").
		BlueText().NewLine().
		OpenItem(0).AddText("Wizard (Fire / Poison)").CloseItem().NewLine().
		OpenItem(1).AddText("Wizard (Ice / Lightning)").CloseItem().NewLine().
		OpenItem(2).AddText("Cleric").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.SecondJobSelection)
}

func (r GrendelTheReallyOld) SecondJobSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.ConfirmSecondJob(job.FirePoisonWizard, "Wizard (Fire / Poison)")
	case 1:
		return r.ConfirmSecondJob(job.IceLightningWizard, "Wizard (Ice / Lightning)")
	case 2:
		return r.ConfirmSecondJob(job.Cleric, "Cleric")
	}
	return nil
}

func (r GrendelTheReallyOld) ConfirmSecondJob(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if character.HasItem(l, span)(c.CharacterId, item.GrendelTheReallyOldsLetter) {
			return script.Exit()(l, span, c)
		}
		m := message.NewBuilder().
			AddText("So you want to make the second job advancement as the ").
			BlueText().AddText(jobName).
			BlackText().AddText("? You know you won't be able to choose a different job for the 2nd job advancement once you make your decision here, right?")
		return script.SendYesNo(l, span, c, m.String(), r.PerformSecondJobAdvancement(jobId, jobName), r.SecondJobChoice)
	}
}

func (r GrendelTheReallyOld) PerformSecondJobAdvancement(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if character.HasItem(l, span)(c.CharacterId, item.ProofOfHero) {
			character.GainItem(l, span)(c.CharacterId, item.ProofOfHero, -1)
		}
		quest.Complete(l)(c.CharacterId, 100008)
		character.ChangeJob(l, span)(c.CharacterId, jobId)

		return r.SecondJobAdvancementSuccess(jobName)(l, span, c)
	}
}

func (r GrendelTheReallyOld) SecondJobAdvancementSuccess(jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Alright, you're the ").
			BlueText().AddText(jobName).
			BlackText().AddText(" from here on out. Magician and wizards are the intelligent bunch with incredible magical prowess, able to pierce the mind and the psychological structure of the monsters with ease... please train yourself each and everyday. I'll help you become even stronger than you already are.")
		return script.SendNext(l, span, c, m.String(), r.SecondJobSkillBook(jobName))
	}
}

func (r GrendelTheReallyOld) SecondJobSkillBook(jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("I have just given you a book that gives you the list of skills you can acquire as a ").
			BlueText().AddText(jobName).
			BlackText().AddText(". Also your etc inventory has expanded by adding another row to it. Your max HP and MP have increased, too. Go check and see for it yourself.")
		return script.SendNextPrevious(l, span, c, m.String(), r.SecondJobSP(jobName), r.SecondJobAdvancementSuccess(jobName))
	}
}

func (r GrendelTheReallyOld) SecondJobSP(jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("I have also given you a little bit of ").
			BlueText().AddText("SP").
			BlackText().AddText(". Open the ").
			BlueText().AddText("Skill Menu").
			BlackText().AddText(" located at the bottom left corner. you'll be able to boost up the newer acquired 2nd level skills. A word of warning, though. You can't boost them up all at once. Some of the skills are only available after you have learned other skills. Make sure you remember that.")
		return script.SendNextPrevious(l, span, c, m.String(), r.NeedToBeStrong(jobName), r.SecondJobSkillBook(jobName))
	}
}

func (r GrendelTheReallyOld) NeedToBeStrong(jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().
			BlueText().AddText(jobName).
			BlackText().AddText(" need to be strong. But remember that you can't abuse that power and use it on a weakling. Please use your enormous power the right way, because... for you to use that the right way, that is much harden than just getting stronger. Please find me after you have advanced much further. I'll be waiting for you.")
		return script.SendNextPrevious(l, span, c, m.String(), script.Exit(), r.SecondJobSP(jobName))
	}
}

func (r GrendelTheReallyOld) GoSeeInstructor(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Go and see the ").
		BlueText().ShowNPC(npc.MagicianJobInstructor).
		BlackText().AddText(".")
	return script.SendOk(l, span, c, m.String())
}

func (r GrendelTheReallyOld) Astonishing(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The progress you have made is astonishing.")
	return script.SendNext(l, span, c, m.String(), r.GoodDecision)
}

func (r GrendelTheReallyOld) GoodDecision(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 100006) {
		quest.Start(l)(c.CharacterId, 100006)
	}
	m := message.NewBuilder().AddText("Good decision. You look strong, but I need to see if you really are strong enough to pass the test, it's not a difficult test, so you'll do just fine. Here, take my letter first... make sure you don't lose it!")
	return script.SendNext(l, span, c, m.String(), r.TakeThisLetter)
}

func (r GrendelTheReallyOld) TakeThisLetter(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItem(l, span)(c.CharacterId, item.GrendelTheReallyOldsLetter) {
		character.GainItem(l, span)(c.CharacterId, item.GrendelTheReallyOldsLetter, 1)
	}
	m := message.NewBuilder().
		AddText("Please get this letter to ").
		BlueText().ShowNPC(npc.MagicianJobInstructor).
		BlackText().AddText(" who's around ").
		BlueText().ShowMap(_map.TheForestNorthOfEllinia).
		BlackText().AddText(" near Ellinia. He is taking care of the job of an instructor in place of me. Give him the letter and he'll test you in place of me. Best of luck to you.")
	return script.SendNextPrevious(l, span, c, m.String(), script.Exit(), r.GoodDecision)
}
