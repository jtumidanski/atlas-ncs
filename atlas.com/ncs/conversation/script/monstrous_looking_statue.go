package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/monster"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MonstrousLookingStatue is located in Hidden Street - Puppeteer's Secret Passage (910510100)
type MonstrousLookingStatue struct {
}

func (r MonstrousLookingStatue) NPCId() uint32 {
	return npc.MonstrousLookingStatue
}

func (r MonstrousLookingStatue) Initial(l logrus.FieldLogger, c Context) State {
	return r.ReadyToFace(l, c)
}

func (r MonstrousLookingStatue) ReadyToFace(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Ahead awaits the Master himself. Are you ready to face him?")
	return SendYesNo(l, c, m.String(), r.Validate, Exit())
}

func (r MonstrousLookingStatue) Validate(l logrus.FieldLogger, c Context) State {
	if _map.CharacterCount(l)(c.WorldId, c.ChannelId, _map.PuppeteersCave) > 0 {
		return r.SomeoneAlreadyChallengine(l, c)
	}
	monster.SpawnMonster(l)(c.WorldId, c.ChannelId, _map.PuppeteersCave, monster.Puppeteer, 95, 200)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.PuppeteersCave, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.PuppeteersCave, c.NPCId)
	}
	return nil
}

func (r MonstrousLookingStatue) SomeoneAlreadyChallengine(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Someone is already challenging the Master. Try again later.")
	return SendOk(l, c, m.String())
}
