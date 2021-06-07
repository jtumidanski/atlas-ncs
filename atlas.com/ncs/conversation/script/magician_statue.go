package script

import (
	"atlas-ncs/character"
	"atlas-ncs/job"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MagicianStatue is located in Victoria Road - Lith Harbor (104000000)
type MagicianStatue struct {
}

func (r MagicianStatue) NPCId() uint32 {
	return npc.MagicianStatue
}

func (r MagicianStatue) Initial(l logrus.FieldLogger, c Context) State {
	if !character.IsJob(l)(c.CharacterId, job.Beginner) {
		m := message.NewBuilder().AddText("You're much stronger now. Keep training!")
		return SendOk(l, c, m.String())
	}

	if !character.MeetsCriteria(l)(c.CharacterId, character.IsLevelCriteria(10), character.HasIntelligenceCriteria(20)) {
		m := message.NewBuilder().
			AddText("If you want to be a ").
			BlueText().AddText("Magician").
			BlackText().AddText(", train yourself further until you reach ").
			BlueText().AddText("level 10 with 20 INT").
			BlackText().AddText(".")
		return SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Hey ").
		ShowCharacterName().
		AddText(", I can send you to ").
		BlueText().ShowMap(_map.MagicLibrary).
		BlackText().AddText(" if you want to be a ").
		BlueText().AddText("Magician").
		BlackText().AddText(". Do you want to go now?")
	return SendYesNoExit(l, c, m.String(), r.Warp, r.ComeBack, r.ComeBack)
}

func (r MagicianStatue) ComeBack(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Come back to me if you decided to be a ").
		BlueText().AddText("Magician").
		BlackText().AddText(".")
	return SendOk(l, c, m.String())
}

func (r MagicianStatue) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.MagicLibrary, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.MagicLibrary, c.NPCId)
	}
	return Exit()(l, c)
}