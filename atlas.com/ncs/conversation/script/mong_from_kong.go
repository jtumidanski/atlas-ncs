package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// MongFromKong is located in Victoria Road - Kerning City (103000000)
type MongFromKong struct {
}

func (r MongFromKong) NPCId() uint32 {
	return npc.MongFromKong
}

func (r MongFromKong) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r MongFromKong) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So, are you going to use the Internet Cafe? There is a fee to use the spaces there, that is ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 5000)).
		BlackText().AddText(". Are you going to enter the Cafe?")
	return SendYesNo(l, c, m.String(), r.Validate, Exit())
}

func (r MongFromKong) Validate(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 5000) {
		return r.NotEnough(l, c)
	}
	return r.Warp(l, c)
}

func (r MongFromKong) NotEnough(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh, you don't have the money, right? Sorry, I can't let you in.")
	return SendOk(l, c, m.String())
}

func (r MongFromKong) Warp(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -5000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
	}
	err = npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.KerningCityInternetCafe, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.KerningCityInternetCafe, c.NPCId)
	}
	return nil
}
