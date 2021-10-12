package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// PowerBForePerion is located in Victoria Road - Perion Street Corner (101040000)
type PowerBForePerion struct {
}

func (r PowerBForePerion) NPCId() uint32 {
	return npc.PowerBForePerion
}

func (r PowerBForePerion) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.IsLevel(l, span)(c.CharacterId, 20) {
		return r.UnderLevel20(l, span, c)
	} else {
		return r.ChooseRoom(l, span, c)
	}
}

func (r PowerBForePerion) UnderLevel20(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This training ground is available only for those under level 20.")
	return script.SendOk(l, span, c, m.String())
}

func (r PowerBForePerion) ChooseRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	format := "Training Center %d (%d/%d)"
	m := message.NewBuilder().
		AddText("Would you like to go into the Training Center?").NewLine()
	for i := 0; i < r.Rooms(); i++ {
		m = m.OpenItem(i).BlueText().AddText(fmt.Sprintf(format, i+1, _map.CharacterCount(l)(c.WorldId, c.ChannelId, r.TrainingMap()+uint32(i)), r.MaxInRoom())).NewLine()
	}
	return script.SendListSelection(l, span, c, m.String(), r.RoomSelection)
}

func (r PowerBForePerion) TrainingMap() uint32 {
	return 910220000
}

func (r PowerBForePerion) Rooms() int {
	return 5
}

func (r PowerBForePerion) MaxInRoom() uint32 {
	return 5
}

func (r PowerBForePerion) RoomSelection(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if selection < 0 || selection >= int32(r.Rooms()) {
			l.Errorf("Invalid selection.")
			return script.Exit()(l, span, c)
		}

		if _map.CharacterCount(l)(c.WorldId, c.ChannelId, c.MapId) >= r.MaxInRoom() {
			return r.RoomFull(l, span, c)
		}

		return script.Warp(r.TrainingMap()+uint32(selection))(l, span, c)
	}
}

func (r PowerBForePerion) RoomFull(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This training center is full.")
	return script.SendNext(l, span, c, m.String(), r.ChooseRoom)
}
