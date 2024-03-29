package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Asia is located in Neo City - <Year 2503> Air Battleship Bow (240070600)
type Asia struct {
}

func (r Asia) NPCId() uint32 {
	return npc.Asia
}

func (r Asia) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 3749) {
		return r.FilledWithDespair(l, span, c)
	} else {
		return r.SeeMySister(l, span, c)
	}
}

func (r Asia) SeeMySister(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("We've already located the enemy's ultimate weapon! Follow along the ship's bow area ahead and you will find my sister ").
		BlueText().ShowNPC(npc.Ashura).
		BlackText().AddText(". Report to her for further instructions on the mission.")
	return script.SendOk(l, span, c, m.String())
}

func (r Asia) FilledWithDespair(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("The future is filled with despair.")
	return script.SendOk(l, span, c, m.String())
}
