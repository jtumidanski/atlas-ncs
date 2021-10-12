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

// AlcandosCabinet is located in Hidden Street - Dark Lab (926120000)
type AlcandosCabinet struct {
}

func (r AlcandosCabinet) NPCId() uint32 {
	return npc.AlcandosCabinet
}

func (r AlcandosCabinet) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3309) && !character.HasItem(l)(c.CharacterId, item.SecretDocument) {
		if character.CanHold(l)(c.CharacterId, item.SecretDocument) {
			character.GainItem(l)(c.CharacterId, item.SecretDocument, 1)
		} else {
			m := message.NewBuilder().AddText("Have a ETC slot available to get the Alcadno's secret document.")
			return script.SendOk(l, span, c, m.String())
		}
	}
	return script.Exit()(l, span, c)
}
