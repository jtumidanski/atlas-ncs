package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9201134 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

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

         EventInstanceManager eim = cm.getEventInstance()
         if (!eim.isEventCleared()) {
            if (status == 0) {
               cm.sendYesNo("9201134_ARE_YOU_SURE")

            } else if (status == 1) {
               cm.warp(551030100, 2)
               cm.dispose()
            }
         } else {
            if (status == 0) {
               cm.sendNext("9201134_PRIZE_FOR_BRAVERY")

            } else if (status == 1) {
               if (!eim.giveEventReward(cm.getPlayer())) {
                  cm.sendNext("9201134_MAKE_INVENTORY_ROOM")

               } else {
                  cm.warp(551030100, 2)
               }

               cm.dispose()
            }
         }
      }
   }
}

NPC9201134 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9201134(cm: cm))
   }
   return (NPC9201134) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }