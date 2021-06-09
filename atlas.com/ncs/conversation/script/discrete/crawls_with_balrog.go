package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// CrawlsWithBalrog is located in Altaire Camp - Tent House 2 (300000002)
type CrawlsWithBalrog struct {
}

func (r CrawlsWithBalrog) NPCId() uint32 {
	return npc.CrawlsWithBalrog
}

func (r CrawlsWithBalrog) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Zzzzz...")
	return script.SendOk(l, c, m.String())
}
