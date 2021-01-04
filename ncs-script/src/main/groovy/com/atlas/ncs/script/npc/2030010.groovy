package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2030010 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      if (cm.getMapId() == 280030000) {
         if (!cm.getEventInstance().isEventCleared()) {
            cm.sendYesNo("2030010_LEAVE_NOW")
         } else {
            cm.sendYesNo("2030010_CONGRATULATIONS")
         }
      } else {
         cm.sendYesNo("2030010_LEAVE_NOW")
      }
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode < 1) {
         cm.dispose()
      } else {
         cm.warp(211042300)
         cm.dispose()
      }
   }
}

NPC2030010 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2030010(cm: cm))
   }
   return (NPC2030010) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }