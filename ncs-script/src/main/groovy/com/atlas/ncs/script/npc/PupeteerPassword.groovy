package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NpcPuppeteerPassword {
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
         if (cm.isQuestStarted(21728)) {
            cm.sendOk("PupeteerPassword_BETTER_RETURN")

            cm.setQuestProgress(21728, 21761, 0)
            cm.dispose()
            return
         }

         cm.sendGetText("PupeteerPassword_SUSPICIOUS_VOICE")

      } else if (status == 1) {
         if (cm.getText() == "Francis is a genius Puppeteer!") {
            if (cm.isQuestStarted(20730) && cm.getQuestProgressInt(20730, 9300285) == 0) {
               cm.warp(910510001, 1)
            } else if (cm.isQuestStarted(21731) && cm.getQuestProgressInt(21731, 9300346) == 0) {
               cm.warp(910510001, 1)
            } else {
               cm.sendPinkText("CORRECT_BUT_FORCES_BLOCKING")
            }

            cm.dispose()
         } else {
            cm.sendOk("PupeteerPassword_WRONG")

         }
      } else if (status == 2) {
         cm.dispose()
      }
   }
}

NpcPuppeteerPassword getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NpcPuppeteerPassword(cm: cm))
   }
   return (NpcPuppeteerPassword) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }