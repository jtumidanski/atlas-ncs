package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9201101 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.getConfiguration().enableCustomNpcScript()) {
         cm.openShopNPC(9201101)
      } else {
         //cm.sendOk("The patrol in New Leaf City is always ready. No creatures are able to break through to the city.");
         cm.sendDefault()
      }

      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC9201101 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201101(cm: cm))
   }
   return (NPC9201101) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }