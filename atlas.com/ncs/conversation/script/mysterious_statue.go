package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// MysteriousStatue is located in 
type MysteriousStatue struct {
}

func (r MysteriousStatue) NPCId() uint32 {
	return npc.MysteriousStatue
}

func (r MysteriousStatue) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r MysteriousStatue) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You feel a mysterious force surrounding this statue.")
	return SendNext(l, c, m.String(), r.ChooseDestination)
}

func (r MysteriousStatue) ChooseDestination(l logrus.FieldLogger, c Context) State {
	zones := 0
	if character.QuestStarted(l)(c.CharacterId, 2052) || character.QuestCompleted(l)(c.CharacterId, 2052) {
		zones = 1
	} else if character.QuestStarted(l)(c.CharacterId, 2053) || character.QuestCompleted(l)(c.CharacterId, 2053) {
		zones = 2
	} else if character.QuestStarted(l)(c.CharacterId, 2054) || character.QuestCompleted(l)(c.CharacterId, 2054) {
		zones = 3
	}
	if zones == 0 {
		return Exit()(l, c)
	}

	m := message.NewBuilder().
		AddText("Its power allows you to will yourself deep inside the forest.").NewLine()
	for i := 0; i < zones; i++ {
		m = m.OpenItem(i).BlueText().AddText(fmt.Sprintf("Deep Forest of Patience %d", i+1)).CloseItem().NewLine()
	}
	return SendListSelectionExit(l, c, m.String(), r.DestinationSelection, r.SeeYouNextTime)
}

func (r MysteriousStatue) DestinationSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Warp(_map.TheDeepForestOfPatienceStep1)
	case 1:
		return r.Warp(_map.TheDeepForestOfPatienceStep3)
	case 2:
		return r.Warp(_map.TheDeepForestOfPatienceStep5)
	}
	return nil
}

func (r MysteriousStatue) SeeYouNextTime(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Alright, see you next time.")
	return SendOk(l, c, m.String())
}

func (r MysteriousStatue) Warp(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return nil
	}
}
