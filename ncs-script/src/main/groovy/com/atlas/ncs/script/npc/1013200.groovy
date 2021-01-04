package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1013200 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (!cm.isQuestStarted(22015)) {
         cm.sendOk("1013200_TOO_FAR_FROM_PIG")
      } else {
         cm.gainItem(4032449, true)
         cm.forceCompleteQuest(22015)
         cm.sendPinkText("1013200_PIGLET_RESCUED")
      }
      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC1013200 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1013200(cm: cm))
   }
   return (NPC1013200) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }