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
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Samuel is located in Leafre - Forest of the Priest (240010501)
type Samuel struct {
}

func (r Samuel) NPCId() uint32 {
	return npc.Samuel
}

func (r Samuel) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.MeetsCriteria(l, span)(c.CharacterId, character.IsLevelCriteria(120), character.IsAJobCriteria(job.Marauder, job.Outlaw, job.Buccaneer, job.Corsair)) {
		return r.DoNotBotherMe(l, span, c)
	}
	if !quest.IsCompleted(l)(c.CharacterId, 6944) {
		return r.NotPassedTrials(l, span, c)
	}

	if character.MeetsCriteria(l, span)(c.CharacterId, character.IsAJobCriteria(job.Marauder, job.Outlaw)) {
		return r.Marevelous(l, span, c)
	}
	return r.IfIMust(l, span, c)
}

func (r Samuel) NotPassedTrials(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You have not yet passed my trials. I can not advance you until you do so.")
	return script.SendOk(l, span, c, m.String())
}

func (r Samuel) DoNotBotherMe(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please don't bother me right now, I am trying to concentrate.")
	return script.SendOk(l, span, c, m.String())
}

func (r Samuel) Marevelous(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You did a marvellous job passing my test. Are you ready to advance to your 4th job?")
	return script.SendYesNo(l, span, c, m.String(), r.Advance, script.Exit())
}

func (r Samuel) IfIMust(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If I must, I can teach you the art of your class.").NewLine().
		OpenItem(0).BlueText().AddText("Teach me the skills of my class.").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r Samuel) Advance(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.SkillBookMapleWarrior) {
		return r.MakeRoom(l, span, c)
	}
	if character.IsJob(l, span)(c.CharacterId, job.Marauder) {
		character.ChangeJob(l, span)(c.CharacterId, job.Buccaneer)
		character.TeachSkill(l)(c.CharacterId, skill.BuccaneerDragonStrike, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.BuccaneerEnergyOrb, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.BuccaneerBarrage, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.BuccaneerSpeedInfusion, 0, 10, -1)
	} else {
		character.ChangeJob(l, span)(c.CharacterId, job.Corsair)
		character.TeachSkill(l)(c.CharacterId, skill.CorsairElementalBoost, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.CorsairWrathOfTheOctopi, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.CorsairRapidFire, 0, 10, -1)
		character.TeachSkill(l)(c.CharacterId, skill.CorsairBullsEye, 0, 10, -1)
	}
	character.GainItem(l, span)(c.CharacterId, item.SkillBookMapleWarrior, 1)
	return r.ItIsDone(l, span, c)
}

func (r Samuel) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please have one slot available on ").
		BlueText().AddText("USE").
		BlackText().AddText(" inventory to receive a skill book.")
	return script.SendOk(l, span, c, m.String())
}

func (r Samuel) Selection(_ int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		//TODO this is temporary until quests get fixed?
		if character.IsJob(l, span)(c.CharacterId, job.Buccaneer) {
			character.TeachSkill(l)(c.CharacterId, skill.BuccaneerSuperTransformation, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.BuccaneerDemolition, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.BuccaneerSnatch, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.BuccaneerTimeLeap, 0, 10, -1)
		} else {
			character.TeachSkill(l)(c.CharacterId, skill.CorsairBattleShip, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.CorsairBattleshipCannon, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.CorsairBattleshipTorpedo, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.CorsairHypnotize, 0, 10, -1)
			character.TeachSkill(l)(c.CharacterId, skill.CorsairAerialStrike, 0, 10, -1)
		}
		return r.ItIsDone(l, span, c)
	}
}

func (r Samuel) ItIsDone(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("It is done. Leave me now.")
	return script.SendOk(l, span, c, m.String())
}
