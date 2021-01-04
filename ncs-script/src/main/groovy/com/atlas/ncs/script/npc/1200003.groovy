package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		Puro
	Map(s): 		Whale Between Lith harbor and Rien
	Description:
*/
class NPC1200003 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   String[] menu = ["Lith Harbor"]

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0 && status == 0) {
            cm.dispose()
            return
         } else if (mode == 0) {
            cm.sendNext("1200003_LET_ME_KNOW")
            cm.dispose()
            return
         }
         status++
         if (status == 0) {
            String display = ""
            for (def i = 0; i < menu.length; i++) {
               display += "\r\n#L" + i + "##b Lith Harbor (800 mesos)#k"
            }
            cm.sendSimple("Are you trying to leave Rien? Board this ship and I'll take you from #bRien#k to #bLith Harbor#k and back. for a #bfee of 800#k Mesos. Would you like to head over to Lith Harbor now? It'll take about a minute to get there.\r\n" + display)

         } else if (status == 1) {
            if (cm.getMeso() < 800) {
               cm.sendNext("1200003_NOT_ENOUGH_MESOS")
               cm.dispose()
            } else {
               cm.gainMeso(-800)
               cm.warp(200090070)
               cm.dispose()
            }
         }
      }
   }
}

NPC1200003 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1200003(cm: cm))
   }
   return (NPC1200003) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }