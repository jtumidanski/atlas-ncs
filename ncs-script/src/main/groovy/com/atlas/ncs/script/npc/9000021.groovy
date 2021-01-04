package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9000021 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 0) {
         cm.dispose()
      } else {
         if (mode == 0 && type > 0) {
            cm.dispose()
            return
         }

         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            cm.sendNext("9000021_HELLO")

         } else if (status == 1) {
            cm.sendNext("9000021_SEQUENTIAL_BOSS_FIGHTS")

         } else if (status == 2) {
            cm.sendAcceptDecline("9000021_IF_YOU_FEEL_POWERFUL_ENOUGH")

         } else if (status == 3) {
            cm.sendOk("9000021_VERY_WELL")

         } else if (status == 4) {
            cm.saveLocation("BOSSPQ")
            cm.warp(970030000, "out00")
            cm.dispose()
         }
      }
   }
}

NPC9000021 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9000021(cm: cm))
   }
   return (NPC9000021) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }