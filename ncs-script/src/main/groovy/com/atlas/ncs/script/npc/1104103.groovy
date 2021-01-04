package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1104103 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int npcId = 1104103
   int baseJob = 14

   def start() {
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else {
         if (mode == 0 && type > 0) {
            cm.dispose()
            return
         }
         if (mode == 1) {
            status++
         } else {
            status--
         }

         if (status == 0) {
            if (Math.floor(cm.getJobId() / 100) != baseJob) {
               cm.sendOk("1104103_INTRUDER_NOT_HERE")
               cm.dispose()
               return
            }

            cm.sendOk("1104103_YOU_FOUND_ME")
         } else if (status == 1) {
            MapleMap map = cm.getMap()
            Point npcPosition = map.getMapObject(cm.getNpcObjectId()).position()

            spawnMob(npcPosition.x, npcPosition.y, 9001009, map)
            map.destroyNPC(npcId)
            cm.dispose()
         }
      }
   }

   static def spawnMob(double x, double y, int id, MapleMap map) {
      if (map.getMonsterById(id) != null) {
         return
      }

      MapleLifeFactory.getMonster(id).ifPresent({ mob -> map.spawnMonsterOnGroundBelow(mob, new Point((int) x, (int) y)) })
   }
}

NPC1104103 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1104103(cm: cm))
   }
   return (NPC1104103) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }