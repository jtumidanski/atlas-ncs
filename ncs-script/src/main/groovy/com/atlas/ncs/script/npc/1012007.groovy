package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1012007 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.haveItem(4031035)) {
         cm.sendNext("1012007_BROTHERS_LETTER")
      } else {
         cm.sendOk("1012007_BROTHER_TOLD_ME")
         cm.dispose()
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 1) {
         cm.dispose()
      } else {
         if (cm.petCount() == 0) {
            cm.sendNextPrev("1012007_DID_YOU_REALLY")
         } else {
            cm.gainItem(4031035, (short) -1)
            cm.gainCloseness(2)
            cm.sendNextPrev("1012007_WHAT_DO_YOU_THINK")
         }
         cm.dispose()
      }
   }
}

NPC1012007 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1012007(cm: cm))
   }
   return (NPC1012007) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }