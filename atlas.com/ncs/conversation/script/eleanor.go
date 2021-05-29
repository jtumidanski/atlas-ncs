package script

import (
	"atlas-ncs/character"
	"atlas-ncs/monster"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Eleanor struct {
}

func (r Eleanor) NPCId() uint32 {
	return npc.Eleanor
}

func (r Eleanor) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 20407) {
		return r.NoChallenging(l, c)
	}
	return r.DoYouWantToFace(l, c)
}

func (r Eleanor) NoChallenging(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("... Knight, you still ").
		BlueText().AddText("seem unsure to face this fight").
		BlackText().AddText(", don't you? There's no grace in challenging someone when they are still not mentally ready for the battle. Talk your peace to that big clumsy bird of yours, maybe it'll put some guts on you.")
	return SendOk(l, c, m.String())
}

func (r Eleanor) DoYouWantToFace(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hahahahaha! This place's Empress is already under my domain, that's surely a great advance on the #bBlack Wings#k' overthrow towards Maple World... And you, there? Still wants to face us? Or, better yet, since you seem strong enough to be quite a supplementary reinforcement at our service, ").
		RedText().AddText("will you meet our expectations and fancy joining us").
		BlackText().AddText(" since there's nothing more you can do?")
	return SendAcceptDecline(l, c, m.String(), r.Start, r.Cowards)
}

func (r Eleanor) Cowards(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Heh, cowards have no place on the ").
		RedText().AddText("Black Magician's").
		BlackText().AddText(" army. Begone!")
	return SendOk(l, c, m.String())
}

func (r Eleanor) Start(l logrus.FieldLogger, c Context) State {
	character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "Eleanor: Oh, lost the Empress and still challenging us? Now you've done it! Prepare yourself!!!")
	monster.SpawnMonster(l)(c.WorldId, c.ChannelId, c.MapId, monster.BlackWitch, 850, 0)
	npc.Destroy(l)(c.WorldId, c.ChannelId, c.MapId, npc.Eleanor)
	return Exit()(l, c)
}
