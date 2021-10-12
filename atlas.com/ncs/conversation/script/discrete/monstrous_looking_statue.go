package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/monster"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// MonstrousLookingStatue is located in Hidden Street - Puppeteer's Secret Passage (910510100)
type MonstrousLookingStatue struct {
}

func (r MonstrousLookingStatue) NPCId() uint32 {
	return npc.MonstrousLookingStatue
}

func (r MonstrousLookingStatue) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.ReadyToFace(l, span, c)
}

func (r MonstrousLookingStatue) ReadyToFace(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ahead awaits the Master himself. Are you ready to face him?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, script.Exit())
}

func (r MonstrousLookingStatue) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if _map.CharacterCount(l)(c.WorldId, c.ChannelId, _map.PuppeteersCave) > 0 {
		return r.SomeoneAlreadyChallengine(l, span, c)
	}
	monster.SpawnMonster(l)(c.WorldId, c.ChannelId, _map.PuppeteersCave, monster.Puppeteer, 95, 200)
	return script.WarpById(_map.PuppeteersCave, 0)(l, span, c)
}

func (r MonstrousLookingStatue) SomeoneAlreadyChallengine(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Someone is already challenging the Master. Try again later.")
	return script.SendOk(l, span, c, m.String())
}
