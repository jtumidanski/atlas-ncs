package discrete

import (
	"atlas-ncs/buff"
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// RussellonsDesk is located in Alcadno Research Institute - Lab - Area B-1 (261020200)
type RussellonsDesk struct {
}

func (r RussellonsDesk) NPCId() uint32 {
	return npc.RussellonsDesk
}

func (r RussellonsDesk) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3314) && !character.HasItem(l, span)(c.CharacterId, item.RusselleonsPill) && RusselleonsPillUsed(l)(c.CharacterId) {
		if character.CanHold(l)(c.CharacterId, item.RusselleonsPill) {
			character.GainItem(l, span)(c.CharacterId, item.RusselleonsPill, 1)
			m := message.NewBuilder().AddText("You took the pills that were laying on the desk.")
			return script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
		} else {
			m := message.NewBuilder().AddText("You don't have a USE slot available to get Russellon's pills.")
			return script.SendOk(l, span, c, m.String(), script.AddSendTalkConfigurator(npc.SetSpeaker(npc.SpeakerCharacterLeft)))
		}
	}
	return script.Exit()(l, span, c)
}

func RusselleonsPillUsed(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return character.BuffSource(l)(characterId, buff.BuffHPRecovery) == item.RusselleonsPill
	}
}
