package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1032109 {
   NPCConversationManager cm
   int status = -1
   int sel = -1

   int mobId = 2220100 //Blue Mushroom

   def start() {
      if (!cm.isQuestStarted(20718)) {
         cm.dispose()
         return
      }

      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1 || (mode == 0 && status == 0)) {
         cm.dispose()
         return
      } else if (mode == 0) {
         status--
      } else {
         status++
      }


      if (status == 0) {
         cm.sendOk("1032109_A_LOT_OF_ANGRY_MONSTERS_SUMMONED")
      } else if (status == 1) {
         for (def i = 0; i < 10; i++) {
            cm.getMonster(mobId).ifPresent({ monster -> cm.spawnMonsterOnGroundBelow(monster, 117, 183) })
         }
         for (def i = 0; i < 10; i++) {
            cm.getMonster(mobId).ifPresent({ monster -> cm.spawnMonsterOnGroundBelow(monster, 4, 183) })
         }
         for (def i = 0; i < 10; i++) {
            cm.getMonster(mobId).ifPresent({ monster -> cm.spawnMonsterOnGroundBelow(monster, -109, 183) })
         }

         cm.completeQuest(20718, 1103003)
         cm.gainExp(4000 * cm.getExpRate())
         cm.dispose()
      }
   }
}

NPC1032109 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1032109(cm: cm))
   }
   return (NPC1032109) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }