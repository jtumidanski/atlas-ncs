package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2093004 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   String menu
   int cost = 10000

   def start() {
      cm.sendYesNo("2093004_MOVE_TO", cost)

   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0) {
            cm.sendNext("2093004_TOO_BUSY")

            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 1) {
            if (cm.getMeso() < cost) {
               cm.sendOk("2093004_NOT_ENOUGH_MESOS")

            } else {
               cm.gainMeso(-cost)
               cm.warp(230000000)
            }
            cm.dispose()
         }
      }
   }
}

NPC2093004 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2093004(cm: cm))
   }
   return (NPC2093004) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }