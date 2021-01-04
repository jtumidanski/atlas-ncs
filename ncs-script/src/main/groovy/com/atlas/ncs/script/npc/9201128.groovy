package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9201128 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int map = 677000004
   int quest = 28179
   int questItem = 4032491

   def start() {
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         status++
      } else {
         cm.dispose()
         return
      }
      if (status == 0) {
         if (cm.isQuestStarted(quest)) {
            if (cm.haveItem(questItem)) {
               cm.sendYesNo("9201128_WOULD_YOU_LIKE_TO_MOVE", map)

            } else {
               cm.sendOk("9201128_ENTRANCE_BLOCKED")

               cm.dispose()
            }
         } else {
            cm.sendOk("9201128_STRANGE_FORCE")

            cm.dispose()
         }
      } else {
         if (cm.haveItem(4001341, 1)) {
            cm.gainItem(4001341, (short) -1)
         }
         if (cm.haveItem(4032478, 1)) {
            cm.gainItem(4032478, (short) -1)
         }

         cm.warp(map, 0)
         cm.dispose()
      }
   }
}

NPC9201128 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201128(cm: cm))
   }
   return (NPC9201128) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }