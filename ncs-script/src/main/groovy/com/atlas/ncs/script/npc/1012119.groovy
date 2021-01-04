package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1012119 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int map = 910060000
   int num = 5
   int maxPlayerCount = 5

   def start() {
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         status++
      } else {
         if (status <= 1) {
            cm.dispose()
            return
         }
         status--
      }
      if (status == 0) {
         if (cm.getLevel() >= 20) {
            cm.sendOk("1012119_UNDER_20_ONLY")
            cm.dispose()
         } else if (cm.isQuestActive(22515) || cm.isQuestActive(22516) || cm.isQuestActive(22517) || cm.isQuestActive(22518)) {
            cm.sendYesNo("1012119_WOULD_YOU_LIKE_TO_GO_SPECIAL")
            status = 1
         } else {
            String selStr = ""
            for (def i = 0; i < num; i++) {
               selStr += "\r\n#b#L" + i + "#Training Center " + i + " (" + cm.getPlayerCount(map + i) + "/" + maxPlayerCount + ")#l#k"
            }
            cm.sendSimple("1012119_WOULD_YOU_LIKE_TO_GO", selStr)
         }

      } else if (status == 1) {
         if (selection < 0 || selection >= num) {
            cm.dispose()
         } else if (cm.getPlayerCount(map + selection) >= maxPlayerCount) {
            cm.sendNext("1012119_MAP_FULL")
            status = -1
         } else {
            cm.warp(map + selection, 0)
            cm.dispose()
         }
      } else if (status == 2) {
         cm.warp(910060100, 0)
         cm.dispose()
      }
   }
}

NPC1012119 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1012119(cm: cm))
   }
   return (NPC1012119) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }