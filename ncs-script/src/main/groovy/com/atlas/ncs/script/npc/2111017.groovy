package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2111017 {
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
            if (cm.isQuestStarted(3339)) {
               int progress = cm.getQuestProgressInt(23339, 1)

               if (progress == 3) {
                  cm.sendGetText("2111017_SECRET_COMPARTMENT_SHOWS_UP")

               } else if (progress == 0) {
                  cm.setQuestProgress(23339, 1, 1)
                  cm.dispose()
               } else if (progress < 3) {
                  cm.setQuestProgress(23339, 1, 0)
                  cm.dispose()
               } else {
                  cm.warp(261000001, 1)
                  cm.dispose()
               }
            } else {
               if (cm.isQuestCompleted(3339)) {
                  cm.warp(261000001, 1)
               }

               cm.dispose()
            }
         } else if (status == 1) {
            if (cm.getText() == "my love Phyllia") {
               cm.setQuestProgress(23339, 1, 4)
               cm.warp(261000001, 1)
               cm.dispose()
            } else {
               cm.sendOk("2111017_WRONG")

               cm.dispose()
            }
         }
      }
   }
}

NPC2111017 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2111017(cm: cm))
   }
   return (NPC2111017) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }