package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2012033 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   String harpNote = 'B'
   String[] harpSounds = ["do", "re", "mi", "pa", "sol", "la", "si"]
   String harpSong = "CCGGAAGFFEEDDC|GGFFEED|GGFFEED|CCGGAAGFFEEDDC|"

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
            cm.playSoundInMap("orbis/" + harpSounds[cm.getNpcId() - 2012027])

            if (cm.isQuestStarted(3114)) {
               int idx = -1 * cm.getQuestProgressInt(3114)

               if (idx > -1) {
                  String nextNote = harpSong[idx]

                  if (harpNote != nextNote) {
                     cm.setQuestProgress(3114, 0)
                     cm.showEffect("quest/party/wrong_kor")
                     cm.playSound("Party1/Failed")
                     cm.sendPinkText("2012027_MISSED_NOTE")
                  } else {
                     nextNote = harpSong[idx + 1]

                     if (nextNote == '|') {
                        idx++

                        if (idx == 45) {     // finished lullaby
                           cm.sendPinkText("2012027_TWINKLE_TWINKLE")
                           cm.setQuestProgress(3114, 42)
                           cm.showEffect("quest/party/clear")
                           cm.playSound("Party1/Clear")
                           cm.dispose()
                           return
                        } else {
                           if (idx == 14) {
                              cm.sendPinkText("2012027_TWINKLE_TWINKLE")
                           } else if (idx == 22) {
                              cm.sendPinkText("2012027_UP_ABOVE")
                           } else if (idx == 30) {
                              cm.sendPinkText("2012027_LIKE_A_DIAMOND")
                           }
                        }
                     }

                     cm.setQuestProgress(3114, -1 * (idx + 1))
                  }
               }
            }

            cm.dispose()
         }
      }
   }
}

NPC2012033 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2012033(cm: cm))
   }
   return (NPC2012033) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }