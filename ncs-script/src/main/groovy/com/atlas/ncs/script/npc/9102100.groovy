package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9102100 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (status == 0 && mode == 0) {
         cm.sendNext("9102100_COVERED_IN_GRASS")

         cm.dispose()
         return
      }
      if (mode == 1) {
         status++
      } else {
         status--
      }
      if (status == 0) {
         if (cm.isQuestStarted(4646)) {
            if (cm.haveItem(4031921)) {
               cm.sendNext("9102100_POOP_WAS_IN_THERE")

               cm.dispose()
            } else {
               cm.sendYesNo("9102100_PULL_IT_OUT")

            }
         } else {
            cm.sendOk("9102100_COULD_NOT_FIND_ANYTHING")

            cm.dispose()
         }
      } else if (status == 1) {
         cm.sendNext("9102100_THIS_NOTE")

         cm.gainItem(4031921, (short) 1)
         cm.dispose()
      }
   }
}

NPC9102100 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9102100(cm: cm))
   }
   return (NPC9102100) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }