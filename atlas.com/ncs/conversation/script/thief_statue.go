package script

import (
	"atlas-ncs/character"
	"atlas-ncs/job"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ThiefStatue is located in Victoria Road - Lith Harbor (104000000)
type ThiefStatue struct {
}

func (r ThiefStatue) NPCId() uint32 {
	return npc.ThiefStatue
}

func (r ThiefStatue) Initial(l logrus.FieldLogger, c Context) State {
	if !character.IsJob(l)(c.CharacterId, job.Beginner) {
		m := message.NewBuilder().AddText("You're much stronger now. Keep training!")
		return SendOk(l, c, m.String())
	}

	if !character.MeetsCriteria(l)(c.CharacterId, character.IsLevelCriteria(10), character.HasDexterityCriteria(25)) {
		m := message.NewBuilder().
			AddText("If you want to be a ").
			BlueText().AddText("Thief").
			BlackText().AddText(", train yourself further until you reach ").
			BlueText().AddText("level 10 with 25 DEX").
			BlackText().AddText(".")
		return SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Hey ").
		ShowCharacterName().
		AddText(", I can send you to ").
		BlueText().ShowMap(_map.ThievesHideout).
		BlackText().AddText(" if you want to be a ").
		BlueText().AddText("Thief").
		BlackText().AddText(". Do you want to go now?")
	return SendYesNoExit(l, c, m.String(), r.Warp, r.ComeBack, r.ComeBack)
}

func (r ThiefStatue) ComeBack(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Come back to me if you decided to be a ").
		BlueText().AddText("Thief").
		BlackText().AddText(".")
	return SendOk(l, c, m.String())
}

func (r ThiefStatue) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ThievesHideout, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ThievesHideout, c.NPCId)
	}
	return Exit()(l, c)
}