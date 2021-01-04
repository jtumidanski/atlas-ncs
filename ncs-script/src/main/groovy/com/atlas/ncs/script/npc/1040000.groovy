package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1040000 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0 && type > 0) {
            cm.dispose()
            return
         }
         if (mode == 1)
            status++
         else
            status--

         if(status == 0) {
            if(cm.isQuestStarted(28177) && !cm.haveItem(4032479)) {
               if(cm.canHold(4032479)) {
                  cm.gainItem(4032479, (short) 1)
                  cm.sendOk("1040000_NOT_THE_SUSPECT_YOU_SEEK")
               } else {
                  cm.sendOk("1040000_MAKE_A_SLOT_AVAILABLE")
               }
            } else {
               cm.sendOk("1040000_ZZZZ")
            }

            cm.dispose()
         }
      }
   }
}

NPC1040000 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1040000(cm: cm))
   }
   return (NPC1040000) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }