package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// SubwayExit is located in Line 3 Construction Site - B1 <Area 1> (103000900) Line 3 Construction Site - B1 <Area 2> (103000901) Line 3 Construction Site - B2 <Area 1> (103000903) Line 3 Construction Site - B2 <Area 2> (103000904) Line 3 Construction Site - B3 <Area 1> (103000906) Line 3 Construction Site - B3 <Area 2> (103000907) Line 3 Construction Site - B3 <Area 3> (103000908)
type SubwayExit struct {
}

func (r SubwayExit) NPCId() uint32 {
	return npc.Exit
}

func (r SubwayExit) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r SubwayExit) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This device is connected to outside.")
	return script.SendNext(l, span, c, m.String(), r.GiveUpAndLeave)
}

func (r SubwayExit) GiveUpAndLeave(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Are you going to give up and leave this place?")
	return script.SendNextPrevious(l, span, c, m.String(), r.Confirm, r.Hello)
}

func (r SubwayExit) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You'll have to start from scratch the next time you come in...")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, r.MaybeNextTime)
}

func (r SubwayExit) MaybeNextTime(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Alright, see you next time.")
	return script.SendOk(l, span, c, m.String())
}

func (r SubwayExit) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.SubwayTicketingBooth, 0)(l, span, c)
}
