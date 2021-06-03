package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// EllinForestMilepost is located in Forest of Poison Haze - Forest of Haze (930000300)
type EllinForestMilepost struct {
}

func (r EllinForestMilepost) NPCId() uint32 {
	return npc.EllinForestMilepost
}

func (r EllinForestMilepost) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to exit this instance? Your party members may have to abandon it as well, so take that in mind.")
	return SendYesNo(l, c, m.String(), r.Yes, Exit())
}

func (r EllinForestMilepost) Yes(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.PurpleStoneOfMagic, -1)
	character.GainItem(l)(c.CharacterId, item.MonsterMarble, -1)
	character.GainItem(l)(c.CharacterId, item.PurificationMarble, -1)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.OuterForestExit, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.OuterForestExit, c.NPCId)
	}
	return Exit()(l, c)
}
