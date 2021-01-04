package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC9020002 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == 1) {
         status++
      } else {
         cm.dispose()
         return
      }
      int mapId = cm.getMapId()
      if (mapId == 103000890) {
         if (status == 0) {
            cm.sendNext("9020002_RETURN_BACK_TO_THE_CITY")

         } else {
            cm.warp(103000000)
            cm.dispose()
         }
      } else {
         if (status == 0) {
            String outText = "Once you leave the map, you'll have to restart the whole quest if you want to try it again.  Do you still want to leave this map?"
            if (mapId == 103000805) {
               outText = "Are you ready to leave this map?"
            }
            cm.sendYesNo(outText)
         } else if (mode == 1) {
            cm.warp(103000890, "st00") // Warp player
            cm.dispose()
         }
      }
   }
}

NPC9020002 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC9020002(cm: cm))
   }
   return (NPC9020002) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }