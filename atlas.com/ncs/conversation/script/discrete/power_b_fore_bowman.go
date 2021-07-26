package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"fmt"
	"github.com/sirupsen/logrus"
)

//TODO review need for this conversation

// PowerBForeBowman is located in Bowman Training Center - Bowman Training Center (910060000)
type PowerBForeBowman struct {
}

func (r PowerBForeBowman) NPCId() uint32 {
	return npc.PowerBForeBowman
}

func (r PowerBForeBowman) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.IsLevel(l)(c.CharacterId, 20) {
		return r.UnderLevel20(l, c)
	} else if quest.AnyActive(l)(c.CharacterId, 22515, 22516, 22517, 22518) {
		return r.GoToSpecial(l, c)
	} else {
		return r.ChooseRoom(l, c)
	}
}

func (r PowerBForeBowman) UnderLevel20(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This training ground is available only for those under level 20.")
	return script.SendOk(l, c, m.String())
}

func (r PowerBForeBowman) GoToSpecial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Would you like to go in the special Training Center?")
	return script.SendYesNo(l, c, m.String(), r.WarpSpecial, script.Exit())
}

func (r PowerBForeBowman) WarpSpecial(l logrus.FieldLogger, c script.Context) script.State {
	return script.Warp(r.TrainingSpecial())(l, c)
}

func (r PowerBForeBowman) ChooseRoom(l logrus.FieldLogger, c script.Context) script.State {
	format := "Training Center %d (%d/%d)"
	m := message.NewBuilder().
		AddText("Would you like to go into the Training Center?").NewLine()
	for i := 0; i < r.Rooms(); i++ {
		m = m.OpenItem(i).BlueText().AddText(fmt.Sprintf(format, i+1, _map.CharacterCount(l)(c.WorldId, c.ChannelId, r.TrainingMap()+uint32(i)), r.MaxInRoom())).NewLine()
	}
	return script.SendListSelection(l, c, m.String(), r.RoomSelection)
}

func (r PowerBForeBowman) TrainingSpecial() uint32 {
	return 910060100
}

func (r PowerBForeBowman) TrainingMap() uint32 {
	return 910060000
}

func (r PowerBForeBowman) Rooms() int {
	return 5
}

func (r PowerBForeBowman) MaxInRoom() uint32 {
	return 5
}

func (r PowerBForeBowman) RoomSelection(selection int32) script.StateProducer {
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

func (r PowerBForeBowman) RoomFull(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This training center is full.")
	return script.SendNext(l, c, m.String(), r.ChooseRoom)
}
