package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2112016 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.isQuestStarted(3367)) {
         int c = cm.getQuestProgressInt(3367, 30)
         if (c >= 30) {
            cm.sendNext("(All files have been organized. Report the found files to Yulete.)", (byte) 2)
            cm.dispose()
            return
         }

         int book = (cm.getNpcObjectId() % 30)
         int prog = cm.getQuestProgressInt(3367, book)
         if (prog == 0) {
            c++

            if (book < 20) {
               if (!cm.canHold(4031797, 1)) {
                  cm.sendNext("2112016_ETC_IS_FULL")

                  cm.dispose()
                  return
               } else {
                  cm.gainItem(4031797, (short) 1)
                  cm.setQuestProgress(3367, 31, cm.getQuestProgressInt(3367, 31) + 1)
               }
            }

            cm.sendNext("(Organized file. #r" + (30 - c) + "#k left.)", (byte) 2)

            cm.setQuestProgress(3367, book, 1)
            cm.setQuestProgress(3367, 30, c)
         }
      }

      cm.dispose()
   }

   def action(Byte mode, Byte type, Integer selection) {

   }
}

NPC2112016 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2112016(cm: cm))
   }
   return (NPC2112016) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }