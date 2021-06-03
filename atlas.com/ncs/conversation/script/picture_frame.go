package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// PictureFrame is located in Magatia - Home of the Missing Alchemist (261000001)
type PictureFrame struct {
}

func (r PictureFrame) NPCId() uint32 {
	return npc.PictureFrame
}

func (r PictureFrame) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3311) {
		progress := character.QuestProgressInt(l)(c.CharacterId, 3311, 0)
		if progress == 4 {
			progress = 7
		} else {
			progress = 5
		}
		character.SetQuestProgress(l)(c.CharacterId, 3311, 0, uint32(progress))
		m := message.NewBuilder().AddText("This is a mug picture of Dr. De Lang. It seems he is adorning a locket with the emblem of the Alcadno academy, he is a retainer of the Alcadno society.")
		return SendOk(l, c, m.String(), AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
	}
	return Exit()(l, c)
}
