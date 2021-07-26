package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/job"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"atlas-ncs/skill"
	"github.com/sirupsen/logrus"
)

// Hellin is located in Leafre - Forest of the Priest (240010501)
type Hellin struct {
}

func (r Hellin) NPCId() uint32 {
	return npc.Hellin
}

func (r Hellin) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.MeetsCriteria(l)(c.CharacterId, character.IsLevelCriteria(120), character.IsAJobCriteria(job.ChiefBandit, job.Shadower, job.Hermit, job.NightLord)) {
		return r.DoNotBotherMe(l, c)
	}
	if !quest.IsCompleted(l)(c.CharacterId, 6934) {
		return r.NotPassedTrials(l, c)
	}

	if character.MeetsCriteria(l)(c.CharacterId, character.IsAJobCriteria(job.ChiefBandit, job.Hermit)) {
		return r.Marevelous(l, c)
	}
	return r.IfIMust(l, c)
}

func (r Hellin) NotPassedTrials(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You have not yet passed my trials. I can not advance you until you do so.")
	return script.SendOk(l, c, m.String())
}

func (r Hellin) DoNotBotherMe(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please don't bother me right now, I am trying to concentrate.")
	return script.SendOk(l, c, m.String())
}

func (r Hellin) Marevelous(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You did a marvellous job passing my test. Are you ready to advance to your 4th job?")
	return script.SendYesNo(l, c, m.String(), r.Advance, script.Exit())
}

func (r Hellin) IfIMust(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If I must, I can teach you the art of your class.").NewLine().
		OpenItem(0).BlueText().AddText("Teach me the skills of my class.").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Hellin) Advance(l logrus.FieldLogger, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.SkillBookMapleWarrior) {
		return r.MakeRoom(l, c)
	}
	if character.IsJob(l)(c.CharacterId, job.ChiefBandit) {
		character.ChangeJob(l)(c.CharacterId, job.Shadower)
		character.TeachSkill(l)(c.CharacterId, skill.ShadowerShadowShifter, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.ShadowerVenemousStab, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.ShadowerBoomerangStep, 0, 10, -1)
	} else {
		character.ChangeJob(l)(c.CharacterId, job.NightLord)
		character.TeachSkill(l)(c.CharacterId, skill.NightLordShadowShifter, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.NightLordVenemousStar, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.NightLordShadowStars, 0, 10, -1)
	}
	character.GainItem(l)(c.CharacterId, item.SkillBookMapleWarrior, 1)
	return r.ItIsDone(l, c)
}

func (r Hellin) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please have one slot available on ").
		BlueText().AddText("USE").
		BlackText().AddText(" inventory to receive a skill book.")
	return script.SendOk(l, c, m.String())
}

func (r Hellin) Selection(_ int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		//TODO this is temporary until quests get fixed?
		if character.IsJob(l)(c.CharacterId, job.Shadower) {
			character.TeachSkill(l)(c.CharacterId, skill.ShadowerNinjaAmbush, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.ShadowerAssassinate, 0, 10, -1)
		} else {
			character.TeachSkill(l)(c.CharacterId, skill.NightLordNinjaStorm, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.NightLordNinjaAmbush, 0, 10, -1)
		}
		return r.ItIsDone(l, c)
	}
}

func (r Hellin) ItIsDone(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("It is done. Leave me now.")
	return script.SendOk(l, c, m.String())
}
