package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/job"
	"atlas-ncs/monster"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Eckhart struct {
}

func (r Eckhart) NPCId() uint32 {
	return npc.Eckhart
}

func (r Eckhart) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.MeetsCriteria(l, span)(c.CharacterId, character.IsAJobCriteria(job.NightWalker1, job.NightWalker2, job.NightWalker3, job.NightWalker4)) {
		return r.NotNightWalker(l, span, c)
	}
	return r.FoundMe(l, span, c)
}

func (r Eckhart) NotNightWalker(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello there, ").ShowCharacterName().
		AddText(". Are you helping us finding the intruder? He is not in this area, I've already searched here.")
	return script.SendOk(l, span, c, m.String())
}

func (r Eckhart) FoundMe(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Darn, you found me! Then, there's only one way out! Let's fight, like ").
		RedText().AddText("Black Wings").
		BlackText().AddText(" should!")
	return script.SendOkTrigger(l, span, c, m.String(), r.Trigger)
}

func (r Eckhart) Trigger(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	monster.SpawnMonsterOnNPC(l)(c.WorldId, c.ChannelId, c.MapId, monster.MasterOfDisguise, c.NPCObjectId)
	npc.Destroy(l)(c.WorldId, c.ChannelId, c.MapId, c.NPCId)
	return script.Exit()(l, span, c)
}