package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Simon is located in Hidden Street - Happyville (209000000) and Hidden Street - Shalom Temple (681000000)
type Simon struct {
}

func (r Simon) NPCId() uint32 {
	return npc.Simon
}

func (r Simon) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.Happyville {
		m := message.NewBuilder().
			AddText("The Shalom Temple is unlike any other place in Happyville, would you like to head to ").
			BlueText().AddText("Shalom Temple").
			BlackText().AddText("?")
		return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.ShalomTemple,0), r.LetMeKnow)
	} else if c.MapId == _map.ShalomTemple {
		m := message.NewBuilder().
			AddText("Would you like to head back to Happyville?")
		return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.Happyville, 0), r.LetMeKnow)
	}
	return script.Exit()(l, span, c)
}

func (r Simon) LetMeKnow(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Let me know if you've changed your mind!")
	return script.SendOk(l, span, c, m.String())
}
