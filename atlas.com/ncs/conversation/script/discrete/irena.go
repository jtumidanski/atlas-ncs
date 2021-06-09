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

type Irena struct {
}

func (r Irena) NPCId() uint32 {
	return npc.Irena
}

func (r Irena) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.MeetsCriteria(l)(c.CharacterId, character.IsAJobCriteria(job.WindArcher1, job.WindArcher2, job.WindArcher3, job.WindArcher4)) {
		return r.NotWindArcher(l, c)
	}
	return r.FoundMe(l, c)
}

func (r Irena) NotWindArcher(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello there, ").ShowCharacterName().
		AddText(". Are you helping us finding the intruder? He is not in this area, I've already searched here.")
	return script.SendOk(l, c, m.String())
}

func (r Irena) FoundMe(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Darn, you found me! Then, there's only one way out! Let's fight, like ").
		RedText().AddText("Black Wings").
		BlackText().AddText(" should!")
	return script.SendOkTrigger(l, c, m.String(), r.Trigger)
}

func (r Irena) Trigger(l logrus.FieldLogger, c script.Context) script.State {
	monster.SpawnMonsterOnNPC(l)(c.WorldId, c.ChannelId, c.MapId, monster.MasterOfDisguise, c.NPCObjectId)
	npc.Destroy(l)(c.WorldId, c.ChannelId, c.MapId, c.NPCId)
	return script.Exit()(l, c)
}