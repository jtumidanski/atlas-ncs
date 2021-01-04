package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2091006 {
   NPCConversationManager cm
   int status = -2
   int sel = -1
   int readNotice = 0

   def start() {
      cm.sendSimple("2091006_COURAGE_TO_CHALLENGE")

   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode == 0 && type == 0) {
         status -= 2
      }
      if (mode >= 0) {
         if (selection == 1 || readNotice == 1) {
            if (status == -1) {
               readNotice = 1
               cm.sendNext("2091006_TAKE_THE_CHALLENGE")

            } else if (status == 0) {
               cm.sendPrev("2091006_CALL_YOUR_FRIENDS")

            } else {
               cm.dispose()
            }
         } else {
            if (status == -1 && mode == 1) {
               cm.sendYesNo("2091006_MYSTERIOUS_ENERGY")

            } else if (status == 0) {
               if (mode == 0) {
                  cm.sendNext("2091006_ENERGY_DISAPPEARED")

               } else {
                  cm.saveLocation("MIRROR")
                  cm.warp(925020000, 4)
               }
               cm.dispose()
            }
         }
      } else {
         cm.dispose()
      }
   }
}

NPC2091006 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2091006(cm: cm))
   }
   return (NPC2091006) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }