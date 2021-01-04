package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9001102 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.getMapId() == 100000000) {
         cm.sendNext("9001102_UFO")
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode > 0) {
         status++
         if (cm.getMapId() == 100000000) {
            if (status == 1) {
               if (cm.getLevel() >= 12) {
                  cm.sendYesNo("9001102_WHAT_DO_WE_DO")

               } else {
                  cm.sendOk("9001102_LEVEL_REQUIREMENT")

               }

            } else if (status == 2) {
               cm.sendNext("9001102_THANK_YOU")

            } else if (status == 3) {
               cm.warp(922240200, 0)
               cm.dispose()
            }
         }
      } else if (mode < 1) {
         cm.dispose()
      }
   }
}

NPC9001102 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9001102(cm: cm))
   }
   return (NPC9001102) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }