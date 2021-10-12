package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// PictureFrame is located in Magatia - Home of the Missing Alchemist (261000001)
type PictureFrame struct {
}

func (r PictureFrame) NPCId() uint32 {
	return npc.PictureFrame
}

func (r PictureFrame) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3311) {
		progress := quest.ProgressInt(l)(c.CharacterId, 3311, 0)
		if progress == 4 {
			progress = 7
		} else {
			progress = 5
		}
		quest.SetProgress(l)(c.CharacterId, 3311, 0, uint32(progress))
		m := message.NewBuilder().AddText("This is a mug picture of Dr. De Lang. It seems he is adorning a locket with the emblem of the Alcadno academy, he is a retainer of the Alcadno society.")
		return script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
	}
	return script.Exit()(l, span, c)
}
