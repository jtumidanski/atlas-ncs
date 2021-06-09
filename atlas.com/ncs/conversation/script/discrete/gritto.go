package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/job"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/skill"
	"github.com/sirupsen/logrus"
)

// Gritto is located in Leafre - Forest of the Priest (240010501)
type Gritto struct {
}

func (r Gritto) NPCId() uint32 {
	return npc.Gritto
}

func (r Gritto) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.MeetsCriteria(l)(c.CharacterId, character.IsLevelCriteria(120), character.IsAJobCriteria(job.FirePoisonMagician, job.FirePoisonArchMagician, job.IceLightningMagician, job.IceLightningArchMagician, job.Priest, job.Bishop)) {
		return r.DoNotBotherMe(l, c)
	}
	if !character.QuestCompleted(l)(c.CharacterId, 6914) {
		return r.NotPassedTrials(l, c)
	}

	if character.MeetsCriteria(l)(c.CharacterId, character.IsAJobCriteria(job.FirePoisonMagician, job.IceLightningMagician, job.Priest)) {
		return r.Marevelous(l, c)
	}
	return r.IfIMust(l, c)
}

func (r Gritto) NotPassedTrials(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You have not yet passed my trials. I can not advance you until you do so.")
	return script.SendOk(l, c, m.String())
}

func (r Gritto) DoNotBotherMe(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please don't bother me right now, I am trying to concentrate.")
	return script.SendOk(l, c, m.String())
}

func (r Gritto) Marevelous(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You did a marvellous job passing my test. Are you ready to advance to your 4th job?")
	return script.SendYesNo(l, c, m.String(), r.Advance, script.Exit())
}

func (r Gritto) IfIMust(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If I must, I can teach you the art of your class.").NewLine().
		OpenItem(0).BlueText().AddText("Teach me the skills of my class.").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Gritto) Advance(l logrus.FieldLogger, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.SkillBookMapleWarrior) {
		return r.MakeRoom(l, c)
	}
	if character.IsJob(l)(c.CharacterId, job.FirePoisonMagician) {
		character.ChangeJob(l)(c.CharacterId, job.FirePoisonArchMagician)
		character.TeachSkill(l)(c.CharacterId, skill.FirePoisonArchMagicianBigBang, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.FirePoisonArchMagicianManaReflection, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.FirePoisonArchMagicianParalyze, 0, 10, -1)
	} else if character.IsJob(l)(c.CharacterId, job.IceLightningMagician) {
		character.ChangeJob(l)(c.CharacterId, job.IceLightningArchMagician)
		character.TeachSkill(l)(c.CharacterId, skill.IceLightningArchMagicianBigBang, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.IceLightningArchMagicianManaReflection, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.IceLightningArchMagicianChainLightning, 0, 10, -1)
	} else {
		character.ChangeJob(l)(c.CharacterId, job.Priest)
		character.TeachSkill(l)(c.CharacterId, skill.BishopBigBang, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.BishopManaReflection, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.BishopHolyShield, 0, 10, -1)
	}
	character.GainItem(l)(c.CharacterId, item.SkillBookMapleWarrior, 1)
	return r.ItIsDone(l, c)
}

func (r Gritto) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please have one slot available on ").
		BlueText().AddText("USE").
		BlackText().AddText(" inventory to receive a skill book.")
	return script.SendOk(l, c, m.String())
}

func (r Gritto) Selection(_ int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		//TODO this is temporary until quests get fixed?
		if character.IsJob(l)(c.CharacterId, job.FirePoisonArchMagician) {
			character.TeachSkill(l)(c.CharacterId, skill.FirePoisonArchMagicianMeteorShower, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.FirePoisonArchMagicianElquines, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.FirePoisonArchMagicianFireDemon, 0, 10, -1)
		} else if character.IsJob(l)(c.CharacterId, job.IceLightningArchMagician) {
			character.TeachSkill(l)(c.CharacterId, skill.IceLightningArchMagicianBlizzard, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.IceLightningArchMagicianIfrit, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.IceLightningArchMagicianIceDemon, 0, 10, -1)
		} else {
			character.TeachSkill(l)(c.CharacterId, skill.BishopGenesis, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.BishopResurrection, 0, 10, -1)
		}
		return r.ItIsDone(l, c)
	}
}

func (r Gritto) ItIsDone(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("It is done. Leave me now.")
	return script.SendOk(l, c, m.String())
}
