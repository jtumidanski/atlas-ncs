package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MetalBucketSnowman is located in Hidden Street - Happyville (209000000)
type MetalBucketSnowman struct {
}

func (r MetalBucketSnowman) NPCId() uint32 {
	return npc.MetalBucketSnowman
}

func (r MetalBucketSnowman) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("We have a beautiful christmas tree.").NewLine().
		AddText("Do you want to see/decorate it?")
	return script.SendYesNo(l, c, m.String(), script.WarpById(_map.TheHillOfChristmas2, 0), script.Exit())
}
