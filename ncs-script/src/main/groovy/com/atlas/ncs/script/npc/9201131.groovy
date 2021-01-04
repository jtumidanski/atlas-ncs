package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9201131 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int map = 677000002
   int quest = 28238
   int questItem = 4032492

   def start() {
      status = -1
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
               cm.sendYesNo("9201131_WOULD_YOU_LIKE_TO_MOVE", map)
            } else {
               cm.sendOk("9201131_HOLD_AN_EMBLEM")
               cm.dispose()
            }
         } else {
            cm.sendOk("9201131_STRANCE_FORCE")
            cm.dispose()
         }
      } else {
         if (cm.haveItem(4032481, 1)) {
            cm.gainItem(4032481, (short) -1)
         }
         if (cm.haveItem(4032482, 1)) {
            cm.gainItem(4032482, (short) -1)
         }
         if (cm.haveItem(4032483, 1)) {
            cm.gainItem(4032483, (short) -1)
         }

         cm.warp(map, 0)
         cm.dispose()
      }
   }
}

NPC9201131 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201131(cm: cm))
   }
   return (NPC9201131) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }