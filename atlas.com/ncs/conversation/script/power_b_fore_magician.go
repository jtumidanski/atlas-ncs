package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// PowerBForeMagician is located in Victoria Road - Magician Training Center (910120000)
type PowerBForeMagician struct {
}

func (r PowerBForeMagician) NPCId() uint32 {
	return npc.PowerBForeMagician
}

func (r PowerBForeMagician) Initial(l logrus.FieldLogger, c Context) State {
	if character.IsLevel(l)(c.CharacterId, 20) {
		return r.UnderLevel20(l, c)
	} else {
		return r.ChooseRoom(l, c)
	}
}

func (r PowerBForeMagician) UnderLevel20(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("This training ground is available only for those under level 20.")
	return SendOk(l, c, m.String())
}

func (r PowerBForeMagician) Warp(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}

func (r PowerBForeMagician) ChooseRoom(l logrus.FieldLogger, c Context) State {
	format := "Training Center %d (%d/%d)"
	m := message.NewBuilder().
		AddText("Would you like to go into the Training Center?").NewLine()
	for i := 0; i < r.Rooms(); i++ {
		m = m.OpenItem(i).BlueText().AddText(fmt.Sprintf(format, i+1, _map.CharacterCount(l)(c.WorldId, c.ChannelId, r.TrainingMap()+uint32(i)), r.MaxInRoom())).NewLine()
	}
	return SendListSelection(l, c, m.String(), r.RoomSelection)
}

func (r PowerBForeMagician) TrainingMap() uint32 {
	return 910120000
}

func (r PowerBForeMagician) Rooms() int {
	return 5
}

func (r PowerBForeMagician) MaxInRoom() uint32 {
	return 5
}

func (r PowerBForeMagician) RoomSelection(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if selection < 0 || selection >= int32(r.Rooms()) {
			l.Errorf("Invalid selection.")
			return Exit()(l, c)
		}

		if _map.CharacterCount(l)(c.WorldId, c.ChannelId, c.MapId) >= r.MaxInRoom() {
			return r.RoomFull(l, c)
		}

		return r.Warp(r.TrainingMap()+uint32(selection))(l, c)
	}
}

func (r PowerBForeMagician) RoomFull(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("This training center is full.")
	return SendNext(l, c, m.String(), r.ChooseRoom)
}