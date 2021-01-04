package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1061007 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      cm.sendYesNo("1061007_WOULD_YOU_LIKE_TO_LEAVE")
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         cm.warp(105040300, 0)
      }
      cm.dispose()
   }
}

NPC1061007 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1061007(cm: cm))
   }
   return (NPC1061007) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }