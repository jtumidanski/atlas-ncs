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

// Legor is located in Leafre - Forest of the Priest (240010501)
type Legor struct {
}

func (r Legor) NPCId() uint32 {
	return npc.Legor
}

func (r Legor) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.MeetsCriteria(l)(c.CharacterId, character.IsLevelCriteria(120), character.IsAJobCriteria(job.Ranger, job.BowMaster, job.Sniper, job.Marksman)) {
		return r.DoNotBotherMe(l, c)
	}
	if !quest.IsCompleted(l)(c.CharacterId, 6924) {
		return r.NotPassedTrials(l, c)
	}

	if character.MeetsCriteria(l)(c.CharacterId, character.IsAJobCriteria(job.Ranger, job.Sniper)) {
		return r.Marevelous(l, c)
	}
	return r.IfIMust(l, c)
}

func (r Legor) NotPassedTrials(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You have not yet passed my trials. I can not advance you until you do so.")
	return script.SendOk(l, c, m.String())
}

func (r Legor) DoNotBotherMe(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please don't bother me right now, I am trying to concentrate.")
	return script.SendOk(l, c, m.String())
}

func (r Legor) Marevelous(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You did a marvellous job passing my test. Are you ready to advance to your 4th job?")
	return script.SendYesNo(l, c, m.String(), r.Advance, script.Exit())
}

func (r Legor) IfIMust(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If I must, I can teach you the art of your class.").NewLine().
		OpenItem(0).BlueText().AddText("Teach me the skills of my class.").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Legor) Advance(l logrus.FieldLogger, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.SkillBookMapleWarrior) {
		return r.MakeRoom(l, c)
	}
	if character.IsJob(l)(c.CharacterId, job.Ranger) {
		character.ChangeJob(l)(c.CharacterId, job.BowMaster)
		character.TeachSkill(l)(c.CharacterId, skill.BowmasterSharpEyes, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.BowmasterBowExpert, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.BowmasterHamstring, 0, 10, -1)
	} else {
		character.ChangeJob(l)(c.CharacterId, job.Marksman)
		character.TeachSkill(l)(c.CharacterId, skill.MarksmanSharpEyes, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.MarksmanBoost, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.MarksmanBlind, 0, 10, -1)
	}
	character.GainItem(l)(c.CharacterId, item.SkillBookMapleWarrior, 1)
	return r.ItIsDone(l, c)
}

func (r Legor) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please have one slot available on ").
		BlueText().AddText("USE").
		BlackText().AddText(" inventory to receive a skill book.")
	return script.SendOk(l, c, m.String())
}

func (r Legor) Selection(_ int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		//TODO this is temporary until quests get fixed?
		if character.IsJob(l)(c.CharacterId, job.BowMaster) {
			character.TeachSkill(l)(c.CharacterId, skill.BowmasterConcentrate, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.BowmasterPhoenix, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.BowmasterHurricane, 0, 10, -1)
		} else {
			character.TeachSkill(l)(c.CharacterId, skill.MarksmanSnipe, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.MarksmanFrostPrey, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.MarksmanPiercingArrow, 0, 10, -1)
		}
		return r.ItIsDone(l, c)
	}
}

func (r Legor) ItIsDone(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("It is done. Leave me now.")
	return script.SendOk(l, c, m.String())
}
