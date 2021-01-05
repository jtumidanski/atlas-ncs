package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9201068 {
   NPCConversationManager cm
   int status = -1
   int oldSelection = -1

   boolean close = false
   EventManager em

   def start() {
      em = cm.getEventManager("Subway").orElseThrow()
      String text = "Here's the ticket reader."
      boolean hasTicket = false
      if (cm.haveItem(4031713) && cm.getMapId() == 600010001) {
         text += "\r\n#b#L0##t4031713#"
         hasTicket = true
      }
      if (!hasTicket) {
         cm.sendOk("9201068_MISSING_TICKET")
         cm.dispose()
      } else {
         cm.sendSimple(text)
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      status++
      if (mode != 1) {
         if (mode == 0) {
            cm.sendNext("9201068_SOME_BUSINESS_HERE")
         }
         cm.dispose()
         return
      }
      if (status == 0) {
         if (selection == 0) {
            if (em.getProperty("entry") == "true") {
               cm.sendYesNo("9201068_PLENTY_OF_ROOM")
            } else {
               cm.sendNext("9201068_BE_PATIENT")
               cm.dispose()
            }
         }
         oldSelection = selection
      } else if (status == 1) {
         if (oldSelection == 0 && cm.haveItem(4031713)) {
            if (em.getProperty("entry") == "true") {
               cm.gainItem(4031713, (short) -1)
               cm.warp(600010002)
            } else {
               cm.sendNext("9201068_BE_PATIENT")
            }
         } else {
            cm.sendNext("9201068_NEED_A_TICKET")
         }

         cm.dispose()
      }
   }
}

NPC9201068 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201068(cm: cm))
   }
   return (NPC9201068) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }