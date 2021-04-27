package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DarkLord is located in Maple Road : Split Road of Destiny (1020000)
type DarkLord struct {
}

func (r DarkLord) NPCId() uint32 {
	return 10203
}

func (r DarkLord) Initial() State {
	return r.ThiefIntroduction
}

func (r DarkLord) ThiefIntroduction(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Thieves are a perfect blend of luck, dexterity, and power that are adept at the surprise attacks against helpless enemies. A high level of avoidability and speed allows Thieves to attack enemies from various angles.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.Demo)
}

func (r DarkLord) Demo(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Thief?")
	conversation.SendYesNo(m.String())
	return YesNo(GenericExit, r.DoDemo, r.SeeMeAgain)
}

func (r DarkLord) DoDemo(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	npc.Processor(l).LockUI()
	npc.Processor(l).Warp(c.WorldId, c.ChannelId, c.CharacterId, 1020400, 0)
	return nil
}

func (r DarkLord) SeeMeAgain(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Thief, come see me again.")
	conversation.SendNext(m.String())
	return nil
}
