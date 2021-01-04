package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9270017 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   def start() {
      cm.sendYesNo("9270017_TAKING_OFF_SOON")

   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode != 1) {
         if (mode == 0) {
            cm.sendOk("9270017_HOLD_ON")

         }
         cm.dispose()
         return
      }
      status++
      if (status == 1) {
         cm.sendNext("9270017_NOT_REFUNDABLE")

      } else if (status == 2) {
         cm.warp(103000000)
         cm.dispose()
      }
   }
}

NPC9270017 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9270017(cm: cm))
   }
   return (NPC9270017) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }