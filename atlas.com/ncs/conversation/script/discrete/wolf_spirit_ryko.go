package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// WolfSpiritRyko is located in Victoria Road - Hall of Bowmen (100000205), Victoria Road - Hall of Magicians (101000005), Victoria Road - Hall of Warriors (102000005), Victoria Road - Hall of Thieves (103000009), The Nautilus - Training Room (120000105), Empress's Road - Knights Chamber (130000101), and Snow Island - Palace of the Master (140010111)

type WolfSpiritRyko struct {
}

func (r WolfSpiritRyko) NPCId() uint32 {
	return npc.WolfSpiritRyko
}

func (r WolfSpiritRyko) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("... I came from distant planes to assist the fight against the ").
		RedText().AddText("Black Magician").
		BlackText().AddText(". Right now I search my master, have you seen him?")
	return script.SendOk(l, span, c, m.String())
}
