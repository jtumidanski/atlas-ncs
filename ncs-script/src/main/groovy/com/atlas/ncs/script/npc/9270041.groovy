package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9270041 {
   NPCConversationManager cm
   int status = -1
   int oldSelection = -1

   def start() {
      cm.sendSimple("9270041_HELLO")

   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode <= 0) {
         oldSelection = -1
         cm.dispose()
      }

      if (status == 0) {
         if (selection == 0) {
            cm.sendYesNo("9270041_TICKET_COST")

         } else if (selection == 1) {
            cm.sendYesNo("9270041_GO_IN_NOW")

         }
         oldSelection = selection
      } else if (status == 1) {
         if (oldSelection == 0) {
            if (cm.getMeso() > 4999 && !cm.haveItem(4031731)) {
               if (cm.canHold(4031731, 1)) {
                  cm.gainMeso(-5000)
                  cm.gainItem(4031731)
                  cm.sendOk("9270041_THANK_YOU")

                  cm.dispose()
               } else {
                  cm.sendOk("9270041_NEED_ETC_SPACE")

                  cm.dispose()
               }
            } else {
               cm.sendOk("9270041_NOT_ENOUGH_MESOS")

               cm.dispose()
            }
         } else if (oldSelection == 1) {
            if (cm.itemQuantity(4031731) > 0) {
               EventManager em = cm.getEventManager("AirPlane")
               if (em.getProperty("entry") == "true") {
                  cm.warp(540010100)
                  cm.gainItem(4031731, (short) -1)
               } else {
                  cm.sendOk("9270041_WAIT_A_FEW_MINUTES")

               }
            } else {
               cm.sendOk("9270041_NEED_A_TICKET")

            }
         }
         cm.dispose()
      }
   }
}

NPC9270041 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9270041(cm: cm))
   }
   return (NPC9270041) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }