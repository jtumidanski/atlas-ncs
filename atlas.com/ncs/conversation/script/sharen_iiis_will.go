package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SharenIIIsWill is located in Sharenian - Underground Waterway (990000600)
type SharenIIIsWill struct {
}

func (r SharenIIIsWill) NPCId() uint32 {
	return npc.SharenIIIsWill
}

func (r SharenIIIsWill) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I fought the Rubian and I lost, and now I am imprisoned in the very gate that blocks my path, my body desecrated. However, my old clothing has holy power within. If you can return the clothing to my body, I should be able to open the gate. Please hurry! ").NewLine().
		AddText("- Sharen III ").NewLine().NewLine().
		AddText("P.S. I know this is rather picky of me, but can you please return the clothes to my body ").
		BlueText().AddText("bottom to top").
		BlackText().AddText("? Thank you for your services.")
	return SendOk(l, c, m.String())
}
