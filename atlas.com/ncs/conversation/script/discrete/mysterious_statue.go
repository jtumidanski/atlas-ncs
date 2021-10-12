package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// MysteriousStatue is located in 
type MysteriousStatue struct {
}

func (r MysteriousStatue) NPCId() uint32 {
	return npc.MysteriousStatue
}

func (r MysteriousStatue) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r MysteriousStatue) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You feel a mysterious force surrounding this statue.")
	return script.SendNext(l, span, c, m.String(), r.ChooseDestination)
}

func (r MysteriousStatue) ChooseDestination(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	zones := 0
	if quest.IsStarted(l)(c.CharacterId, 2052) || quest.IsCompleted(l)(c.CharacterId, 2052) {
		zones = 1
	} else if quest.IsStarted(l)(c.CharacterId, 2053) || quest.IsCompleted(l)(c.CharacterId, 2053) {
		zones = 2
	} else if quest.IsStarted(l)(c.CharacterId, 2054) || quest.IsCompleted(l)(c.CharacterId, 2054) {
		zones = 3
	}
	if zones == 0 {
		return script.Exit()(l, span, c)
	}

	m := message.NewBuilder().
		AddText("Its power allows you to will yourself deep inside the forest.").NewLine()
	for i := 0; i < zones; i++ {
		m = m.OpenItem(i).BlueText().AddText(fmt.Sprintf("Deep Forest of Patience %d", i+1)).CloseItem().NewLine()
	}
	return script.SendListSelectionExit(l, span, c, m.String(), r.DestinationSelection, r.SeeYouNextTime)
}

func (r MysteriousStatue) DestinationSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return script.WarpById(_map.TheDeepForestOfPatienceStep1, 0)
	case 1:
		return script.WarpById(_map.TheDeepForestOfPatienceStep3, 0)
	case 2:
		return script.WarpById(_map.TheDeepForestOfPatienceStep5, 0)
	}
	return nil
}

func (r MysteriousStatue) SeeYouNextTime(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Alright, see you next time.")
	return script.SendOk(l, span, c, m.String())
}
