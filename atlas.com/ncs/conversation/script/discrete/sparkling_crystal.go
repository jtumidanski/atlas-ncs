package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// SparklingCrystal is located in Hidden Street - The Other Dimension (108010101) Hidden Street - The Other Dimension (108010201) Hidden Street - The Other Dimension (108010301) Hidden Street - The Other Dimension (108010401) Shadow Zone - The Other Dimension (108010501)

type SparklingCrystal struct {
}

func (r SparklingCrystal) NPCId() uint32 {
	return npc.SparklingCrystal
}

func (r SparklingCrystal) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Would you like to leave?")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, script.Exit())
}

func (r SparklingCrystal) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	destination := c.MapId
	switch c.MapId {
	case _map.TheOtherDimension1:
		destination = _map.SleepyDungeonV
	case _map.TheOtherDimension2:
		destination = _map.TheForestOfEvilII
	case _map.TheOtherDimension3:
		destination = _map.AntTunnelPark
	case _map.TheOtherDimension4:
		destination = _map.MonkeySwampII
	case _map.TheOtherDimension5:
		destination = _map.TheCaveOfEvilEyeII
	}
	return script.WarpById(destination, 0)(l, span, c)
}
