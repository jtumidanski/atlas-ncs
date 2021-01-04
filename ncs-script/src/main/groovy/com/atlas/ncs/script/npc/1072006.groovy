package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1072006 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   boolean completed

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
            if (cm.haveItem(4031013, 30)) {
               completed = true
               cm.sendOk("1072006_TRUE_HERO")
            } else {
               completed = false
               cm.sendSimple("1072006_COLLECT_MARBLES")
            }
         } else if (status == 1) {
            if (completed) {
               cm.removeAll(4031013)
               cm.completeQuest(100001)
               cm.startQuest(100002)
               cm.gainItem(4031012)
            }

            cm.warp(106010000, 9)
            cm.dispose()
         }
      }
   }
}

NPC1072006 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1072006(cm: cm))
   }
   return (NPC1072006) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }