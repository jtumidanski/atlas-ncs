package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2111025 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      cm.sendNext ("2111025_SEEMS_TO_BE_MONITORING_SOMETHING")

      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC2111025 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2111025(cm: cm))
   }
   return (NPC2111025) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }