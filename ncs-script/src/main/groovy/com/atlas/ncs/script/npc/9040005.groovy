package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9040005 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      cm.sendYesNo("9040005_WOULD_YOU_LIKE_TO_EXIT")

   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         cm.warp(101030104)
      }
      cm.dispose()
   }
}

NPC9040005 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9040005(cm: cm))
   }
   return (NPC9040005) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }