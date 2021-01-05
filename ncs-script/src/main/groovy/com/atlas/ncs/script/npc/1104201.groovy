package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC1104201 {
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
            if (!(cm.isQuestCompleted(20407) || cm.isQuestStarted(20407) && cm.getQuestProgressInt(20407, 9001010) != 0) && cm.getMapMonsterCount(9001010) == 0 && cm.getNpcById(1104002).isEmpty()) {
               cm.sendOk("1104201_SHES_ALREADY_HERE")
               cm.spawnNpc(1104002, 850, 0)
            } else {
               cm.sendOk("...")
            }

            cm.dispose()
         }
      }
   }
}

NPC1104201 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC1104201(cm: cm))
   }
   return (NPC1104201) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }