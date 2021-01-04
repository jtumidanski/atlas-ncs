package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC11000 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      MapleShopProcessor.getInstance().sendShop(MapleShopFactory.getInstance().getShop(11000), cm.getClient())
      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {
   }
}

NPC11000 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC11000(cm: cm))
   }
   return (NPC11000) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }