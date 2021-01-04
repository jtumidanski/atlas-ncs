package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2133004 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
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
            if (!cm.haveItem(4001163) || !cm.isEventLeader()) {
               cm.sendYesNo("2133004_LET_YOUR_PARTY_LEADER")

            } else {
               cm.sendNext("2133004_I_SHALL_SHOW_YOU_THE_PATH")

            }
         } else if (status == 1) {
            if (!cm.haveItem(4001163)) {
               cm.warp(930000800, 0)
            } else {
               cm.getEventInstance().warpEventTeam(930000600)
            }

            cm.dispose()
         }
      }
   }
}

NPC2133004 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2133004(cm: cm))
   }
   return (NPC2133004) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }