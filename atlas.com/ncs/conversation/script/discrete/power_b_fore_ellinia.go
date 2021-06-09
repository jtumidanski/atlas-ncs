package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// PowerBForeEllinia is located in Victoria Road - The Field South of Ellinia (100050000)
type PowerBForeEllinia struct {
}

func (r PowerBForeEllinia) NPCId() uint32 {
	return npc.PowerBForeEllinia
}

func (r PowerBForeEllinia) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.IsLevel(l)(c.CharacterId, 20) {
		return r.UnderLevel20(l, c)
	} else {
		return r.ChooseRoom(l, c)
	}
}

func (r PowerBForeEllinia) UnderLevel20(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This training ground is available only for those under level 20.")
	return script.SendOk(l, c, m.String())
}

func (r PowerBForeEllinia) ChooseRoom(l logrus.FieldLogger, c script.Context) script.State {
	format := "Training Center %d (%d/%d)"
	m := message.NewBuilder().
		AddText("Would you like to go into the Training Center?").NewLine()
	for i := 0; i < r.Rooms(); i++ {
		m = m.OpenItem(i).BlueText().AddText(fmt.Sprintf(format, i+1, _map.CharacterCount(l)(c.WorldId, c.ChannelId, r.TrainingMap()+uint32(i)), r.MaxInRoom())).NewLine()
	}
	return script.SendListSelection(l, c, m.String(), r.RoomSelection)
}

func (r PowerBForeEllinia) TrainingMap() uint32 {
	return 910120000
}

func (r PowerBForeEllinia) Rooms() int {
	return 5
}

func (r PowerBForeEllinia) MaxInRoom() uint32 {
	return 5
}

func (r PowerBForeEllinia) RoomSelection(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		if selection < 0 || selection >= int32(r.Rooms()) {
			l.Errorf("Invalid selection.")
			return script.Exit()(l, c)
		}

		if _map.CharacterCount(l)(c.WorldId, c.ChannelId, c.MapId) >= r.MaxInRoom() {
			return r.RoomFull(l, c)
		}

		return script.Warp(r.TrainingMap()+uint32(selection))(l, c)
	}
}

func (r PowerBForeEllinia) RoomFull(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This training center is full.")
	return script.SendNext(l, c, m.String(), r.ChooseRoom)
}