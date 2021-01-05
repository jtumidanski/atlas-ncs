package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9270033 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   EventInstanceManager eim

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0 && status == 0) {
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         eim = cm.getEventInstance()
         if (status == 0) {
            if (!eim.isEventCleared()) {
               cm.sendYesNo("9270033_READY_TO_LEAVE")
            } else {
               cm.sendYesNo("9270033_DEFEATED")
            }
         } else if (status == 1) {
            if (eim.isEventCleared()) {
               if (!eim.giveEventReward(cm.getCharacterId())) {
                  cm.sendOk("9270033_MAKE_INVENTORY_ROOM")
                  cm.dispose()
                  return
               }
            }

            cm.warp(541010110, 0)
            cm.dispose()
         }
      }
   }
}

NPC9270033 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9270033(cm: cm))
   }
   return (NPC9270033) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }