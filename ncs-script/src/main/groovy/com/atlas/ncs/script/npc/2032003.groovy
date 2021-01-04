package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

/*
	NPC Name: 		Lira
	Map(s): 		Adobis's Mission I : Breath of Lava <Level 2>
	Description: 	
*/
class NPC2032003 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 1) {
         cm.dispose()
      } else {
         status++

         if (status == 0) {
            cm.sendNext("2032003_CONGRATULATIONS")
         } else if (status == 1) {
            if (!cm.canHold(4031062)) {
               cm.sendOk("2032003_FREE_A_SLOT")
               cm.dispose()
               return
            }

            cm.sendNext("2032003_TIME_FOR_YOU_TO_HEAD_OFF")
         } else if (status == 2) {
            cm.gainItem(4031062, (short) 1)
            cm.gainExp(10000 * cm.getExpRate())
            cm.warp(211042300)

            cm.dispose()
         }
      }
   }
}

NPC2032003 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2032003(cm: cm))
   }
   return (NPC2032003) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }