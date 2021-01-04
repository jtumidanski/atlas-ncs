package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPCThiefPassword {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1 || (mode == 0 && status == 0)) {
         cm.dispose()
         return
      } else if (mode == 0) {
         status--
      } else {
         status++
      }



      if (status == 0) {
         cm.sendGetText("ThiefPassword_SUSPICIOUS_VOICE")
      } else if (status == 1) {
         if (cm.getText() == "Open Sesame") {
            if (cm.isQuestCompleted(3925)) {
               cm.warp(260010402, 1)
            } else {
               cm.sendPinkText("CORRECT_BUT_FORCES_BLOCKING")
            }
            cm.dispose()
         } else {
            cm.sendOk("ThiefPassword_WRONG")
         }
      } else if (status == 2) {
         cm.dispose()
      }
   }
}

NPCThiefPassword getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPCThiefPassword(cm: cm))
   }
   return (NPCThiefPassword) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }