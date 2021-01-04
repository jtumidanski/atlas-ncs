package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2060100 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if(cm.isQuestStarted(6301)) {
         if (cm.haveItem(4000175)) {
            cm.gainItem(4000175, (short) -1)
            cm.warp(923000000, 0)
         } else {
            cm.sendOk("2060100_POSSESS_ONE_MINIATURE_PIANUS")

         }
      } else {
         cm.sendOk("2060100_DONT_FOOL_AROUND")

      }

      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC2060100 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2060100(cm: cm))
   }
   return (NPC2060100) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }