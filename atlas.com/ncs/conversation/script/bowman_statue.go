package script

import (
	"atlas-ncs/character"
	"atlas-ncs/job"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// BowmanStatue is located in Victoria Road - Lith Harbor (104000000)
type BowmanStatue struct {
}

func (r BowmanStatue) NPCId() uint32 {
	return npc.BowmanStatue
}

func (r BowmanStatue) Initial(l logrus.FieldLogger, c Context) State {
	if !character.IsJob(l)(c.CharacterId, job.Beginner) {
		m := message.NewBuilder().AddText("You're much stronger now. Keep training!")
		return SendOk(l, c, m.String())
	}

	if !character.MeetsCriteria(l)(c.CharacterId, character.IsLevelCriteria(10), character.HasDexterityCriteria(25)) {
		m := message.NewBuilder().
			AddText("If you want to be a ").
			BlueText().AddText("Bowman").
			BlackText().AddText(", train yourself further until you reach ").
			BlueText().AddText("level 10 with 25 DEX").
			BlackText().AddText(".")
		return SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Hey ").
		ShowCharacterName().
		AddText(", I can send you to ").
		BlueText().ShowMap(_map.BowmanInstructionalSchool).
		BlackText().AddText(" if you want to be a ").
		BlueText().AddText("Bowman").
		BlackText().AddText(". Do you want to go now?")
	return SendYesNoExit(l, c, m.String(), WarpById(_map.BowmanInstructionalSchool, 0), r.ComeBack, r.ComeBack)
}

func (r BowmanStatue) ComeBack(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Come back to me if you decided to be a ").
		BlueText().AddText("Bowman").
		BlackText().AddText(".")
	return SendOk(l, c, m.String())
}
