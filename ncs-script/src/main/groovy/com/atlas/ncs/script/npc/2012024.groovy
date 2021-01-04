package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2012024 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   def start() {
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (status == 0) {
         cm.sendYesNo("2012024_LEAVE_THE_GENIE")
         status++
      } else {
         if (mode < 1) {
            cm.dispose()
         } else {
            if (status == 1) {
               cm.sendNext("2012024_ALRIGHT")
               status++
            } else if (status == 2) {
               cm.warp(200000151, 0)
               cm.dispose()
            }
         }
      }
   }
}

NPC2012024 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2012024(cm: cm))
   }
   return (NPC2012024) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }