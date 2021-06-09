package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// CrumblingStatue is located in Hidden Street - The Deep Forest of Patience <Step 1> (105040310) Hidden Street - The Deep Forest of Patience <Step 2> (105040311) Hidden Street - The Deep Forest of Patience <Step 3> (105040312) Hidden Street - The Deep Forest of Patience <Step 4> (105040313) Hidden Street - The Deep Forest of Patience <Step 5> (105040314) Hidden Street - The Deep Forest of Patience <Step 6> (105040315) Hidden Street - The Deep Forest of Patience <Step 7> (105040316)
type CrumblingStatue struct {
}

func (r CrumblingStatue) NPCId() uint32 {
	return npc.CrumblingStatue
}

func (r CrumblingStatue) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Would you like to leave?")
	return script.SendYesNo(l, c, m.String(), script.WarpById(_map.Sleepywood, 0), script.Exit())
}