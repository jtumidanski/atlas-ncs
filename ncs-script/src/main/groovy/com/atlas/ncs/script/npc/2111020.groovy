package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2111020 {
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
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            if (cm.isQuestStarted(3345)) {
               int progress = cm.getQuestProgressInt(3345)

               if (progress == 0) {
                  cm.setQuestProgress(3345, 1)
                  cm.dispose()
               } else if (progress < 4) {
                  cm.setQuestProgress(3345, 0)
                  cm.dispose()
               } else {
                  cm.dispose()
               }
            }
         }
      }
   }
}

NPC2111020 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2111020(cm: cm))
   }
   return (NPC2111020) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }