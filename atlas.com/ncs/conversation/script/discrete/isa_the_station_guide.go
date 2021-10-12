package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// IsaTheStationGuide is located in Orbis - Orbis Station Enterence (200000100)
type IsaTheStationGuide struct {
}

func (r IsaTheStationGuide) NPCId() uint32 {
	return npc.IsaTheStationGuide
}

func (r IsaTheStationGuide) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Orbis Station has lots of platforms available to choose from. You need to choose the one that'll take you to the destination of your choice. Which platform will you take?").NewLine().
		OpenItem(0).BlueText().AddText("The platform to the ship that heads to Ellinia.").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("The platform to the train that heads to Ludibrium.").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("The platform to the bird that heads to Leafre.").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("The platform to Hak that heads to Mu Lung.").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("The platform to Genie that heads to Ariant.").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("The platform to the ship that heads to Ereve.").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r IsaTheStationGuide) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.SendYouTo(_map.StationTunnelToEllinia)
	case 1:
		return r.SendYouTo(_map.StationPathwayLudibrium)
	case 2:
		return r.SendYouTo(_map.CabinPathToLeafre)
	case 3:
		return r.SendYouTo(_map.CabinPathToMuLung)
	case 4:
		return r.SendYouTo(_map.StationTunnelToAriant)
	case 5:
		return r.SendYouTo(_map.StationHall)
	}
	return nil
}

func (r IsaTheStationGuide) SendYouTo(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText("Ok ").
			ShowCharacterName().AddText(", I will send you to the platform for ").
			BlueText().ShowMap(mapId).
			BlackText().AddText(".")
		return script.SendNext(l, span, c, m.String(), r.Warp(mapId))
	}
}

func (r IsaTheStationGuide) Warp(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		return script.WarpByName(mapId, "west00")(l, span, c)
	}
}
