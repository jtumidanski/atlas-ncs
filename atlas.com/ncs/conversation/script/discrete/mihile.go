package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/job"
	"atlas-ncs/monster"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Mihile struct {
}

func (r Mihile) NPCId() uint32 {
	return npc.Mihile
}

func (r Mihile) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.MeetsCriteria(l)(c.CharacterId, character.IsAJobCriteria(job.DawnWarrior1, job.DawnWarrior2, job.DawnWarrior3, job.DawnWarrior4)) {
		return r.NotDawnWarrior(l, c)
	}
	return r.FoundMe(l, c)
}

func (r Mihile) NotDawnWarrior(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello there, ").ShowCharacterName().
		AddText(". Are you helping us finding the intruder? He is not in this area, I've already searched here.")
	return script.SendOk(l, c, m.String())
}

func (r Mihile) FoundMe(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Darn, you found me! Then, there's only one way out! Let's fight, like ").
		RedText().AddText("Black Wings").
		BlackText().AddText(" should!")
	return script.SendOkTrigger(l, c, m.String(), r.Trigger)
}

func (r Mihile) Trigger(l logrus.FieldLogger, c script.Context) script.State {
	monster.SpawnMonsterOnNPC(l)(c.WorldId, c.ChannelId, c.MapId, monster.MasterOfDisguise, c.NPCObjectId)
	npc.Destroy(l)(c.WorldId, c.ChannelId, c.MapId, c.NPCId)
	return script.Exit()(l, c)
}
