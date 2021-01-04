package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9000010 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      cm.sendNext("9000010_TRY_AGAIN_SOME_OTHER_TIME")

   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      }
      cm.warp(cm.getSavedLocation("EVENT"))
      cm.dispose()
   }
}

NPC9000010 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9000010(cm: cm))
   }
   return (NPC9000010) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }