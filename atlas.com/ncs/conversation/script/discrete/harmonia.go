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

// Harmonia is located in Leafre - Forest of the Priest (240010501)
type Harmonia struct {
}

func (r Harmonia) NPCId() uint32 {
	return npc.Harmonia
}

func (r Harmonia) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.MeetsCriteria(l)(c.CharacterId, character.IsLevelCriteria(120), character.IsAJobCriteria(job.DragonKnight, job.Crusader, job.WhiteKnight, job.Hero, job.Paladin, job.DarkKnight)) {
		return r.DoNotBotherMe(l, c)
	}
	if !quest.IsCompleted(l)(c.CharacterId, 6904) {
		return r.NotPassedTrials(l, c)
	}

	if character.MeetsCriteria(l)(c.CharacterId, character.IsAJobCriteria(job.DragonKnight, job.Crusader, job.WhiteKnight)) {
		return r.Marevelous(l, c)
	}
	return r.IfIMust(l, c)
}

func (r Harmonia) NotPassedTrials(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You have not yet passed my trials. I can not advance you until you do so.")
	return script.SendOk(l, c, m.String())
}

func (r Harmonia) DoNotBotherMe(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please don't bother me right now, I am trying to concentrate.")
	return script.SendOk(l, c, m.String())
}

func (r Harmonia) Marevelous(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You did a marvellous job passing my test. Are you ready to advance to your 4th job?")
	return script.SendYesNo(l, c, m.String(), r.Advance, script.Exit())
}

func (r Harmonia) IfIMust(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If I must, I can teach you the art of your class.").NewLine().
		OpenItem(0).BlueText().AddText("Teach me the skills of my class.").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Harmonia) Advance(l logrus.FieldLogger, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.SkillBookMapleWarrior) {
		return r.MakeRoom(l, c)
	}
	if character.IsJob(l)(c.CharacterId, job.Crusader) {
		character.ChangeJob(l)(c.CharacterId, job.Hero)
		character.TeachSkill(l)(c.CharacterId, skill.HeroMonsterMagnet, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.HeroAchilles, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.HeroBrandish, 0, 10, -1)
	} else if character.IsJob(l)(c.CharacterId, job.WhiteKnight) {
		character.ChangeJob(l)(c.CharacterId, job.Paladin)
		character.TeachSkill(l)(c.CharacterId, skill.PaladinMonsterMagnet, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.PaladinAchilles, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.PaladinBlast, 0, 10, -1)
	} else {
		character.ChangeJob(l)(c.CharacterId, job.DarkKnight)
		character.TeachSkill(l)(c.CharacterId, skill.DarkKnightMonsterMagnet, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.DarkKnightAchilles, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.DarkKnightBeholder, 0, 10, -1)
	}
	character.GainItem(l)(c.CharacterId, item.SkillBookMapleWarrior, 1)
	return r.ItIsDone(l, c)
}

func (r Harmonia) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please have one slot available on ").
		BlueText().AddText("USE").
		BlackText().AddText(" inventory to receive a skill book.")
	return script.SendOk(l, c, m.String())
}

func (r Harmonia) Selection(_ int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		//TODO this is temporary until quests get fixed?
		if character.IsJob(l)(c.CharacterId, job.Hero) {
			character.TeachSkill(l)(c.CharacterId, skill.HeroEnrage, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.HeroGuardian, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.HeroStance, 0, 10, -1)
		} else if character.IsJob(l)(c.CharacterId, job.Paladin) {
			character.TeachSkill(l)(c.CharacterId, skill.PaladinStance, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.PaladinSwordHolyCharge, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.PaladinBluntWeaponHolyCharge, 0, 10, -1)
		} else {
			character.TeachSkill(l)(c.CharacterId, skill.DarkKnightStance, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.DarkKnightAuraOfBeholder, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.DarkKnightHexOfBeholder, 0, 10, -1)
		}
		return r.ItIsDone(l, c)
	}
}

func (r Harmonia) ItIsDone(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("It is done. Leave me now.")
	return script.SendOk(l, c, m.String())
}
