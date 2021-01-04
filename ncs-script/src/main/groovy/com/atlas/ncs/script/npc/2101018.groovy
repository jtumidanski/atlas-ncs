package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2101018 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if ((cm.getLevel() < 19 || cm.getLevel() > 30) && !cm.isGM()) {
         cm.sendNext("2101018_LEVEL_RANGE")

         cm.dispose()
         return
      }
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (status == 4) {
         cm.saveLocation("MIRROR")
         cm.warp(980010000, 3)
         cm.dispose()
      }
      if (mode != 1) {
         if (mode == 0 && type == 0) {
            status -= 2
         } else {
            cm.dispose()
            return
         }
      }
      if (status == 0) {
         cm.sendNext("2101018_HUGE_FESTIVAL")

      } else if (status == 1) {
         cm.sendNextPrev("2101018_ARIANT_COLISEUM_CHALLENGE_EXPLAINED")

      } else if (status == 2) {
         cm.sendSimple("2101018_ARE_YOU_INTERESTED")

      } else if (status == 3) {
         cm.sendNext("2101018_EMERGE_VICTORIOUS")

      }
   }
}

NPC2101018 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2101018(cm: cm))
   }
   return (NPC2101018) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }