package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/job"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/party"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// InsignificantBeing is located in Dungeon - Another Entrance (105090200)
type InsignificantBeing struct {
}

func (r InsignificantBeing) NPCId() uint32 {
	return npc.InsignificantBeing
}

func (r InsignificantBeing) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 6107) && !quest.IsStarted(l)(c.CharacterId, 6108) {
		return r.NotAllowedToEnter(l, c)
	}

	if !party.HasParty(l)(c.CharacterId) {
		return r.FormParty(l, c)
	}

	p, err := party.GetForCharacter(l)(c.CharacterId)
	if err != nil {
		l.WithError(err).Errorf("Unable to retieve party information for character %d.", c.CharacterId)
		return script.Exit()(l, c)
	}

	if len(p.Members()) != 2 {
		return r.PartySize(l, c)
	}

	for _, m := range p.Members() {
		if !character.MeetsCriteria(l)(m.Id(), character.IsAJobCriteria(job.BowMaster, job.Marksman, job.GM)) {
			return r.NotElligible(l, c)
		}
		if !character.IsLevel(l)(m.Id(), 120) {
			return r.LowLevel(l, c)
		}
	}

	return r.CheckEvent(l, c)
}

func (r InsignificantBeing) NotAllowedToEnter(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You're not allowed to enter the other world with unknown reason.")
	return script.SendOk(l, c, m.String())
}

func (r InsignificantBeing) FormParty(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please form a party and talk to me again.")
	return script.SendOk(l, c, m.String())
}

func (r InsignificantBeing) PartySize(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Please make sure that your party is a size of 2.")
	return script.SendOk(l, c, m.String())
}

func (r InsignificantBeing) NotElligible(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("One of your party member's job is not eligible for entering the other world.")
	return script.SendOk(l, c, m.String())
}

func (r InsignificantBeing) LowLevel(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("One of your party member's level is not eligible for entering the other world.")
	return script.SendOk(l, c, m.String())
}

func (r InsignificantBeing) CheckEvent(l logrus.FieldLogger, c script.Context) script.State {
	//TODO
	return script.Exit()(l, c)
}
