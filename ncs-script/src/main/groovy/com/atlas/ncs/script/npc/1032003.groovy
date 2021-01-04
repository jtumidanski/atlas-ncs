package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1032003 {
   NPCConversationManager cm
   int status = 0
   int sel = -1

   int check = 0

   def start() {
      if (cm.getLevel() < 25) {
         cm.sendOk("1032003_LEVEL_REQUIREMENT")
         cm.dispose()
         check = 1
      } else {
         cm.sendYesNo("1032003_FOREST_OF_PATIENCE_FEE")
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0) {
            cm.sendOk("1032003_SEE_YOU_NEXT_TIME")
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 1) {
            if (check != 1) {
               if (cm.getMeso() < 5000) {
                  cm.sendOk("1032003_NOT_ENOUGH_MESO")
                  cm.dispose()
               } else {
                  if (cm.isQuestStarted(2050)) {
                     cm.warp(101000100, 0)
                  } else if (cm.isQuestStarted(2051)) {
                     cm.warp(101000102, 0)
                  } else if (cm.getLevel() >= 25 && cm.getLevel() < 50) {
                     cm.warp(101000100, 0)
                  } else if (cm.getLevel() >= 50) {
                     cm.warp(101000102, 0)
                  }
                  cm.gainMeso(-5000)
                  cm.dispose()
               }
            }
         }
      }
   }
}

NPC1032003 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1032003(cm: cm))
   }
   return (NPC1032003) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }