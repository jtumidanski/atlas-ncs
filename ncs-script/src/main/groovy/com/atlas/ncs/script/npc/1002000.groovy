package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1002000 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int[] iMaps = [104000000, 102000000, 100000000, 101000000, 103000000, 120000000, 105040300]
   String[] towns = ["LITH_HARBOR", "PERION", "ELLINIA", "HENESYS", "KERNING_CITY", "NAUTALIS", "SLEEPYWOOD"]
   int[] maps = [102000000, 100000000, 101000000, 103000000, 120000000]
   int[] cost = [1000, 1000, 800, 1000, 800]
   int selectedMap = -1
   boolean town = false

   def start() {
      cm.sendNext("1002000_INIT")
   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode != 1) {
         if ((mode == 0 && !town) || mode == -1) {
            if (type == 1 && mode != -1) {
               cm.sendNext("1002000_MORE")
            }
            cm.dispose()
            return
         } else {
            status -= 2

            if (status < 1) {
               cm.dispose()
               return
            }
         }

      }
      if (status == 1) {
         cm.sendSimple("1002000_CONFUSED")
      } else if (status == 2) {
         if (selection == 0) {
            town = true
            def text = cm.evaluateToken("1002000_7_BIG_TOWN_BASE")
            for (def i = 0; i < iMaps.length; i++) {
               text += "\r\n#L" + i + "##m" + iMaps[i] + "##l"
            }
            cm.sendSimple(text)
         } else if (selection == 1) {
            def selStr = cm.evaluateToken(cm.getJobId() == 0 ? "1002000_BEGINNER_WHERE" : "1002000_NON_BEGINNER_WHERE")
            for (def i = 0; i < maps.length; i++) {
               selStr += "\r\n#L" + i + "##m" + maps[i] + "# (" + (cost[i] / (cm.getJobId() == 0 ? 10 : 1)) + " mesos)#l"
            }
            cm.sendSimple(selStr)
         }
      } else if (town) {
         if (selectedMap == -1) {
            selectedMap = selection
         }
         if (status == 3) {
            cm.sendNext("1002000_" + towns[selectedMap] + "_" + [status - 3])
         } else {
            if (cm.evaluateToken("1002000_" + towns[selectedMap] + "_" + [status - 3]) == null) {
               cm.dispose()
            } else {
               cm.sendNextPrev("1002000_" + towns[selectedMap] + "_" + [status - 3])
            }
         }
      } else if (status == 3) {
         selectedMap = selection
         cm.sendYesNo("1002000_DO_YOU_REALLY", maps[selection], (cost[selection] / (cm.getJobId() == 0 ? 10 : 1)))
      } else if (status == 4) {
         if (cm.getMeso() < (cost[selectedMap] / (cm.getJobId() == 0 ? 10 : 1))) {
            cm.sendNext("1002000_NOT_ENOUGH_MESOS")
         } else {
            cm.gainMeso((int) -(cost[selectedMap] / (cm.getJobId() == 0 ? 10 : 1)))
            cm.warp(maps[selectedMap])
         }
         cm.dispose()
      }
   }
}

NPC1002000 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1002000(cm: cm))
   }
   return (NPC1002000) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }