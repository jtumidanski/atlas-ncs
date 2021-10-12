package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// CenterOfTheMagicPentagram is located in Hidden Street - Black Magician's Lab (261040000)
type CenterOfTheMagicPentagram struct {
}

func (r CenterOfTheMagicPentagram) NPCId() uint32 {
	return npc.CenterOfTheMagicPentagram
}

func (r CenterOfTheMagicPentagram) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3345) {
		progress := quest.ProgressInt(l)(c.CharacterId, 3345, 0)

		if progress == 3 && character.HasItem(l, span)(c.CharacterId, item.MagicStoneOfHumility) && character.HasItem(l, span)(c.CharacterId, item.MagicStoneOfHonesty) && character.HasItem(l, span)(c.CharacterId, item.MagicStoneOfTrust) {
			quest.SetProgress(l)(c.CharacterId, 3345, 0, 4)
			character.GainItem(l, span)(c.CharacterId, item.MagicStoneOfHumility, -1)
			character.GainItem(l, span)(c.CharacterId, item.MagicStoneOfHonesty, -1)
			character.GainItem(l, span)(c.CharacterId, item.MagicStoneOfTrust, -1)

			m := message.NewBuilder().AddText("(As you place the shards a light shines over the circle, repelling whatever omens were brewing inside the artifact.)")
			return script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
		} else if progress < 4 {
			quest.SetProgress(l)(c.CharacterId, 3345, 0, 0)
		}
	}
	return script.Exit()(l, span, c)
}
