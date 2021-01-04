package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1002007 {
   NPCConversationManager cm
   int status = 0
   int sel = -1
   int[] maps = [100000000, 102000000, 101000000, 103000000, 120000000]
   int[] cost = [1000, 1000, 800, 1000, 800]
   int selectedMap = -1
   int mesos

   def start() {
      if (cm.hasItem(4032313, 1)) {
         cm.sendNext("1002007_HENESYS_COUPON")
      } else {
         cm.sendNext("1002007_HELLO")
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (status == 1 && mode == 0) {
            cm.dispose()
            return
         } else if (status >= 2 && mode == 0) {
            cm.sendNext("1002007_A_LOT_TO_SEE")
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 1) {
            if (cm.hasItem(4032313, 1)) {
               cm.gainItem(4032313, (short) -1)
               cm.warp(maps[0], 0)
               cm.dispose()
               return
            }

            def selStr = ""
            if (cm.getJobId() == 0) {
               selStr += cm.evaluateToken("1002007_BEGINNER_DISCOUNT")
            }
            selStr += cm.evaluateToken("1002007_CHOOSE")
            for (def i = 0; i < maps.length; i++) {
               selStr += "\r\n#L" + i + "##m" + maps[i] + "# (" + (cm.getJobId() == 0 ? cost[i] / 10 : cost[i]) + " mesos)#l"
            }
            cm.sendSimple(selStr)
         } else if (status == 2) {
            cm.sendYesNo("1002007_NOTHING_ELSE_TO_DO", maps[selection], (cm.getJobId() == 0 ? cost[selection] / 10 : cost[selection]))
            selectedMap = selection
         } else if (status == 3) {
            if (cm.getJobId() == 0) {
               mesos = (int) (cost[selectedMap] / 10)
            } else {
               mesos = cost[selectedMap]
            }

            if (cm.getMeso() < mesos) {
               cm.sendNext("1002007_NOT_ENOUGH_MESO")
               cm.dispose()
               return
            }

            cm.gainMeso(-mesos)
            cm.warp(maps[selectedMap], 0)
            cm.dispose()
         }
      }
   }
}

NPC1002007 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1002007(cm: cm))
   }
   return (NPC1002007) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }