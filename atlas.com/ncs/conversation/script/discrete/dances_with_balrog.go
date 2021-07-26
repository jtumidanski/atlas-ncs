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
	"github.com/sirupsen/logrus"
)

// DancesWithBalrog is located in Victoria Road - Warriors' Sanctuary (102000003)
type DancesWithBalrog struct {
}

func (r DancesWithBalrog) NPCId() uint32 {
	return npc.DancesWithBalrog
}

func (r DancesWithBalrog) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.IsJob(l)(c.CharacterId, job.Beginner) {
		return r.FirstJob(l, c)
	}

	if character.MeetsCriteria(l)(c.CharacterId, character.IsJobCriteria(job.Warrior), character.IsLevelCriteria(30)) {
		return r.SecondJob(l, c)
	}
	//TODO third job
	return r.ChosenWisely(l, c)
}

func (r DancesWithBalrog) ChosenWisely(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You have chosen wisely.")
	return script.SendOk(l, c, m.String())
}

func (r DancesWithBalrog) FirstJob(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Do you want to become a ").
		RedText().AddText("warrior").
		BlackText().AddText("? You need to meet some criteria in order to do so.").
		BlueText().AddText(" You should be at least in level 10, and at least STR 35").
		BlackText().AddText(". Let's see...")
	return script.SendNext(l, c, m.String(), r.ValidateFirstJobCriteria)
}

func (r DancesWithBalrog) ValidateFirstJobCriteria(l logrus.FieldLogger, c script.Context) script.State {
	if !character.MeetsCriteria(l)(c.CharacterId, character.IsJobCriteria(job.Beginner), character.IsLevelCriteria(10), character.HasStrengthCriteria(35)) {
		return r.TrainMore(l, c)
	}
	return r.ImportantChoice(l, c)
}

func (r DancesWithBalrog) TrainMore(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Train a bit more until you reach the base requirements and I can show you the way of the ").
		RedText().AddText("Warrior").
		BlackText().AddText(".")
	return script.SendOk(l, c, m.String())
}

func (r DancesWithBalrog) ImportantChoice(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It is an important and final choice. You will not be able to turn back.")
	return script.SendNextPreviousExit(l, c, m.String(), r.ProcessFirstJob, r.FirstJob, r.MakeUpYourMind)
}

func (r DancesWithBalrog) MakeUpYourMind(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Make up your mind and visit me again.")
	return script.SendOk(l, c, m.String())
}

func (r DancesWithBalrog) ProcessFirstJob(l logrus.FieldLogger, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.BeginnerWarriorsSword) {
		return r.MakeInventoryRoom(l, c)
	}
	if character.IsJob(l)(c.CharacterId, job.Beginner) {
		character.ChangeJob(l)(c.CharacterId, job.Warrior)
		character.GainItem(l)(c.CharacterId, item.BeginnerWarriorsSword, 1)
		character.ResetAP(l)(c.CharacterId)
	}
	return r.FirstJobAdvance(l, c)
}

func (r DancesWithBalrog) MakeInventoryRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Make some room in your inventory and talk back to me.")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r DancesWithBalrog) FirstJobAdvance(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("From here on out, you are going to the Warrior path. This is not an easy job, but if you have discipline and confidence in your own body and skills, you will overcome any difficulties in your path. Go, young Warrior!")
	return script.SendNext(l, c, m.String(), r.GottenStronger)
}

func (r DancesWithBalrog) GottenStronger(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You've gotten much stronger now. Plus every single one of your inventories have added slots. A whole row, to be exact. Go see for it yourself. I just gave you a little bit of #bSP#k. When you open up the #bSkill#k menu on the lower left corner of the screen, there are skills you can learn by using SP's. One warning, though: You can't raise it all together all at once. There are also skills you can acquire only after having learned a couple of skills first.")
	return script.SendNextPrevious(l, c, m.String(), r.Reminder, r.FirstJobAdvance)
}

func (r DancesWithBalrog) Reminder(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Now a reminder. Once you have chosen, you cannot change up your mind and try to pick another path. Go now, and live as a proud Warrior.")
	return script.SendNextPrevious(l, c, m.String(), script.Exit(), r.GottenStronger)
}

func (r DancesWithBalrog) SecondJob(l logrus.FieldLogger, c script.Context) script.State {
	if character.HasItem(l)(c.CharacterId, item.ProofOfHero) {
		return r.ChooseAPath(l, c)
	} else if character.HasItem(l)(c.CharacterId, item.DancesWithBalrogsLetter) {
		return r.GoAndSee(l, c)
	}
	return r.Astonishing(l, c)
}

func (r DancesWithBalrog) ChooseAPath(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Oh... you came back safe! I knew you'd breeze through. I'll admit, you are a strong, formidable Warrior! Alright, I'll make you an even stronger Warrior than you already are. But before that, you need to choose one of the three paths that you'll be given. It isn't going to be easy, so if you have and questions, feel free to ask.")
	return script.SendNext(l, c, m.String(), r.PathInfo)
}

func (r DancesWithBalrog) GoAndSee(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Go and see the ").
		BlueText().ShowNPC(npc.WarriorJobInstructor).
		BlackText().AddText(".")
	return script.SendOk(l, c, m.String())
}

func (r DancesWithBalrog) Astonishing(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The progress you have made is astonishing.")
	return script.SendNext(l, c, m.String(), r.GoodDecision)
}

func (r DancesWithBalrog) GoodDecision(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Good decision. You look strong, but I need to see if you really are strong enough to pass the test, it's not a difficult test, so you'll do just fine. Here, take my letter first... make sure you don't lose it!")
	if !quest.IsStarted(l)(c.CharacterId, 100003) {
		quest.Start(l)(c.CharacterId, 100003)
	}
	return script.SendNext(l, c, m.String(), r.GiveLetter)
}

func (r DancesWithBalrog) GiveLetter(l logrus.FieldLogger, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.DancesWithBalrogsLetter) {
		return r.MakeSpace(l, c)
	}

	if !character.HasItem(l)(c.CharacterId, item.DancesWithBalrogsLetter) {
		character.GainItem(l)(c.CharacterId, item.DancesWithBalrogsLetter, 1)
	}

	m := message.NewBuilder().
		AddText("Please get this letter to ").
		BlueText().ShowNPC(npc.WarriorJobInstructor).
		BlackText().AddText(" who's around ").
		BlueText().ShowMap(_map.WestRockyMountainIV).
		BlackText().AddText(" near Perion. He is taking care of the job of an instructor in place of me. Give him the letter and he'll test you in place of me. Best of luck to you.")
	return script.SendNextPrevious(l, c, m.String(), script.Exit(), r.GoodDecision)
}

func (r DancesWithBalrog) MakeSpace(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please, make some space in your inventory.")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r DancesWithBalrog) PathInfo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright, when you have made your decision, click on [I'll choose my occupation] at the bottom.").NewLine().
		OpenItem(0).BlueText().AddText("Please explain to me what being the Fighter is all about.").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Please explain to me what being the Page is all about.").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Please explain to me what being the Spearman is all about.").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("I'll choose my occupation!").CloseItem()
	return script.SendListSelectionExit(l, c, m.String(), r.PathInfoSelected, r.MakeUpYourMind)
}

func (r DancesWithBalrog) PathInfoSelected(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.FighterInfo
	case 1:
		return r.PageInfo
	case 2:
		return r.SpearmanInfo
	case 3:
		return r.PathSelection
	}
	return nil
}

func (r DancesWithBalrog) FighterInfo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Warriors that master ").
		RedText().AddText("Swords or Axes").
		BlackText().AddText(".").NewLine().NewLine().
		RedText().AddText("Fighters").
		BlackText().AddText(" get ").
		BlueText().AddText("Rage").
		BlackText().AddText(", which boosts your party's weapon attack by 10. During 2nd job this is strongly appreciated, as it is free (except for -10 wep def, which is not going to impact the damage you take much at all), takes no Use slots and increases each party member's damage (except Magicians) by several hundreds. The other classes can give themselves a weapon attack boost as well, but need items to do so. ").
		RedText().AddText("Fighters").
		BlackText().AddText(" also get ").
		BlueText().AddText("Power Guard").
		BlackText().AddText(", reducing touch damage by 40% and deals it back to the monster. This is the main reason why ").
		RedText().AddText("Fighters").
		BlackText().AddText(" are considered soloers is because this reduces pot costs immensely.")
	return script.SendNextExit(l, c, m.String(), r.PathInfo, r.MakeUpYourMind)
}

func (r DancesWithBalrog) PageInfo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Warriors that master ").
		RedText().AddText("Swords or Maces/Blunt weapons").
		BlackText().AddText(".").NewLine().NewLine().
		RedText().AddText("Pages").
		BlackText().AddText(" get ").
		BlueText().AddText("Threaten").
		BlackText().AddText(", a skill that lowers the enemies' weapon defense and weapon attack by 20; this is mostly used to lower damage dealt to you. Pages also get ").
		BlueText().AddText("Power Guard").
		BlackText().AddText(", reducing touch damage by 40% and deals it back to the monster. This is one of the main reason why ").
		BlueText().AddText("Pages/WKs").
		BlackText().AddText(" are considered solo players, that's because this reduces pot costs immensely. Of course, constant KB and ").
		BlueText().AddText("Ice Charge").
		BlackText().AddText(" helps also to the soloing factor.")
	return script.SendNextExit(l, c, m.String(), r.PathInfo, r.MakeUpYourMind)
}

func (r DancesWithBalrog) SpearmanInfo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Warriors that master ").
		RedText().AddText("Spears or Polearms").
		BlackText().AddText(".").NewLine().NewLine().
		RedText().AddText("Spearmen").
		BlackText().AddText(" get ").
		BlueText().AddText("Hyper Body").
		BlackText().AddText(", which boosts your max HP/MP and that of your party by 60% when maxed. This skill is particularly useful for helping partied Thieves, Archers, and Magicians to survive more hits from enemies and/or PQ bosses. They also get ").
		BlueText().AddText("Iron Will").
		BlackText().AddText(" which gives +20 wep def and +20 mag def for 300 sec. It is basically a nerfed Bless with 100 seconds more duration but gives no accuracy or avoidability bonus. Even with this skill maxed, it isn't even close to being in the same league as Power Guard and is why Spearmen/Dark Knights are not considered a soloing class.")
	return script.SendNextExit(l, c, m.String(), r.PathInfo, r.MakeUpYourMind)
}

func (r DancesWithBalrog) PathSelection(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Now... have you made up your mind? Please choose the job you'd like to select for your 2nd job advancement. ").NewLine().
		OpenItem(0).BlueText().AddText("Fighter").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Page").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("Spearman").CloseItem()
	return script.SendListSelectionExit(l, c, m.String(), r.PathSelected, r.MakeUpYourMind)
}

func (r DancesWithBalrog) PathSelected(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.SecondJobConfirmation(job.Fighter, "Fighter")
	case 1:
		return r.SecondJobConfirmation(job.Page, "Page")
	case 2:
		return r.SecondJobConfirmation(job.Spearman, "Spearman")
	}
	return nil
}

func (r DancesWithBalrog) SecondJobConfirmation(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("So you want to make the second job advancement as the ").
			AddText(jobName).
			AddText("? You know you won't be able to choose a different job for the 2nd job advancement once you make your decision here, right? Are you sure about this?")
		return script.SendYesNoExit(l, c, m.String(), r.SecondJobAdvance(jobId, jobName), r.PathSelection, r.MakeUpYourMind)
	}
}

func (r DancesWithBalrog) SecondJobAdvance(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if character.HasItem(l)(c.CharacterId, item.ProofOfHero) {
			character.GainItem(l)(c.CharacterId, item.ProofOfHero, -1)
		}

		quest.Complete(l)(c.CharacterId, 100005)

		if !character.IsJob(l)(c.CharacterId, jobId) {
			character.ChangeJob(l)(c.CharacterId, jobId)
		}

		return r.SecondJobSuccess(jobId, jobName)(l, c)
	}
}

func (r DancesWithBalrog) SecondJobSuccess(jobId uint16, jobName string) script.StateProducer {
	switch jobId {
	case job.Fighter:
		return r.FighterSuccess(jobId, jobName)
	case job.Page:
		return r.PageSuccess(jobId, jobName)
	case job.Spearman:
		return r.SpearmanSuccess(jobId, jobName)
	}
	return script.Exit()
}

func (r DancesWithBalrog) FighterSuccess(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Alright, you have now become the ").
			BlueText().AddText("Fighter").
			BlackText().AddText(". A fighter strives to become the strongest of the strong, and never stops fighting. Don't ever lose that will to fight, and push forward 24/7. I'll help you become even stronger than you already are.")
		return script.SendNext(l, c, m.String(), r.BookGiven(jobId, jobName))
	}
}

func (r DancesWithBalrog) PageSuccess(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Alright, you have now become a ").
			BlueText().AddText("Page").
			BlackText().AddText("! Pages have high intelligence and bravery, which I hope you'll employ throughout your journey to the right path. I'll help you become much stronger than you already are.")
		return script.SendNext(l, c, m.String(), r.BookGiven(jobId, jobName))
	}
}

func (r DancesWithBalrog) SpearmanSuccess(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Alright, you have now become the ").
			BlueText().AddText("Spearman").
			BlackText().AddText(". The Spearman use the power of darkness to take out the enemies, always in shadows... Please believe in yourself and your awesome power as you go in your journey. I'll help you become much stronger than you are right now.")
		return script.SendNext(l, c, m.String(), r.BookGiven(jobId, jobName))
	}
}

func (r DancesWithBalrog) BookGiven(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("I have just given you a book that gives you the list of skills you can acquire as a ").
			AddText(jobName).
			AddText(". Also your etc inventory has expanded by adding another row to it. Your max HP and MP have increased, too. Go check and see for it yourself.")
		return script.SendNextPrevious(l, c, m.String(), r.SPGiven(jobId, jobName), r.SecondJobSuccess(jobId, jobName))
	}
}

func (r DancesWithBalrog) SPGiven(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("I have also given you a little bit of ").
			BlueText().AddText("SP").
			BlackText().AddText(". Open the ").
			BlueText().AddText("Skill Menu").
			BlackText().AddText(" located at the bottom left corner. you'll be able to boost up the newer acquired 2nd level skills. A word of warning, though. You can't boost them up all at once. Some of the skills are only available after you have learned other skills. Make sure yo remember that.")
		return script.SendNextPrevious(l, c, m.String(), r.BecomeStrong(jobId, jobName), r.BookGiven(jobId, jobName))
	}
}

func (r DancesWithBalrog) BecomeStrong(jobId uint16, jobName string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText(jobName).
			AddText(" need to be strong. But remember that you can't abuse that power and use it on a weakling. Please use your enormous power the right way, because... for you to use that the right way, that is much harden than just getting stronger. Please find me after you have advanced much further. I'll be waiting for you.")
		return script.SendNextPrevious(l, c, m.String(), script.Exit(), r.SPGiven(jobId, jobName))
	}
}
