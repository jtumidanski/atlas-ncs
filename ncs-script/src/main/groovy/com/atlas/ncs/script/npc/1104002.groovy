package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1104002 {
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
         MapleMap map = cm.getMap()

         if (mode == 0 && type > 0) {
            MapleLifeFactory.getMonster(9001010).ifPresent({ monster ->
               cm.sendPinkText("ELEANOR_STILL_CHALLENGING")
               map.spawnMonsterOnGroundBelow(monster, new Point(850, 0))
               map.destroyNPC(1104002)
            })

            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            if (!cm.isQuestStarted(20407)) {
               cm.sendOk("1104002_NO_CHALLENGING")
               cm.dispose()
               return
            }

            cm.sendAcceptDecline("1104002_STILL_WANT_TO_FACE_US")
         } else if (status == 1) {
            cm.sendOk("1104002_COWARDS")
            cm.dispose()
         }
      }
   }
}

NPC1104002 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1104002(cm: cm))
   }
   return (NPC1104002) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }