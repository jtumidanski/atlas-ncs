package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Cliff is located in Hidden Street - Happyville (209000000)
type Cliff struct {
}

func (r Cliff) NPCId() uint32 {
	return npc.Cliff
}

func (r Cliff) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you see a bunch of snowmen standing around there? Go talk to one of them, and it'll take you to the famous Christmas tree here that is just humongous. The tree can be decorated using various kinds of ornaments. What do you think? Sounds fun, right?")
	return SendNext(l, c, m.String(), r.TreeRules)
}

func (r Cliff) TreeRules(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Only 6 can be at the place where the tree is at once, and you can't ").
		BlueText().AddText("trade or open store").
		BlackText().AddText(" there. The ornaments that you drop can only be picked back up by yourself, so don't worry about losing your ornaments here.")
	return SendNextPrevious(l, c, m.String(), r.NoExpiration, r.Initial)
}

func (r Cliff) NoExpiration(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Of course, the items that are dropped in there will never disappear. Once you get out of there through the snowman that's inside, all the items you've dropped at that map will come back to you, so you won't have to pick all those items up before leaving the place. Isn't that sweet?")
	return SendNextPrevious(l, c, m.String(), r.BuyOrnaments, r.TreeRules)
}

func (r Cliff) BuyOrnaments(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Well then, go see ").
		ShowNPC(npc.Rudi).
		AddText(", buy some Christmas ornaments there, and then decorate the tree with those~ Oh yeah! The biggest, the most beautiful ornament cannot be bought from him. It's probably ... taken by a monster ... huh huh ..")
	return SendPrevious(l, c, m.String(), r.NoExpiration)
}
