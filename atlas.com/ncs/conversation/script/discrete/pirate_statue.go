package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/job"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// PirateStatue is located in Victoria Road - Lith Harbor (104000000)
type PirateStatue struct {
}

func (r PirateStatue) NPCId() uint32 {
	return npc.PirateStatue
}

func (r PirateStatue) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.IsJob(l)(c.CharacterId, job.Beginner) {
		m := message.NewBuilder().AddText("You're much stronger now. Keep training!")
		return script.SendOk(l, c, m.String())
	}

	if !character.MeetsCriteria(l)(c.CharacterId, character.IsLevelCriteria(10), character.HasDexterityCriteria(20)) {
		m := message.NewBuilder().
			AddText("If you want to be a ").
			BlueText().AddText("Pirate").
			BlackText().AddText(", train yourself further until you reach ").
			BlueText().AddText("level 10 with 20 DEX").
			BlackText().AddText(".")
		return script.SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Hey ").
		ShowCharacterName().
		AddText(", I can send you to ").
		BlueText().ShowMap(_map.NavigationRoom).
		BlackText().AddText(" if you want to be a ").
		BlueText().AddText("Pirate").
		BlackText().AddText(". Do you want to go now?")
	return script.SendYesNoExit(l, c, m.String(), r.Warp, r.ComeBack, r.ComeBack)
}

func (r PirateStatue) ComeBack(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Come back to me if you decided to be a ").
		BlueText().AddText("Pirate").
		BlackText().AddText(".")
	return script.SendOk(l, c, m.String())
}

func (r PirateStatue) Warp(l logrus.FieldLogger, c script.Context) script.State {
	return script.WarpById(_map.NavigationRoom, 0)(l, c)
}