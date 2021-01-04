package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC10200 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      cm.sendNext("10200_BOWMAN_INTRO")
   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode != 1) {
         if (mode == 0) {
            cm.sendNext("10200_DEMO_NOTE")
         }
         cm.dispose()
         return
      }
      if (status == 0) {
         cm.sendYesNo("10200_DEMO_PROMPT")
      } else if (status == 1) {
         cm.lockUI()
         cm.warp(1020300, 0)
         cm.dispose()
      }
   }
}

NPC10200 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC10200(cm: cm))
   }
   return (NPC10200) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }