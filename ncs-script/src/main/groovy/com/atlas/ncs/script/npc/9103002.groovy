package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.EventInstanceManager
import com.atlas.ncs.processor.NPCConversationManager

class NPC9103002 {
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
      }
      if (mode == 0) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         } else {
            status--
         }
         if (status == 0) {
            cm.sendYesNo("9103002_STELLAR_EFFORT")
         } else if (status == 1) {
            EventInstanceManager eim = cm.getEventInstance()
            if (!eim.giveEventReward(cm.getCharacterId())) {
               cm.sendNext("9103002_NEED_FREE_SLOT")
            } else {
               cm.warp(809050017)
            }

            cm.dispose()
         }
      }
   }
}

NPC9103002 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9103002(cm: cm))
   }
   return (NPC9103002) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }