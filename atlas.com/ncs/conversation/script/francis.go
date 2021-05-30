package script

import (
	"atlas-ncs/event"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"strconv"
)

// Francis is located in Hidden Street - Puppeteer's Cave (910510001)
type Francis struct {
}

func (r Francis) NPCId() uint32 {
	return npc.Francis
}

func (r Francis) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("What the... you don't belong here!")
	return SendNext(l, c, m.String(), r.StartEvent)
}

func (r Francis) StartEvent(l logrus.FieldLogger, c Context) State {
	event.SetProperty(l)("Puppeteer", "player", strconv.Itoa(int(c.CharacterId)))
	event.StartEvent(l)(c.CharacterId, "Puppeteer")
	return Exit()(l, c)
}
