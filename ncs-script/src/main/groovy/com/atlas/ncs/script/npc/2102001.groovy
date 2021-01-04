package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2102001 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   def start() {
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (status == 0) {
         cm.sendYesNo("2102001_WISH_TO_LEAVE_THE_GENIE")

         status++
      } else {
         if (mode < 1) {
            cm.dispose()
         } else {
            if (status == 1) {
               cm.sendNext("2102001_SEE_YOU_NEXT_TIME")

               status++
            } else if (status == 2) {
               cm.warp(260000100, 0)
               cm.dispose()
            }
         }
      }
   }
}

NPC2102001 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2102001(cm: cm))
   }
   return (NPC2102001) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }