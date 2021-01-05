package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC1092019 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int seagullProgress
   int seagullIdx = -1
   String[] seagullQuestion = ["One day, I went to the ocean and caught 62 Octopi for dinner. But then some kid came by and gave me 10 Octopi as a gift! How many Octopi do I have then, in total?"]
   String[] seagullAnswer = ["72"]

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
            if (!cm.isQuestStarted(6400)) {
               cm.sendOk("1092019_GO_BOTHER_SOMEBODY_ELSE")
               cm.dispose()
            } else {
               seagullProgress = cm.getQuestProgressInt(6400, 1)

               if (seagullProgress == 0) {
                  seagullIdx = Math.floor(Math.random() * seagullQuestion.length).intValue()
                  cm.sendNext("1092019_FIRST_QUESTION")
               } else if (seagullProgress == 1) {
                  cm.sendNext("1092019_NEXT_QUESTION")
               } else {
                  cm.sendNext("1092019_IMPRESSIVE")
               }
            }
         } else if (status == 1) {
            if (seagullProgress == 0) {
               cm.sendGetText(seagullQuestion[seagullIdx])
            } else if (seagullProgress == 1) {
               cm.sendNextPrev("1092019_TEST_OF_WILL")
            } else {
               cm.sendNextPrev("1092019_NOTIFY_US_USING_AIR_STRIKE")
            }
         } else if (status == 2) {
            if (seagullIdx > -1) {
               String answer = cm.getText()
               if (answer == seagullAnswer[seagullIdx]) {
                  cm.sendNext("1092019_INCREDIBLY_SMART")
                  cm.setQuestProgress(6400, 1, 1)
                  cm.dispose()
               } else {
                  cm.sendOk("1092019_NOT_RIGHT")
                  cm.dispose()
               }
            } else if (seagullProgress != 2) {
               cm.sendNextPrev("1092019_TRUE_PIRATE")
            } else {
               cm.sendNextPrev("1092019_PASSED")
               cm.dispose()
            }
         } else if (status == 3) {
            EventManager em = cm.getEventManager("4jaerial").orElseThrow()
            if (!em.startInstance(cm.getCharacterId())) {
               cm.sendOk("1092019_ANOTHER_PLAYER_TAKING_CHALLENGE")
            }

            cm.dispose()
         }
      }
   }
}

NPC1092019 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1092019(cm: cm))
   }
   return (NPC1092019) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }