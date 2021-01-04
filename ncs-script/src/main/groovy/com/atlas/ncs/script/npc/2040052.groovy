package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2040052 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int[] questId = [3615, 3616, 3617, 3618, 3630, 3633, 3639, 3920]
   int[] questItem = [4031235, 4031236, 4031237, 4031238, 4031270, 4031280, 4031298, 4031591]
   int counter = 0
   String books
   int i

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 0) {
            if (counter == 0) {
               books = ""
               for (i = 0; i < questId.length; i++) {
                  if (cm.isQuestCompleted(questId[i])) {
                     counter += 1
                     books += "\r\n#v" + questItem[i] + "# #b#t" + questItem[i] + "##k"
                  }
               }
               if (counter == 0) {
                  counter = 99
               }
            }
            if (counter == 99) {
               cm.sendOk("2040052_NO_STORYBOOK")

               cm.dispose()
            } else {
               cm.sendNext("Let's see.. #b#h ##k have returned a total of #b" + counter + "#k books. The list of returned books is as follows:" + books)
            }
         } else if (status == 1) {
            cm.sendNextPrev("2040052_SETTLING_DOWN")

         } else if (status == 2) {
            cm.dispose()
         }
      }
   }
}

NPC2040052 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040052(cm: cm))
   }
   return (NPC2040052) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }