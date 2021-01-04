package com.atlas.ncs.script.npc

import com.atlas.ncs.processor.NPCConversationManager

class NPC2040037 {
   NPCConversationManager cm
   int status = 0
   int sel = -1
   int curMap, stage

   def start() {
      curMap = cm.getMapId()
      stage = Math.floor((curMap - 922010100) / 100).intValue() + 1
      status = -1
      action((byte) 1, (byte) 0, 0)
   }

   static def clearStage(int stage, EventInstanceManager eim, int curMap) {
      eim.setProperty(stage + "stageclear", "true")
      eim.showClearEffect(true)

      eim.linkToNextStage(stage, "lpq", curMap)  //opens the portal to the next map
   }

   def action(Byte mode, Byte type, Integer selection) {
      if (mode == -1) {
         cm.dispose()
      } else if (mode == 0) {
         cm.dispose()
      } else {
         if (mode == 1) {
            status++
         } else {
            status--
         }

         EventInstanceManager eim = cm.getPlayer().getEventInstance()

         if (eim.getProperty(stage.toString() + "stageclear") != null) {
            cm.sendNext("2040037_GO_TO_NEXT_STAGE")
         } else {
            if (eim.isEventLeader(cm.getPlayer())) {
               int state = eim.getIntProperty("statusStg" + stage)

               if (state == -1) {           // preamble
                  cm.sendOk("2040037_WELCOME", stage)
                  eim.setProperty("statusStg" + stage, 0)
               } else {       // check stage completion
                  if (cm.haveItem(4001022, 15)) {
                     cm.sendOk("2040037_GOOD_JOB")
                     cm.gainItem(4001022, (short) -15)

                     eim.setProperty("statusStg" + stage, 1)
                     clearStage(stage, eim, curMap)
                  } else {
                     cm.sendNext("2040037_SORRY")
                  }
               }
            } else {
               cm.sendNext("2040037_PARTY_LEADER_MUST_TALK")
            }
         }

         cm.dispose()
      }
   }
}

NPC2040037 getNPC() {
   if (!getBinding().hasVariable("npc")) {
      NPCConversationManager cm = (NPCConversationManager) getBinding().getVariable("cm")
      getBinding().setVariable("npc", new NPC2040037(cm: cm))
   }
   return (NPC2040037) getBinding().getVariable("npc")
}

def start() {
   getNPC().start()
}

def action(Byte mode, Byte type, Integer selection) { getNPC().action(mode, type, selection) }