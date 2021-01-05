package com.atlas.ncs.script.npc

import com.atlas.ncs.model.MapObject
import com.atlas.ncs.processor.NPCConversationManager

class NPC1104102 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int npcId = 1104102
   int baseJob = 13

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
               cm.sendOk("1104102_INTRUDER_NOT_HERE")
               cm.dispose()
               return
            }

            cm.sendOk("1104102_YOU_FOUND_ME")
         } else if (status == 1) {
            MapObject mapObject = cm.getMapObject(cm.getNpcObjectId())
            cm.getMonster(9001009).ifPresent({ monster -> cm.spawnMonsterOnGroundBelow(monster, mapObject.x(), mapObject.y()) })
            cm.destroyNPC(npcId)
            cm.dispose()
         }
      }
   }
}

NPC1104102 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1104102(cm: cm))
   }
   return (NPC1104102) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }